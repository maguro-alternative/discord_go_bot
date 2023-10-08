package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	botRouter "github.com/maguro-alternative/discord_go_bot/bot_handler/bot_router"
	"github.com/maguro-alternative/discord_go_bot/commands"
	"github.com/maguro-alternative/discord_go_bot/server_handler/router"
	"github.com/maguro-alternative/discord_go_bot/model/envconfig"

	"github.com/bwmarrin/discordgo"
)

// セッションの定義
var discord *discordgo.Session

func main() {
	//Discordのセッションを作成
	env,err := envconfig.NewEnv()
	if err != nil {
		fmt.Println("error loading env")
		env = &envconfig.Env{
			TOKEN:            os.Getenv("TOKEN"),
			ServerPort:       os.Getenv("PORT"),
		}
	}
	Token := "Bot " + env.TOKEN //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	discord, err := discordgo.New(Token)

	// 権限追加
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}
	// websocketを開いてlistening開始
	if err = discord.Open(); err != nil {
		fmt.Println(err)
		panic("Error while opening session")
	}

	// ハンドラーの登録
	botRouter.RegisterHandlers(discord)

	var commandHandlers []*botRouter.Handler
	// 所属しているサーバすべてにスラッシュコマンドを追加する
	// NewCommandHandlerの第二引数を空にすることで、グローバルでの使用を許可する
	commandHandler := botRouter.NewCommandHandler(discord, "")
	// 追加したいコマンドをここに追加
	commandHandler.CommandRegister(commands.PingCommand())
	commandHandler.CommandRegister(commands.RecordCommand())
	commandHandler.CommandRegister(commands.DisconnectCommand())
	commandHandlers = append(commandHandlers, commandHandler)

	fmt.Println("Discordに接続しました。")
	fmt.Println("終了するにはCtrl+Cを押してください。")

	// サーバーの待ち受けを開始(ゴルーチンで非同期処理)
	// ここでサーバーを起動すると、Ctrl+Cで終了するまでサーバーが起動し続ける
	go func() {
		const (
			defaultPort   = "8080"
		)

		port := env.ServerPort
		if port == "" {
			port = defaultPort
		}
		port = ":" + port

		mux := router.NewRouter(discord)
		log.Printf("Serving HTTP port: %s\n", port)
		log.Fatal(http.ListenAndServe(port, mux))
	}()

	// Ctrl+Cを受け取るためのチャンネル
	sc := make(chan os.Signal, 1)
	// Ctrl+Cを受け取る
	signal.Notify(sc, os.Interrupt)
	<-sc //プログラムが終了しないようロック

	fmt.Println("Removing commands...")

	// コマンドを削除
	for i := range commandHandlers {
		// すべてのコマンドを取得
		commands := commandHandlers[i].GetCommands()
		for _, command := range commands {
			err := commandHandlers[i].CommandRemove(command)
			if err != nil {
				panic("error removing command")
			}
		}
	}

	// websocketを閉じる
	err = discord.Close()
	if err != nil {
		panic("error closing connection")
	}
	fmt.Println("Disconnected")
}

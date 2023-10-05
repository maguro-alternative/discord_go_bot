package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	handlers "github.com/maguro-alternative/discord_go_bot/bot_handler"
	"github.com/maguro-alternative/discord_go_bot/commands"
	"github.com/maguro-alternative/discord_go_bot/db"
	"github.com/maguro-alternative/discord_go_bot/server_handler/router"
	"github.com/maguro-alternative/discord_go_bot/model"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     = "Bot " + os.Getenv("TOKEN") //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	stopBot   = make(chan bool)
	vcsession *discordgo.VoiceConnection
)

// セッションの定義
var discord *discordgo.Session

func main() {
	//Discordのセッションを作成
	env,err := model.NewEnv()
	Token = "Bot " + env.TOKEN //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
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
		panic("Error while opening session")
	}

	// ハンドラーの登録
	handlers.RegisterHandlers(discord)

	var commandHandlers []*handlers.Handler
	// 所属しているサーバすべてにスラッシュコマンドを追加する
	commandHandler := handlers.NewCommandHandler(discord, "")
	// 追加したいコマンドをここに追加
	commandHandler.CommandRegister(commands.PingCommand())
	commandHandler.CommandRegister(commands.RecordCommand())
	commandHandler.CommandRegister(commands.DisconnectCommand())
	commandHandlers = append(commandHandlers, commandHandler)

	fmt.Println("Discordに接続しました。")
	fmt.Println("終了するにはCtrl+Cを押してください。")

	// サーバーの待ち受けを開始
	go func() {
		const (
			defaultPort   = ":8080"
			defaultDBPath = ".sqlite3/todo.db"
		)

		port := env.ServerPort
		if port == "" {
			port = defaultPort
		}
		dbPath := "host=" + env.DatabaseHost + " port=" + env.DatabasePort + " user=" + env.DatabaseUser + " dbname=" + env.DatabaseName + " password=" + env.DatabasePassword //+ " sslmode=disable"

		indexDB, err := db.NewPostgresDB(dbPath)
		if err != nil {
			fmt.Println(err)
		}

		mux := router.NewRouter(indexDB)
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

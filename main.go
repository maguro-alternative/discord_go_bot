package main

import (
	"fmt"
	"os"

	"github.com/maguro-alternative/discord_go_bot/commands"
	"github.com/maguro-alternative/discord_go_bot/handlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
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
	errr := godotenv.Load()
	Token = "Bot " + os.Getenv("TOKEN") //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	if errr != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
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

	// 所属しているサーバすべてを取得
	guilds, err := discord.UserGuilds(100, "", "")
	var commandHandlers []*handlers.Handler
	// 所属しているサーバすべてにスラッシュコマンドを追加する
	for _, guild := range guilds {
		commandHandler := handlers.NewCommandHandler(discord, guild.ID)
		commandHandler.CommandRegister(commands.PingCommand())
		commandHandlers = append(commandHandlers, commandHandler)
	}

	fmt.Println("Discordに接続しました。")
	fmt.Println("終了するにはCtrl+Cを押してください。")
	//sc := make(chan os.Signal, 1)
	//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot //プログラムが終了しないようロック

	fmt.Println("Removing commands...")

	// コマンドを削除
	for i := range guilds {
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

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
	Token = "Bot " + os.Getenv("TOKEN") //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	stopBot   = make(chan bool)
	vcsession *discordgo.VoiceConnection
)

// スラッシュコマンドの追加
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

	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer discord.Close()

	fmt.Println("Discordに接続しました。")
	fmt.Println("終了するにはCtrl+Cを押してください。")
	//sc := make(chan os.Signal, 1)
	//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot //プログラムが終了しないようロック

	err = discord.Close()

	fmt.Println("Removing commands...")

	for i := range guilds {
		commands := commandHandlers[i].GetCommands()
		//fmt.Println(guild)
		for _, command := range commands {
			err := commandHandlers[i].CommandRemove(command)
			if err != nil {
				panic("error removing command")
			}
		}
	}

	return
}

package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/maguro-alternative/discord_go_bot/handlers"
)

func PingCommand() *handlers.Command {
	/*
		pingコマンドの定義

		コマンド名: ping
		説明: Pong!
		オプション: なし
	*/
	return &handlers.Command{
		Name:        "ping",
		Description: "Pong!",
		Options:     []*discordgo.ApplicationCommandOption{},
		Executor:    handlePing,
	}
}

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	/*
		pingコマンドの実行

		コマンドの実行結果を返す
	*/
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println(err)
	}
	for _, guild := range guilds {
		if guild.ID == i.GuildID {
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong",
				},
			})
			if err != nil {
				fmt.Printf("error responding to ping command: %v\n", err)
			}
		}
	}
}

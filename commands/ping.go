package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	botRouter "github.com/maguro-alternative/discord_go_bot/bot_handler/bot_router"
)

func PingCommand() *botRouter.Command {
	/*
		pingコマンドの定義

		コマンド名: ping
		説明: Pong!
		オプション: なし
	*/
	return &botRouter.Command{
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
	if i.Interaction.ApplicationCommandData().Name != "ping" {
		return
	}
	if i.Interaction.GuildID != i.GuildID {
		return
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})
	if err != nil {
		fmt.Printf("error responding to ping command: %v\n", err)
	}

}

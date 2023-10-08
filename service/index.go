package service

import (
	"github.com/bwmarrin/discordgo"
)

type IndexService struct {
	DiscordSession *discordgo.Session
}

// IndexServiceを返す
func NewIndexService(discordSession *discordgo.Session) *IndexService {
	return &IndexService{
		DiscordSession : discordSession,
	}
}

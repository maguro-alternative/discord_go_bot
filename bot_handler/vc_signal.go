package botHandler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func OnVoiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate ) {
    fmt.Print("hoge")
    fmt.Printf("%+v", vs.VoiceState)
}
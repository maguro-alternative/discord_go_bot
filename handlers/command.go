package handlers

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Options     []*discordgo.ApplicationCommandOption
	AppCommand  *discordgo.ApplicationCommand
	Executor    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (c *Command) AddApplicationCommand(appCmd *discordgo.ApplicationCommand) {
	c.AppCommand = appCmd
}
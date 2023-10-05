package botRouter

import "github.com/bwmarrin/discordgo"

/*
スラッシュコマンドのハンドラ

スラッシュコマンドのハンドラは、
discordgo.Session.AddHandler()で登録する必要があります。

discordgo.Session.AddHandler()の引数には、
discordgo.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command.Executor(s, i)
}
のように、
discordgo.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command.Executor(s, i)
}
を渡す必要があります。
*/

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

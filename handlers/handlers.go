// handlers.go
package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	session  *discordgo.Session
	commands map[string]*Command
	guild    string
}

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(OnMessageCreate)
}

func NewCommandHandler(session *discordgo.Session, guildID string) *Handler {
	return &Handler{
		session:  session,
		commands: make(map[string]*Command),
		guild:    guildID,
	}
}

func (h *Handler) CommandRegister(command *Command) error {
	if _, exists := h.commands[command.Name]; exists {
		return fmt.Errorf("command with name `%s` already exists", command.Name)
	}

	appCmd, err := h.session.ApplicationCommandCreate(
		h.session.State.User.ID,
		h.guild,
		&discordgo.ApplicationCommand{
			ID:            command.Name,
			ApplicationID: h.session.State.User.ID,
			Name:          command.Name,
			Description:   command.Description,
			Options:       command.Options,
		},
	)

	if err != nil {
		return err
	}

	command.AddApplicationCommand(appCmd)

	h.commands[command.Name] = command

	h.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		command.Executor(s, i)
	})

	return nil
}

func (h *Handler) CommandRemove(command *Command) error {
	err := h.session.ApplicationCommandDelete(h.session.State.User.ID, h.guild, command.AppCommand.ID)
	if err != nil {
		return fmt.Errorf("error while deleting application command: %v", err)
	}

	delete(h.commands, command.Name)

	return nil
}

func (h *Handler) GetCommands() []*Command {
	var commands []*Command

	for _, v := range h.commands {
		commands = append(commands, v)
	}

	return commands
}

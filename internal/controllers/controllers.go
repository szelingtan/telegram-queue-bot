package controllers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/josh1248/nusc-queue-bot/internal/commandtypes"
	"github.com/josh1248/nusc-queue-bot/internal/handlers"
)

// receives a user message and returns with message to be sent.
func ReceiveCommand(userMessage tgbotapi.Update) (reply tgbotapi.MessageConfig) {
	reply = tgbotapi.NewMessage(userMessage.Message.Chat.ID, "")
	username := userMessage.SentFrom().UserName

	if userMessage.Message.Document != nil {
		log.Printf("Invalid non-text message received from user @%s\n", username)
		reply.Text = handlers.NonTextHandler(userMessage)
	} else if !userMessage.Message.IsCommand() {
		log.Printf("Invalid output of %s received from user @%s\n", userMessage.Message.Text, username)
		reply.Text = handlers.NonCommandHandler(userMessage)
	} else {
		log.Printf("Processing command %s from @%s...\n", userMessage.Message.Text, username)
		for _, command := range commandtypes.AvailableCommands {
			if userMessage.Message.Command() == command.Command {
				reply.Text = command.Handler(userMessage)
				break
			}
		}

		if reply.Text == "" {
			reply.Text = handlers.InvalidCommand(userMessage)
		}

		log.Printf("Processed command %s from @%s.\n", userMessage.Message.Text, username)
	}

	return reply
}

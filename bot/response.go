package inibot

import (

	//"go-discord-bot/scraper"

	"go-discord-bot/scraper"
	tools "go-discord-bot/utilitys"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//empeza todo con traer el link del que solicitó
	link := tools.Link(m.Author.ID)
	if link == "" {
		// Manda un mensaje indicando que no hay link
		_, _ = s.ChannelMessageSend(m.ChannelID, "Usted no tiene link D:")
		return
	}
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "<monday" && m.ChannelID == "1312559101380657152" {
		model := scraper.InicScraper(link)

		//esta weá crea una especie de canal privado de bot a tu, por qué lo complican tanto D:
		channel, _ := s.UserChannelCreate(m.Author.ID)
		order := tools.Send(model)
		for _, x := range order {
			_, _ = s.ChannelMessageSendEmbed(channel.ID, &x)
		}

	}
}

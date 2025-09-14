package inibot

import (

	//"go-discord-bot/scraper"

	"fmt"
	"go-discord-bot/scraper"
	tools "go-discord-bot/utilitys"
	"log"

	"github.com/bwmarrin/discordgo"
)

// aquí va según lo que se interactua en el dc, esto es por los mensajes
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "<monday" {
		////////hay que ordenar xdd, la idea es tenmer todo dentro de una funcion///////
		//empeza todo con traer el link del que solicitó
		//tools.link no va a irse a github, ya que acá tenemos los id y las url de los compañeros de trabajo
		//básicamente lo que tiene es un diccionario con key: user_id, value: link monday con filtros
		link := tools.Link(m.Author.ID)

		//esta weá crea una especie de canal privado de bot a tu, por qué lo complican tanto D:
		channel, _ := s.UserChannelCreate(m.Author.ID)
		if link == "" {
			// Si no está en la lista, mandamos un mensaje diciendo que no tiene link
			_, _ = s.ChannelMessageSend(channel.ID, "Usted no tiene link D:")
			return
		}
		//si tiene link, empezamos la scrapeada xd
		model := scraper.InicScraper(link, s, m, "", m.Author.ID)

		order := tools.Send(model)
		for _, x := range order {
			_, _ = s.ChannelMessageSendEmbed(channel.ID, &x)
		}

	} else if m.Content == "<mondayy" {
		tools.AwaitName(s, m)
	}
}

// acá cuando apretan algún botón
func ButtonInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}
	////////hay que ordenar xdd, la idea es tenmer todo dentro de uan funcion
	//que empiece a pensar
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Println("Error respondiendo interacción:", err)
	}
	var userID string

	if i.Member != nil {
		userID = i.Member.User.ID
	} else if i.User != nil {
		userID = i.User.ID
	}

	idButton := i.MessageComponentData().CustomID
	link := tools.Link(idButton)

	channel, _ := s.UserChannelCreate(userID)
	if link == "" {
		// Si no está en la lista, mandamos un mensaje diciendo que no tiene link
		_, _ = s.ChannelMessageSend(i.ChannelID, "Usted no tiene link D:")
		return
	}
	model := scraper.InicScraper(link, s, nil, idButton, userID)
	order := tools.Send(model)
	for _, x := range order {
		_, _ = s.ChannelMessageSendEmbed(channel.ID, &x)

	}
	//aca deja de pensar
	msg := "redi"
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &msg,
	})
	if err != nil {
		fmt.Println("Error edit:", err)
	}
}

func Voice(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {

	fmt.Println("no entre")

}

package tools

import (
	"fmt"
	types "go-discord-bot/types"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Send(model []types.Developer) []discordgo.MessageEmbed {
	listembed := []discordgo.MessageEmbed{}
	num := 1
	for _, sep := range model {
		randomColor := rand.Intn(0xFFFFFF)
		a := discordgo.MessageEmbed{
			Title:       "Tarea: " + strconv.Itoa(num),
			Description: "[Click aquí](" + sep.Link + ")",
			Color:       randomColor,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Tarea",
					Value:  sep.Tarea,
					Inline: true,
				},
				{
					Name:   "Cliente",
					Value:  sep.Cliente,
					Inline: false,
				},
				{
					Name:   "Prioridad",
					Value:  sep.Prioridad,
					Inline: true,
				},
				{
					Name:   "Estado",
					Value:  sep.Estado,
					Inline: true,
				},
				{
					Name:   "EstadoKanban",
					Value:  sep.Estadokanvan,
					Inline: true,
				},
				{
					Name:   "Fecha Requerimiento",
					Value:  sep.FechaRe,
					Inline: true,
				},
				{
					Name:   "Fecha Inicio",
					Value:  sep.FechaIni,
					Inline: true,
				},
				{
					Name:   "Feca fin Tarea",
					Value:  sep.FechaFin,
					Inline: true,
				},
				{
					Name:   "Días de la tarea",
					Value:  sep.DiasDesarro,
					Inline: true,
				},
				{
					Name:   "Desarrolador",
					Value:  sep.Developer,
					Inline: true,
				},
			},
		}
		num++
		listembed = append(listembed, a)
	}

	return listembed
}

func AwaitName(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, _ := s.UserChannelCreate(m.Author.ID)

	//traemos botones con los nombres/id
	button := PersoButton()
	//agregamos los botones primera fila
	actionRow0 := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button[0], button[1], button[2], button[4], button[5]},
	}
	//agregamos los botones segunda fila
	actionRow1 := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button[3], button[6]},
	}
	//creamos el embed
	embeds := []*discordgo.MessageEmbed{{
		Title:       "Clickea un nombre",
		Description: "Te traerá las tareas pendientes junto al link de su filtro.",
		Color:       0x1ABC9C,
	}}

	//agregamos los botones
	_, err := s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Embeds:     embeds,
		Content:    "Clickea el nombre de la persona que necesites.",
		Components: []discordgo.MessageComponent{actionRow0, actionRow1},
	})

	if err != nil {
		fmt.Println("algo falló al crear la wea de los botones")
	}

}

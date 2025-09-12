package tools

import (
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

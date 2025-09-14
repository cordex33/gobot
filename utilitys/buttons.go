package tools

import "github.com/bwmarrin/discordgo"

func PersoButton() []discordgo.Button {
	names := []string{"Matías", "Vladi", "Claudio", "Gianinno", "Juan", "Ricardo", "Sin Desarrolador"}
	ids := []string{"1303685483607494739", "1222916217636126745", "1154144508490043482", "1266507782446649436", "1313159721061974158", "1067834582029844521", "S/D"}
	button := []discordgo.Button{}
	for x := 0; x < len(names); x++ {
		a := discordgo.Button{
			CustomID: ids[x],
			Label:    names[x],
			Style:    discordgo.SecondaryButton,
		}
		button = append(button, a)
	}

	return button
}

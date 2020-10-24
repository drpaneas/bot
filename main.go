package main

import (
	"github.com/ecnepsnai/discord"
)

func postDiscord(text, author, link, title, desc, imageLink string) {
	var pic discord.Image
	pic.URL = imageLink
	discord.WebhookURL = "PUT THE WEBHOOK URL TAKEN FROM THE ROOM"
	discord.Post(discord.PostOptions{
		// Content: text,
		Embeds: []discord.Embed{
			{
				Color:       16777215,
				URL:         link,
				Title:       title,
				Description: desc,
				Thumbnail:   &pic,
			},
		},
	})
}

func main() {
	text := "This is text"
	author := "drpaneas"
	link := "https://www.alfavita.gr/epistimi/335725_i-gi-exoplanitis-pos-tha-eblepan-ton-kosmo-mas-exogiinoi-apo-1000-astra"
	title := "Η Γη εξωπλανήτης; Πώς θα έβλεπαν τον κόσμο μας εξωγήινοι από 1.000 άστρα"
	desc := "Ερευνητές αντιστρέφουν μία κλασική απορία, ρωτώντας πώς θα φαινόταν ο πλανήτης Γη σε εξωγήινα όντα 1.000 και πλέον άστρων	"
	imageLink := "https://starlord.gr/images/post/exoplanets.jpg"
	postDiscord(text, author, link, title, desc, imageLink)
}

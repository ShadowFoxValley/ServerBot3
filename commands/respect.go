package commands

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)



func (data Commands) respect() {
	args := strings.Split(data.message.Content, " ")
	var title string

	if cap(data.message.Mentions) > 0 {
		title = "Pay respect for " + data.message.Mentions[0].Username
	} else if cap(args) > 1 {
		title = "Pay respect for " + strings.Join(args[1:], " ")
	} else {
		title = "Pay respect for " + data.message.Author.Username
	}
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: data.message.Author.AvatarURL("50x50"),
			Name:    data.message.Author.Username,
		},
		Title: title,
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Payed respect",
				Value:  data.mainSession.State.User.Username,
				Inline: false,
			},
		},
	}
	message, err := data.mainSession.ChannelMessageSendEmbed(data.message.ChannelID, embed)
	if err != nil {
		log.Println(err)
	}
	// Добавляем эмоут на сообщение
	go func(m *discordgo.Message, s *discordgo.Session) {
		s.MessageReactionAdd(m.ChannelID, m.ID, "🇫")
	}(message, data.mainSession)

	trackReactions[message.ID] = message
	time.AfterFunc(time.Duration(10)*time.Minute, func() {
		data.mainSession.MessageReactionsRemoveAll(trackReactions[message.ID].ChannelID, trackReactions[message.ID].ID)
		// Выставляем дефолтный цвет для эмбеда
		trackReactions[message.ID].Embeds[0].Color = 0x0
		data.mainSession.ChannelMessageEditEmbed(message.ChannelID, message.ID, trackReactions[message.ID].Embeds[0])
		delete(trackReactions, message.ID)
	})
}

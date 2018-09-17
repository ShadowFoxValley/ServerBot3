package commands

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"log"
	)

var emojiPoll = []string{
	"0⃣",
	"1⃣",
	"2⃣",
	"3⃣",
	"4⃣",
	"5⃣",
	"6⃣",
	"7⃣",
	"8⃣",
	"9⃣",
}

func (data Commands) poll() {
	var (
		messageData = strings.Split(data.message.Content, "\n")
		variants = messageData[1:]
		VariantField string
	)

	for i:=range variants{
		VariantField += strconv.Itoa(i) + ") " + variants[i] + "\n"
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{Name: data.message.Author.Username},
		Color:  0x00ff00,
		Fields: []*discordgo.MessageEmbedField{{
				Name:   "Варианты",
				Value:  VariantField,
				Inline: true,
			},
		},
	}

	message, sendError := data.mainSession.ChannelMessageSendEmbed(data.channelId, embed)
	if sendError != nil {
		log.Print(sendError.Error())
	}

	go func(s *discordgo.Session, m *discordgo.Message, num int){
		for i:=0; i < num; i++{
			s.MessageReactionAdd(m.ChannelID, m.ID, emojiPoll[i])
		}
	}(data.mainSession, message, len(variants))

	data.mainSession.ChannelMessageDelete(data.channelId, data.message.ID)
}
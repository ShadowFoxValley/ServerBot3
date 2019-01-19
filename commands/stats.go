package commands

import (
	"ServerBot3/pointSystem"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func (data Commands) stats() {
	var elements = strings.Split(data.message.Content, " ")

	if len(elements) == 2 {
		var points, check = pointSystem.GetUserPoints(data.message.Author.ID)

		if !check {
			data.mainSession.ChannelMessageSend(data.message.ChannelID, "Error while getting data")
			return
		}

		data.mainSession.ChannelMessageSend(data.message.ChannelID, "You have "+strconv.Itoa(points)+" stars")
	} else if len(elements) == 3 {
		if elements[2] == "top" {
			var userTopList, check = pointSystem.GetStarPointTop()
			if !check {
				data.mainSession.ChannelMessageSend(data.message.ChannelID, "Error while getting data")
				return
			}

			var username string
			var points int
			var inline = false
			var counter = 0

			var fields []*discordgo.MessageEmbedField

			for userTopList.Next() {
				userTopList.Scan(&username, &points)
				counter += 1
				if counter > 1 {
					inline = true
				} else {
					username += "⭐"
				}

				var tmp = &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("%d - %s", counter, username),
					Value:  strconv.Itoa(points),
					Inline: inline,
				}
				fields = append(fields, tmp)
			}

			embed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{Name: data.message.Author.Username},
				Color:  0x00ff00,
				Fields: fields,
			}
			data.mainSession.ChannelMessageSendEmbed(data.channelId, embed)
		}
	}

}

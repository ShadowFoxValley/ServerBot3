package commands

import (
	"ServerBot3/members"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (data Commands) muteManipulate(typeCommand string) {

	// Start check permissions
	var authorAllInfo, errorGetUser = data.mainSession.GuildMember(data.guild.ID, data.message.Author.ID)
	if errorGetUser != nil {
		data.mainSession.ChannelMessageSend(data.channelId, errorGetUser.Error())
	}

	var checkFunction = func(member *discordgo.Member) bool { // Ну, вроде удобно
		for i := range member.Roles {
			if member.Roles[i] == "374707108764975109" || member.Roles[i] == "485056864309084160" {
				return true
			}
		}
		return false
	}
	// End check permissions

	if checkFunction(authorAllInfo) {
		var targets = data.message.Mentions

		if len(targets) == 0 {
			data.mainSession.ChannelMessageSend(data.channelId, "Выбери цель")
			return
		}

		if typeCommand == "mute" {

			for i := range targets {
				data.mainSession.GuildMemberRoleAdd(data.guild.ID, targets[i].ID, "375707226800652290")
			}

		} else if typeCommand == "unmute" {

			for i := range targets {
				data.mainSession.GuildMemberRoleRemove(data.guild.ID, targets[i].ID, "375707226800652290")
			}

		}

	} else {
		data.mainSession.ChannelMessageSend(data.channelId, "You must be root for this")
	}

}

func (data Commands) root() {

	var elements = strings.Split(data.message.Content, " ")
	if len(elements) < 3 {
		data.mainSession.ChannelMessageSend(data.channelId, "Syntax error! Usage: ``sudo root mute/unmute mention``")
	}

	if elements[2] == "mute" || elements[2] == "unmute" {
		data.muteManipulate(elements[2])
		return
	} else if elements[2] == "users" {
		if elements[3] == "update" {
			members.UpdateUsers(data.mainSession, data.guild.ID)
		}
	}

}

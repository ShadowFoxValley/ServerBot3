package commands

import (
	"ServerBot3/common"
	"ServerBot3/members"
	"strconv"
	"strings"
)

func (data Commands) muteManipulate(typeCommand string) {

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

}

func (data Commands) root() {
	// Start check permissions
	var authorAllInfo, errorGetUser = data.mainSession.GuildMember(data.guild.ID, data.message.Author.ID)
	if errorGetUser != nil {
		data.mainSession.ChannelMessageSend(data.channelId, errorGetUser.Error())
	}

	if !common.CheckPermit(authorAllInfo) {
		data.mainSession.ChannelMessageSend(data.channelId, "You must be root for this")
		return
	}

	var elements = strings.Split(data.message.Content, " ")
	if len(elements) < 3 {
		data.mainSession.ChannelMessageSend(data.channelId, "Syntax error! Usage: ``sudo root mute/unmute mention``")
		return
	}

	if elements[2] == "mute" || elements[2] == "unmute" {

		data.muteManipulate(elements[2])
		return

	} else if elements[2] == "users" {

		if elements[3] == "update" {
			members.UpdateUsers(data.mainSession, data.guild.ID)
		}

	} else if elements[2] == "clear" {

		if len(elements) < 4 {
			data.mainSession.ChannelMessageSend(data.channelId, "Syntax error! Usage: ``sudo root clear NUM``")
			return
		}

		var num, err = strconv.Atoi(elements[3])
		if err != nil {
			data.mainSession.ChannelMessageSend(data.channelId, err.Error())
		}
		if num > 98 {
			data.mainSession.ChannelMessageSend(data.channelId, "max 100")
		}

		var messagesForDelete, errorGetMessages = data.mainSession.ChannelMessages(data.channelId, num+1, "", "", "")
		if errorGetMessages != nil {
			data.mainSession.ChannelMessageSend(data.channelId, errorGetMessages.Error())
		}

		for i:=range messagesForDelete{
			go data.mainSession.ChannelMessageDelete(data.channelId, messagesForDelete[i].ID)
		}
	}

}

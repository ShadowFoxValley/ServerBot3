package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (data Commands) roles() {
	var elements = strings.Split(data.message.Content, " ")
	if len(elements) < 4 {
		data.mainSession.ChannelMessageSend(data.channelId, "Syntax error! Usage: ``sudo roles get/remove %rolename%``")
		return
	}

	var rolesList, errorGetRoles = data.mainSession.GuildRoles(data.guild.ID)
	if errorGetRoles != nil {
		data.mainSession.ChannelMessageSend(data.channelId, errorGetRoles.Error())
	}

	var targetRoleString = strings.Join(elements[3:], " ")

	if targetRoleString == "Герой" {
		data.mainSession.ChannelMessageSend(data.channelId, "Ты не можешь удалить основную роль")
		return
	}

	var roleId, check, _ = func(rolesList []*discordgo.Role, target string) (string, bool, int) {
		var heroPosition int

		for i := range rolesList {
			if rolesList[i].ID == "374901853332176898" {
				heroPosition = rolesList[i].Position
				break
			}
		}

		for i := range rolesList {
			if strings.ToLower(rolesList[i].Name) == strings.ToLower(target) && rolesList[i].Position < heroPosition {
				return rolesList[i].ID, true, rolesList[i].Position
			}
		}

		return "", false, 0
	}(rolesList, targetRoleString)

	if check == false {
		data.mainSession.ChannelMessageSend(data.channelId, "Роль не найдена")
		return
	}

	//data.mainSession.ChannelMessageSend(data.channelId, strconv.Itoa(position)+" ``"+targetRoleString+"``")
	//return

	if elements[2] == "get" {
		//data.mainSession.ChannelMessageSend(data.channelId, "You wanna get role?")
		var errorRoleAdd = data.mainSession.GuildMemberRoleAdd(data.guild.ID, data.message.Author.ID, roleId)

		if errorRoleAdd != nil {
			data.mainSession.ChannelMessageSend(data.channelId, errorRoleAdd.Error())
		} else {
			data.mainSession.ChannelMessageSend(data.channelId, "Готово, выдал роль")
		}
	}

	if elements[2] == "remove" {
		var errorRoleAdd = data.mainSession.GuildMemberRoleRemove(data.guild.ID, data.message.Author.ID, roleId)

		if errorRoleAdd != nil {
			data.mainSession.ChannelMessageSend(data.channelId, errorRoleAdd.Error())
		} else {
			data.mainSession.ChannelMessageSend(data.channelId, "Готово, убрал роль")
		}
	}

}

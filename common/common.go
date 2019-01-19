package common

import "github.com/bwmarrin/discordgo"

func CheckPermit(member *discordgo.Member) bool { // Ну, вроде удобно
	for i := range member.Roles {
		if member.Roles[i] == "374707108764975109" || member.Roles[i] == "485056864309084160" {
			return true
		}
	}
	return false
}

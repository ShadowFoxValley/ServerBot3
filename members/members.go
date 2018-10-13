package members

import (
	"ServerBot3/pointSystem"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func MemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Println("New user: ", m.User.Username, m.Nick)

	_, err := pointSystem.Database.Exec("INSERT INTO users (username, discord_id, starpoint) VALUES(?, ?, ?)", m.User.Username, m.User.ID, 0)
	if err != nil {
		log.Println(err)
	}
}

func UpdateUsers(s *discordgo.Session, guildId string) {
	var userList, error = s.GuildMembers(guildId, "", 1000)

	if error != nil {
		log.Println(error)
	}

	for i := range userList {
		_, err := pointSystem.Database.Exec("INSERT INTO users (username, discord_id, starpoint) VALUES(?, ?, ?)", userList[i].User.Username, userList[i].User.ID, 0)
		if err != nil {
			log.Println(err)
		}
	}
}

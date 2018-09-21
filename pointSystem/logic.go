package pointSystem

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func StartPointSystemPreparing(){
	var databaseTmp, err = sql.Open("mysql", "discord:truePass@tcp(localhost:3306)/discord")
	if err != nil {
		log.Print(err.Error())
	} else {
		database = databaseTmp
	}
}

func GetUserPoints(userId string) (int, bool){
	points, err := database.Query("SELECT starpoint FROM users WHERE discord_id=?", userId)
	if err != nil {
		log.Println(err.Error())
		return 0, false
	}
	var userPoints int
	points.Next()
	points.Scan(&userPoints)

	return userPoints, true
}

func ReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}

	if r.Emoji.Name == "⭐" {
		var targetMessage, errorGetMessage= s.ChannelMessage(r.ChannelID, r.MessageID)
		if errorGetMessage != nil {
			return
		}

		if r.UserID == targetMessage.Author.ID {
			return
		}

		database.Query("UPDATE users SET starpoint = `starpoint`+1 WHERE discord_id=?", targetMessage.Author.ID)
	}
}

func ReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove){
	if r.UserID == s.State.User.ID {
		return
	}

	if r.Emoji.Name == "⭐" {
		var targetMessage, errorGetMessage= s.ChannelMessage(r.ChannelID, r.MessageID)
		if errorGetMessage != nil {
			return
		}

		if r.UserID == targetMessage.Author.ID {
			return
		}

		database.Query("UPDATE users SET starpoint = `starpoint`-1 WHERE discord_id=?", targetMessage.Author.ID)
	}
}
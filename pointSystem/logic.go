package pointSystem

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Database *sql.DB
)

func StartPointSystemPreparing() {
	var databaseTmp, err = sql.Open("mysql", "discord:truePass@tcp(localhost:3306)/discord")
	if err != nil {
		log.Print(err.Error())
	} else {
		Database = databaseTmp
	}
}

func GetUserPoints(userId string) (int, bool) {
	points, err := Database.Query("SELECT starpoint FROM users WHERE discord_id=?", userId)
	if err != nil {
		log.Println(err.Error())
		return 0, false
	}
	var userPoints int
	points.Next()
	points.Scan(&userPoints)

	return userPoints, true
}

func GetStarPointTop() (*sql.Rows, bool) {
	points, err := Database.Query("SELECT username, starpoint FROM users ORDER BY starpoint DESC LIMIT 10")
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}
	return points, true
}

func ReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}
	
	if r.Emoji.ID == "🇫" || r.Emoji.ID == "♿" {
		return
	}

	var targetMessage, errorGetMessage = s.ChannelMessage(r.ChannelID, r.MessageID)
	if errorGetMessage != nil {
		return
	}

	var messageTime, _ = targetMessage.Timestamp.Parse()
	var timeDiff = time.Now().Sub(messageTime).Seconds()

	if timeDiff > 180 {
		return
	}

	if r.UserID == targetMessage.Author.ID {
		return
	}

	var _, errAddStar = Database.Exec("UPDATE users SET starpoint = `starpoint`+1 WHERE discord_id=?", targetMessage.Author.ID)

	if errAddStar != nil {
		log.Println(errAddStar)
	}

}

func ReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.UserID == s.State.User.ID {
		return
	}

	var targetMessage, errorGetMessage = s.ChannelMessage(r.ChannelID, r.MessageID)
	if errorGetMessage != nil {
		return
	}

	var messageTime, _ = targetMessage.Timestamp.Parse()
	var timeDiff = time.Now().Sub(messageTime).Seconds()

	if timeDiff > 180 {
		return
	}

	if r.UserID == targetMessage.Author.ID {
		return
	}

	var _, errRemoveStar = Database.Exec("UPDATE users SET starpoint = `starpoint`-1 WHERE discord_id=?", targetMessage.Author.ID)

	if errRemoveStar != nil {
		log.Println(errRemoveStar)
	}
}

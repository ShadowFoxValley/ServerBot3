package commands

import (
	"ServerBot3/pointSystem"
	"strconv"
)

func (data Commands) stats(){
	var points, check = pointSystem.GetUserPoints(data.message.Author.ID)

	if !check {
		data.mainSession.ChannelMessageSend(data.message.ChannelID, "Error while getting data")
		return
	}

	data.mainSession.ChannelMessageSend(data.message.ChannelID, "You have " + strconv.Itoa(points) + " stars ⭐")

}

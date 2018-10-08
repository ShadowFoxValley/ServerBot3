package commands

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"fmt"
	"strings"
)

func (data Commands) throw() {
	var target = data.message.Mentions
	if len(target) == 0 {
		return
	}

	var targetAllInfo, _ = data.mainSession.GuildMember(data.guild.ID, target[0].ID)
	var authorInfo, _ = data.mainSession.GuildMember(data.guild.ID, data.message.Author.ID)

	var targetNick, authorNick = targetAllInfo.Nick, authorInfo.Nick

	if targetNick == "" {
		targetNick = target[0].Username
	}
	if authorNick == "" {
		authorNick = data.message.Author.Username
	}

	var allEmoji = data.guild.Emojis
	var staticEmoji []*discordgo.Emoji
	for emoji := range allEmoji {
		if allEmoji[emoji].Animated == false {
			staticEmoji = append(staticEmoji, allEmoji[emoji])
		}
	}

	var targetEmoji = staticEmoji[rand.Intn(len(staticEmoji))]
	var emojiString = fmt.Sprintf("<:%s:%s>", targetEmoji.Name, targetEmoji.ID)

	var messageText = fmt.Sprintf("**%s** threw %s at **%s**", authorNick, emojiString, targetNick)

	if strings.Contains(messageText, "@everyone") || strings.Contains(messageText, "@here") {
		return
	}

	data.mainSession.ChannelMessageSend(data.channelId, messageText)
}
package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

const commandStart = "sudo"

var (
	trackReactions = make(map[string]*discordgo.Message)
)

type Commands struct {
	mainSession *discordgo.Session
	message *discordgo.MessageCreate
	channelId string
	guild *discordgo.Guild
}

func (data Commands) help(){
	const help = "```markdown\n" +
		"1. Commands\n" +
		"[roll](roll %long long int% %long long int%)\n" +
		"[throw](throw {mention})\n" +
		"[spank](spank {mention})\n" +
		"[roles](roles get/remove {rolename})\n" +
		"[respect](respect anything)\n" +
		"[wheelchair](wheelchair %mention%)\n" +
		"< Everything inside %% is optional >\n\n" +
		"[poll](poll\n" +
		"option 1\n" +
		"option 10\n" +
		")\n" +
		"**Стражи:**\n" +
		"[root](root mute/unmute {mention})" +
		"```"

	data.mainSession.ChannelMessageSend(data.channelId, help)
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	var channel, _ = s.State.Channel(m.ChannelID)
	var guild, _ = s.State.Guild(channel.GuildID)

	var commandWorker = Commands{
		s,
		m,
		m.ChannelID,
		guild,
	}

	var contentElements = strings.Split(strings.Replace(m.Content, "\n", " ", -1), " ")

	if len(contentElements) <= 1 || contentElements[0] != commandStart {
		return
	}

	switch contentElements[1] {
	case "roll":
		commandWorker.roll()
		break

	case "throw":
		commandWorker.throw()
		break

	case "poll":
		commandWorker.poll()
		break

	case "spank":
		commandWorker.spank()
		break
	case "respect":
		commandWorker.respect()
		break

	case "roles":
		commandWorker.roles()
		break

	case "wheelchair":
		commandWorker.wheelchair()
		break

	case "root":
		commandWorker.root()
		break

	case "help":
		commandWorker.help()
		break
	}
}

func ReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}


	if elem, ok := trackReactions[r.MessageID]; ok && (r.Emoji.Name == "🇫" || r.Emoji.Name == "♿")  {
		user, err := s.User(r.UserID)
		if err != nil {
			log.Println(err)
			return
		}

		if strings.Contains(elem.Embeds[0].Fields[0].Value, user.Username) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				err := s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
				if err != nil {
					log.Println(err, r.MessageID, elem.ID)
				}
			}()
			return
		}

		var fieldInfo = "Payed respect"

		if r.Emoji.Name == "♿" {
			fieldInfo = "Скинулись на коляску"
		}


		embed := &discordgo.MessageEmbed{
			Title:  elem.Embeds[0].Title,
			Author: elem.Embeds[0].Author,
			Color:  elem.Embeds[0].Color,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   fieldInfo,
					Value:  elem.Embeds[0].Fields[0].Value + "\n" + user.Username,
					Inline: false,
				},
			},
		}
		message, err := s.ChannelMessageEditEmbed(r.ChannelID, r.MessageID, embed)
		if err != nil {
			log.Println(err)
			return
		}
		trackReactions[r.MessageID] = message
		go func() {
			time.Sleep(100 * time.Millisecond)
			err := s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
			if err != nil {
				log.Println(err, r.MessageID, elem.ID)
			}
		}()
	}


}

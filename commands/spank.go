package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/thanhpk/randstr"
	"image"
	"image/color"
	"log"
	"os"
)

var (
	srcImage, _ = gg.LoadImage("spank.jpg")
)

func (data Commands) spank() {
	var realTarget *discordgo.User
	if len(data.message.Mentions) == 0 {
		data.mainSession.ChannelMessageSend(data.channelId, "Тебе нужно выбрать кого-нибудь")
		return
	} else {
		realTarget = data.message.Mentions[0]
	}

	var tmpName = randstr.RandomString(7) + ".jpg"

	const imageSize int = 400

	var spanked, _ = data.mainSession.UserAvatarDecode(realTarget)
	var spanker, _ = data.mainSession.UserAvatarDecode(data.message.Author)
	spanker = imaging.Resize(spanker, imageSize, imageSize, imaging.Lanczos)
	spanked = imaging.Resize(spanked, imageSize, imageSize, imaging.Lanczos)

	dst := imaging.New(2000, 1333, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, srcImage, image.Pt(0, 0))
	dst = imaging.Paste(dst, spanked, image.Pt(1470, 660))
	dst = imaging.Paste(dst, spanker, image.Pt(970, 0))
	err := imaging.Save(dst, tmpName)
	if err != nil {
		log.Println(err.Error())
	}
	file, _ := os.Open(tmpName)

	data.mainSession.ChannelFileSend(data.channelId, "spank.jpg", file)
	os.Remove(tmpName)
}

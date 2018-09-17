package main

import (
	"ServerBot3/commands"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"

	"syscall"
)

var (

	dgv *discordgo.VoiceConnection
	mainSessionEndpoint *discordgo.Session
)


func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot "+token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	mainSessionEndpoint = dg // Copy variable for voice modules and news module

	// Команды
	dg.AddHandler(commands.MessageCreate)

	//Отлов реакций для команд wheelchair и respect
	dg.AddHandler(commands.ReactionAdd)


	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}


	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

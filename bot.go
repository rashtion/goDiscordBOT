package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	AvatarFile string
	AvatarURL string
)

const prefix = "!"

func init() {

	flag.StringVar(&Token, "t", "", "bot token")
	
	
	
	flag.Parse()
}

func main() {


	dg, err := discordgo.New("Bot " + "bot token")
	if err != nil {
		fmt.Println("Discord'da oturum açılamadı!,", err)
		return
	}


	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		fmt.Println("bağlantı hatası,", err)
		return
	}

	fmt.Println("Bot çalışıyor!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {


	if m.Author.ID == s.State.User.ID {
		return
	}


	if m.Content == prefix + "info" {
		s.ChannelMessageSend(m.ChannelID, "Merhaba, bu bot GO dili ile yazılmış bir demodur.")
	}


	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "selam" {
		s.ChannelMessageSend(m.ChannelID, "selam")
	}


	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

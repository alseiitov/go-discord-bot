package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const token string = "YOUR_DISCORD_BOT_TOKEN"
const auth = "AUTH_KEY_ON_DASHBOARD.ALEM.SCHOOL"
const URL = "https://api.alem.school/leaderboard"
const infoURL = "https://api.alem.school/user/"

var botID string

func main() {
	db, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := db.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	botID = u.ID

	db.AddHandler(messageHandler)

	err = db.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running...")

	<-make(chan struct{})

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botID {
		return
	}
	info := makeRequest(strings.ToLower(m.Content))
	var inSchool bool
	switch info.InSchoolStatus {
	case 1:
		inSchool = true
	case 2:
		inSchool = false
	}
	msg := fmt.Sprintf("First name: %v\nLast name: %v\nIn school: %v\nLast activity: %v", info.FirstName, info.LastName, inSchool, info.Last)
	imgBytes, _ := base64.StdEncoding.DecodeString(info.Avatar)
	avatar := bytes.NewReader(imgBytes)
	if info.FirstName != "" {
		_, _ = s.ChannelFileSendWithMessage(m.ChannelID, msg, "info.png", avatar)
	}
}

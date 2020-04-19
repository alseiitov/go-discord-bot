package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	botToken   string = "DISCORD_BOT_TOKEN"
	adminID    string = "ADMINS_DISCORD_ID"
	infoURL    string = "https://api.alem.school/user/"
	refreshURL string = "https://email-trigger.alem.school/auth/refresh"
)

var (
	jwtToken string
	botID    string
	users    = map[string]string{}
)

func init() {
	readUsers()
	readJWT()
}

func main() {
	db, err := discordgo.New("Bot " + botToken)
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

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)

	go autoUpdateToken()

	<-make(chan struct{})
}

func autoUpdateToken() {
	for {
		check := parseInfo("4050")
		time.Sleep(10 * time.Second)
		if check.FirstName != "Dossan" {
			refreshToken()
		}
		check.FirstName = ""
		time.Sleep(5 * time.Minute)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botID {
		return
	}

	if strings.HasPrefix(m.Content, "/token ") && m.Author.ID == adminID {
		updateJWT(m.Content[7:])
		log.Println("Token updated by admin: ", jwtToken)
		return
	}

	if strings.ToLower(m.Content) == "help" {
		msg := fmt.Sprintf("```user username\nget info abount user\n```")
		_, _ = s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	sendUserInfo(s, m)
	return
}

func sendUserInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	id, ok := users[strings.ToLower(m.Content)]
	if !ok {
		return
	}

	info := parseInfo(id)
	inSchool := ":red_circle: Not in school "
	if info.InSchoolStatus == 1 {
		inSchool = ":green_circle: In school"
	}

	msg := fmt.Sprintf(":bookmark_tabs: %v %v\n%v\n:alarm_clock: %v\n", info.FirstName, info.LastName, inSchool, info.Last)

	var avatar io.Reader

	imgBytes, err := base64.StdEncoding.DecodeString(info.Avatar)
	if err == nil {
		avatar = bytes.NewReader(imgBytes)
	}

	if info.FirstName != "" {
		_, _ = s.ChannelFileSendWithMessage(m.ChannelID, msg, "info.png", avatar)
	}
}

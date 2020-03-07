package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func makeRequest(username string) user {
	var tempUser user
	var allUsers Users
	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Add("Authorization", auth)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &allUsers)
	for _, user := range allUsers.Data.EventUser {
		if strings.ToLower(user.User.GithubLogin) == username {
			tempUser = parseInfo(user.User.ID)
		}
	}
	return tempUser
}

func parseInfo(ID int) user {
	var tempUser user
	var info UserInfo
	req, err := http.NewRequest("GET", infoURL+strconv.Itoa(ID), nil)
	req.Header.Add("Authorization", auth)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &info)
	tempUser.FirstName = info.User.Data.User[0].FirstName
	tempUser.LastName = info.User.Data.User[0].LastName
	tempUser.InSchoolStatus = info.Attendance.Data[len(info.Attendance.Data)-1].Status
	tempUser.Last = info.Attendance.Data[len(info.Attendance.Data)-1].Date
	tempUser.Avatar = info.Image.Data[0].Face
	return tempUser
}

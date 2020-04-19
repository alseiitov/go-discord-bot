package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func refreshToken() {
	data := fmt.Sprintf("{\"jwt_token\":\"%v\"}", jwtToken)
	req, _ := http.NewRequest("POST", refreshURL, strings.NewReader(data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var newToken jwt
	json.Unmarshal(body, &newToken)
	if newToken.JwtToken != "" {
		jwtToken = newToken.JwtToken
		updateJWT(jwtToken)
		log.Println("Token refreshed: ", jwtToken)
	} else {
		log.Println(string(body))
	}
}

func parseInfo(ID string) user {
	var tempUser user
	var info UserInfo
	req, err := http.NewRequest("GET", infoURL+ID, nil)
	if err != nil {
		log.Println(err.Error())
		return tempUser
	}
	req.Header.Add("Authorization", jwtToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return tempUser
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return tempUser
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Println(err.Error())
		return tempUser
	}

	if len(info.User.Data.User) != 0 {
		tempUser.Login = info.User.Data.User[0].GithubLogin
		tempUser.FirstName = info.User.Data.User[0].FirstName
		tempUser.LastName = info.User.Data.User[0].LastName
	}

	if len(info.Attendance.Data) != 0 {
		tempUser.InSchoolStatus = info.Attendance.Data[len(info.Attendance.Data)-1].Status
		tempUser.Last = info.Attendance.Data[len(info.Attendance.Data)-1].Date

	}

	if len(info.Image.Data) != 0 {
		tempUser.Avatar = info.Image.Data[0].Face
	}

	if info.Progress.Data.User != nil {
		for _, subject := range info.Progress.Data.User[0].Progresses {
			tempUser.DoneProjects = append(tempUser.DoneProjects, subject.Object.Name)
		}
	}

	return tempUser
}

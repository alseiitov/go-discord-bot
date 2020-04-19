package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func readUsers() {
	file, _ := os.Open("./users.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "|")
		id := arr[0]
		name := strings.ToLower(arr[1])
		users[name] = id
	}
}

func readJWT() {
	bytes, _ := ioutil.ReadFile("./jwt.txt")
	jwtToken = string(bytes)
}

func updateJWT(jwt string) {
	jwtToken = jwt
	ioutil.WriteFile("./jwt.txt", []byte(jwt), 0644)
}

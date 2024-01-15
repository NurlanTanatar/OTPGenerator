package main

import (
	"OTPGenerator/tools"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	email := flag.String("email", "diable102@gmail.com", "provide email")
	username := flag.String("username", "jastime", "provide name of your company")
	password := flag.String("password", "1Abcdefg.", "provide password")
	flag.Parse()
	fmt.Println(password)
	user := &tools.BasicInfo{Email: *email, Name: *username}
	user.GenOTPCLI()

	fmt.Println(tools.GenerateTokenCLI(*username, *password))
	passwordB64 := base64.StdEncoding.EncodeToString([]byte(*username + ":" + *password))
	fmt.Println(passwordB64)
}

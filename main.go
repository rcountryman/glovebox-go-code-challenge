package main

import (
	"fmt"
	"os"
)

func main() {

	username := os.Getenv("GLOVEBOX_USER")
	password := os.Getenv("GLOVEBOX_PASS")

	yproxy := yproxy{}
	err := yproxy.Login(username, password)

	if err != nil {
		fmt.Println("No auth fail.")
	}

	page := yproxy.browser.MustPage("https://news.ycombinator.com/newest")
	el := page.MustElement("#me")

	if el.MustText() != "rcountry21" {
		fmt.Println("Cannot find correct user after login.")
	}

	fmt.Println("Success!!")
}
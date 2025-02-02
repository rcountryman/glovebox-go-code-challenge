# carrierproxy-poc

This repository is used for creating a proof of concept around the `PolicyProvider` interface.

## Task

- [ ] Fork this repository and implement the `PolicyProvider`'s `Login` method against a web property of your choice using one of the scraping libraries listed below. 
- [ ] Implement tests that accept environment variables for the credentials (login & password). 
- [ ] Document your code and usage instructions.
- [ ] Send in a pull request for code review.

## Scraping Libraries

* [go-rod](https://github.com/go-rod/rod)
* [chromedp](https://github.com/chromedp/chromedp)

## Usage

The `yproxy` struct is used for logging in and scraping data from ycombinator (web property of choice)
`yproxy` object must be created, then the `Login` called with username and password passed to it

Example in main.go:

	username := os.Getenv("GLOVEBOX_USER")
	password := os.Getenv("GLOVEBOX_PASS")

	yproxy := yproxy{}
	err := yproxy.Login(username, password)




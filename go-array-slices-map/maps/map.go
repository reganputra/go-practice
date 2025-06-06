package main

import "fmt"

func main() {

	webUrl := map[string]string{
		"google":   "https://www.google.com",
		"facebook": "https://www.facebook.com",
	}

	webUrl["twitter"] = "https://www.twitter.com"
	delete(webUrl, "facebook")
	fmt.Println(webUrl)
}

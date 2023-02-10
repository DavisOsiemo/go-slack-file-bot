package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-424914707189-4807004404192-Rd0fW68leFPOZKKfebjvqAbd")
	os.Setenv("CHANNEL_ID", "CCGFWP928")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"BRITAM.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}

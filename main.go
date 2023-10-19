package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main(){
	//add slack token
	os.Setenv("SLACK_BOT_TOKEN", "")
	//add slack channel_id
	os.Setenv("CHANNEL_ID", "")
	api:= slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr:= []string{"non-OECD-countries copy.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params:= slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("Name: %v, Url:  %v", file.Name, file.URL)
		
	}
}
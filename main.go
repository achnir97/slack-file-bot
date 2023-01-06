package main 
import (
    "fmt"
    "os"
    "github.com/slack-go/slack"
    )

func main () {
    os.Setenv("SLACK_BOT_TOKEN", " ")
    os.Setenv("CHANNEL_ID","")
    api:= slack.New(os.Getenv("SLACK_BOT_TOKEN"))
    channelArr:=[]string{
        os.Getenv("CHANNEL_ID")
    }
    fileArr:=[]string{""}
}
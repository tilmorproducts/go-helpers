package gohelpers

import goteamsnotify "github.com/atc0005/go-teams-notify/v2"
import messageCard "github.com/atc0005/go-teams-notify/v2/messagecard"

func TeamsSuccess(message string, title string, channel string) {
	mstClient := goteamsnotify.NewTeamsClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = message
	msgCard.ThemeColor = "#00FF04"
	mstClient.Send(channel, msgCard)
}

func TeamsError(err error, title string, channel string) {
	mstClient := goteamsnotify.NewTeamsClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = "Error message: " + err.Error()
	msgCard.ThemeColor = "#FF0000"
	mstClient.Send(channel, msgCard)
}

func TeamsInfo(message string, title string, channel string) {
	mstClient := goteamsnotify.NewTeamsClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = message
	msgCard.ThemeColor = "#ffe600"
	mstClient.Send(channel, msgCard)
}

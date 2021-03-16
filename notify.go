package gohelpers

import goteamsnotify "github.com/atc0005/go-teams-notify"

func teamsSuccess(message string, title string, errorChannel string) {
	mstClient := goteamsnotify.NewClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = message
	msgCard.ThemeColor = "#00FF04"
	mstClient.Send(errorChannel, msgCard)
}

func teamsError(err error, title string, errorChannel string) {
	mstClient := goteamsnotify.NewClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = "Error message: " + err.Error()
	msgCard.ThemeColor = "#FF0000"
	mstClient.Send(errorChannel, msgCard)
}

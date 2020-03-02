package internal

import (
	"log"

	gosxnotifier "github.com/kevindoveton/gosx-notifier"
)

const title = "Slack"
const group = "com.kdoveton.slack-notifications"
const standardIcon = "assets/icon.png" // this doesn't work when bundled

/*
NotifyWithDeepLink will send a message to the notification tray
and allow the user to specify a link
*/
func NotifyWithDeepLink(message string, link string) {
	note := gosxnotifier.NewNotification(message)
	//Optionally, set a title
	note.Title = title
	// Optionally, set a subtitle
	// note.Subtitle = "My subtitle"
	note.Group = group
	// not sure why this isn't working..
	note.AppIcon = standardIcon
	note.Link = link

	err := note.Push()
	if err != nil {
		log.Println("Uh oh!")
	}
}

/*
Notify will send a message to the notification tray
*/
func Notify(message string) {
	note := gosxnotifier.NewNotification(message)
	//Optionally, set a title
	note.Title = title
	// Optionally, set a subtitle
	// note.Subtitle = "My subtitle"
	note.Group = group
	// not sure why this isn't working..
	note.AppIcon = standardIcon

	err := note.Push()
	if err != nil {
		log.Println("Uh oh!")
	}
}

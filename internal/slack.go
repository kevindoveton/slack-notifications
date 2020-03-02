package internal

import (
	"strings"

	"github.com/slack-go/slack"
)

type user struct {
	userID   string
	username string
	email    string
}

// GetUsername will convert a userid to username
func GetUsername(api *slack.Client, userID string) string {
	data, _ := api.GetUserInfo(userID)
	return data.Name
}

// GetGroupName converts a group id to name
func GetGroupName(api *slack.Client, groupID string) string {
	data, _ := api.GetGroupInfo(groupID)
	return data.Name
}

// GetChannelName converts a group id to name
func GetChannelName(api *slack.Client, channelID string) string {
	data, _ := api.GetChannelInfo(channelID)
	return data.Name
}

// GetByID get something by id
func GetByID(api *slack.Client, channelID string, userID string) string {
	if strings.HasPrefix(channelID, "G") {
		// group
		return GetGroupName(api, channelID)
	} else if strings.HasPrefix(channelID, "C") {
		//channel
		return GetChannelName(api, channelID)
	} else if strings.HasPrefix(channelID, "U") {
		return GetUsername(api, userID)
		// user
	} else if strings.HasPrefix(channelID, "D") {
		// direct message
		return GetUsername(api, userID)
	}

	return ""
}

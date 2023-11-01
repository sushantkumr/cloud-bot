package resources

import (
	"cloud-bot/constants"
	"cloud-bot/cloud"
)

var (
	apiKey      string = ""
	accessToken string = ""
)

func getCredentials() {
	apiKey = cloud.getParamtersFromSSM(constants.API_KEY)
	accessToken = cloud.getParamtersFromSSM(constants.ACCESS_TOKEN)
}

func login() {
	getCredentials()
}

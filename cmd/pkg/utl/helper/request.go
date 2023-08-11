package helper

import (
	"fmt"

	requestHandler "github.com/onskycloud/request-handler"
)

func GetHealth() (string, error) {
	var result string
	requestModel := requestHandler.RequestModel{
		URL:       fmt.Sprintf("http://115.79.27.36:8083/medic/"),
		TokenType: requestHandler.Bearer,
		Token:     "",
		Body:      "",
	}
	if err := requestHandler.GetV2(requestModel, result); err != nil {
		return "", err
	}
	return result, nil
}
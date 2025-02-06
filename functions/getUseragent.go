package functions

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)


func GetUserAgents() ([]string, error) {
	file, err := os.ReadFile(assetsDir + "useragents.json")
	if err != nil {
		return nil, err
	}

	var userAgents []string
	err = json.Unmarshal(file, &userAgents)
	if err != nil {
		return nil, err
	}
	return userAgents, nil
}

func GetRandomUserAgent() (string, error) {
	userAgents, err := GetUserAgents()
	if err != nil {
		return "", err
	}

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomUserAgent := userAgents[randGen.Intn(len(userAgents))]
	return randomUserAgent, nil
}

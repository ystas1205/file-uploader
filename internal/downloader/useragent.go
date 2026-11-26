package downloader

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type UserAgents struct {
	User_agents []string
}

// ReadUserAgent читает JSON-файл user_agents.json и возвращает
// случайный User-Agent и ошибку.
func ReadUserAgent() (string, error) {

	data, err := os.ReadFile("internal/downloader/user_agents.json")
	if err != nil {
		return "", fmt.Errorf("прочитать файл: %w", err)
	}
	var ua UserAgents

	err = json.Unmarshal(data, &ua)
	if err != nil {
		return "", fmt.Errorf("разобрать JSON: %w", err)
	}
	return ua.User_agents[rand.Intn(len(ua.User_agents))], nil

}

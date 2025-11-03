package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type ConfigButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Config struct {
	ClientID      string         `json:"clientId"`
	State         string         `json:"state"`
	Details       string         `json:"details"`
	LargeImage    string         `json:"largeImage"`
	LargeText     string         `json:"largeText"`
	SmallImage    string         `json:"smallImage"`
	SmallText     string         `json:"smallText"`
	ShowTimestamp bool           `json:"showTimestamp"`
	Buttons       []ConfigButton `json:"buttons"`
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config.json: %v", err)
	}

	if config.ClientID == "YOUR_DISCORD_CLIENT_ID" {
		log.Fatalln("Error: Please update config.json with your actual Discord Client ID.")
	}

	err = client.Login(config.ClientID)
	if err != nil {
		log.Fatalf("Failed to log in to Discord. Is the desktop client running?\nError: %v", err)
	}
	log.Println("Successfully connected to Discord RPC.")

	activity := client.Activity{
		State:      config.State,
		Details:    config.Details,
		LargeImage: config.LargeImage,
		LargeText:  config.LargeText,
		SmallImage: config.SmallImage,
		SmallText:  config.SmallText,
	}

	if config.ShowTimestamp {
		now := time.Now()
		activity.Timestamps = &client.Timestamps{
			Start: &now,
		}
	}

	if len(config.Buttons) > 0 {
		var rpcButtons []*client.Button
		for i, b := range config.Buttons {
			if i >= 2 {
				log.Println("Warning: Discord only supports 2 buttons. Truncating extra buttons.")
				break
			}
			rpcButtons = append(rpcButtons, &client.Button{
				Label: b.Label,
				Url:   b.URL,
			})
		}
		activity.Buttons = rpcButtons
	}

	err = client.SetActivity(activity)
	if err != nil {
		log.Fatalf("Failed to set activity: %v", err)
	}
	log.Println("Rich Presence activity successfully set.")

	log.Println("Application is running. Press Ctrl+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Shutting down...")
	client.Logout()
}

func loadConfig(filepath string) (*Config, error) {
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

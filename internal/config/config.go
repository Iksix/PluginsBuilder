package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	PluginName string `json:"plugin_name"`
	HasApi     bool   `json:"has_api"`
	ApiName    string `json:"api_name"`
}

var Instance Settings = Settings{}

func Read() {
	bytes, err := os.ReadFile("./settings.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(bytes, &Instance)
	fmt.Println("Settings readed...")
}

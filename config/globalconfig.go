package config

import (
	"encoding/json"
	"io/ioutil"
)

type GlobalConfig struct {
	/* Server configuration */
	ServerName string
	ServerHost string
	ServerPort int
	MaxPackageSize int

	/* Log configuration */
	EnableLog bool
	LogToConsole bool
	LogPath string

	/* Database configuration */
	DBType string
	DBHost string
	DBPort int
	DBUser string
	DBPassword string
	DBSchemaName string
}

// global configuration
var GlobalConfiguration *GlobalConfig

// init() returns a global object
func init() {
	GlobalConfiguration = &GlobalConfig{
		ServerName: "Mint Server",
			ServerHost: "127.0.0.1",
			ServerPort: 30000,
			MaxPackageSize: 2048,
			EnableLog: true,
			LogToConsole: true,
			LogPath: "",
			DBType: "MySQL",
			DBHost: "127.0.0.1",
			DBPort: 20000,
			DBUser: "u1",
			DBPassword: "123456",
			DBSchemaName: "minigame"}
}

// Load the configuration file into server
func (gc *GlobalConfig) Load() {
	data, err := ioutil.ReadFile("./config-example.json")
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &GlobalConfiguration)
	if err != nil {
		panic(err.Error())
	}
}
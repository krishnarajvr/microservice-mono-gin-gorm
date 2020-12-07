/*
 * @File: common.common.go
 * @Description: Defines common information of the service
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads configuration from the config file
func LoadEnv() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("Not able to get current working director")
	}
	// loads values from .env into the system
	fmt.Println("Current directory : " + dir + "/.env")
	if err := godotenv.Load(dir + "/.env"); err != nil {
		fmt.Print("No .env file found")
	}
}

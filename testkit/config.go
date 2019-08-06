package testkit

import (
	"fmt"
	"os"
	"taskulu/internal"
)

func InitTestConfig(relativePath string) *internal.Config {
	configPath := fmt.Sprintf("%s/%s", os.Getenv("PWD"), relativePath)
	fmt.Println("Config path: ", configPath)
	return internal.NewConfig(configPath)
}

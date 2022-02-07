package helper

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// GetEnv finds an env variable or the given fallback.
func GetEnv(key, fallback string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	fmt.Println(value)
	return value
}

func GetStatus(status uint) *bool {
	newStatus := false
	if status == 1 {
		newStatus = true
		return &newStatus
	}
	return &newStatus
}

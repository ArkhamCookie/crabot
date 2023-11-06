package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(target string, file string) (string, error) {
	if file == "" {
		file = ".env"
	}

	err := godotenv.Load(file)
	if err != nil {
		log.Fatalln("Some error occured. Error:", err)
	}

	value := os.Getenv(target)
	if value == "" {
		log.Printf("%s value not found\n", target)
		return "", errors.New("value not found")
	}
	return value, nil
}

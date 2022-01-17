package gakujo

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}

func TestScrape(t *testing.T) {
	username := os.Getenv("GAKUJOUSERNAME")
	password := os.Getenv("GAKUJOUPASS")

	client := NewClient()

	err := client.Login(username, password)

	if err != nil {
		log.Fatal(err)
	}

	result := client.getTask()

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

}

package helpers

import (
	"log"
	"os"
)

func Log(text string) {
	file, err := os.OpenFile("service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(text)
}

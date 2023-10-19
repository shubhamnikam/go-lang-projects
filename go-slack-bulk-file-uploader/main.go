package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main(){
	fmt.Println("Welcome to slack-bulk-file-uploader :)")

	godotenv.Load(".env")
	
	filePaths :=  GetFilesToUpload()
	SetupFileUploader(filePaths)
	
	fmt.Println("Exiting :)")
}

func GetFilesToUpload() ([]string) {
	fmt.Println("Enter full file path to upload (Press 'enter' to add new entry & '/stop' to end)")
	scanner := bufio.NewScanner(os.Stdin)

	var filePaths []string
	for {
		scanner.Scan()
		filePath := scanner.Text()
		if len(filePath) == 0 || strings.TrimSpace(filePath) == "/stop" {
			break
		} 

		filePaths = append(filePaths, filePath)
	}

	fmt.Printf("Total files to upload: %v, All filepaths: %v\n", len(filePaths), filePaths)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return filePaths
}

func SetupFileUploader(filePaths []string) {
	if len(filePaths) == 0 {
		log.Fatalln("No file path found")		
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))

	// upload logic
	for i:=0; i<len(filePaths); i++ {
		params := slack.FileUploadParameters{
			Channels: strings.Split(os.Getenv("SLACK_CHANNELS_ID"), ","),
			File: filePaths[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Uploaded File No: %v, Name: %s\n", i+1, file.Name)
	}
}

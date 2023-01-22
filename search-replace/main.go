package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*NOTE => 
# run go:
go run main.go data.txt is AAAA

# explain:
fileName = data.txt
searchText = is
replaceText = AAAA

*/

func main() {
	//check arg validity
	if (len(os.Args[1:]) > 3) {
		log.Fatal("Enter data in following order: file-name search-text replace-text")
	}

	fileName  := os.Args[1]
	searchText  := os.Args[2]
	replaceText  := os.Args[3]

	SearchAndReplace(fileName, searchText, replaceText)
}

func SearchAndReplace(fileName, searchText, replaceText string){
	fmt.Printf("Your Input:\nfileName: %v\nsearchText: %v\nreplaceText: %v\n", fileName, searchText, replaceText)
	
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	
	//search
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), searchText, replaceText, -1)
		lines = append(lines, line)
	}
	
	err = ReplaceText(fileName, lines)
	checkError(err)
	
}

func ReplaceText(fileName string, lines []string) (err error){
	
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY, 0644)
	checkError(err)
	defer file.Close()

	//replace
	for _, line := range lines {		
		_, writeErr := file.WriteString(line+"\n")
		if writeErr != nil{
			return err
		}
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
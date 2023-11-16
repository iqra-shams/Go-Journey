package utils

import (
	"io/ioutil"
	"log"
)


func ReadFile(filepath string)string{
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	FileData := string(content)
	return FileData
}
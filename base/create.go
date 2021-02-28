package base

import (
	"fmt"
	"os"
	//"strings"
)

func Create(name string) {
	if _, err := os.Stat("database"); os.IsNotExist(err) {
		fmt.Printf("Database Folder doesn't exist.\n")
		os.MkdirAll("database", 0777)
		fmt.Printf("Created Database Folder.\n")
	}
	os.MkdirAll(name,0777)
}
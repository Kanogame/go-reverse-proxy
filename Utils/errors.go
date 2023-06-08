package utils

import (
	"fmt"
	"os"
)

func HandleUserError(err error) bool {
	return err == nil
}

func HandleAppError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintError(errorMessage string) {
	fmt.Println("Error: " + errorMessage)
	os.Exit(0)
}

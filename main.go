package main

import (
	"fmt"
	"os"

	"github.com/arfeloreed/password-gen/internal"
)

func main() {
	fmt.Print("Welcome to the password generator!\n")
	mode, email, passLen := internal.Mode()

	if mode == 1 {
		// search the json file with provided email
		data, err := internal.SearchJson(email)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if data.Email == "" {
			fmt.Print("No record found\n")
		} else {
			fmt.Printf("Password for %s is %v\n", email, data.Password)
		}
	} else if mode == 2 {
		var password string = internal.GeneratePassword(passLen)
		// creating a json file
		message, err := internal.CreateJson(email, password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else if message == "success" {
			internal.Screen()
			fmt.Print("Your password is successfully generated\n")
		}
	} else if mode == 3 {
		var password string = internal.GeneratePassword(passLen)
		// search the json file with provided email and update the password
		message, err := internal.EditJson(email, password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if message == "success" {
			internal.Screen()
			fmt.Print("Your password is successfully updated\n")
		} else if message == "" {
			internal.Screen()
			fmt.Print("No record found\n")
		}
	} else if mode == 4 {
		// search the json file with provided email and delete the record
		message, err := internal.DeleteEmail(email)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if message == "success" {
			internal.Screen()
			fmt.Print("Your record was successfully deleted\n")
		} else if message == "" {
			internal.Screen()
			fmt.Print("No record found\n")
		}
	} else {
		internal.Screen()
		fmt.Print("Invalid mode\n")
		os.Exit(1)
	}
}

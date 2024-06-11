package internal

import (
	"fmt"
	"os"
)

func Mode() (int, string, int) {
	var mode int
	var email string
	var passLen int

	// ask for mode
	fmt.Print(`Choose mode:
	1. Search for password for an existing account.
	2. Generate new password for a new account.
	3. Edit a password for an existing account.
	4. Delete account.
	5. Exit
	Enter mode:
	`)
	fmt.Scanln(&mode)

	if mode > 5 {
		Screen()
		fmt.Print("Invalid mode\n")
		os.Exit(1)
	}

	if mode == 1 {
		Screen()
		// ask for user email or username
		fmt.Print("Enter your email or username: ")
		fmt.Scanln(&email)

		return mode, email, 0
	} else if mode == 2 {
		Screen()
		// ask for user email or username
		fmt.Print("Enter your email or username: ")
		fmt.Scanln(&email)
		// ask for a user input of how many char is the password
		fmt.Print("Password should be at least 15 characters long\n")
		fmt.Print("Enter how long you want your password: ")
		fmt.Scanln(&passLen)
		if passLen < 15 {
			Screen()
			fmt.Print("Invalid password length\n")
			os.Exit(1)
		}

		return mode, email, passLen
	} else if mode == 3 {
		Screen()
		// ask for user email or username
		fmt.Print("Enter your email or username: ")
		fmt.Scanln(&email)
		// ask for a user input of how many char is the password
		fmt.Print("Password should be at least 15 characters long\n")
		fmt.Print("Enter length of new password: ")
		fmt.Scanln(&passLen)
		if passLen < 15 {
			Screen()
			fmt.Print("Invalid password length\n")
			os.Exit(1)
		}

		return mode, email, passLen
	} else if mode == 4 {
		Screen()
		fmt.Print("Enter email or username of account to delete: ")
		fmt.Scanln(&email)
		return mode, email, 0
	} else if mode == 5 {
		Screen()
		fmt.Print("Exiting...\n")
		fmt.Print("See ya later.\n")
		os.Exit(0)
	}

	return mode, "", 0
}

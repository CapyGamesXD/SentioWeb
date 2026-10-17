package main

import (
	"fmt"
	"os"
)

func enlistEmail(email string) {
	//Send email

}

func main() {

	//Print my awesome text thing!
	fmt.Println(` ____             _   _       
/ ___|  ___ _ __ | |_(_) ___  
\___ \ / _ \ '_ \| __| |/ _ \ 
 ___) |  __/ | | | |_| | (_) |
|____/ \___|_| |_|\__|_|\___/ ` + "\n")

	emails, err := os.ReadFile("sentioUsers.json")

	if os.IsNotExist(err) {
		fmt.Println("\033[33m" + "Warning: the users file does not exist. A new file will be made upon recieving a request." + "\033[0m")
	} else if err != os.ErrNotExist && err != nil {
		fmt.Println("\033[33m"+"Something went wrong when reading the users list:", err, "\033[0m")
	}

	if len(emails) == 0 && !os.IsNotExist(err) {
		fmt.Println("\033[33m" + "Warning: no emails in record" + "\033[0m")
	}

	enlistEmail("capy@capydesigns.com")
}

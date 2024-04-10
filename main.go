package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fmt.Println("Welcome to Key-Value Store...")

	kv := NewKeyValueStore()

	// Use sync.WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for {
		options := []string{"Put", "Get", "Delete", "Exit"}
		var selectedOption string
		prompt := &survey.Select{
			Message: "Select an option",
			Options: options,
		}
		survey.AskOne(prompt, &selectedOption)

		switch selectedOption {
		case "Put":
			fmt.Println("Put servic starting")
			var key, value string
			survey.AskOne(&survey.Input{Message: "Enter key:"}, &key)
			survey.AskOne(&survey.Input{Message: "Enter value:"}, &value)

			wg.Add(1)

			go func(key, value string) {
				defer wg.Done()
				handleError(kv.Put(key, value))
				fmt.Println("Put Operation Completed.")
			}(key, value)

		case "Get":
			var key string
			survey.AskOne(&survey.Input{Message: "Enter key:"}, &key)
			wg.Add(1)

			go func(key string) {
				defer wg.Done()
				value, err := kv.Get(key)
				handleError(err)
				fmt.Println("Value:", value)
			}(key)

		case "Delete":
			var key string
			survey.AskOne(&survey.Input{Message: "Enter Key to Delete: "}, &key)

			wg.Add(1)
			go func(key string) {
				defer wg.Done()
				handleError(kv.Delete(key))
				fmt.Println("Delete operation completed.")
			}(key)

		case "Exit":
			fmt.Println("Exiting...")
			wg.Wait() //wait for all goroutines to finish before exiting
			os.Exit(0)
		}
	}
}

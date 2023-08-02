package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	createKeyName := createCommand.String("key", "", "Key name (Required)")

	copyCommand := flag.NewFlagSet("copy", flag.ExitOnError)
	copyKeyName := copyCommand.String("key", "", "Key name (Required)")
	copyOutputPath := copyCommand.String("output", "", "Output path (Required)")

	if len(os.Args) < 2 {
		fmt.Println("create or copy subcommand is required")
		os.Exit(1)
	}

	// Create KeyManager and DatabaseManager instances
	km := &KeyManager{}
	dm, err := NewDatabaseManager(km)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		os.Exit(1)
	}

	// Initialize the database
	err = dm.InitializeDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCommand.Parse(os.Args[2:])
		if *createKeyName == "" {
			createCommand.PrintDefaults()
			os.Exit(1)
		}

		// Insert a new key
		err := dm.InsertKey(*createKeyName)
		if err != nil {
			fmt.Println("Error creating key:", err)
			os.Exit(1)
		}

	case "copy":
		copyCommand.Parse(os.Args[2:])
		if *copyKeyName == "" || *copyOutputPath == "" {
			copyCommand.PrintDefaults()
			os.Exit(1)
		}

		// Fetch the key from the database
		privateKey, publicKey, err := dm.GetKey(*copyKeyName)
		if err != nil {
			fmt.Println("Error fetching key:", err)
			os.Exit(1)
		}

		if privateKey == "" || publicKey == "" {
			fmt.Println("Key does not exist")
			os.Exit(1)
		}

		// Write the keys to the output path
		err = os.WriteFile(*copyOutputPath+"/private.pem", []byte(privateKey), 0644)
		if err != nil {
			fmt.Println("Error writing private key:", err)
			os.Exit(1)
		}

		err = os.WriteFile(*copyOutputPath+"/public.pem", []byte(publicKey), 0644)
		if err != nil {
			fmt.Println("Error writing public key:", err)
			os.Exit(1)
		}

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

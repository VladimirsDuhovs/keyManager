package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/VladimirsDuhovs/keyManager/key_manager"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// Handle the creation of a new key

func createKey(dm *key_manager.DatabaseManager, keyName string) {
	fmt.Printf("Creating key '%s'...\n", keyName)
	err := dm.InsertKey(keyName)
	if err != nil {
		fmt.Println("Error creating key:", err)
		os.Exit(1)
	}

	fmt.Printf("Key '%s' successfully created\n", keyName)
}

// Handle the copying of a key
func copyKey(dm *key_manager.DatabaseManager, keyName string, outputPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	absoluteOutputPath := filepath.Join(currentDir, outputPath)
	fmt.Printf("Copying key '%s' to '%s'...\n", keyName, absoluteOutputPath)

	privateKey, publicKey, err := dm.GetKey(keyName)
	if err != nil {
		fmt.Println("Error fetching key:", err)
		os.Exit(1)
	}

	if privateKey == "" || publicKey == "" {
		fmt.Println("Key does not exist")
		os.Exit(1)
	}

	err = os.WriteFile(absoluteOutputPath+"/private.pem", []byte(privateKey), 0644)
	if err != nil {
		fmt.Println("Error writing private key:", err)
		os.Exit(1)
	}

	err = os.WriteFile(absoluteOutputPath+"/public.pem", []byte(publicKey), 0644)
	if err != nil {
		fmt.Println("Error writing public key:", err)
		os.Exit(1)
	}

	err = dm.AddCopyRecord(keyName, absoluteOutputPath)
	if err != nil {
		fmt.Println("Error recording copy operation:", err)
		os.Exit(1)
	}

	fmt.Println("Key copied successfully to", absoluteOutputPath)
}

func printCopyHistory(dm *key_manager.DatabaseManager, keyName string) {
	fmt.Printf("Fetching copy history for key '%s'...\n\n", keyName)
	copies, err := dm.GetCopyData(keyName)
	if err != nil {
		fmt.Println("Error fetching key:", err)
		os.Exit(1)
	}

	if len(copies) == 0 {
		fmt.Printf("No copy history for key '%s'\n", keyName)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Copy History for Key '" + keyName + "'")
	t.AppendHeader(table.Row{"#", "Timestamp", "Path", "Username"})

	for i, copy := range copies {
		t.AppendRow([]interface{}{i + 1, copy.Timestamp.Format(time.RFC3339), copy.Path, copy.Username})
	}
	t.AppendFooter(table.Row{"", "", "Total", len(copies)})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}

func deleteKey(dm *key_manager.DatabaseManager, keyName string) {
	fmt.Printf("Deleting key '%s'...\n", keyName)
	err := dm.DeleteKey(keyName)
	if err != nil {
		fmt.Println("Error deleting key:", err)
		os.Exit(1)
	}

	fmt.Printf("Key '%s' successfully deleted\n", keyName)
}

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			os.Exit(1)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func main() {
	km := &key_manager.KeyManager{}

	// Handle errors properly when creating DatabaseManager
	dm, err := key_manager.NewDatabaseManager(km)
	if err != nil {
		fmt.Println("Error creating DatabaseManager:", err)
		os.Exit(1)
	}

	var rootCmd = &cobra.Command{
		Use:   "keyManager",
		Short: "KeyManager is a CLI for managing keys",
		Long:  `A Fast and Flexible key manager built with love by Vladimirs Duhovs in Go.`,
	}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new key",
		Run: func(cmd *cobra.Command, args []string) {
			keyName, _ := cmd.Flags().GetString("key")
			createKey(dm, keyName)
		},
	}
	createCmd.Flags().StringP("key", "k", "", "Key name (Required)")
	createCmd.MarkFlagRequired("key")

	var copyCmd = &cobra.Command{
		Use:   "copy",
		Short: "Copy a key",
		Run: func(cmd *cobra.Command, args []string) {
			keyName, _ := cmd.Flags().GetString("key")
			outputPath, _ := cmd.Flags().GetString("output")
			copyKey(dm, keyName, outputPath)
		},
	}
	copyCmd.Flags().StringP("key", "k", "", "Key name (Required)")
	copyCmd.MarkFlagRequired("key")
	copyCmd.Flags().StringP("output", "o", "", "Output path (Required)")
	copyCmd.MarkFlagRequired("output")

	var historyCmd = &cobra.Command{
		Use:   "history",
		Short: "Print copy history for a key",
		Run: func(cmd *cobra.Command, args []string) {
			keyName, _ := cmd.Flags().GetString("key")
			printCopyHistory(dm, keyName)
		},
	}
	historyCmd.Flags().StringP("key", "k", "", "Key name (Required)")
	historyCmd.MarkFlagRequired("key")

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a key",
		Run: func(cmd *cobra.Command, args []string) {
			keyName, _ := cmd.Flags().GetString("key")
			if askForConfirmation(fmt.Sprintf("Do you really want to delete the key %s?", keyName)) {
				deleteKey(dm, keyName)
			} else {
				fmt.Println("Deletion cancelled.")
			}
		},
	}

	deleteCmd.Flags().StringP("key", "k", "", "Key name (Required)")
	deleteCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(createCmd, copyCmd, historyCmd, deleteCmd)

	// Remove the completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.Execute()
}

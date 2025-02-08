package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/olekukonko/tablewriter"
)

type KeyData struct {
	Keys []map[string]int `json:"keys"`
}

func readJSON(filename string) (KeyData, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return KeyData{}, err
	}
	defer file.Close()

	var data KeyData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return KeyData{}, nil
	}
	return data, nil
}

func writeJSON(filename string, data KeyData) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// Function to ensure .cache directory exists in the user's home directory
func ensureCacheDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %v", err)
	}

	cacheDir := filepath.Join(usr.HomeDir, ".cache")

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err := os.MkdirAll(cacheDir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create .cache directory: %v", err)
		}
	}

	return filepath.Join(cacheDir, "api-keys.json"), nil
}

func addKeys(filename string, newKeys []string) error {
	data, err := readJSON(filename)
	if err != nil {
		return err
	}

	for _, newKey := range newKeys {
		// Check if key already exists
		keyExists := false
		for _, key := range data.Keys {
			if _, exists := key[newKey]; exists {
				keyExists = true
				break
			}
		}
		if keyExists {
			fmt.Printf("Key '%s' already exists\n", newKey)
			continue
		}

		newMap := map[string]int{newKey: 0}
		data.Keys = append(data.Keys, newMap)
		fmt.Printf("Key '%s' added\n", newKey)
	}

	err = writeJSON(filename, data)
	if err != nil {
		return err
	}
	return nil
}

// List all keys and their usage (values)
func listKeys(filename string) error {
	data, err := readJSON(filename)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"API Key", "Usage Count"})

	for _, key := range data.Keys {
		for keyName, usageCount := range key {
			table.Append([]string{keyName, fmt.Sprintf("%d", usageCount)})
		}
	}

	table.Render()

	return nil
}

func main() {
	var filename string

	var rootCmd = &cobra.Command{
		Use:   "keymanager",
		Short: "A CLI tool to manage keys with default values.",
	}

	// Define the add command (add multiple keys)
	var addCmd = &cobra.Command{
		Use:   "add [keys]",
		Short: "Add new keys with a default value of 0.",
		Args:  cobra.MinimumNArgs(1), // Requires at least 1 key
		Run: func(cmd *cobra.Command, args []string) {
			err := addKeys(filename, args)
			if err != nil {
				fmt.Println("cannot add keys:", err)
			}
		},
	}

	// Define the list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all keys with their values.",
		Run: func(cmd *cobra.Command, args []string) {
			err := listKeys(filename)
			if err != nil {
				fmt.Println("cannot list keys:", err)
			}
		},
	}

	var err error
	filename, err = ensureCacheDir()
	if err != nil {
		fmt.Println("cannot ensure cache directory:", err)
		os.Exit(1)
	}

	// Add flags for the filename
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", filename, "Path to the JSON file")

	// Add the add and list commands to the root command
	rootCmd.AddCommand(addCmd, listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cannot execute command:", err)
		os.Exit(1)
	}
}

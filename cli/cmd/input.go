package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"scheduler/v2/models"
	"scheduler/v2/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	inputFile    string
	scheduleName string
)

// Sample input command for testing: ./scheduler input -p tests/valid.json -n school

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Takes a JSON file as input for the scheduling engine",
	Long:  `Processes a given JSON file that defines the constraints, resources, and tasks required to generate a schedule.`,

	Run: func(cmd *cobra.Command, args []string) {

		// Check if the input file is JSON
		utils.CheckJSON(inputFile)
		fmt.Printf("Reading data from file...\n")

		if inputFile == "" {
			fmt.Printf("No file provided! Please provide a path..")
			os.Exit(1)
		}

		/* Handle the name of the file.
		   If provided, then okay.
		   If not, then ask the user if they want to use a default name. */
		if scheduleName == "" {
			if utils.Conformation("Name flag not provided. Use default name?") {
				scheduleName = utils.RemoveDirectory(inputFile)
				fmt.Printf("Using default name: %s\n", scheduleName)
			} else {
				fmt.Println("Operation cancelled. Please provide a name using the -n or --name flag.")
				os.Exit(1)
			}
		} else {
			fmt.Printf("Schedule name provided: %s\n", scheduleName)
		}

		fileContent, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Error reading file: \n%v\n", err)
			os.Exit(1)
		}

		// Get and do a sample operation on the file contents
		// For example, checking the number of characters..

		var problemDef models.ProblemDefinition
		err = json.Unmarshal(fileContent, &problemDef)
		if err != nil {
			fmt.Printf("Error: Malformed JSON in %s: %v\n", inputFile, err)
			os.Exit(1)
		}

		utils.CheckValues(problemDef)

		fmt.Printf("Successfully read %d bytes from %s\n", len(fileContent), inputFile)

		// Storing the input constraint and requirements file locally
		storageDir := ".constraints"
		if err := os.MkdirAll(storageDir, 0755); err != nil {
			fmt.Println("Error creating storage directory!")
			os.Exit(1)
		}

		outputPath := filepath.Join(storageDir + scheduleName + ".json")

		// Check if the name already exists
		if _, err := os.Stat(outputPath); err == nil {
			color.Red("Error: A schedule named '%s' already exists!", scheduleName)
			os.Exit(1)
		}

		outputData, err := json.MarshalIndent(problemDef, "", "  ")
		if err != nil {
			color.Red("Error preparing data for saving: %v", err)
			os.Exit(1)
		}

		// Make the new JSON file
		if err := os.WriteFile(outputPath, outputData, 0644); err != nil {
			color.Red("Error saving schedule file: %v", err)
			os.Exit(1)
		}
		color.Green("Schedule '%s' was successfully saved.", scheduleName)
	},
}

func init() {
	rootCmd.AddCommand(inputCmd)

	inputCmd.Flags().StringVarP(&inputFile, "path", "p", "", "Path to the input JSON file")
	inputCmd.MarkFlagRequired("path")

	inputCmd.Flags().StringVarP(&scheduleName, "name", "n", "", "Optional flag to set name")
}

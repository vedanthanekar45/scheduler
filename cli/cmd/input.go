package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"scheduler/v2/models"
	"scheduler/v2/utils"

	"github.com/spf13/cobra"
)

var (
	inputFile    string
	scheduleName string
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Takes a JSON file as input for the scheduling engine",
	Long:  `Processes a given JSON file that defines the constraints, resources, and tasks required to generate a schedule.`,

	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckJSON(inputFile)
		fmt.Printf("Reading data from file...\n")

		if inputFile == "" {
			fmt.Printf("No file provided! Please provide a path..")
			os.Exit(1)
		}

		name_res := utils.Conformation("Name flag -n or --name not used. Use default name? [Y/n]: ")
		if name_res {
			if scheduleName == "" {
				scheduleName = utils.RemoveDirectory(inputFile)
				fmt.Printf("--name flag not provided. Using default name: %s\n", scheduleName)
			} else {
				fmt.Printf("Schedule name provided: %s\n", scheduleName)
			}
		}
		os.Exit(1)

		fileContent, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Error reading file: \n%v\n", err)
			os.Exit(1)
		}

		var problemDef models.ProblemDefinition
		err = json.Unmarshal(fileContent, &problemDef)
		if err != nil {
			fmt.Printf("Error: Malformed JSON in %s: %v\n", inputFile, err)
			os.Exit(1)
		}

		utils.CheckValues(problemDef)

		fmt.Printf("Successfully read %d bytes from %s\n", len(fileContent), inputFile)
	},
}

func init() {
	rootCmd.AddCommand(inputCmd)

	inputCmd.Flags().StringVarP(&inputFile, "path", "p", "", "Path to the input JSON file")
	inputCmd.MarkFlagRequired("path")

	inputCmd.Flags().StringVarP(&scheduleName, "name", "n", "", "Optional flag to set name")
}

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"scheduler/v2/models"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command

// Sample command to run it: ./scheduler generate [scheduleName]

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a schedule from your previously saved constraints file.",
	Long: `Reads a saved schedule by its name, runs the solver engine, and outputs a valid 
	timetable.`,
	Run: func(cmd *cobra.Command, args []string) {
		scheduleName := args[0]
		storageDir := "schedules"

		filePath := filepath.Join(storageDir, scheduleName+".json")

		// Looking for the file
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			color.Red("Error: No schedule named '%s' found.", scheduleName)
			fmt.Printf("Looked for file at: %s\n", filePath)
			os.Exit(1)
		}

		// Try to parse the data inside the file
		var problemDef models.ProblemDefinition
		if err := json.Unmarshal(fileContent, &problemDef); err != nil {
			color.Red("Error parsing saved schedule file: %v", err)
			os.Exit(1)
		}

		color.Green("Data parsed successfully. Starting the solver...")

		/* COMMANDS TO START THE SOLVER GOES HERE */
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

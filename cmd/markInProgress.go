package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ReflectionShade/task-cli/tasksHandlers"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "update status task - `in-progress`",
	Long:  "update status task - `in-progress",
	Run: func(cmd *cobra.Command, args []string) {
		textInput := pterm.DefaultInteractiveTextInput
		textInput.DefaultText = "Write id of the task you want to update-status"
		strTaskID, err := textInput.Show()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		taskID, err := strconv.Atoi(strTaskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = tasksHandlers.UpdateTaskStatus(2, taskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

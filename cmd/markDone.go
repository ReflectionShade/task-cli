package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ReflectionShade/task-cli/tasksHandlers"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "update status task - is done",
	Long:  "update status task - is done",
	Run: func(cmd *cobra.Command, args []string) {
		textInput := pterm.DefaultInteractiveTextInput
		textInput.DefaultText = "Write ID of the task you want to update-status"
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

		err = tasksHandlers.UpdateTaskStatus(3, taskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

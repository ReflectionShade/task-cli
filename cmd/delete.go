package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ReflectionShade/task-cli/tasksHandlers"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task by ID",
	Long:  `Its command delete command by his ID, you need to print ID for task for delete`,
	Run: func(cmd *cobra.Command, args []string) {
		textInput := pterm.DefaultInteractiveTextInput
		textInput.DefaultText = "Write id of the task you want to delete:"
		strTaskID, _ := textInput.Show()
		taskID, err := strconv.Atoi(strTaskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = tasksHandlers.DeleteTask(taskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Task successfully deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

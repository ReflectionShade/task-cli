package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/ReflectionShade/task-cli/tasksHandlers"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Its command for task-cli, he is add your task",
	Long:  `Its command for task-cli, it is add your task in task-list`,
	Run: func(cmd *cobra.Command, args []string) {
		textInput := pterm.DefaultInteractiveTextInput
		textInput.DefaultText = "Write the description of the task you want to add:"
		descriptionTask, _ := textInput.Show()
		task := tasksHandlers.TaskConstructor(descriptionTask)
		err := tasksHandlers.AppendTask(task)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

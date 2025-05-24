package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ReflectionShade/task-cli/tasksHandlers"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update your task description",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// получаем описание задачи из терминала
		descriptionInput := pterm.DefaultInteractiveTextInput
		descriptionInput.DefaultText = "Write the updated description of the task you want to update:"
		descriptionTask, _ := descriptionInput.Show()

		// получаем и конвертируем ID
		IDInput := pterm.DefaultInteractiveTextInput
		IDInput.DefaultText = "Okey, write ID for task"
		strtaskID, err := IDInput.Show()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		taskID, err := strconv.Atoi(strtaskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = tasksHandlers.UpdateTaskDescription(descriptionTask, taskID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Task successfully updated!!!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

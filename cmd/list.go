package cmd

import (
	"fmt"
	"os"

	"github.com/ReflectionShade/task-cli/tasksHandlers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "prints all tasks",
	Long: `list have a subcommands:
	in-progress (prints all tasks where status in progress) 
	done (prints all tasks where status is done)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasksHandlers.ReadTasks()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, task := range tasks {
			var status string
			// 1 - это todo, 2 - in-progress, 3 - done
			if task.Status == 1 {
				status = "todo"
			} else if task.Status == 2 {
				status = "in progress"
			} else {
				status = "is done"
			}
			fmt.Printf("ID: %d, description task:%s, status - %s, created at: %s, updated at: %s\n",
				task.ID,
				task.Description,
				status,
				task.CreatedAt.Format("Mon Jan 2 15:04:05 MST 2006"),
				task.UpdatedAt.Format("Mon Jan 2 15:04:05 MST 2006"),
			)
		}

	},
}

var listDoneCmd = &cobra.Command{
	Use:   "done",
	Short: "List tasks with status 'done'",
	Long:  `List all tasks that have the status 'done'.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasksHandlers.ReadTasks()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, task := range tasks {
			if task.Status == 3 {
				status := "is done"
				// print in format
				fmt.Printf("ID: %d, description task:%s, status - %s, created at: %s, updated at: %s\n",
					task.ID,
					task.Description,
					status,
					task.CreatedAt.Format("Mon Jan 2 15:04:05 MST 2006"),
					task.UpdatedAt.Format("Mon Jan 2 15:04:05 MST 2006"),
				)
			}
		}
	},
}

var listInProgress = &cobra.Command{
	Use:   "in-progress",
	Short: "List tasks with status 'in progress'",
	Long:  `List all tasks that have the status 'in-progress'.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasksHandlers.ReadTasks()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, task := range tasks {
			if task.Status == 2 {
				status := "in progress"

				fmt.Printf("ID: %d, description task:%s, status - %s, created at: %s, updated at: %s\n",
					task.ID,
					task.Description,
					status,
					task.CreatedAt.Format("Mon Jan 2 15:04:05 MST 2006"),
					task.UpdatedAt.Format("Mon Jan 2 15:04:05 MST 2006"),
				)
			}
		}
	},
}

func init() {
	// Adding listCmd to rootCmd
	rootCmd.AddCommand(listCmd)

	// adding the subcommands
	listCmd.AddCommand(listDoneCmd)
	listCmd.AddCommand(listInProgress)

}

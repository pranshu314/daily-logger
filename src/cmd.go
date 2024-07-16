package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	terminal_red              = "\033[1;31m"
	terminal_yellow           = "\033[1;33m"
	terminal_reset_formatting = "\033[0m"
)

var rootCmd = &cobra.Command{
	Use:   "lg",
	Short: "A cli application to log your thoughts while coding.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Where your logs are stored.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Println(setupPath())
		return err
	},
}

var addLogCmd = &cobra.Command{
	Use:   "project [Project Name]",
	Short: "Adds a log to the specified project.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ldb, err := openDB(setupPath())
		if err != nil {
			return err
		}
		defer ldb.db.Close()

		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		var lines string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			if len(strings.TrimSpace(line)) == 0 {
				break
			}
			lines = lines + line
		}

		if err := ldb.insert(args[0], lines); err != nil {
			return err
		}

		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Deletes log entry by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ldb, err := openDB(setupPath())
		if err != nil {
			return err
		}
		defer ldb.db.Close()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		return ldb.delete(uint(id))
	},
}

var listCmd = &cobra.Command{
	Use:   "list [Project Name]",
	Short: "List all your logs for the given project.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ldb, err := openDB(setupPath())
		if err != nil {
			return err
		}
		defer ldb.db.Close()

		logs, err := ldb.getProjectLogs(args[0])
		if err != nil {
			return err
		}

		for _, v := range logs {
			fmt.Printf("%s%s: (%d) [%s]%s \n%s\n", terminal_yellow, v.CREATED, v.ID, v.PROJECT, terminal_reset_formatting, v.LOG_ENTRY)
		}

		return nil
	},
}

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List all the projects.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ldb, err := openDB(setupPath())
		if err != nil {
			return err
		}
		defer ldb.db.Close()

		projects, err := ldb.getProjects()
		if err != nil {
			return err
		}

		for _, v := range projects {
			fmt.Println(v)
		}

		return nil
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [ID]",
	Short: "Update a log entry by ID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ldb, err := openDB(setupPath())
		if err != nil {
			return err
		}
		defer ldb.db.Close()

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		log_entry, err := cmd.Flags().GetString("log_entry")
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		newLog := lg{time.Time{}, log_entry, project, uint(id)}

		return ldb.update(newLog)
	},
}

func init() {
	updateCmd.Flags().StringP(
		"project",
		"p",
		"",
		"specify the project name",
	)
	updateCmd.Flags().StringP(
		"log_entry",
		"l",
		"",
		"specify the log entry",
	)

	rootCmd.AddCommand(whereCmd)
	rootCmd.AddCommand(addLogCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(projectsCmd)
}

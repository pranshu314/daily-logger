package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

func setupPath() string {
	scope := gap.NewScope(gap.User, "daily_logs")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}

	var logDir string
	if len(dirs) > 0 {
		logDir = dirs[0]
	} else {
		logDir, _ = os.UserHomeDir()
	}

	if err := initLogDir(logDir); err != nil {
		log.Fatal(err)
	}

	return logDir
}

func openDB(path string) (*logDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "dailyl_logs.db"))
	if err != nil {
		return nil, err
	}

	lg := logDB{db, path}
	if !lg.tableExists("logs") {
		// fmt.Println("Inside tableExists")
		err := lg.createTable("logs")
		if err != nil {
			return nil, err
		}
	}

	return &lg, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// type proj struct {
// 	CREATED     time.Time
// 	NAME        string
// 	DESCRIPTION string
// 	REPO_LINK   string
// 	ID          uint
// }

type lg struct {
	CREATED   time.Time
	LOG_ENTRY string
	PROJECT   string
	ID        uint
}

type logDB struct {
	db      *sql.DB
	dataDir string
}

// func (p proj) Name() string {
// 	return p.NAME
// }

// func (p proj) Description() string {
// 	return p.DESCRIPTION
// }

func (l lg) Get() string {
	return fmt.Sprintf(
		"%s: {%s} %s",
		l.CREATED.Format("2006-01-02"),
		l.PROJECT,
		l.LOG_ENTRY,
	)
}

func (ldb *logDB) tableExists(name string) bool {
	if _, err := ldb.db.Exec("SELECT * FROM ?", name); err == nil {
		return true
	}
	return false
}

func (ldb *logDB) createTable(name string) error {
	switch name {
	// case "projects":
	// 	_, err := ldb.db.Exec(`CREATE TABLE "projects" ( "id" INTEGER, "created" DATETIME, "name" TEXT NOT NULL, "description" TEXT, "repo_link" TEXT, PRIMARY KEY("id" AUTOINCREMENT) )`)
	// 	return err
	case "logs":
		_, err := ldb.db.Exec(`CREATE TABLE "logs" ( "id" INTEGER, "created" DATETIME, "project" TEXT, "log_entry" BLOB, PRIMARY KEY("id" AUTOINCREMENT) )`)
		return err
	}
	return errors.New("The program doesnot support creation of table with the name " + name)
}

func dummy() {
	logg := lg{time.Now(), "Log entry", "Dummy Project", 456}
	fmt.Println(logg.Get())
}

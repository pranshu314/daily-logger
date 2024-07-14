package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type proj struct {
	CREATED     time.Time
	NAME        string
	DESCRIPTION string
	REPO_LINK   string
	ID          uint
}

type lg struct {
	CREATED    time.Time
	LOG_ENTRY  string
	PROJECT_ID uint
	ID         uint
}

type logDB struct {
	db      *sql.DB
	dataDir string
}

func (p proj) Name() string {
	return p.NAME
}

func (p proj) Description() string {
	return p.DESCRIPTION
}

func (l lg) Get() string {
	return fmt.Sprintf(
		"%s: {pid %d} %s",
		l.CREATED.Format("2006-01-02"),
		l.PROJECT_ID,
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
	case "projects":
		// TODO: Complete the SQL Query
		_, err := ldb.db.Exec(`CREATE TABLE "projects"`)
		return err
	case "logs":
		// TODO: Complete the SQL Query
		_, err := ldb.db.Exec(`CREATE TABLE "logs"`)
		return err
	}
	return errors.New("The program doesnot support creation of table with the name " + name)
}

func dummy() {
	logg := lg{time.Now(), "Log entry", 123, 456}
	fmt.Println(logg.Get())
}

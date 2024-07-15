package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"reflect"
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

func initLogDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
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

func (ldb *logDB) insert(project string, log_entry string) error {
	_, err := ldb.db.Exec(
		"INSERT INTO logs(created, project, log_entry) VALUES (?, ?, ?)",
		time.Now(),
		project,
		log_entry,
	)

	return err
}

func (ldb *logDB) delete(id uint) error {
	_, err := ldb.db.Exec(
		"DELETE FROM logs WHERE id = ?",
		id,
	)

	return err
}

func (orij *lg) merge(lg lg) {
	uValues := reflect.ValueOf(&lg).Elem()
	oValues := reflect.ValueOf(orij).Elem()

	for i := 0; i < uValues.NumField(); i++ {
		uField := uValues.Field(i).Interface()
		if oValues.CanSet() {
			if v, ok := uField.(int64); ok && uField != 0 {
				oValues.Field(i).SetInt(v)
			}
			if v, ok := uField.(string); ok && uField != "" {
				oValues.Field(i).SetString(v)
			}
		}
	}
}

func (ldb *logDB) update(lg lg) error {
	orij, err := ldb.getLogEntry(lg.ID)
	if err != nil {
		return err
	}

	orij.merge(lg)
	_, err = ldb.db.Exec(
		"UPDATE logs SET project = ?, log_entry = ? WHERE id = ?",
		orij.PROJECT,
		orij.LOG_ENTRY,
		orij.ID,
	)

	return err
}

func (ldb *logDB) getLogEntry(id uint) (lg, error) {
	var lg lg
	err := ldb.db.QueryRow("SELECT * FROM logs WHERE id = ?", id).Scan(
		&lg.ID,
		&lg.PROJECT,
		&lg.LOG_ENTRY,
		&lg.CREATED,
	)

	return lg, err
}

func (ldb *logDB) getProjectLogs(project string) ([]lg, error) {
	var lgs []lg
	rows, err := ldb.db.Query("SELECT * FROM logs WHERE project = ? ORDER BY created ASC", project)
	if err != nil {
		return lgs, fmt.Errorf("unable to get values: %w", err)
	}

	for rows.Next() {
		var lg lg
		err = rows.Scan(
			&lg.ID,
			&lg.CREATED,
			&lg.PROJECT,
			&lg.LOG_ENTRY,
		)
		if err != nil {
			return lgs, err
		}
		lgs = append(lgs, lg)
	}

	return lgs, err
}

package project

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"strings"
)

type (
	PGHandler struct{}
)

func (pg *PGHandler) Init() error {
	return pg.Run(`CREATE TABLE MG_VERSION (version text);`)
}

func (pg *PGHandler) connect() (*sql.DB, error) {
	cStr := fmt.Sprintf("dbname=%s user='%s' password='%s' port='%s' host='%s' sslmode='%s'",
		viper.GetString("database"), viper.GetString("user"), viper.GetString("password"),
		viper.GetString("port"), viper.GetString("host"), viper.GetString("pgssl"))
	return sql.Open("postgres", cStr)
}

func (pg *PGHandler) CheckMigration(v string) (bool, error) {
	db, err := pg.connect()
	if err != nil {
		return false, err
	}
	var count int
	err = db.
		QueryRow(`SELECT COUNT(*) FROM MG_VERSION WHERE version = $1;`, v).
		Scan(&count)
	return (count > 0), err
}

func (pg *PGHandler) MarkMigration(version string) error {
	return pg.Run(`INSERT INTO MG_VERSION (version) VALUES ('` + version + `');`)
}

func (pg *PGHandler) UnMarkMigration(version string) error {
	return pg.Run(`DELETE FROM MG_VERSION WHERE version='` + version + `';`)
}

func (pg *PGHandler) GetBackwardsVersions(lastVersion string) ([]string, error) {
	db, err := pg.connect()
	if err != nil {
		return []string{}, err
	}
	var versions []string
	rows, err := db.Query("SELECT version FROM MG_VERSION ORDER BY version DESC")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return []string{}, err
		}
		if strings.TrimSpace(lastVersion) == "" {
			versions = append(versions, version)
			return versions, nil
		}
		if strings.TrimSpace(lastVersion) == strings.TrimSpace(version) {
			versions = append(versions, version)
			return versions, nil
		} else {
			versions = append(versions, version)
		}
	}
	if strings.TrimSpace(lastVersion) == "" {
		return []string{}, errors.New("Nothing to revert.")
	} else {
		return []string{}, errors.New("Version not found.")
	}
}

func (pg *PGHandler) Run(script string) error {
	commands := strings.Split(script, ";")
	db, err := pg.connect()
	if err != nil {
		return err
	}
	for _, data := range commands {
		txn, err := db.Begin()
		if err != nil {
			return err
		}

		stmt, err := txn.Prepare(data)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			return err
		}

		err = stmt.Close()
		if err != nil {
			return err
		}

		err = txn.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

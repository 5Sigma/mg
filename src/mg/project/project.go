package project

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"mg/console"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func createPath(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, os.ModePerm)
	}
}

func CreateProject(filePath string) error {
	createPath(filePath)
	handler, err := GetHandler()
	if err != nil {
		return err
	}
	return handler.Init()
}

func NewMigration(name string) MigrationFile {
	ts := time.Now().Format("20060102150405")
	re := regexp.MustCompile("\\s+")
	desc := re.ReplaceAllString(name, "_")
	filename := ts + "_" + strings.ToLower(desc) + ".sql"
	filepath := path.Join(viper.GetString("migration-path"), filename)
	os.Create(filepath)
	return MigrationFile{FilePath: filepath}
}

func migrateFile(fp string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	if filepath.Ext(fp) != ".sql" {
		return nil
	}
	mFile := MigrationFile{FilePath: fp}

	contents, err := mFile.GetUp()
	if err != nil {
		return err
	}

	handler, err := GetHandler()
	if err != nil {
		return err
	}

	check, err := handler.CheckMigration(mFile.Version())
	if err != nil {
		return err
	}

	if check == true {
		return nil
	}

	if err != nil {
		return err
	}

	console.Alertf("Migrating %s", mFile.FilePath)
	start := time.Now()
	err = handler.Run(contents)
	elapsed := time.Since(start)
	console.Printf("Completed in %f seconds.", elapsed.Seconds())
	if err != nil {
		return err
	}
	return handler.MarkMigration(mFile.Version())
}

func Migrate() error {
	start := time.Now()
	err := filepath.Walk(viper.GetString("migration-path"), migrateFile)
	elapsed := time.Since(start)
	if err == nil {
		console.Successf("\nCompleted in %f seconds.", elapsed.Seconds())
	}
	return err
}

func revertVersion(version string) error {
	matches, err := filepath.Glob(path.Join(viper.GetString("migration-path"), version+"_*.sql"))
	if len(matches) == 0 {
		return errors.New(fmt.Sprintf("Could not find migration file for %s", version))
	}
	fp := matches[0]
	mf := MigrationFile{FilePath: fp}
	contents, err := mf.GetDown()
	if err != nil {
		return err
	}
	handler, err := GetHandler()
	if err != nil {
		return err
	}
	err = handler.Run(contents)
	if err != nil {
		console.Fatalf("Could not revert: %s\n%s", mf.FilePath, err.Error())
	}
	err = handler.UnMarkMigration(version)
	if err != nil {
		console.Fatalf("Could not remove migration:\n%s", err.Error())
	}
	return nil
}

func Revert(version string) error {
	handler, err := GetHandler()
	if err != nil {
		return err
	}
	versions, err := handler.GetBackwardsVersions(version)
	if err != nil {
		console.Fatalf("Could not get version list:\n%s", err)
		return err
	}
	if len(versions) == 0 {
		console.Fatalf("Could not find any versions to revert.")
	}
	console.Alertf("Rolling back %d migrations.", len(versions))
	for _, v := range versions {
		err = revertVersion(v)
		if err != nil {
			return err
		}
	}
	return nil
}

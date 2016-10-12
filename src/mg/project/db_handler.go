package project

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

type (
	DBHandler interface {
		Run(filepath string) error
		Init() error
		CheckMigration(version string) (bool, error)
		MarkMigration(version string) error
		UnMarkMigration(version string) error
		GetBackwardsVersions(version string) ([]string, error)
	}
)

func GetHandler() (DBHandler, error) {
	if strings.ToUpper(viper.GetString("driver")) == "POSTGRES" {
		return &PGHandler{}, nil
	}
	return nil, errors.New("Invalid driver")
}

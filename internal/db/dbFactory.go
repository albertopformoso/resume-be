package db

import (
	"log"

	"github.com/albertopformoso/resume-be/internal/model"
)

// Database supported driver list
const (
	DriverUndefined model.DBDriver = 0
	DriverPG        model.DBDriver = 1
)

// DB database interface
type DB interface {
}

func NewDB(driver model.DBDriver) DB {
	switch driver {
	case DriverPG:
		return &postgres{}
	case DriverUndefined:
		log.Printf("ERROR: driver %q not supported", driver)
		return nil
	default:
		log.Printf("ERROR: driver %q not supported", driver)
		return nil
	}
}

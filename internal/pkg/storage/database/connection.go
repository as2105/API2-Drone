package database

import (
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/config"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // SQL adapters needed for Gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // SQL adapters needed for Gorm
)

// DB ...
type DB = gorm.DB

type gormLogger struct {
	log *logging.Logger
}

func (l *gormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		l.log.WithFields(logging.Fields{"module": "gorm", "type": "sql"}).Info(v[3])
	}
	if v[0] == "log" {
		l.log.WithFields(logging.Fields{"module": "gorm", "type": "log"}).Info(v[2])
	}
}

// NewConnection ...
func NewConnection(log *logging.Logger, config *config.Config) (*DB, error) {
	db, err := gorm.Open(config.DatabaseType, config.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetLogger(&gormLogger{log})
	db.LogMode(true)

	// we will manage these manually
	db.Callback().Create().Remove("gorm:update_time_stamp")
	db.Callback().Update().Remove("gorm:update_time_stamp")

	return db, nil
}

package model

import (
	"bookmarks/server/common/logger"
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

// ConstructDSN constructs a DSN string for PostgreSQL connection.
func ConstructDSN(host, user, password, dbname, port string) string {
	return "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable TimeZone=Asia/Kolkata"
}

// InitializeDatabase initializes the database connection using GORM and returns a DB instance.
func InitializeDatabase(dsn string) (*DB, error) {
	log := logger.GetLogger()
	ctx := context.Background()

	config := &gorm.Config{}
	if log.GetLogLevel() == logger.DEBUG {
		// cl = gLogger.Config{
		// 	SlowThreshold:             time.Second,
		// 	LogLevel:                  gLogger.Info,
		// 	IgnoreRecordNotFoundError: true,
		// 	Colorful:                  true,
		// }
		// config.Logger = gLogger.New(log, cl)
		// TODO: Add this.
		// Spent too much time on this.
		// Using default logger for now.
		// Later, can refactor to make this work.
		config.Logger = gLogger.Default.LogMode(gLogger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Error(ctx, "Failed to connect to database: "+err.Error())
		return nil, err
	}

	if err := db.Raw("SELECT 1").Error; err != nil {
		log.Error(ctx, "Failed to ping database: "+err.Error())
		return nil, err
	}

	log.Info(ctx, "Connection established with "+db.Dialector.Name()+" database")

	return &DB{Conn: db}, nil
}

// TODO: Make this flavor agnostic

package model

import (
	"bookmarks/server/common/logger"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

type DSNConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func (c *DSNConfig) Validate() error {
	if strings.TrimSpace(c.Host) == "" {
		return errors.New("DB host is required")
	}
	if strings.TrimSpace(c.User) == "" {
		return errors.New("DB user is required")
	}
	if strings.TrimSpace(c.DBName) == "" {
		return errors.New("DB name is required")
	}

	if c.Port != "" {
		port, err := strconv.Atoi(c.Port)
		if err != nil || port <= 0 || port > 65535 {
			return errors.New("DB port must be a valid number between 1 and 65535")
		}
	} else {
		c.Port = "5432" // default PostgreSQL port
	}

	validSSLModes := map[string]bool{
		"disable":     true,
		"require":     true,
		"verify-ca":   true,
		"verify-full": true,
	}
	if c.SSLMode != "" && !validSSLModes[c.SSLMode] {
		return errors.New("DB SSL mode must be one of: disable, require, verify-ca, verify-full")
	} else if c.SSLMode == "" {
		c.SSLMode = "disable" // default SSL mode
	}

	if c.TimeZone == "" {
		c.TimeZone = "Asia/Kolkata" // default time zone
	}

	return nil
}

func cleanedValue(value string) string {
	if strings.Contains(value, " ") || strings.ContainsAny(value, `'\"\\=`) {
		// postgres uses single quotes for escaping
		return "'" + strings.ReplaceAll(value, "'", "''") + "'"
	}
	return value
}

// ConstructDSN constructs a DSN string for PostgreSQL connection.
func ConstructDSN(config DSNConfig) (string, error) {
	if err := config.Validate(); err != nil {
		return "", err
	}

	parts := []string{
		fmt.Sprintf("host=%s", cleanedValue(config.Host)),
		fmt.Sprintf("user=%s", cleanedValue(config.User)),
		fmt.Sprintf("password=%s", cleanedValue(config.Password)),
		fmt.Sprintf("dbname=%s", cleanedValue(config.DBName)),
		fmt.Sprintf("port=%s", cleanedValue(config.Port)),
		fmt.Sprintf("sslmode=%s", cleanedValue(config.SSLMode)),
		fmt.Sprintf("TimeZone=%s", cleanedValue(config.TimeZone)),
	}
	return strings.Join(parts, " "), nil
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

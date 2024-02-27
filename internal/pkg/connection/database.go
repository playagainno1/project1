package connection

import (
	"runtime"
	"taylor-ai-server/internal/config"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var db *sqlx.DB

var _ = mysql.Config{}

func DB() *sqlx.DB {
	return db
}

func InitDB() {
	c := config.Config.Database
	if c.Driver != "mysql" {
		logrus.Fatal("Only mysql is supported for now")
	}
	dsn := c.DSN

	logrus.WithFields(logrus.Fields{
		"driver": c.Driver,
		"dsn":    dsn,
	}).Info("Connecting to database...")

	conn, err := newDB(c.Driver, dsn)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}
	db = conn
}

func newDB(driver, dsn string) (*sqlx.DB, error) {
	conn, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(runtime.NumCPU() * 2)
	return conn.Unsafe(), nil
}

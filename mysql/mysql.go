package mysql

import (
	"database/sql"
	"fmt"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

// New returns a new sql.DB with the provided parameters. If the Ping to the DB fails
// due to not existing DB it'll create the DB
func New(host string, port int, user, password string, ops Options) (*sql.DB, error) {
	if host == "" {
		return nil, xerrors.New("host is a required parameter")
	} else if port == 0 {
		return nil, xerrors.New("port is a required parameter")
	} else if user == "" {
		return nil, xerrors.New("user is a required parameter")
	} else if password == "" {
		return nil, xerrors.New("password is a required parameter")
	}

	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?clientFoundRows=%t&parseTime=%t&multiStatements=%t",
		user, password, host, port, ops.DBName, ops.ClientFoundRows, ops.ParseTime, ops.MultiStatements,
	)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, xerrors.Errorf("could not connect to the MySQL database: %w", err)
	}

	// If we get an error of ER_BAD_DB_ERROR means that the DB was not found, so not created
	// so we have to create it, which means to start a new connection without the DBName specified
	// and we create the DB and then "retry"
	var sqlerr *mysql.MySQLError
	if err := db.Ping(); err != nil {
		if xerrors.As(err, &sqlerr) {
			if sqlerr.Number == mysqlerr.ER_BAD_DB_ERROR {
				ndns := fmt.Sprintf(
					"%s:%s@tcp(%s:%d)/%s?clientFoundRows=%t&parseTime=%t&multiStatements=%t",
					user, password, host, port, "", ops.ClientFoundRows, ops.ParseTime, ops.MultiStatements,
				)

				ndb, err := sql.Open("mysql", ndns)
				if err != nil {
					return nil, xerrors.Errorf("could not connect to the MySQL database to create database: %w", err)
				}
				defer ndb.Close()

				if err := ndb.Ping(); err != nil {
					return nil, xerrors.Errorf("could not ping DB to create database: %w", err)
				}

				_, err = ndb.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", ops.DBName))
				if err != nil {
					return nil, xerrors.Errorf("could not create DB %s: %w", ops.DBName, err)
				}

				if err := db.Ping(); err != nil {
					return nil, xerrors.Errorf("could not ping DB to check database created: %w", err)
				}

			}
		} else {
			return nil, xerrors.Errorf("could not ping DB: %w", err)
		}
	}

	return db, nil
}

// Options list of options that can be assigned to the New function
type Options struct {
	DBName          string
	ClientFoundRows bool
	ParseTime       bool
	MultiStatements bool
}

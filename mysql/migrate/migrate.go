package migrate

import (
	"database/sql"

	"github.com/lopezator/migrator"
	"github.com/xescugc/chaoswall/mysql/migrate/migrations"
	"golang.org/x/xerrors"
)

// Migrate runs the migrations on the provided db
func Migrate(db *sql.DB) error {
	ms := make([]interface{}, 0, len(migrations.Migrations))
	for _, m := range migrations.Migrations {
		val := m
		ms = append(ms, &migrator.Migration{
			Name: val.Name,
			Func: func(tx *sql.Tx) error {
				if _, err := tx.Exec(val.SQL); err != nil {
					return err
				}
				return nil
			},
		})
	}

	m, err := migrator.New(migrator.Migrations(ms...))
	if err != nil {
		return xerrors.Errorf("error while creating the migratior: %w", err)
	}

	if err := m.Migrate(db); err != nil {
		return xerrors.Errorf("error while migrating: %w", err)
	}

	return nil
}

package e2e_test

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/xescugc/chaoswall/mysql"
	"github.com/xescugc/chaoswall/mysql/migrate"
	"github.com/xescugc/chaoswall/service"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	// TODO: For some reason this cfg is empty but the
	// viper 'viper.GetString("db-host")' returns the
	// expected ENV value
	//cfg, err := config.New(viper.GetViper())
	//if err != nil {
	//log.Fatal(err)
	//}
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	dropDB, err := mysql.New(viper.GetString("db-host"), viper.GetInt("db-port"), viper.GetString("db-user"), viper.GetString("db-password"), mysql.Options{
		MultiStatements: true,
		ClientFoundRows: true,
	})
	if err != nil {
		panic(err)
	}

	_, err = dropDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", viper.GetString("db-name")))
	if err != nil {
		panic(err)
	}
	defer dropDB.Close()

	migrateDB, err := mysql.New(viper.GetString("db-host"), viper.GetInt("db-port"), viper.GetString("db-user"), viper.GetString("db-password"), mysql.Options{
		DBName:          viper.GetString("db-name"),
		MultiStatements: true,
		ClientFoundRows: true,
	})
	if err != nil {
		panic(err)
	}

	err = migrate.Migrate(migrateDB)
	if err != nil {
		panic(err)
	}

	defer migrateDB.Close()

	db, err = mysql.New(viper.GetString("db-host"), viper.GetInt("db-port"), viper.GetString("db-user"), viper.GetString("db-password"), mysql.Options{
		DBName:          viper.GetString("db-name"),
		MultiStatements: true,
		ClientFoundRows: true,
	})
	defer db.Close()

	exitVal := m.Run()

	os.Exit(exitVal)
}

func newService(t *testing.T) service.Service {

	gr := mysql.NewGymRepository(db)
	wr := mysql.NewWallRepository(db)
	rr := mysql.NewRouteRepository(db)
	suow := mysql.StartUnitOfWork(db)

	return service.New(nil, gr, wr, nil, rr, suow)
}

package cmd

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	kitlog "github.com/go-kit/kit/log"
	"github.com/markbates/pkger"
	"golang.org/x/xerrors"

	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xescugc/chaoswall/config"
	"github.com/xescugc/chaoswall/mysql"
	"github.com/xescugc/chaoswall/mysql/migrate"
	"github.com/xescugc/chaoswall/service"
	serviceHTTP "github.com/xescugc/chaoswall/service/transport/http"
)

type opener struct{}

func (o opener) Open(path string) (http.File, error) {
	return pkger.Open(path)
}

var (
	op       = opener{}
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "ChaosWall serve",
		Long:  "ChaosWall serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(viper.GetViper())
			if err != nil {
				return xerrors.Errorf("failed initializing the config: %w", err)
			}

			if err = cfg.Validate(); err != nil {
				return xerrors.Errorf("failed validating the config: %w", err)
			}

			logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stdout))
			logger = kitlog.With(logger, "ts", kitlog.TimestampFormat(time.Now, time.RFC3339), "caller", kitlog.DefaultCaller)

			// Initializing DB
			logger.Log("msg", "MariaDB starting ...")
			db, err := mysql.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, mysql.Options{
				DBName:          cfg.DBName,
				MultiStatements: true,
				ClientFoundRows: true,
			})
			if err != nil {
				panic(err)
			}
			logger.Log("msg", "MariaDB started")

			logger.Log("msg", "migrations running ...")
			err = migrate.Migrate(db)
			if err != nil {
				panic(err)
			}
			logger.Log("msg", "migrations finished")

			// Initializint Repositories
			gr := mysql.NewGymRepository(db)
			wr := mysql.NewWallRepository(db)
			hr := mysql.NewHoldRepository(db)
			rr := mysql.NewRouteRepository(db)
			suow := mysql.StartUnitOfWork(db)

			s := service.New(nil, gr, wr, hr, rr, suow)

			mux := http.NewServeMux()

			mux.Handle("/", serviceHTTP.MakeHandler(s))

			mux.Handle("/assets/", http.FileServer(op))

			log := handlers.CustomLoggingHandler(os.Stdout, mux, func(writer io.Writer, params handlers.LogFormatterParams) {
				username := "-"
				if params.URL.User != nil {
					if name := params.URL.User.Username(); name != "" {
						username = name
					}
				}

				host, _, err := net.SplitHostPort(params.Request.RemoteAddr)
				if err != nil {
					host = params.Request.RemoteAddr
				}

				uri := params.Request.RequestURI

				// Requests using the CONNECT method over HTTP/2.0 must use
				// the authority field (aka r.Host) to identify the target.
				// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
				if params.Request.ProtoMajor == 2 && params.Request.Method == "CONNECT" {
					uri = params.Request.Host
				}
				if uri == "" {
					uri = params.URL.RequestURI()
				}
				logger.Log(
					"host", host,
					"username", username,
					"method", params.Request.Method,
					"uri", uri,
					"status", strconv.Itoa(params.StatusCode),
					"size", strconv.Itoa(params.Size),
				)
			})

			logger.Log("port", cfg.Port, "msg", "started server")
			return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), removeTrailingSlash(log))
		},
	}
)

// removeTrailingSlash removes the end (trailing) `/` from the
// URL so it can match `/items/` and `/items` equally
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func init() {
	serveCmd.PersistentFlags().IntP("port", "p", config.DefaultPort, "Destination port")
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))

	serveCmd.PersistentFlags().String("db-host", "", "The database server IP or host name to connect")
	viper.BindPFlag("db-host", serveCmd.PersistentFlags().Lookup("db-host"))

	serveCmd.PersistentFlags().Uint32("db-port", 0, "The port where the database server is listening")
	viper.BindPFlag("db-port", serveCmd.PersistentFlags().Lookup("db-port"))

	serveCmd.PersistentFlags().String("db-user", "", "The user to use to connect to the database server")
	viper.BindPFlag("db-user", serveCmd.PersistentFlags().Lookup("db-user"))

	serveCmd.PersistentFlags().String("db-password", "", "The password of the database user to use")
	viper.BindPFlag("db-password", serveCmd.PersistentFlags().Lookup("db-password"))

	serveCmd.PersistentFlags().String("db-name", "", "The name of the database to use")
	viper.BindPFlag("db-name", serveCmd.PersistentFlags().Lookup("db-name"))
}

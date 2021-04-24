package cli

import (
	"github.com/schidstorm/ffmpeg-jobs/api/config"
	"github.com/schidstorm/ffmpeg-jobs/api/dependencies"
	"github.com/schidstorm/ffmpeg-jobs/api/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := cobra.Command{
		RunE: runApplication,
	}

	rootCmd.PersistentFlags().String("dialector", "sqlite", "Database dialector (postgres, sqlite)")
	rootCmd.PersistentFlags().String("dsn", "test.db", "Database connection string")
	rootCmd.PersistentFlags().String("listen", "0.0.0.0:8081", "Url of the API server")
	rootCmd.PersistentFlags().String("logLevel", logrus.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

	err := rootCmd.Execute()
	if err != nil {
		logrus.Error(err)
	}
}

func runApplication(cmd *cobra.Command, _ []string) error {
	logLevelString, err := cmd.PersistentFlags().GetString("logLevel")
	if err != nil {
		return err
	}

	logLevel, err := logrus.ParseLevel(logLevelString)
	if err != nil {
		return err
	}

	logrus.SetLevel(logLevel)

	cfg, err := parseConfig(cmd)
	if err != nil {
		return err
	}

	dependencies.InitConnection(cfg.DatabaseDialector, cfg.DatabaseDsn)
	srv := server.NewServer()
	server.InitializeServer(srv)
	return srv.Serve(cfg.ListenAddress)
}

func parseConfig(cmd *cobra.Command) (*config.Config, error) {
	dialector, err := cmd.PersistentFlags().GetString("dialector")
	if err != nil {
		return nil, err
	}

	dsn, err := cmd.PersistentFlags().GetString("dsn")
	if err != nil {
		return nil, err
	}

	listen, err := cmd.PersistentFlags().GetString("listen")
	if err != nil {
		return nil, err
	}


	return &config.Config{
		ListenAddress:  listen,
		DatabaseDialector:        dialector,
		DatabaseDsn: dsn,
	}, nil
}

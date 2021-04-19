package cli

import (
	"github.com/schidstorm/ffmpeg-jobs/worker/config"
	"github.com/schidstorm/ffmpeg-jobs/worker/worker"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := cobra.Command{
		RunE: runApplication,
	}

	rootCmd.PersistentFlags().String("apiServer", "http://localhost:8080", "Url of the API server")
	rootCmd.PersistentFlags().String("logLevel", logrus.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

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
	return worker.NewLoop(cfg).Run(ApplicationContext())
}

func parseConfig(cmd *cobra.Command) (*config.Config, error) {
	apiServerString, err := cmd.PersistentFlags().GetString("apiServer")
	if err != nil {
		return nil, err
	}

	return &config.Config{ApiServerUrl: apiServerString}, nil
}

package cli

import (
	"github.com/schidstorm/ffmpeg-jobs/synchronizer/config"
	"github.com/schidstorm/ffmpeg-jobs/synchronizer/synchronizer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := cobra.Command{
		RunE: runApplication,
	}

	rootCmd.PersistentFlags().String("inputFileDirectory", "/input", "Directory where to look for optimizations")
	rootCmd.PersistentFlags().String("outputFileDirectory", "/output", "Directory where to put the optimizations")
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
	return synchronizer.Run(cfg, ApplicationContext())
}

func parseConfig(cmd *cobra.Command) (*config.Config, error) {
	inputFileDirectory, err := cmd.PersistentFlags().GetString("inputFileDirectory")
	if err != nil {
		return nil, err
	}

	outputFileDirectory, err := cmd.PersistentFlags().GetString("outputFileDirectory")
	if err != nil {
		return nil, err
	}

	apiServerUrl, err := cmd.PersistentFlags().GetString("apiServer")
	if err != nil {
		return nil, err
	}

	return &config.Config{
		InputFileDirectory: inputFileDirectory,
		ApiServerUrl: apiServerUrl,
		OutputFileDirectory: outputFileDirectory,
	}, nil
}

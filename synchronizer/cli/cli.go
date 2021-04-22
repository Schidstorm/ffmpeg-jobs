package cli

import (
	"github.com/schidstorm/ffmpeg-jobs/synchronizer/config"
	"github.com/schidstorm/ffmpeg-jobs/synchronizer/synchronizer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path"
	"time"
)

func Run() {
	rootCmd := cobra.Command{
		RunE: runApplication,
	}

	wd, err := os.Getwd()
	if err != nil {
		logrus.Warn(err)
		wd = "."
	}
	inputDirPath := path.Join(wd, "test/input")
	outputDirPath := path.Join(wd, "test/output")

	rootCmd.PersistentFlags().String("inputFileDirectory", inputDirPath, "Directory where to look for optimizations")
	rootCmd.PersistentFlags().String("outputFileDirectory", outputDirPath, "Directory where to put the optimizations")
	rootCmd.PersistentFlags().String("apiServer", "http://localhost:8081", "Url of the API server")
	rootCmd.PersistentFlags().Duration("watcherLoopWait", 5*time.Second, "Time to wait between file syncs")
	rootCmd.PersistentFlags().String("logLevel", logrus.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

	err = rootCmd.Execute()
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

	apiServer, err := cmd.PersistentFlags().GetString("apiServer")
	if err != nil {
		return nil, err
	}

	watcherLoopWait, err := cmd.PersistentFlags().GetDuration("watcherLoopWait")
	if err != nil {
		return nil, err
	}

	return &config.Config{
		InputFileDirectory:  inputFileDirectory,
		ApiServerUrl:        apiServer,
		OutputFileDirectory: outputFileDirectory,
		WatcherLoopWait:     watcherLoopWait,
	}, nil
}

package config

import "time"

type Config struct {
	InputFileDirectory  string
	OutputFileDirectory string
	ApiServerUrl        string
	WatcherLoopWait     time.Duration
}

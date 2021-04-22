package synchronizer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/schidstorm/ffmpeg-jobs/api/domain"
	"github.com/schidstorm/ffmpeg-jobs/synchronizer/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type inputFileState struct {
	InputFilePath    string
	OutputFilePath   string
	OutputFileExists bool
}

func RunWatcher(config *config.Config, applicationContext context.Context) error {
	for {
		select {
		case <-applicationContext.Done():
			return nil
		case <-time.After(config.WatcherLoopWait):
		}

		apiJobs := getApiJobs(config.ApiServerUrl)
		fileStates := scanDirectory(config.InputFileDirectory, config.OutputFileDirectory)

		for _, fileState := range fileStates {
			if _, ok := apiJobs[fileState.InputFilePath]; ok {
				continue
			}

			if !fileState.OutputFileExists {
				addJobForFile(config.ApiServerUrl, config.OutputFileDirectory, fileState.InputFilePath)
			}
		}
	}
}

func scanDirectory(dirPath string, outputDir string) []inputFileState {
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logrus.Error(err)
		return []inputFileState{}
	}

	var fileStates []inputFileState
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && !strings.HasPrefix(fileInfo.Name(), ".") {
			inputFilePath := path.Join(dirPath, fileInfo.Name())
			outputFilePath := outputFileFromInputFilePath(outputDir, inputFilePath)
			_, err := os.Stat(outputFilePath)
			outputFileExists := err == nil
			fileStates = append(fileStates, inputFileState{
				InputFilePath:    inputFilePath,
				OutputFilePath:   outputFilePath,
				OutputFileExists: outputFileExists,
			})
		}
	}

	return fileStates
}

func getApiJobs(apiServerUrl string) map[string]domain.Job {
	resp, err := http.Get(fmt.Sprintf("%s/job", apiServerUrl))
	if err != nil {
		logrus.Error(err)
		return map[string]domain.Job{}
	}
	defer resp.Body.Close()

	responseJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return map[string]domain.Job{}
	}

	var jobResults struct {
		Data []domain.Job
	}

	err = json.Unmarshal(responseJson, &jobResults)
	if err != nil {
		logrus.Error(err)
		return map[string]domain.Job{}
	}

	jobsMap := map[string]domain.Job{}
	for _, apiJob := range jobResults.Data {
		jobsMap[apiJob.InputFile] = apiJob
	}
	return jobsMap
}

func addJobForFile(apiServerUrl, outputDir, inputPath string) {
	outputPath := outputFileFromInputFilePath(outputDir, inputPath)
	postUrl := fmt.Sprintf("%s/job", apiServerUrl)
	postDataMap := map[string]interface{}{
		"InputFile":  inputPath,
		"OutputFile": outputPath,
	}

	postData, _ := json.Marshal(postDataMap)
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(postData))

	if err != nil {
		logrus.Error(err)
	} else {
		defer resp.Body.Close()
		logrus.Infof("added job for %s -> %s", inputPath, outputPath)
	}
}

func outputFileFromInputFilePath(outputDir, inputFilePath string) string {
	ext := path.Ext(inputFilePath)
	return path.Join(outputDir, path.Base(strings.TrimSuffix(inputFilePath, ext)+".mp4"))
}

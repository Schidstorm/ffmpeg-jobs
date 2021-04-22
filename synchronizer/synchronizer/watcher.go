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
	"path"
	"strings"
	"time"
)

func RunWatcher(config *config.Config, applicationContext context.Context) error  {
	for {
		select {
		case <-applicationContext.Done():
			return nil
		case <- time.After(1 * time.Second):
		}


	}
}

func scanDirectory(dirPath string) []string {
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logrus.Error(err)
		return []string{}
	}

	var filePaths []string
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && !strings.HasPrefix(fileInfo.Name(), ".") {
			filePaths = append(filePaths, path.Join(dirPath, fileInfo.Name()))
		}
	}

	return filePaths
}

func apiFilesExist(apiServerUrl string, paths []string) []string {
	resp, err := http.Get(fmt.Sprintf("%s/job"))
	if err != nil {
		logrus.Error(err)
		return []string{}
	}
	defer resp.Body.Close()

	responseJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return []string{}
	}

	var job domain.Job

}

func addJobForFile(apiServerUrl, outputDir, inputPath string) {
	ext := path.Ext(inputPath)
	outputPath := path.Join(outputDir, strings.TrimSuffix(path.Base(inputPath), ext) + ".mp4")
	postUrl := fmt.Sprintf("%s/job", apiServerUrl)
	postDataMap := map[string]string{
		"InputFile": inputPath,
		"OutputFile": outputPath,
	}

	postData, _ := json.Marshal(postDataMap)
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(postData))
	defer resp.Body.Close()

	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("added job for %s -> %s", inputPath, outputPath)
	}
}
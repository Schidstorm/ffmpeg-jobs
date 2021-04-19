package worker

import (
	"context"
	"encoding/json"
	"github.com/phayes/freeport"
	"github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"net"
	"regexp"
	"strconv"
)

type ProgressCallback = func(p float64)

type ffmpegProbe struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func RunFfmpegJob(job Job, callback ProgressCallback, ctx context.Context) error {
	probeString, err := ffmpeg.Probe(job.Data.InputFile)
	if err != nil {
		return err
	}

	var probe ffmpegProbe
	err = json.Unmarshal([]byte(probeString), &probe)
	if err != nil {
		return err
	}
	totalDuration, err := strconv.ParseFloat(probe.Format.Duration, 64)
	if err != nil {
		return err
	}
	logrus.Info(totalDuration)

	socketPath := ProgressSocket(totalDuration, callback)

	compiledStream := ffmpeg.Input(job.Data.InputFile).
		Output(job.Data.OutputFile, ffmpeg.KwArgs{"c:v": "libx264", "preset": "veryslow"}).
		GlobalArgs("-progress", socketPath).
		OverWriteOutput().
		Compile()

	err = compiledStream.Start()
	if err != nil {
		return err
	}

	done := make(chan error)
	go func() {
		done <- compiledStream.Wait()
	}()

	select {
	case err = <-done:
		break
	case <-ctx.Done():
		err = compiledStream.Process.Kill()
		break
	}

	if err != nil {
		return err
	}

	callback(1)
	return nil
}

func ProgressSocket(totalDuration float64, callback ProgressCallback) string {
	port, err := freeport.GetFreePort()
	if err != nil {
		panic(err)
	}
	socketPath := "127.0.0.1:" + strconv.Itoa(port)
	l, err := net.Listen("tcp", socketPath)
	if err != nil {
		panic(err)
	}

	go func() {
		re := regexp.MustCompile(`out_time_ms=(\d+)`)
		fd, err := l.Accept()
		defer l.Close()
		if err != nil {
			logrus.Fatal("accept error:", err)
		}
		buf := make([]byte, 16)
		data := ""
		var lastProgress float64
		for {
			_, err := fd.Read(buf)
			if err != nil {
				return
			}
			data += string(buf)
			matches := re.FindAllStringSubmatch(data, -1)
			if len(matches) > 0 && len(matches[len(matches)-1]) > 0 {
				c, _ := strconv.Atoi(matches[len(matches)-1][len(matches[len(matches)-1])-1])
				p := float64(c) / totalDuration / 1000000
				if p > lastProgress {
					callback(p)
					lastProgress = p
				}
			}
		}
	}()

	return "tcp://" + socketPath
}

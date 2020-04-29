package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Extract(category string) (string, error) {
	basePath := Config.DataBasePath(category)
	extractUrl := Config.DataUrl(category)

	os.MkdirAll(basePath, 0755)
	now := time.Now()
	nowString := fmt.Sprintf("%04d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	fileName := fmt.Sprintf("%s_%s.csv", nowString, category)
	fullFilePath := filepath.Join(basePath, fileName)

	return fullFilePath, download(fullFilePath, extractUrl)
}

func download(extractFilePath string, dataUrl string) error {
	logger.Infof("Download from %s", dataUrl)
	resp, err := http.Get(dataUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(extractFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	logger.Infof("Stream output to %s", extractFilePath)
	_, err = io.Copy(out, resp.Body)
	return err
}

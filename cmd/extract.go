package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

func Extract(category string) (string, error) {
	basePath := viper.GetString(fmt.Sprintf("data.%s.base_path", category))
	extractUrl := viper.GetString(fmt.Sprintf("data.%s.url", category))

	os.MkdirAll(basePath, 0755)
	now := time.Now()
	nowString := fmt.Sprintf("%04d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	fileName := fmt.Sprintf("%s_%s.csv", nowString, category)
	fullFilePath := filepath.Join(basePath, fileName)

	return fullFilePath, download(fullFilePath, extractUrl)
}

func download(extractFilePath string, dataUrl string) error {
	log.Printf("Download from %s", dataUrl)
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

	log.Printf("Stream output to %s", extractFilePath)
	_, err = io.Copy(out, resp.Body)
	return err
}

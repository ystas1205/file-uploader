package main

import (
	"file-uploader/internal/downloader"
	"file-uploader/internal/path"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	URL     string
	Out     string
	Verbose bool
}

func main() {
	cfg := &Config{}

	cmd := os.Args

	if len(cmd) < 2 {
		fmt.Fprintln(os.Stderr, "использование: app -url=https://...")
		os.Exit(1)
	}

	flag.StringVar(&cfg.URL, "url", "", "URL файла для загрузки (обязательно)")
	flag.StringVar(&cfg.Out, "out", "", "Путь для сохранения файла")
	flag.BoolVar(&cfg.Verbose, "v", false, "Выводить подробную информацию")

	flag.Parse()

	if strings.TrimSpace(cfg.URL) == "" {
		fmt.Fprintln(os.Stderr, "обязательный флаг --url не указан")
		os.Exit(1)
	}
	userAgent, errRead := downloader.ReadUserAgent()
	if errRead != nil {
		fmt.Fprintln(os.Stderr, "ошибка", errRead)
		os.Exit(1)
	}

	body, resp, errFetch := downloader.FetchURL(cfg.URL, userAgent)
	if errFetch != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", errFetch)
		os.Exit(1)
	}

	filePath := cfg.Out
	if filePath == "" {
		var errBuild error
		filePath, errBuild = path.BuildFileName(resp, cfg.URL)
		if errBuild != nil {
			fmt.Fprintln(os.Stderr, "ошибка:", errBuild)
			os.Exit(1)
		}
	} else if filepath.Ext(filePath) == "" {
		fmt.Fprintln(os.Stderr, "файл указан без расширения, используйте -out=file.jpg")
		os.Exit(1)
	}

	size, errSave := downloader.SaveStream(body, resp, filePath, cfg.Verbose)
	if errSave != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", errSave)
		os.Exit(1)
	}
	fmt.Printf("файл успешно загружен: %q, размер: %d байт\n", filePath, size)

}

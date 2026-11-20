package main

import (
	"file-uploader/internal/downloader"
	"file-uploader/internal/path"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	cmd, _, ok := strings.Cut(strings.Join(os.Args[1:], " "), "=")
	if len(os.Args) < 2 {
		fmt.Println("использование: app -url=https://...")
		os.Exit(1)
	}
	if !ok {
		fmt.Println("должен быть знак равно")
		os.Exit(1)
	}

	switch cmd {
	case "--url":
		url := flag.String("url", "", "url")
		flag.Parse()
		if strings.TrimSpace(*url) == "" {
			fmt.Println("Обязательный флаг --url не указан")
			os.Exit(1)

		}

		body, resp, errFetch := downloader.FetchURL(*url)
		if errFetch != nil {
			fmt.Println("ошибка:", errFetch)
			os.Exit(1)
		}

		path := path.NamePath(resp, *url)

		size, errSave := downloader.SaveStream(body, resp, path)
		if errSave != nil {
			fmt.Println("ошибка")
			os.Exit(1)
		}

		fmt.Println("файл успешно загружен размер", size)

	default:
		fmt.Println("неизвестная команда:", cmd)
		fmt.Println("доступные команды: --url")
		os.Exit(1)

	}

}

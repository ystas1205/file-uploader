package main

import (
	"file-uploader/internal/downloader"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	url := flag.String("url", "", "url")
	flag.Parse()
	if strings.TrimSpace(*url) == "" {
		fmt.Println("Обязательный флаг -url не указан")
		os.Exit(1)

	}

	body, extension, base, errFetch := downloader.FetchURL(*url)
	if errFetch != nil {
		fmt.Println("ошибка:", errFetch)
		os.Exit(1)
	}

	// errCreate := downloader.CreateDir()
	// if errCreate != nil {
	// 	fmt.Println("ошибка")
	// 	os.Exit(1)
	// }

	size, errSave := downloader.SaveStream(body, extension, base)
	if errSave != nil {
		fmt.Println("ошибка")
		os.Exit(1)
	}

	fmt.Println("файл успешно загружен размер", size)

}

package main

import (
	"file-uploader/internal/downloader"
	"file-uploader/internal/path"
	// "file-uploader/internal/progress"
	

	// "file-uploader/internal/progress"
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
	fmt.Println(len(cmd))

	if len(cmd) < 2 {
		fmt.Println("использование: app -url=https://...")
		os.Exit(1)
	}

	flag.StringVar(&cfg.URL, "url", "", "URL файла для загрузки (обязательно)")
	flag.StringVar(&cfg.Out, "out", "", "Путь для сохранения файла")
	flag.BoolVar(&cfg.Verbose, "v", false, "Выводить подробную информацию")

	flag.Parse()

	if strings.TrimSpace(cfg.URL) == "" {
		fmt.Println("обязательный флаг --url не указан")
		os.Exit(1)

	}

	body, resp, errFetch := downloader.FetchURL(cfg.URL)
	if errFetch != nil {
		fmt.Println("ошибка:", errFetch)
		os.Exit(1)
	}



	if strings.TrimSpace(cfg.Out) != "" {

		if filepath.Ext(cfg.Out) == "" {
			fmt.Println("файл указан без расширения, используйте -out=file.jpg")
			os.Exit(1)

		}

		if len(cmd) >4 {
			fmt.Println("использование: app -url=https://... --out=file.text")
			os.Exit(1)
		}
		
			

		size, errSave := downloader.SaveStream(body, resp, cfg.Out,cfg.Verbose)
		if errSave != nil {
			fmt.Println("ошибка")
			os.Exit(1)
		}

		
		fmt.Printf("файл успешно загружен: %q, размер: %d байт\n", cfg.Out, size)
		
		
	} else {

		if len(cmd) > 3 {
			fmt.Println("использование: app -url=https://...")
			os.Exit(1)
		}

		filePath := path.NamePath(resp, cfg.URL)

		size, errSave := downloader.SaveStream(body, resp, filePath,cfg.Verbose)
		if errSave != nil {
			fmt.Println("ошибка")
			os.Exit(1)
		}

		fmt.Printf("файл успешно загружен: %q, размер: %d байт\n", filePath, size)
	}




	// if cfg.Verbose{
	// 	if len(cmd) < 3 {
	// 		fmt.Println("использование: app -v")
	// 		os.Exit(1)
	// 	}

	// 	_, errSave := downloader.SaveStream(body, resp, cfg.Out,true)
	// 	if errSave != nil {
	// 		fmt.Println("ошибка")
	// 		os.Exit(1)
	// 	}
		


	// }
	
}



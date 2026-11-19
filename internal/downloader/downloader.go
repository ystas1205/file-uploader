package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// func CreateDir() error {
// 	err := os.Mkdir("folder-file", 0755)
// 	if err != nil {
// 		return fmt.Errorf("не удалось создать директорию")
// 	}
// 	return nil
// }

func SaveStream(body io.ReadCloser, extension string, base string) (int64, error) {
	defer body.Close()

	file := fmt.Sprintf("%s.%s", base, extension)
	
	output, err := os.Create(file)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания файла %w", err)

	}
	defer output.Close()

	size, err := io.Copy(output, body)
	if err != nil {
		panic(err)
	}
	return size, nil

}

func FetchURL(url string) (io.ReadCloser, string, string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, "", "", fmt.Errorf("неуспешный ответ от сервера: %s", response.Status)
	}
	extension := strings.Split(response.Header.Get("Content-Type"), "/")[1]
	base := filepath.Base(url)

	return response.Body, extension, base, nil

}

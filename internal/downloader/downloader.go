package downloader

import (
	"file-uploader/internal/progress"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SaveStream(body io.ReadCloser, resp *http.Response, path string) (int64, error) {
	defer body.Close()
	output, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания файла %w", err)
	}
	defer output.Close()
	size, err := io.Copy(io.MultiWriter(output, progress.ProgressBar(resp)), body)
	if err != nil {
		panic(err)
	}
	return size, nil

}

func FetchURL(url string) (io.ReadCloser, *http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("создания запроса %q: %w", url, err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("выполнения запроса %q: %w", req.URL, err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("неуспешный ответ от сервера: %s", response.Status)
	}
	return response.Body, response, nil

}

package downloader

import (
	"file-uploader/internal/progress"
	"fmt"
	"io"
	"net/http"
	"os"
)
// SaveStream записывает файл из тела ответа и возвращает
// количество записанных байт и ошибку.
func SaveStream(
	body io.ReadCloser,
	resp *http.Response,
	path string,
	bar bool) (int64, error) {

	defer body.Close()
	output, errCreate := os.Create(path)
	if errCreate != nil {
		return 0, fmt.Errorf("создать файл %q: %w", path, errCreate)
	}
	defer output.Close()

	if !bar {
		size, err := io.Copy(output, body)
		if err != nil {
			return size, fmt.Errorf("скопировать данные в файл: %w", err)
		}
		return size, nil
	} else {

		size, err := io.Copy(io.MultiWriter(output, progress.ProgressBar(resp)), body)
		if err != nil {
			panic(err)
		}
		return size, nil
	}
}

// FetchURL делает GET-запрос по указанному URL и возвращает
// тело ответа и *http.Response для дальнейшей обработки.
func FetchURL(url string, userAgent string) (io.ReadCloser, *http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("создать запрос %q: %w", url, err)
	}

	req.Header.Set("User-Agent", userAgent)
	response, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("выполнить запрос %q: %w", req.URL, err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("неожиданный статус ответа: %q", response.Status)
	}
	return response.Body, response, nil

}

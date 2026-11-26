package progress

import (
	"github.com/schollz/progressbar/v3"
	"net/http"
)

// ProgressBar создаёт прогресс-бар для скачивания файла
// по размеру тела ответа resp.ContentLength.
func ProgressBar(resp *http.Response) *progressbar.ProgressBar {
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	return bar
}

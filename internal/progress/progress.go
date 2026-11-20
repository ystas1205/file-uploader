package progress

import (
	"github.com/schollz/progressbar/v3"
	"net/http"
)

func ProgressBar(resp *http.Response) *progressbar.ProgressBar {
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	return bar
}

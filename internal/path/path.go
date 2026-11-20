package path

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func NamePath(resp *http.Response, url string) string {

	extension := strings.Split(resp.Header.Get("Content-Type"), "/")[1]
	base := filepath.Base(url)

	return fmt.Sprintf("%s.%s", base, extension)

}

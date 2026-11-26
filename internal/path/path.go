package path

import (
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"path"
)

// extFromMediaType возвращает расширение файла (с точкой) для MIME-типа.
// Если тип неизвестен, возвращает пустую строку.
func extFromMediaType(mediaType string) string {
	switch mediaType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	case "image/gif":
		return ".gif"
	case "video/mp4":
		return ".mp4"
	case "application/pdf":
		return ".pdf"
	case "text/plain":
		return ".txt"
	default:
		return ""
	}
}

// BuildFileName строит имя файла для сохранения по URL и заголовкам ответа.
// Если в URL уже есть расширение, оно используется как есть. Иначе
// расширение определяется по Content-Type из resp. Возвращает ошибку,
// если URL не удалось разобрать, имя получить не удалось или
// Content-Type неизвестен.
func BuildFileName(resp *http.Response, rawURL string) (string, error) {
	u, err := url.Parse(rawURL)

	if err != nil {
		return "", fmt.Errorf("разобрать URL %q: %w", rawURL, err)
	}

	base := path.Base(u.Path)

	if base == "." || base == "/" || base == "" {
		return "", fmt.Errorf("не удалось получить имя из URL %q", rawURL)
	}

	if path.Ext(base) != "" {
		return base, nil
	}

	ct := resp.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(ct)
	if err != nil {
		return "", fmt.Errorf("разобрать Content-Type %q: %w", ct, err)
	}

	ext := extFromMediaType(mediaType)
	if ext == "" {
		return "", fmt.Errorf("неизвестный Content-Type %q", mediaType)
	}

	return base + ext, nil
}

package gate

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var extensionToContentType = map[string]string{
	".html": "text/html; charset=utf-8",
	".css":  "text/css; charset=utf-8",
	".js":   "application/javascript",
	".xml":  "text/xml; charset=utf-8",
	".jpg":  "image/jpeg",
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，允许所有来源的跨域请求
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

	wd, _ := os.Getwd()
	path := fmt.Sprintf("%s%s", wd, r.URL.Path)

	f, err := os.Open(path)
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}

	if d.IsDir() {
		Error(w, toHTTPError(err))
		return
	}

	data, err := io.ReadAll(f)
	if err != nil {
		Error(w, toHTTPError(err))
		return
	}

	ext := filepath.Ext(path)
	if contentType := extensionToContentType[ext]; contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	w.Header().Set("Content-Length", strconv.FormatInt(d.Size(), 10))
	w.Write(data)
}

func toHTTPError(err error) int {
	if os.IsNotExist(err) {
		return http.StatusNotFound
	}
	if os.IsPermission(err) {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}

func Error(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

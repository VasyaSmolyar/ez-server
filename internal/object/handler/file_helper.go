package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("get file type: %s", err)
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func FileUploadToLocal(fname string, r *http.Request) (string, string, error) {
	file, handler, err := r.FormFile(fname)
	if err != nil {
		return "", "", fmt.Errorf("file from request: %s", err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", "", fmt.Errorf("file tmp: %s", err)
	}

	io.Copy(f, file)

	f, err = os.OpenFile(handler.Filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", "", fmt.Errorf("file tmp read: %s", err)
	}

	content, err := getFileContentType(f)
	if err != nil {
		return "", "", err
	}

	defer f.Close()
	return handler.Filename, content, nil
}

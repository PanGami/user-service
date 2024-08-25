package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	DateFormatRFC3339 = time.RFC3339
	StatusInActive    = 0
	StatusActive      = 1
	StatusForgot      = 2
	StatusDraft       = "10"
	PaidStatusPro     = "pro"
	PaidStatusBasic   = "basic"
	SubscribePending  = "pending"
	SubscribePaid     = "paid"
	SubscribeExpired  = "expired"
	SubscribeFailed   = "failed"
	DefaultPage       = 1
	DefaultCount      = 15
	limitPage         = 9999
)

func ParseStringToTime(dateTimeStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimePointer(t time.Time) *time.Time {
	return &t
}

func ConvertBytesToMultipartFile(fileBytes []byte, fileName string) (*multipart.FileHeader, error) {
	// Create a buffer to write the multipart form data
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Create a form file with the provided filename
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	// Write the byte slice to the form file
	_, err = io.Copy(part, bytes.NewReader(fileBytes))
	if err != nil {
		return nil, err
	}

	// Close the writer to finalize the multipart form data
	writer.Close()

	// Parse the multipart form
	req := &http.Request{
		Body:   io.NopCloser(buf),
		Header: make(http.Header),
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	err = req.ParseMultipartForm(int64(buf.Len()))
	if err != nil {
		return nil, err
	}

	// Extract the multipart file header
	file, header, err := req.FormFile("file")
	if err != nil {
		return nil, err
	}
	file.Close() // Close the file as we just need the header

	return header, nil
}

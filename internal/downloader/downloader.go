package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func DownloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	size := resp.ContentLength

	bar := progressbar.DefaultBytes(
		size,
		"Downloading",
	)

	_, err = io.Copy(
		io.MultiWriter(out, bar),
		resp.Body,
	)

	if err != nil {
		return err
	}

	fmt.Println("\nDownload complete")

	return nil

}

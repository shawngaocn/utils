package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil && err != io.EOF {
			return nil, err
		}
		return data, nil
	}
	return nil, err
}

func HttpPostFormURLEncoded(url string, params string) ([]byte, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(params))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpDownload(url string, filePath string) (int64, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}

	written, err := io.Copy(file, res.Body)
	if err != nil {
		file.Close()
		return 0, err
	}

	err = file.Close()
	if err != nil {
		return written, err
	}

	return written, nil
}

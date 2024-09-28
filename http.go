package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("User-Agent", "cli "+runtime.GOOS+"/"+runtime.GOARCH+" version "+version)
}

func do(method, path string, body []byte) (int, []byte) {
	client := &http.Client{}
	bs := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, config.Endpoint+path, bs)
	if err != nil {
		fmt.Fprintln(os.Stderr,
			"error: please try again later or contact support")
		os.Exit(1)
	}
	setHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr,
			"error: unable to connect, check your internet connection or try again later")
		os.Exit(1)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr,
			"error: please try again later or contact support")
		os.Exit(1)
	}
	if res.StatusCode != 200 {
		fmt.Fprintf(os.Stderr,
			"error: %s\n", string(b))
		os.Exit(1)
	}
	return res.StatusCode, b
}

func get(path string) (int, []byte) {
	return do("GET", path, nil)
}

func delete(path string) (int, []byte) {
	return do("DELETE", path, nil)
}

func put(path string, body []byte) (int, []byte) {
	return do("PUT", path, body)
}

func post(path string, body []byte) (int, []byte) {
	return do("POST", path, body)
}

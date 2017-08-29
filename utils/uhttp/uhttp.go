package uhttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Request request
type Request struct {
	Method string
	URL    string
	Body   []byte
	Header http.Header
}

// Cmd fetch command
func Cmd(args Request) ([]byte, error) {
	client := &http.Client{}
	// set request
	req, err := http.NewRequest(args.Method, args.URL,
		bytes.NewReader(
			args.Body,
		),
	)
	if err != nil {
		return nil, nil
	}
	req.Close = true
	req.Header = args.Header
	// get response
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func PostForm(url string, data url.Values) ([]byte, error) {
	println(url)
	resp, err := http.PostForm(url, data)
	if err != nil {
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

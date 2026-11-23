package api

import (
	"Glowl/internal/common/entities"
	"bytes"
	"io"
	"net/http"
)

func createUrl(url entities.ApiUrl, port *entities.ApiPort) string {
	var res string
	if port == nil {
		res = "http://" + string(url)
	} else {
		res = "http://" + string(url) + ":" + string(*port)
	}
	return res
}

func newOption(Method entities.ApiMethod, Url string, Body entities.ApiBody) fetchOption {
	return fetchOption{
		Method: Method,
		Url:    Url,
		Body:   Body,
	}
}

func fetch(option fetchOption) (io.ReadCloser, error) {
	req, err := http.NewRequest(string(option.Method), option.Url, bytes.NewReader(option.Body))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

package api

import (
	"Glowl/internal/common/entities"
	"bytes"
	"context"
	"io"
	"net"
	"net/http"

	"github.com/coder/websocket"
)

func createUrl(protocol entities.ApiProtocol, url entities.ApiUrl, port *entities.ApiPort) string {
	var res string
	if port == nil {
		res = string(protocol) + "://" + string(url)
	} else {
		res = string(protocol) + "://" + string(url) + ":" + string(*port)
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

func wsConnect(ctx context.Context, url string) (*websocket.Conn, error) {
	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(1 << 20)
	return c, nil
}

func createHttpServer(network string, address string) (*server, error) {
	mux := http.NewServeMux()
	adr, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	handler := &http.Server{Handler: mux}
	return &server{
		mux:  mux,
		adr:  adr,
		http: handler,
	}, nil
}

func NewServer(network string, address string) (*server, error) {
	server, err := constructorHttpServer(network, address)
	if err != nil {
		return nil, err
	}
	service := NewHttpService()
	RegisterRoutes(server.mux, service)
	RegisterWsRoutes(server.mux, service)
	return server, nil
}

func (s *server) Start() error {
	return s.http.Serve(s.adr)
}

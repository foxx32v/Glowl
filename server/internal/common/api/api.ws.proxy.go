package api

import (
	"context"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func wsProxy(dial func(ctx context.Context) (*websocket.Conn, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
		if err != nil {
			return
		}
		defer client.Close(websocket.StatusInternalError, "done")
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		server, err := dial(ctx)
		if err != nil {
			client.Close(websocket.StatusInternalError, err.Error())
			return
		}
		defer server.Close(websocket.StatusNormalClosure, "done")
		pipe(r.Context(), client, server)
	}
}

func pipe(ctx context.Context, client, server *websocket.Conn) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	errCh := make(chan error, 2)
	go forward(ctx, client, server, errCh)
	go forward(ctx, server, client, errCh)
	<-errCh
}

func forward(ctx context.Context, src, dst *websocket.Conn, errCh chan<- error) {
	for {
		typ, data, err := src.Read(ctx)
		if err != nil {
			errCh <- err
			return
		}
		if err := dst.Write(ctx, typ, data); err != nil {
			errCh <- err
			return
		}
	}
}

func RegisterWsRoutes(mux *http.ServeMux, svc *Service) {
	mux.HandleFunc("GET /ws/llm", wsProxy(svc.LlmWsConnect))
	mux.HandleFunc("GET /ws/asr", wsProxy(svc.AsrWsConnect))
	mux.HandleFunc("GET /ws/tts", wsProxy(svc.TtsWsConnect))

	mux.HandleFunc("GET /ws/llm/generate", wsProxy(svc.LlmGenerateWsConnect))
	mux.HandleFunc("GET /ws/llm/chat", wsProxy(svc.LlmChatWsConnect))
	mux.HandleFunc("GET /ws/tts/process", wsProxy(svc.TtsProcessWsConnect))
}

package api

import (
	"context"

	"github.com/coder/websocket"
)

/*-----------------------------------------------------\
|				Standard connects (WS)					|
\-----------------------------------------------------*/

func (s *Service) LlmWsConnect(ctx context.Context) (*websocket.Conn, error) {
	return wsConnect(ctx, s.opts.WS.LLM.NIL.Url)
}

func (s *Service) AsrWsConnect(ctx context.Context) (*websocket.Conn, error) {
	return wsConnect(ctx, s.opts.WS.ASR.NIL.Url)
}

func (s *Service) TtsWsConnect(ctx context.Context) (*websocket.Conn, error) {
	return wsConnect(ctx, s.opts.WS.TTS.NIL.Url)
}

/*-----------------------------------------------------\
|				Custom connects (WS)					|
\-----------------------------------------------------*/

func (s *Service) LlmGenerateWsConnect(ctx context.Context) (*websocket.Conn, error) {
	return wsConnect(ctx, s.opts.WS.LLMGenerate.NIL.Url)
}

func (s *Service) LlmChatWsConnect(ctx context.Context) (*websocket.Conn, error) {
	return wsConnect(ctx, s.opts.WS.LLMChat.NIL.Url)
}

func (s *Service) TtsProcessWsConnect(ctx context.Context) (*websocket.Conn, error) {
	return wsConnect(ctx, s.opts.WS.TTSProcess.NIL.Url)
}

/*-----------------------------------------------------\
|					Custom connect						|
\-----------------------------------------------------*/

func (s *Service) CustomWsConnect(ctx context.Context, url string) (*websocket.Conn, error) {
	return wsConnect(ctx, url)
}

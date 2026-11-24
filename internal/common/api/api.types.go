package api

import (
	"Glowl/internal/common/entities"
	"io"
)

type Urls struct {
	LLM              string
	ASR              string
	TTS              string
	LLMGenerate      string
	LLMChat          string
	TTSProcess       string
	LLMHttps         string
	ASRHttps         string
	TTSHttps         string
	LLMGenerateHttps string
	LLMChatHttps     string
	TTSProcessHttps  string
	LLMWs            string
	ASRWs            string
	TTSWs            string
	LLMGenerateWs    string
	LLMChatWs        string
	TTSProcessWs     string
}

type fetchResponse struct {
	Body io.ReadCloser
	Err  error
}

type fetchOption struct {
	Method entities.ApiMethod
	Url    string
	Body   entities.ApiBody
}

type endpointOptions[T any] struct {
	GET     T
	POST    T
	PUT     T
	PATCH   T
	DELETE  T
	OPTIONS T
	NIL     T
}
type schemaOptions[T any] struct {
	LLM         endpointOptions[T]
	ASR         endpointOptions[T]
	TTS         endpointOptions[T]
	LLMGenerate endpointOptions[T]
	LLMChat     endpointOptions[T]
	TTSProcess  endpointOptions[T]
}

type fetchType[T any] struct {
	HTTP  schemaOptions[T]
	HTTPS schemaOptions[T]
	WS    schemaOptions[T]
}

type ServiceRes = fetchResponse
type Service struct {
	opts fetchType[fetchOption]
}

package api

import (
	"Glowl/internal/common/entities"
	"io"
)

type Urls struct {
	LLM         string
	ASR         string
	TTS         string
	LLMGenerate string
	LLMChat     string
	TTSProcess  string
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

type fetchType[T any] struct {
	LLM_GET     T
	LLM_POST    T
	LLM_PUT     T
	LLM_PATCH   T
	LLM_DELETE  T
	LLM_OPTIONS T

	ASR_GET     T
	ASR_POST    T
	ASR_PUT     T
	ASR_PATCH   T
	ASR_DELETE  T
	ASR_OPTIONS T

	TTS_GET     T
	TTS_POST    T
	TTS_PUT     T
	TTS_PATCH   T
	TTS_DELETE  T
	TTS_OPTIONS T

	LLM_GENERATE_GET     T
	LLM_GENERATE_POST    T
	LLM_GENERATE_PUT     T
	LLM_GENERATE_PATCH   T
	LLM_GENERATE_DELETE  T
	LLM_GENERATE_OPTIONS T

	LLM_CHAT_GET     T
	LLM_CHAT_POST    T
	LLM_CHAT_PUT     T
	LLM_CHAT_PATCH   T
	LLM_CHAT_DELETE  T
	LLM_CHAT_OPTIONS T

	TTS_PROCESS_GET     T
	TTS_PROCESS_POST    T
	TTS_PROCESS_PUT     T
	TTS_PROCESS_PATCH   T
	TTS_PROCESS_DELETE  T
	TTS_PROCESS_OPTIONS T
}

type ServiceRes = fetchResponse
type Service struct {
	opts fetchType[fetchOption]
}

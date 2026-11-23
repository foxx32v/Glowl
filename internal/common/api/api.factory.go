package api

import (
	"Glowl/internal/common/entities"
)

func urlFactory() Urls {
	urls := entities.NewApiUrl()
	ports := entities.NewApiPort()
	llm := createUrl(urls.Localhost, &ports.Llm)
	tts := createUrl(urls.Localhost, &ports.Tts)
	asr := createUrl(urls.Localhost, &ports.Asr)
	return Urls{
		LLM:         llm,
		ASR:         asr,
		TTS:         tts,
		LLMGenerate: llm + "/api/generate",
		LLMChat:     llm + "/api/chat",
		TTSProcess:  tts + "/process",
	}
}

func optionFactory() fetchType[fetchOption] {
	urls := urlFactory()
	methods := entities.NewApiMethods()
	return fetchType[fetchOption]{
		LLM_GET:     newOption(methods.GET, urls.LLM, nil),
		LLM_POST:    newOption(methods.POST, urls.LLM, nil),
		LLM_PUT:     newOption(methods.PUT, urls.LLM, nil),
		LLM_PATCH:   newOption(methods.PATCH, urls.LLM, nil),
		LLM_DELETE:  newOption(methods.DELETE, urls.LLM, nil),
		LLM_OPTIONS: newOption(methods.OPTIONS, urls.LLM, nil),

		ASR_GET:     newOption(methods.GET, urls.ASR, nil),
		ASR_POST:    newOption(methods.POST, urls.ASR, nil),
		ASR_PUT:     newOption(methods.PUT, urls.ASR, nil),
		ASR_PATCH:   newOption(methods.PATCH, urls.ASR, nil),
		ASR_DELETE:  newOption(methods.DELETE, urls.ASR, nil),
		ASR_OPTIONS: newOption(methods.OPTIONS, urls.ASR, nil),

		TTS_GET:     newOption(methods.GET, urls.TTS, nil),
		TTS_POST:    newOption(methods.POST, urls.TTS, nil),
		TTS_PUT:     newOption(methods.PUT, urls.TTS, nil),
		TTS_PATCH:   newOption(methods.PATCH, urls.TTS, nil),
		TTS_DELETE:  newOption(methods.DELETE, urls.TTS, nil),
		TTS_OPTIONS: newOption(methods.OPTIONS, urls.TTS, nil),

		LLM_GENERATE_GET:     newOption(methods.GET, urls.LLMGenerate, nil),
		LLM_GENERATE_POST:    newOption(methods.POST, urls.LLMGenerate, nil),
		LLM_GENERATE_PUT:     newOption(methods.PUT, urls.LLMGenerate, nil),
		LLM_GENERATE_PATCH:   newOption(methods.PATCH, urls.LLMGenerate, nil),
		LLM_GENERATE_DELETE:  newOption(methods.DELETE, urls.LLMGenerate, nil),
		LLM_GENERATE_OPTIONS: newOption(methods.OPTIONS, urls.LLMGenerate, nil),

		LLM_CHAT_GET:     newOption(methods.GET, urls.LLMChat, nil),
		LLM_CHAT_POST:    newOption(methods.POST, urls.LLMChat, nil),
		LLM_CHAT_PUT:     newOption(methods.PUT, urls.LLMChat, nil),
		LLM_CHAT_PATCH:   newOption(methods.PATCH, urls.LLMChat, nil),
		LLM_CHAT_DELETE:  newOption(methods.DELETE, urls.LLMChat, nil),
		LLM_CHAT_OPTIONS: newOption(methods.OPTIONS, urls.LLMChat, nil),

		TTS_PROCESS_GET:     newOption(methods.GET, urls.TTSProcess, nil),
		TTS_PROCESS_POST:    newOption(methods.POST, urls.TTSProcess, nil),
		TTS_PROCESS_PUT:     newOption(methods.PUT, urls.TTSProcess, nil),
		TTS_PROCESS_PATCH:   newOption(methods.PATCH, urls.TTSProcess, nil),
		TTS_PROCESS_DELETE:  newOption(methods.DELETE, urls.TTSProcess, nil),
		TTS_PROCESS_OPTIONS: newOption(methods.OPTIONS, urls.TTSProcess, nil),
	}
}

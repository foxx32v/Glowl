package api

import (
	"Glowl/internal/common/entities"
)

func urlFactory() Urls {
	protocols := entities.NewApiProtocol()
	urls := entities.NewApiUrl()
	ports := entities.NewApiPort()
	llm := createUrl(protocols.Http, urls.Localhost, &ports.Llm)
	tts := createUrl(protocols.Http, urls.Localhost, &ports.Tts)
	asr := createUrl(protocols.Http, urls.Localhost, &ports.Asr)
	llmHttps := createUrl(protocols.Https, urls.Localhost, &ports.Llm)
	ttsHttps := createUrl(protocols.Https, urls.Localhost, &ports.Tts)
	asrHttps := createUrl(protocols.Https, urls.Localhost, &ports.Asr)
	llmWs := createUrl(protocols.Ws, urls.Localhost, &ports.Llm)
	ttsWs := createUrl(protocols.Ws, urls.Localhost, &ports.Tts)
	asrWs := createUrl(protocols.Ws, urls.Localhost, &ports.Asr)
	return Urls{
		LLM:              llm,
		ASR:              asr,
		TTS:              tts,
		LLMGenerate:      llm + "/api/generate",
		LLMChat:          llm + "/api/chat",
		TTSProcess:       tts + "/process",
		LLMHttps:         llmHttps,
		ASRHttps:         asrHttps,
		TTSHttps:         ttsHttps,
		LLMGenerateHttps: llmHttps + "/api/generate",
		LLMChatHttps:     llmHttps + "/api/chat",
		TTSProcessHttps:  ttsHttps + "/process",
		LLMWs:            llmWs,
		ASRWs:            asrWs,
		TTSWs:            ttsWs,
		LLMGenerateWs:    llmWs + "/api/generate",
		LLMChatWs:        llmWs + "/api/chat",
		TTSProcessWs:     ttsWs + "/process",
	}
}

func optionFactory() fetchType[fetchOption] {
	urls := urlFactory()
	methods := entities.NewApiMethods()
	return fetchType[fetchOption]{
		HTTP: schemaOptions[fetchOption]{
			LLM: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.LLM, nil),
				POST:    newOption(methods.POST, urls.LLM, nil),
				PUT:     newOption(methods.PUT, urls.LLM, nil),
				PATCH:   newOption(methods.PATCH, urls.LLM, nil),
				DELETE:  newOption(methods.DELETE, urls.LLM, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.LLM, nil),
			},
			ASR: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.ASR, nil),
				POST:    newOption(methods.POST, urls.ASR, nil),
				PUT:     newOption(methods.PUT, urls.ASR, nil),
				PATCH:   newOption(methods.PATCH, urls.ASR, nil),
				DELETE:  newOption(methods.DELETE, urls.ASR, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.ASR, nil),
			},
			TTS: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.TTS, nil),
				POST:    newOption(methods.POST, urls.TTS, nil),
				PUT:     newOption(methods.PUT, urls.TTS, nil),
				PATCH:   newOption(methods.PATCH, urls.TTS, nil),
				DELETE:  newOption(methods.DELETE, urls.TTS, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.TTS, nil),
			},
			LLMGenerate: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.LLMGenerate, nil),
				POST:    newOption(methods.POST, urls.LLMGenerate, nil),
				PUT:     newOption(methods.PUT, urls.LLMGenerate, nil),
				PATCH:   newOption(methods.PATCH, urls.LLMGenerate, nil),
				DELETE:  newOption(methods.DELETE, urls.LLMGenerate, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.LLMGenerate, nil),
			},
			LLMChat: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.LLMChat, nil),
				POST:    newOption(methods.POST, urls.LLMChat, nil),
				PUT:     newOption(methods.PUT, urls.LLMChat, nil),
				PATCH:   newOption(methods.PATCH, urls.LLMChat, nil),
				DELETE:  newOption(methods.DELETE, urls.LLMChat, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.LLMChat, nil),
			},
			TTSProcess: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.TTSProcess, nil),
				POST:    newOption(methods.POST, urls.TTSProcess, nil),
				PUT:     newOption(methods.PUT, urls.TTSProcess, nil),
				PATCH:   newOption(methods.PATCH, urls.TTSProcess, nil),
				DELETE:  newOption(methods.DELETE, urls.TTSProcess, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.TTSProcess, nil),
			},
		},
		HTTPS: schemaOptions[fetchOption]{
			LLM: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.LLMHttps, nil),
				POST:    newOption(methods.POST, urls.LLMHttps, nil),
				PUT:     newOption(methods.PUT, urls.LLMHttps, nil),
				PATCH:   newOption(methods.PATCH, urls.LLMHttps, nil),
				DELETE:  newOption(methods.DELETE, urls.LLMHttps, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.LLMHttps, nil),
			},
			ASR: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.ASRHttps, nil),
				POST:    newOption(methods.POST, urls.ASRHttps, nil),
				PUT:     newOption(methods.PUT, urls.ASRHttps, nil),
				PATCH:   newOption(methods.PATCH, urls.ASRHttps, nil),
				DELETE:  newOption(methods.DELETE, urls.ASRHttps, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.ASRHttps, nil),
			},
			TTS: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.TTSHttps, nil),
				POST:    newOption(methods.POST, urls.TTSHttps, nil),
				PUT:     newOption(methods.PUT, urls.TTSHttps, nil),
				PATCH:   newOption(methods.PATCH, urls.TTSHttps, nil),
				DELETE:  newOption(methods.DELETE, urls.TTSHttps, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.TTSHttps, nil),
			},
			LLMGenerate: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.LLMGenerateHttps, nil),
				POST:    newOption(methods.POST, urls.LLMGenerateHttps, nil),
				PUT:     newOption(methods.PUT, urls.LLMGenerateHttps, nil),
				PATCH:   newOption(methods.PATCH, urls.LLMGenerateHttps, nil),
				DELETE:  newOption(methods.DELETE, urls.LLMGenerateHttps, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.LLMGenerateHttps, nil),
			},
			LLMChat: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.LLMChatHttps, nil),
				POST:    newOption(methods.POST, urls.LLMChatHttps, nil),
				PUT:     newOption(methods.PUT, urls.LLMChatHttps, nil),
				PATCH:   newOption(methods.PATCH, urls.LLMChatHttps, nil),
				DELETE:  newOption(methods.DELETE, urls.LLMChatHttps, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.LLMChatHttps, nil),
			},
			TTSProcess: endpointOptions[fetchOption]{
				GET:     newOption(methods.GET, urls.TTSProcessHttps, nil),
				POST:    newOption(methods.POST, urls.TTSProcessHttps, nil),
				PUT:     newOption(methods.PUT, urls.TTSProcessHttps, nil),
				PATCH:   newOption(methods.PATCH, urls.TTSProcessHttps, nil),
				DELETE:  newOption(methods.DELETE, urls.TTSProcessHttps, nil),
				OPTIONS: newOption(methods.OPTIONS, urls.TTSProcessHttps, nil),
			},
		},
		WS: schemaOptions[fetchOption]{
			LLM: endpointOptions[fetchOption]{
				NIL: newOption(methods.NIL, urls.LLMWs, nil),
			},
			ASR: endpointOptions[fetchOption]{
				NIL: newOption(methods.NIL, urls.ASRWs, nil),
			},
			TTS: endpointOptions[fetchOption]{
				NIL: newOption(methods.NIL, urls.TTSWs, nil),
			},
			LLMGenerate: endpointOptions[fetchOption]{
				NIL: newOption(methods.NIL, urls.LLMGenerateWs, nil),
			},
			LLMChat: endpointOptions[fetchOption]{
				NIL: newOption(methods.NIL, urls.LLMChatWs, nil),
			},
			TTSProcess: endpointOptions[fetchOption]{
				NIL: newOption(methods.NIL, urls.TTSProcessWs, nil),
			},
		},
	}
}

func serverFactory() (*servers, error) {
	url := "127.0.0.1:8000"
	tcpLocalhost, err := constructorHttpServer("tcp", url)
	if err != nil {
		return nil, err
	}
	return &servers{
		TcpLocalhost: tcpLocalhost,
	}, nil
}

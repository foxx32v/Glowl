package api

import (
	"Glowl/internal/common/entities"
)

func TestApi() {
	s := NewHttpService()

	llm := "http://127.0.0.1:11434"
	asr := "http://127.0.0.1:2700"
	tts := "http://127.0.0.1:9898"

	m := entities.NewApiMethods()

	s.opts.HTTP.LLM.GET = newOption(m.GET, llm+"/api/tags", nil)
	s.opts.HTTP.LLM.POST = newOption(m.POST, llm+"/api/chat", nil)
	s.opts.HTTP.LLM.PUT = newOption(m.PUT, llm+"/api/chat", nil)
	s.opts.HTTP.LLM.PATCH = newOption(m.PATCH, llm+"/api/chat", nil)
	s.opts.HTTP.LLM.DELETE = newOption(m.DELETE, llm+"/api/delete", nil)
	s.opts.HTTP.LLM.OPTIONS = newOption(m.OPTIONS, llm+"/api/tags", nil)

	s.opts.HTTP.LLMGenerate.GET = newOption(m.GET, llm+"/api/tags", nil)
	s.opts.HTTP.LLMGenerate.POST = newOption(m.POST, llm+"/api/generate", nil)
	s.opts.HTTP.LLMGenerate.PUT = newOption(m.PUT, llm+"/api/generate", nil)
	s.opts.HTTP.LLMGenerate.PATCH = newOption(m.PATCH, llm+"/api/generate", nil)
	s.opts.HTTP.LLMGenerate.DELETE = newOption(m.DELETE, llm+"/api/delete", nil)
	s.opts.HTTP.LLMGenerate.OPTIONS = newOption(m.OPTIONS, llm+"/api/tags", nil)

	s.opts.HTTP.LLMChat.GET = newOption(m.GET, llm+"/api/tags", nil)
	s.opts.HTTP.LLMChat.POST = newOption(m.POST, llm+"/api/chat", nil)
	s.opts.HTTP.LLMChat.PUT = newOption(m.PUT, llm+"/api/chat", nil)
	s.opts.HTTP.LLMChat.PATCH = newOption(m.PATCH, llm+"/api/chat", nil)
	s.opts.HTTP.LLMChat.DELETE = newOption(m.DELETE, llm+"/api/delete", nil)
	s.opts.HTTP.LLMChat.OPTIONS = newOption(m.OPTIONS, llm+"/api/tags", nil)

	s.opts.HTTP.ASR.GET = newOption(m.GET, asr+"/client/ws/speech", nil)
	s.opts.HTTP.ASR.POST = newOption(m.POST, asr+"/client/ws/speech", nil)
	s.opts.HTTP.ASR.PUT = newOption(m.PUT, asr+"/client/ws/speech", nil)
	s.opts.HTTP.ASR.PATCH = newOption(m.PATCH, asr+"/client/ws/speech", nil)
	s.opts.HTTP.ASR.DELETE = newOption(m.DELETE, asr+"/client/ws/speech", nil)
	s.opts.HTTP.ASR.OPTIONS = newOption(m.OPTIONS, asr+"/client/ws/speech", nil)

	s.opts.HTTP.TTS.GET = newOption(m.GET, tts+"/voices", nil)
	s.opts.HTTP.TTS.POST = newOption(m.POST, tts+"/tts", nil)
	s.opts.HTTP.TTS.PUT = newOption(m.PUT, tts+"/tts", nil)
	s.opts.HTTP.TTS.PATCH = newOption(m.PATCH, tts+"/tts", nil)
	s.opts.HTTP.TTS.DELETE = newOption(m.DELETE, tts+"/tts", nil)
	s.opts.HTTP.TTS.OPTIONS = newOption(m.OPTIONS, tts+"/voices", nil)

	s.opts.HTTP.TTSProcess.GET = newOption(m.GET, tts+"/voices", nil)
	s.opts.HTTP.TTSProcess.POST = newOption(m.POST, tts+"/tts", nil)
	s.opts.HTTP.TTSProcess.PUT = newOption(m.PUT, tts+"/tts", nil)
	s.opts.HTTP.TTSProcess.PATCH = newOption(m.PATCH, tts+"/tts", nil)
	s.opts.HTTP.TTSProcess.DELETE = newOption(m.DELETE, tts+"/tts", nil)
	s.opts.HTTP.TTSProcess.OPTIONS = newOption(m.OPTIONS, tts+"/voices", nil)

	s.opts.HTTPS.LLM.GET = newOption(m.GET, "https://127.0.0.1:11434/api/tags", nil)
	s.opts.HTTPS.LLM.POST = newOption(m.POST, "https://127.0.0.1:11434/api/chat", nil)

	var body entities.ApiBody

	s.LlmGET()
	s.LlmPOST(body)
	s.LlmPUT(body)
	s.LlmPATCH(body)
	s.LlmDELETE()
	s.LlmOPTIONS()

	s.AsrGET()
	s.AsrPOST(body)
	s.AsrPUT(body)
	s.AsrPATCH(body)
	s.AsrDELETE()
	s.AsrOPTIONS()

	s.TtsGET()
	s.TtsPOST(body)
	s.TtsPUT(body)
	s.TtsPATCH(body)
	s.TtsDELETE()
	s.TtsOPTIONS()

	s.LlmGenerateGET()
	s.LlmGeneratePOST(body)
	s.LlmGeneratePUT(body)
	s.LlmGeneratePATCH(body)
	s.LlmGenerateDELETE()
	s.LlmGenerateOPTIONS()

	s.LlmChatGET()
	s.LlmChatPOST(body)
	s.LlmChatPUT(body)
	s.LlmChatPATCH(body)
	s.LlmChatDELETE()
	s.LlmChatOPTIONS()

	s.TtsProcessGET()
	s.TtsProcessPOST(body)
	s.TtsProcessPUT(body)
	s.TtsProcessPATCH(body)
	s.TtsProcessDELETE()
	s.TtsProcessOPTIONS()

	s.LlmHttpsGET()
	s.LlmHttpsPOST(body)
	s.LlmHttpsPUT(body)
	s.LlmHttpsPATCH(body)
	s.LlmHttpsDELETE()
	s.LlmHttpsOPTIONS()

	s.AsrHttpsGET()
	s.AsrHttpsPOST(body)
	s.AsrHttpsPUT(body)
	s.AsrHttpsPATCH(body)
	s.AsrHttpsDELETE()
	s.AsrHttpsOPTIONS()

	s.TtsHttpsGET()
	s.TtsHttpsPOST(body)
	s.TtsHttpsPUT(body)
	s.TtsHttpsPATCH(body)
	s.TtsHttpsDELETE()
	s.TtsHttpsOPTIONS()

	s.LlmGenerateHttpsGET()
	s.LlmGenerateHttpsPOST(body)
	s.LlmGenerateHttpsPUT(body)
	s.LlmGenerateHttpsPATCH(body)
	s.LlmGenerateHttpsDELETE()
	s.LlmGenerateHttpsOPTIONS()

	s.LlmChatHttpsGET()
	s.LlmChatHttpsPOST(body)
	s.LlmChatHttpsPUT(body)
	s.LlmChatHttpsPATCH(body)
	s.LlmChatHttpsDELETE()
	s.LlmChatHttpsOPTIONS()

	s.TtsProcessHttpsGET()
	s.TtsProcessHttpsPOST(body)
	s.TtsProcessHttpsPUT(body)
	s.TtsProcessHttpsPATCH(body)
	s.TtsProcessHttpsDELETE()
	s.TtsProcessHttpsOPTIONS()

	check := func(name string, res ServiceRes) {
	}
	check("LlmGET", s.LlmGET())
	check("TtsGET", s.TtsGET())
	check("Custom", s.Custom(m.GET, llm+"/api/tags", nil))
}

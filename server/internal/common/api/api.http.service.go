package api

import "Glowl/internal/common/entities"

func NewHttpService() *Service {
	return (&Service{}).constructorService()
}

func (s *Service) constructorService() *Service {
	s.opts = optionFactory()
	return s
}

func (s *Service) do(opt fetchOption) ServiceRes {
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) doBody(opt fetchOption, body entities.ApiBody) ServiceRes {
	opt.Body = body
	return s.do(opt)
}

/*-----------------------------------------------------\
|				Standard KRUDs (HTTP)					|
\-----------------------------------------------------*/

func (s *Service) LlmGET() ServiceRes { return s.do(s.opts.HTTP.LLM.GET) }
func (s *Service) LlmPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLM.POST, body)
}
func (s *Service) LlmPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLM.PUT, body)
}
func (s *Service) LlmPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLM.PATCH, body)
}
func (s *Service) LlmDELETE() ServiceRes  { return s.do(s.opts.HTTP.LLM.DELETE) }
func (s *Service) LlmOPTIONS() ServiceRes { return s.do(s.opts.HTTP.LLM.OPTIONS) }

func (s *Service) AsrGET() ServiceRes { return s.do(s.opts.HTTP.ASR.GET) }
func (s *Service) AsrPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.ASR.POST, body)
}
func (s *Service) AsrPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.ASR.PUT, body)
}
func (s *Service) AsrPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.ASR.PATCH, body)
}
func (s *Service) AsrDELETE() ServiceRes  { return s.do(s.opts.HTTP.ASR.DELETE) }
func (s *Service) AsrOPTIONS() ServiceRes { return s.do(s.opts.HTTP.ASR.OPTIONS) }

func (s *Service) TtsGET() ServiceRes { return s.do(s.opts.HTTP.TTS.GET) }
func (s *Service) TtsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.TTS.POST, body)
}
func (s *Service) TtsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.TTS.PUT, body)
}
func (s *Service) TtsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.TTS.PATCH, body)
}
func (s *Service) TtsDELETE() ServiceRes  { return s.do(s.opts.HTTP.TTS.DELETE) }
func (s *Service) TtsOPTIONS() ServiceRes { return s.do(s.opts.HTTP.TTS.OPTIONS) }

/*-----------------------------------------------------\
|				Custom CRUD (HTTP)						|
\-----------------------------------------------------*/

func (s *Service) LlmGenerateGET() ServiceRes { return s.do(s.opts.HTTP.LLMGenerate.GET) }
func (s *Service) LlmGeneratePOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLMGenerate.POST, body)
}
func (s *Service) LlmGeneratePUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLMGenerate.PUT, body)
}
func (s *Service) LlmGeneratePATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLMGenerate.PATCH, body)
}
func (s *Service) LlmGenerateDELETE() ServiceRes  { return s.do(s.opts.HTTP.LLMGenerate.DELETE) }
func (s *Service) LlmGenerateOPTIONS() ServiceRes { return s.do(s.opts.HTTP.LLMGenerate.OPTIONS) }

func (s *Service) LlmChatGET() ServiceRes { return s.do(s.opts.HTTP.LLMChat.GET) }
func (s *Service) LlmChatPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLMChat.POST, body)
}
func (s *Service) LlmChatPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLMChat.PUT, body)
}
func (s *Service) LlmChatPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.LLMChat.PATCH, body)
}
func (s *Service) LlmChatDELETE() ServiceRes  { return s.do(s.opts.HTTP.LLMChat.DELETE) }
func (s *Service) LlmChatOPTIONS() ServiceRes { return s.do(s.opts.HTTP.LLMChat.OPTIONS) }

func (s *Service) TtsProcessGET() ServiceRes { return s.do(s.opts.HTTP.TTSProcess.GET) }
func (s *Service) TtsProcessPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.TTSProcess.POST, body)
}
func (s *Service) TtsProcessPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.TTSProcess.PUT, body)
}
func (s *Service) TtsProcessPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTP.TTSProcess.PATCH, body)
}
func (s *Service) TtsProcessDELETE() ServiceRes  { return s.do(s.opts.HTTP.TTSProcess.DELETE) }
func (s *Service) TtsProcessOPTIONS() ServiceRes { return s.do(s.opts.HTTP.TTSProcess.OPTIONS) }

/*-----------------------------------------------------\
|				Standard KRUDs (HTTPS)					|
\-----------------------------------------------------*/

func (s *Service) LlmHttpsGET() ServiceRes { return s.do(s.opts.HTTPS.LLM.GET) }
func (s *Service) LlmHttpsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLM.POST, body)
}
func (s *Service) LlmHttpsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLM.PUT, body)
}
func (s *Service) LlmHttpsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLM.PATCH, body)
}
func (s *Service) LlmHttpsDELETE() ServiceRes  { return s.do(s.opts.HTTPS.LLM.DELETE) }
func (s *Service) LlmHttpsOPTIONS() ServiceRes { return s.do(s.opts.HTTPS.LLM.OPTIONS) }

func (s *Service) AsrHttpsGET() ServiceRes { return s.do(s.opts.HTTPS.ASR.GET) }
func (s *Service) AsrHttpsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.ASR.POST, body)
}
func (s *Service) AsrHttpsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.ASR.PUT, body)
}
func (s *Service) AsrHttpsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.ASR.PATCH, body)
}
func (s *Service) AsrHttpsDELETE() ServiceRes  { return s.do(s.opts.HTTPS.ASR.DELETE) }
func (s *Service) AsrHttpsOPTIONS() ServiceRes { return s.do(s.opts.HTTPS.ASR.OPTIONS) }

func (s *Service) TtsHttpsGET() ServiceRes { return s.do(s.opts.HTTPS.TTS.GET) }
func (s *Service) TtsHttpsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.TTS.POST, body)
}
func (s *Service) TtsHttpsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.TTS.PUT, body)
}
func (s *Service) TtsHttpsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.TTS.PATCH, body)
}
func (s *Service) TtsHttpsDELETE() ServiceRes  { return s.do(s.opts.HTTPS.TTS.DELETE) }
func (s *Service) TtsHttpsOPTIONS() ServiceRes { return s.do(s.opts.HTTPS.TTS.OPTIONS) }

func (s *Service) LlmGenerateHttpsGET() ServiceRes { return s.do(s.opts.HTTPS.LLMGenerate.GET) }
func (s *Service) LlmGenerateHttpsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLMGenerate.POST, body)
}
func (s *Service) LlmGenerateHttpsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLMGenerate.PUT, body)
}
func (s *Service) LlmGenerateHttpsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLMGenerate.PATCH, body)
}
func (s *Service) LlmGenerateHttpsDELETE() ServiceRes  { return s.do(s.opts.HTTPS.LLMGenerate.DELETE) }
func (s *Service) LlmGenerateHttpsOPTIONS() ServiceRes { return s.do(s.opts.HTTPS.LLMGenerate.OPTIONS) }

func (s *Service) LlmChatHttpsGET() ServiceRes { return s.do(s.opts.HTTPS.LLMChat.GET) }
func (s *Service) LlmChatHttpsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLMChat.POST, body)
}
func (s *Service) LlmChatHttpsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLMChat.PUT, body)
}
func (s *Service) LlmChatHttpsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.LLMChat.PATCH, body)
}
func (s *Service) LlmChatHttpsDELETE() ServiceRes  { return s.do(s.opts.HTTPS.LLMChat.DELETE) }
func (s *Service) LlmChatHttpsOPTIONS() ServiceRes { return s.do(s.opts.HTTPS.LLMChat.OPTIONS) }

func (s *Service) TtsProcessHttpsGET() ServiceRes { return s.do(s.opts.HTTPS.TTSProcess.GET) }
func (s *Service) TtsProcessHttpsPOST(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.TTSProcess.POST, body)
}
func (s *Service) TtsProcessHttpsPUT(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.TTSProcess.PUT, body)
}
func (s *Service) TtsProcessHttpsPATCH(body entities.ApiBody) ServiceRes {
	return s.doBody(s.opts.HTTPS.TTSProcess.PATCH, body)
}
func (s *Service) TtsProcessHttpsDELETE() ServiceRes  { return s.do(s.opts.HTTPS.TTSProcess.DELETE) }
func (s *Service) TtsProcessHttpsOPTIONS() ServiceRes { return s.do(s.opts.HTTPS.TTSProcess.OPTIONS) }

/*-----------------------------------------------------\
|					Сustom request							|
\-----------------------------------------------------*/

func (s *Service) Custom(
	method entities.ApiMethod,
	url string,
	body entities.ApiBody,
) ServiceRes {
	return s.do(newOption(method, url, body))
}

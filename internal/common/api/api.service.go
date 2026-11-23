package api

import "Glowl/internal/common/entities"

// Standard KRUDs

func (s *Service) LlmGET() ServiceRes {
	res, err := fetch(s.opts.LLM_GET)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmPOST(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_POST
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmPUT(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_PUT
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmPATCH(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_PATCH
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmDELETE() ServiceRes {
	res, err := fetch(s.opts.LLM_DELETE)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmOPTIONS() ServiceRes {
	res, err := fetch(s.opts.LLM_OPTIONS)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) AsrGET() ServiceRes {
	res, err := fetch(s.opts.ASR_GET)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) AsrPOST(body entities.ApiBody) ServiceRes {
	opt := s.opts.ASR_POST
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) AsrPUT(body entities.ApiBody) ServiceRes {
	opt := s.opts.ASR_PUT
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) AsrPATCH(body entities.ApiBody) ServiceRes {
	opt := s.opts.ASR_PATCH
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) AsrDELETE() ServiceRes {
	res, err := fetch(s.opts.ASR_DELETE)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) AsrOPTIONS() ServiceRes {
	res, err := fetch(s.opts.ASR_OPTIONS)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsGET() ServiceRes {
	res, err := fetch(s.opts.TTS_GET)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsPOST(body entities.ApiBody) ServiceRes {
	opt := s.opts.TTS_POST
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsPUT(body entities.ApiBody) ServiceRes {
	opt := s.opts.TTS_PUT
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsPATCH(body entities.ApiBody) ServiceRes {
	opt := s.opts.TTS_PATCH
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsDELETE() ServiceRes {
	res, err := fetch(s.opts.TTS_DELETE)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsOPTIONS() ServiceRes {
	res, err := fetch(s.opts.TTS_OPTIONS)
	return ServiceRes{Body: res, Err: err}
}

// Custom CRUD

func (s *Service) LlmGenerateGET() ServiceRes {
	res, err := fetch(s.opts.LLM_GENERATE_GET)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmGeneratePOST(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_GENERATE_POST
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmGeneratePUT(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_GENERATE_PUT
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmGeneratePATCH(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_GENERATE_PATCH
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmGenerateDELETE() ServiceRes {
	res, err := fetch(s.opts.LLM_GENERATE_DELETE)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmGenerateOPTIONS() ServiceRes {
	res, err := fetch(s.opts.LLM_GENERATE_OPTIONS)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmChatGET() ServiceRes {
	res, err := fetch(s.opts.LLM_CHAT_GET)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmChatPOST(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_CHAT_POST
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmChatPUT(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_CHAT_PUT
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmChatPATCH(body entities.ApiBody) ServiceRes {
	opt := s.opts.LLM_CHAT_PATCH
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmChatDELETE() ServiceRes {
	res, err := fetch(s.opts.LLM_CHAT_DELETE)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) LlmChatOPTIONS() ServiceRes {
	res, err := fetch(s.opts.LLM_CHAT_OPTIONS)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsProcessGET() ServiceRes {
	res, err := fetch(s.opts.TTS_PROCESS_GET)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsProcessPOST(body entities.ApiBody) ServiceRes {
	opt := s.opts.TTS_PROCESS_POST
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsProcessPUT(body entities.ApiBody) ServiceRes {
	opt := s.opts.TTS_PROCESS_PUT
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsProcessPATCH(body entities.ApiBody) ServiceRes {
	opt := s.opts.TTS_PROCESS_PATCH
	opt.Body = body
	res, err := fetch(opt)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsProcessDELETE() ServiceRes {
	res, err := fetch(s.opts.TTS_PROCESS_DELETE)
	return ServiceRes{Body: res, Err: err}
}

func (s *Service) TtsProcessOPTIONS() ServiceRes {
	res, err := fetch(s.opts.TTS_PROCESS_OPTIONS)
	return ServiceRes{Body: res, Err: err}
}

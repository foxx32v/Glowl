package entities

func NewApiMethods() apiMethods {
	return apiMethods{
		GET:     methodGet,
		POST:    methodPost,
		PUT:     methodPut,
		PATCH:   methodPatch,
		OPTIONS: methodOptions,
		DELETE:  methodDelete,
	}
}

func NewApiUrl() apiUrl {
	return apiUrl{
		Localhost: urlLocalhost,
	}
}

func NewApiPort() apiPort {
	return apiPort{
		Llm: portLlm,
		Tts: portTts,
		Asr: portAsr,
	}
}

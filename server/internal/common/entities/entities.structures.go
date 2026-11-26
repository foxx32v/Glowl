package entities

type apiMethods struct {
	GET     ApiMethod
	POST    ApiMethod
	PUT     ApiMethod
	PATCH   ApiMethod
	OPTIONS ApiMethod
	DELETE  ApiMethod
	NIL     ApiMethod
}

type apiUrl struct {
	Localhost ApiUrl
}

type apiPort struct {
	Llm ApiPort
	Asr ApiPort
	Tts ApiPort
}

type apiProtocol struct {
	Http  ApiProtocol
	Https ApiProtocol
	Ws    ApiProtocol
}

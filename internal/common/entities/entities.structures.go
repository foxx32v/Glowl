package entities

type apiMethods struct {
	GET     ApiMethod
	POST    ApiMethod
	PUT     ApiMethod
	PATCH   ApiMethod
	OPTIONS ApiMethod
	DELETE  ApiMethod
}

type apiUrl struct {
	Localhost ApiUrl
}

type apiPort struct {
	Llm ApiPort
	Asr ApiPort
	Tts ApiPort
}

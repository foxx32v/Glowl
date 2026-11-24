package entities

const (
	methodGet     ApiMethod = "GET"
	methodPost    ApiMethod = "POST"
	methodPut     ApiMethod = "PUT"
	methodPatch   ApiMethod = "PATCH"
	methodOptions ApiMethod = "OPTIONS"
	methodDelete  ApiMethod = "DELETE"
	methodNil     ApiMethod = ""
)

const (
	urlLocalhost ApiUrl = "localhost"
)

const (
	portTts ApiPort = 9898
	portAsr ApiPort = 2700
	portLlm ApiPort = 11434
)

const (
	protocolHTTP  ApiProtocol = "http"
	protocolHTTPS ApiProtocol = "https"
	protocolWS    ApiProtocol = "ws"
)

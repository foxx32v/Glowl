package entities

import (
	"fmt"
)

func ParseApiMethod(str string) (ApiMethod, error) {
	switch str {
	case string(methodGet):
		return methodGet, nil
	case string(methodPost):
		return methodPost, nil
	case string(methodPut):
		return methodPut, nil
	case string(methodPatch):
		return methodPatch, nil
	case string(methodDelete):
		return methodDelete, nil
	case string(methodOptions):
		return methodOptions, nil
	default:
		return "", fmt.Errorf("Error parse API method")
	}
}

func ParseApiUrl(str string) (ApiUrl, error) {
	switch str {
	case string(urlLocalhost):
		return urlLocalhost, nil
	default:
		return "", fmt.Errorf("Error Parse API url")
	}
}

func ParseApiPort(port string) (ApiPort, error) {
	switch port {
	case string(portTts):
		return portTts, nil
	case string(portAsr):
		return portAsr, nil
	case string(portLlm):
		return portLlm, nil
	default:
		return "", fmt.Errorf("Error parse API port")
	}
}

func ParseApiProtocol(str string) (ApiProtocol, error) {
	switch str {
	case string(protocolHTTP):
		return protocolHTTP, nil
	case string(protocolHTTPS):
		return protocolHTTPS, nil
	case string(protocolWS):
		return protocolWS, nil
	default:
		return "", fmt.Errorf("Error parse API protocol")
	}
}

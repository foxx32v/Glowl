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
		return "", fmt.Errorf("Error HTTP method")
	}
}

func ParseApiUrl(str string) (ApiUrl, error) {
	switch str {
	case string(urlLocalhost):
		return urlLocalhost, nil
	default:
		return "", fmt.Errorf("")
	}
}

func ParseApiPort(port int) (ApiPort, error) {
	switch port {
	case int(portTts):
		return portTts, nil
	case int(portAsr):
		return portAsr, nil
	case int(portLlm):
		return portLlm, nil
	default:
		return 0, fmt.Errorf("")
	}
}

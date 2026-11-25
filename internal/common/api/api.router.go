package api

import (
	"Glowl/internal/common/entities"
	"encoding/json"
	"io"
	"net/http"
)

func constructorHttpServer(network string, address string) (*server, error) {
	server, err := createHttpServer(network, address)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (s *server) SetHandler(h http.Handler) {
	s.http.Handler = h
}

func handler(fn func() ServiceRes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeRes(w, fn())
	}
}
func handlerBody(fn func(entities.ApiBody) ServiceRes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()
		writeRes(w, fn(entities.ApiBody(raw)))
	}
}

func writeRes(w http.ResponseWriter, res ServiceRes) {
	if res.Err != nil {
		writeErr(w, http.StatusInternalServerError, res.Err)
		return
	}
	if res.Body == nil {
		w.WriteHeader(http.StatusOK)
		return
	}
	defer res.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, res.Body)
}

func writeErr(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func RegisterRoutes(mux *http.ServeMux, svc *Service) {
	mux.HandleFunc("GET /llm", handler(svc.LlmGET))
	mux.HandleFunc("POST /llm", handlerBody(svc.LlmPOST))
	mux.HandleFunc("PUT /llm", handlerBody(svc.LlmPUT))
	mux.HandleFunc("PATCH /llm", handlerBody(svc.LlmPATCH))
	mux.HandleFunc("DELETE /llm", handler(svc.LlmDELETE))
	mux.HandleFunc("OPTIONS /llm", handler(svc.LlmOPTIONS))

	mux.HandleFunc("GET /asr", handler(svc.AsrGET))
	mux.HandleFunc("POST /asr", handlerBody(svc.AsrPOST))
	mux.HandleFunc("PUT /asr", handlerBody(svc.AsrPUT))
	mux.HandleFunc("PATCH /asr", handlerBody(svc.AsrPATCH))
	mux.HandleFunc("DELETE /asr", handler(svc.AsrDELETE))
	mux.HandleFunc("OPTIONS /asr", handler(svc.AsrOPTIONS))

	mux.HandleFunc("GET /tts", handler(svc.TtsGET))
	mux.HandleFunc("POST /tts", handlerBody(svc.TtsPOST))
	mux.HandleFunc("PUT /tts", handlerBody(svc.TtsPUT))
	mux.HandleFunc("PATCH /tts", handlerBody(svc.TtsPATCH))
	mux.HandleFunc("DELETE /tts", handler(svc.TtsDELETE))
	mux.HandleFunc("OPTIONS /tts", handler(svc.TtsOPTIONS))

	mux.HandleFunc("GET /llm/generate", handler(svc.LlmGenerateGET))
	mux.HandleFunc("POST /llm/generate", handlerBody(svc.LlmGeneratePOST))
	mux.HandleFunc("PUT /llm/generate", handlerBody(svc.LlmGeneratePUT))
	mux.HandleFunc("PATCH /llm/generate", handlerBody(svc.LlmGeneratePATCH))
	mux.HandleFunc("DELETE /llm/generate", handler(svc.LlmGenerateDELETE))
	mux.HandleFunc("OPTIONS /llm/generate", handler(svc.LlmGenerateOPTIONS))

	mux.HandleFunc("GET /llm/chat", handler(svc.LlmChatGET))
	mux.HandleFunc("POST /llm/chat", handlerBody(svc.LlmChatPOST))
	mux.HandleFunc("PUT /llm/chat", handlerBody(svc.LlmChatPUT))
	mux.HandleFunc("PATCH /llm/chat", handlerBody(svc.LlmChatPATCH))
	mux.HandleFunc("DELETE /llm/chat", handler(svc.LlmChatDELETE))
	mux.HandleFunc("OPTIONS /llm/chat", handler(svc.LlmChatOPTIONS))

	mux.HandleFunc("GET /tts/process", handler(svc.TtsProcessGET))
	mux.HandleFunc("POST /tts/process", handlerBody(svc.TtsProcessPOST))
	mux.HandleFunc("PUT /tts/process", handlerBody(svc.TtsProcessPUT))
	mux.HandleFunc("PATCH /tts/process", handlerBody(svc.TtsProcessPATCH))
	mux.HandleFunc("DELETE /tts/process", handler(svc.TtsProcessDELETE))
	mux.HandleFunc("OPTIONS /tts/process", handler(svc.TtsProcessOPTIONS))

	mux.HandleFunc("GET /https/llm", handler(svc.LlmHttpsGET))
	mux.HandleFunc("POST /https/llm", handlerBody(svc.LlmHttpsPOST))
	mux.HandleFunc("PUT /https/llm", handlerBody(svc.LlmHttpsPUT))
	mux.HandleFunc("PATCH /https/llm", handlerBody(svc.LlmHttpsPATCH))
	mux.HandleFunc("DELETE /https/llm", handler(svc.LlmHttpsDELETE))
	mux.HandleFunc("OPTIONS /https/llm", handler(svc.LlmHttpsOPTIONS))

	mux.HandleFunc("GET /https/asr", handler(svc.AsrHttpsGET))
	mux.HandleFunc("POST /https/asr", handlerBody(svc.AsrHttpsPOST))
	mux.HandleFunc("PUT /https/asr", handlerBody(svc.AsrHttpsPUT))
	mux.HandleFunc("PATCH /https/asr", handlerBody(svc.AsrHttpsPATCH))
	mux.HandleFunc("DELETE /https/asr", handler(svc.AsrHttpsDELETE))
	mux.HandleFunc("OPTIONS /https/asr", handler(svc.AsrHttpsOPTIONS))

	mux.HandleFunc("GET /https/tts", handler(svc.TtsHttpsGET))
	mux.HandleFunc("POST /https/tts", handlerBody(svc.TtsHttpsPOST))
	mux.HandleFunc("PUT /https/tts", handlerBody(svc.TtsHttpsPUT))
	mux.HandleFunc("PATCH /https/tts", handlerBody(svc.TtsHttpsPATCH))
	mux.HandleFunc("DELETE /https/tts", handler(svc.TtsHttpsDELETE))
	mux.HandleFunc("OPTIONS /https/tts", handler(svc.TtsHttpsOPTIONS))

	mux.HandleFunc("GET /https/llm/generate", handler(svc.LlmGenerateHttpsGET))
	mux.HandleFunc("POST /https/llm/generate", handlerBody(svc.LlmGenerateHttpsPOST))
	mux.HandleFunc("PUT /https/llm/generate", handlerBody(svc.LlmGenerateHttpsPUT))
	mux.HandleFunc("PATCH /https/llm/generate", handlerBody(svc.LlmGenerateHttpsPATCH))
	mux.HandleFunc("DELETE /https/llm/generate", handler(svc.LlmGenerateHttpsDELETE))
	mux.HandleFunc("OPTIONS /https/llm/generate", handler(svc.LlmGenerateHttpsOPTIONS))

	mux.HandleFunc("GET /https/llm/chat", handler(svc.LlmChatHttpsGET))
	mux.HandleFunc("POST /https/llm/chat", handlerBody(svc.LlmChatHttpsPOST))
	mux.HandleFunc("PUT /https/llm/chat", handlerBody(svc.LlmChatHttpsPUT))
	mux.HandleFunc("PATCH /https/llm/chat", handlerBody(svc.LlmChatHttpsPATCH))
	mux.HandleFunc("DELETE /https/llm/chat", handler(svc.LlmChatHttpsDELETE))
	mux.HandleFunc("OPTIONS /https/llm/chat", handler(svc.LlmChatHttpsOPTIONS))

	mux.HandleFunc("GET /https/tts/process", handler(svc.TtsProcessHttpsGET))
	mux.HandleFunc("POST /https/tts/process", handlerBody(svc.TtsProcessHttpsPOST))
	mux.HandleFunc("PUT /https/tts/process", handlerBody(svc.TtsProcessHttpsPUT))
	mux.HandleFunc("PATCH /https/tts/process", handlerBody(svc.TtsProcessHttpsPATCH))
	mux.HandleFunc("DELETE /https/tts/process", handler(svc.TtsProcessHttpsDELETE))
	mux.HandleFunc("OPTIONS /https/tts/process", handler(svc.TtsProcessHttpsOPTIONS))
}

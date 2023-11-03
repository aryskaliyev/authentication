package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) infoLog(message string) {
	h.logger.InfoLog.Printf(message)
}

func (h *Handler) errorLog(message string) {
	h.logger.ErrorLog.Printf("%v", message)
}

func (h *Handler) handleError(w http.ResponseWriter, status int, message string) {
	response := map[string]interface{}{
		"message": message,
	}

	h.respondWithJSON(w, status, response)
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

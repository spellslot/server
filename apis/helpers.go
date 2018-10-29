package apis

import "net/http"

func sendResponse(w http.ResponseWriter, contentType string, payload []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

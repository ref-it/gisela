package common

import "net/http"

func WriteJsonData(w http.ResponseWriter, data []byte) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return w
}

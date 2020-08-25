package response

import (
	"encoding/json"
	"net/http"
)

// JSON emite una respuesta exitosa en formato JSON
func JSON(data interface{}, statusCode int, w http.ResponseWriter) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		JSONError("error on create response", http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(jsonData)
}

// JSONError emite una respuesta errónea simple en formato JSON
func JSONError(message string, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	msg := map[string]interface{}{"msg": message}
	_ = json.NewEncoder(w).Encode(msg)
}

// JSONErrors emite una respuesta errónea detallada en formato JSON
func JSONErrors(message string, details []string, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	msg := map[string]interface{}{"msg": message, "details": details}
	_ = json.NewEncoder(w).Encode(msg)
}

// PlainText emite una respuesta en formato plano
func PlainText(content string, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(content))
}

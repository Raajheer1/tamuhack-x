package utils

import "net/http"

func Respond(statusCode int, msg string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}

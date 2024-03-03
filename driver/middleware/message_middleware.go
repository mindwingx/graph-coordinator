package middleware

import (
	"context"
	"encoding/json"
	"github.com/mindwingx/graph-coordinator/helper"
	"net/http"
)

type resp struct {
	Message string `json:"message"`
}

func MsgMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var res resp
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			helper.HandleResponse(rw, http.StatusUnprocessableEntity, "invalid payload")
			return
		}

		if !isValidMessageSize(res.Message) {
			//fmt.Println("[api-coordinator] middleware: message rejected by size")
			helper.HandleResponse(rw, http.StatusUnprocessableEntity, "invalid message size")
			return
		}

		// retrieve the request json decoded body for the related handler function
		ctx := context.WithValue(r.Context(), "decodedMsg", res.Message)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

// HELPER METHODS

func isValidMessageSize(message string) bool {
	size := uintptr(len(message))
	return size >= helper.MinLegalByteAmount && size <= helper.MaxLegalByteAmount
}

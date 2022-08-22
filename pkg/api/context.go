package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HandlerFunc func(w Writer, r *http.Request) error

func (h HandlerFunc) ServeHTTP(w Writer, r *http.Request) {
	err := h(w, r)
	if err != nil {
		w.JSON(500, Response{Code: InternalError})
		fmt.Println("An unhandled error occured... ", err)
	}
}

type Writer interface {
	JSON(status int, body Response) error
}

type writer struct {
	http.ResponseWriter
	wasWritten bool
}

func (w *writer) JSON(status int, body Response) error {
	if w.wasWritten {
		return errors.New("headers were aleady written")
	}
	w.wasWritten = true
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	x, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.Write(x)
	return nil
}

func NewWriter(w http.ResponseWriter) Writer {
	return &writer{wasWritten: false, ResponseWriter: w}
}

func W(h HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dw := &writer{wasWritten: false, ResponseWriter: w}
		start := time.Now()
		h.ServeHTTP(dw, r)
		end := time.Now()
		fmt.Println(end.Sub(start))
		if dw.wasWritten {
			return
		}
		dw.WriteHeader(http.StatusOK)
	})
}

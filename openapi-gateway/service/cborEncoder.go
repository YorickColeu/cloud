package service

import (
	"net/http"

	"github.com/go-ocf/ocf-cloud/openapi-connector/events"

	"github.com/go-ocf/kit/codec/cbor"
)

func newCBORResponseWriterEncoder(contentType string) func(w http.ResponseWriter, v interface{}) error {
	return func(w http.ResponseWriter, v interface{}) error {
		if v == nil {
			return nil
		}
		w.Header().Set(events.ContentTypeKey, contentType)
		return cbor.WriteTo(w, v)
	}
}

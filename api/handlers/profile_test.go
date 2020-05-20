package handlers

import (
	"net/http"
	"testing"
)

func TestHandlerData_ProfileLostHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		hd   *HandlerData
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hd.ProfileLostHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_ProfileLostOpeningHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		hd   *HandlerData
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hd.ProfileLostOpeningHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_ProfileFoundHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		hd   *HandlerData
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hd.ProfileFoundHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_ProfileFoundOpeningHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		hd   *HandlerData
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hd.ProfileFoundOpeningHandler(tt.args.w, tt.args.r)
		})
	}
}

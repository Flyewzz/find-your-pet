package handlers

import (
	"net/http"
	"testing"
)

func TestHandlerData_FoundHandler(t *testing.T) {
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
			tt.hd.FoundHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_FoundByIdGetHandler(t *testing.T) {
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
			tt.hd.FoundByIdGetHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_AddFoundHandler(t *testing.T) {
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
			tt.hd.AddFoundHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_RemoveFoundHandler(t *testing.T) {
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
			tt.hd.RemoveFoundHandler(tt.args.w, tt.args.r)
		})
	}
}

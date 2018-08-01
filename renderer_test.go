package mulate

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewRenderer(t *testing.T) {
	type args struct {
		mlt  *Template
		page *Page
	}
	tests := []struct {
		name string
		args args
		want *Renderer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewRenderer(tt.args.mlt, tt.args.page); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewRenderer() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRenderer_Render(t *testing.T) {
	type fields struct {
		page *Page
		mlt  *Template
	}
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := Renderer{
			page: tt.fields.page,
			mlt:  tt.fields.mlt,
		}
		if err := r.Render(tt.args.w); (err != nil) != tt.wantErr {
			t.Errorf("%q. Renderer.Render() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestRenderer_WriteContentType(t *testing.T) {
	type fields struct {
		page *Page
		mlt  *Template
	}
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := Renderer{
			page: tt.fields.page,
			mlt:  tt.fields.mlt,
		}
		r.WriteContentType(tt.args.w)
	}
}

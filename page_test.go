package mulate

import (
	"html/template"
	"net/http"
	"testing"
)

type fields struct {
	Title       string
	Status      int
	ContentType string
	Layout      string
	JS          []string
	CSS         []string
	Request     *http.Request
	content     template.HTML
	funcs       template.FuncMap
	errLayout   string
}

func TestPage_SetLayout(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Page{
			Title:       tt.fields.Title,
			Status:      tt.fields.Status,
			ContentType: tt.fields.ContentType,
			Layout:      tt.fields.Layout,
			JS:          tt.fields.JS,
			CSS:         tt.fields.CSS,
			Request:     tt.fields.Request,
			content:     tt.fields.content,
			funcs:       tt.fields.funcs,
			errLayout:   tt.fields.errLayout,
		}
		got, err := p.SetLayout(tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Page.SetLayout() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Page.SetLayout() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPage_SetTitle(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Page{
			Title:       tt.fields.Title,
			Status:      tt.fields.Status,
			ContentType: tt.fields.ContentType,
			Layout:      tt.fields.Layout,
			JS:          tt.fields.JS,
			CSS:         tt.fields.CSS,
			Request:     tt.fields.Request,
			content:     tt.fields.content,
			funcs:       tt.fields.funcs,
			errLayout:   tt.fields.errLayout,
		}
		got, err := p.SetTitle(tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Page.SetTitle() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Page.SetTitle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPage_Raise(t *testing.T) {
	type args struct {
		status  int
		title   string
		message string
		abort   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Page{
			Title:       tt.fields.Title,
			Status:      tt.fields.Status,
			ContentType: tt.fields.ContentType,
			Layout:      tt.fields.Layout,
			JS:          tt.fields.JS,
			CSS:         tt.fields.CSS,
			Request:     tt.fields.Request,
			content:     tt.fields.content,
			funcs:       tt.fields.funcs,
			errLayout:   tt.fields.errLayout,
		}
		got, err := p.Raise(tt.args.status, tt.args.title, tt.args.message, tt.args.abort)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Page.Raise() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Page.Raise() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPage_RedirectFound(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Page{
			Title:       tt.fields.Title,
			Status:      tt.fields.Status,
			ContentType: tt.fields.ContentType,
			Layout:      tt.fields.Layout,
			JS:          tt.fields.JS,
			CSS:         tt.fields.CSS,
			Request:     tt.fields.Request,
			content:     tt.fields.content,
			funcs:       tt.fields.funcs,
			errLayout:   tt.fields.errLayout,
		}
		got, err := p.RedirectFound(tt.args.uri)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Page.RedirectFound() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Page.RedirectFound() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

package mulate

import (
	"html/template"
	"net/http"
	"reflect"
	"testing"

	"github.com/oxtoacart/bpool"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want *Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := New(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTemplate_DisableCache(t *testing.T) {
	type fields struct {
		Funcs        template.FuncMap
		config       Config
		includes     []string
		pages        map[string]pageDef
		layouts      map[string]layoutDef
		bufpool      *bpool.BufferPool
		disableCache bool
	}
	type args struct {
		flag bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t := &Template{
			Funcs:        tt.fields.Funcs,
			config:       tt.fields.config,
			includes:     tt.fields.includes,
			pages:        tt.fields.pages,
			layouts:      tt.fields.layouts,
			bufpool:      tt.fields.bufpool,
			disableCache: tt.fields.disableCache,
		}
		t.DisableCache(tt.args.flag)
	}
}

func TestTemplate_Pages(t *testing.T) {
	type fields struct {
		Funcs        template.FuncMap
		config       Config
		includes     []string
		pages        map[string]pageDef
		layouts      map[string]layoutDef
		bufpool      *bpool.BufferPool
		disableCache bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		mlt := &Template{
			Funcs:        tt.fields.Funcs,
			config:       tt.fields.config,
			includes:     tt.fields.includes,
			pages:        tt.fields.pages,
			layouts:      tt.fields.layouts,
			bufpool:      tt.fields.bufpool,
			disableCache: tt.fields.disableCache,
		}
		if got := mlt.Pages(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Template.Pages() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTemplate_Layouts(t *testing.T) {
	type fields struct {
		Funcs        template.FuncMap
		config       Config
		includes     []string
		pages        map[string]pageDef
		layouts      map[string]layoutDef
		bufpool      *bpool.BufferPool
		disableCache bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		mlt := &Template{
			Funcs:        tt.fields.Funcs,
			config:       tt.fields.config,
			includes:     tt.fields.includes,
			pages:        tt.fields.pages,
			layouts:      tt.fields.layouts,
			bufpool:      tt.fields.bufpool,
			disableCache: tt.fields.disableCache,
		}
		if got := mlt.Layouts(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Template.Layouts() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTemplate_LoadTemplates(t *testing.T) {
	type fields struct {
		Funcs        template.FuncMap
		config       Config
		includes     []string
		pages        map[string]pageDef
		layouts      map[string]layoutDef
		bufpool      *bpool.BufferPool
		disableCache bool
	}
	type args struct {
		funcs template.FuncMap
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
		mlt := &Template{
			Funcs:        tt.fields.Funcs,
			config:       tt.fields.config,
			includes:     tt.fields.includes,
			pages:        tt.fields.pages,
			layouts:      tt.fields.layouts,
			bufpool:      tt.fields.bufpool,
			disableCache: tt.fields.disableCache,
		}
		if err := mlt.LoadTemplates(tt.args.funcs); (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.LoadTemplates() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestTemplate_RenderPage(t *testing.T) {
	type fields struct {
		Funcs        template.FuncMap
		config       Config
		includes     []string
		pages        map[string]pageDef
		layouts      map[string]layoutDef
		bufpool      *bpool.BufferPool
		disableCache bool
	}
	type args struct {
		uri   string
		funcs template.FuncMap
		r     *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantP   *Page
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		mlt := &Template{
			Funcs:        tt.fields.Funcs,
			config:       tt.fields.config,
			includes:     tt.fields.includes,
			pages:        tt.fields.pages,
			layouts:      tt.fields.layouts,
			bufpool:      tt.fields.bufpool,
			disableCache: tt.fields.disableCache,
		}
		gotP, err := mlt.RenderPage(tt.args.uri, tt.args.funcs, tt.args.r)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.RenderPage() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotP, tt.wantP) {
			t.Errorf("%q. Template.RenderPage() = %v, want %v", tt.name, gotP, tt.wantP)
		}
	}
}

func TestTemplate_RenderLayout(t *testing.T) {
	type fields struct {
		Funcs        template.FuncMap
		config       Config
		includes     []string
		pages        map[string]pageDef
		layouts      map[string]layoutDef
		bufpool      *bpool.BufferPool
		disableCache bool
	}
	type args struct {
		w http.ResponseWriter
		p *Page
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
		mlt := &Template{
			Funcs:        tt.fields.Funcs,
			config:       tt.fields.config,
			includes:     tt.fields.includes,
			pages:        tt.fields.pages,
			layouts:      tt.fields.layouts,
			bufpool:      tt.fields.bufpool,
			disableCache: tt.fields.disableCache,
		}
		if err := mlt.RenderLayout(tt.args.w, tt.args.p); (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.RenderLayout() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

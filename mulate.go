// Package mulate represents multiple layouts template functionality
package mulate

// https://stackoverflow.com/questions/42747183/how-to-render-templates-to-multiple-layouts-in-go

import (
	"fmt"
	"github.com/oxtoacart/bpool"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Config holds config variables and its defaults
type Config struct {
	Root        string   `long:"templates" default:"tmpl/" description:"Templates root path"`
	Ext         string   `long:"mask" default:".tmpl" description:"Templates filename mask"`
	Includes    string   `long:"includes" default:"inc/" description:"Includes path"`
	Layouts     string   `long:"layouts" default:"layout/" description:"Layouts path"`
	Pages       string   `long:"pages" default:"page/" description:"Pages path"`
	DefLayout   string   `long:"def_layout" default:"base" description:"Default layout template"`
	ErrLayout   string   `long:"err_layout" default:"error" description:"Error page layout template"`
	ContentType []string `long:"content-type" default:"text/html; charset=utf-8" description:"Default content type"`
	BufferSize  int      `long:"buffer" default:"64" description:"Template buffer size"`
}

// pageDef holds single template data
type pageDef struct {
	file string             // real template filename
	name string             // name for Execute call
	tmpl *template.Template // template object
}

// layoutDef holds single layout data
type layoutDef struct {
	file string
	tmpl *template.Template
}

// Template holds all internally used template attributes
type Template struct {
	Funcs        template.FuncMap
	config       Config
	includes     []string
	pages        map[string]pageDef   // pages[uri]
	layouts      map[string]layoutDef // layouts[name]
	bufpool      *bpool.BufferPool
	disableCache bool // reload template on each request
}

// New creates mulate object
func New(cfg Config) *Template {
	return &Template{config: cfg}
}

// DisableCache disables template caching
func (t *Template) DisableCache(flag bool) {
	t.disableCache = flag
}

// Pages returns a list of page templates founded
func (t *Template) Pages() []string {
	keys := make([]string, len(t.pages))

	i := 0
	for k := range t.pages {
		keys[i] = k
		i++
	}
	return keys
}

// Layouts returns a list of page layouts founded
func (t *Template) Layouts() []string {
	keys := make([]string, len(t.layouts))
	i := 0
	for k := range t.layouts {
		keys[i] = k
		i++
	}
	return keys
}

// LoadTemplates parses template files
func (t *Template) LoadTemplates(funcs template.FuncMap) error {
	pages := make(map[string]pageDef)
	layouts := make(map[string]layoutDef)

	pageFiles := []string{}
	root := t.config.Root + t.config.Pages
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "walk error")
		}
		if !f.IsDir() {
			pageFiles = append(pageFiles, path)
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "walk failed")
	}

	layoutFiles, err := filepath.Glob(t.config.Root + t.config.Layouts + "*" + t.config.Ext)
	if err != nil {
		return errors.Wrap(err, "layout globbing failed")
	}

	inc, err := filepath.Glob(t.config.Root + t.config.Includes + "*" + t.config.Ext)
	if err != nil {
		return errors.Wrap(err, "includes globbing failed")
	}
	t.includes = inc
	// TODO: .Delims(e.config.Delims.Left, e.config.Delims.Right)

	for _, file := range layoutFiles {
		name := strings.TrimSuffix(filepath.Base(file), t.config.Ext)
		td := layoutDef{file: file}
		if !t.disableCache {
			files := append([]string{file}, t.includes...)
			td.tmpl, err = template.New(name).Funcs(funcs).ParseFiles(files...)
			if err != nil {
				return errors.Wrap(err, "process layout")
			}
		}
		layouts[name] = td
	}
	if _, ok := layouts[t.config.DefLayout]; !ok {
		return fmt.Errorf("Default layout (%s) not found", t.config.DefLayout)
	}
	if _, ok := layouts[t.config.ErrLayout]; !ok {
		return fmt.Errorf("Error layout (%s) not found", t.config.ErrLayout)
	}

	for _, file := range pageFiles {
		name := filepath.Base(file)
		td := pageDef{file: file, name: name}
		if !t.disableCache {
			files := append([]string{file}, t.includes...)
			td.tmpl, err = template.New(name).Funcs(funcs).ParseFiles(files...)
			if err != nil {
				return errors.Wrap(err, "process page")
			}
		}
		uri := strings.TrimPrefix(file, (t.config.Root + t.config.Pages))
		uri = strings.TrimSuffix(uri, t.config.Ext)
		uri = strings.TrimSuffix(uri, "index")
		//		uri = strings.TrimSuffix(uri, "/")
		pages[uri] = td
	}

	// we called from main, no goroutines yet
	//e.tplMutex.Lock()
	t.bufpool = bpool.NewBufferPool(t.config.BufferSize)
	t.pages = pages
	t.layouts = layouts
	t.Funcs = funcs
	//e.tplMutex.Unlock()
	return nil
}

// RenderPage renders content for given uri
func (t *Template) RenderPage(uri string, funcs template.FuncMap, r *http.Request) (p *Page, err error) {
	p = &Page{
		Status:      http.StatusOK,
		ContentType: t.config.ContentType,
		Layout:      t.config.DefLayout,
		funcs:       funcs,
		Request:     r,
		errLayout:   t.config.ErrLayout,
	}

	tmplDef, ok := t.pages[uri]
	if !ok {
		err = fmt.Errorf("The page %s does not exist.", uri)
		p.Raise(http.StatusInternalServerError, "NOT FOUND", err.Error(), false)
		return
	}

	buf := t.bufpool.Get()
	var tmpl *template.Template
	if t.disableCache {
		files := append([]string{tmplDef.file}, t.includes...)
		tmpl, err = template.New(tmplDef.name).Funcs(p.funcs).ParseFiles(files...)
		if err != nil {
			return
		}
	} else {
		tmpl = tmplDef.tmpl
		tmpl.Funcs(p.funcs)
	}
	err = tmpl.ExecuteTemplate(buf, tmplDef.name, p)
	if err != nil {
		return
	}
	p.content = template.HTML(buf.Bytes())
	t.bufpool.Put(buf)
	return
}

// RenderLayout renders page in given layout
func (t *Template) RenderLayout(w http.ResponseWriter, p *Page) (err error) {
	tmplDef, ok := t.layouts[p.Layout]
	if !ok {
		return fmt.Errorf("The layout %s does not exist.", p.Layout)
	}

	buf := t.bufpool.Get()
	var tmpl *template.Template
	if t.disableCache {
		files := append([]string{tmplDef.file}, t.includes...)
		tmpl, err = template.New(p.Layout).Funcs(p.funcs).ParseFiles(files...)
		if err != nil {
			return errors.Wrap(err, "parse layout")
		}
	} else {
		tmpl = tmplDef.tmpl
		tmpl.Funcs(p.funcs)
	}
	err = tmpl.ExecuteTemplate(buf, p.Layout+t.config.Ext, struct {
		Content template.HTML
		*Page
	}{p.content, p})
	if err != nil {
		return errors.Wrap(err, "exec layout")
	}

	w.WriteHeader(p.Status)
	buf.WriteTo(w)
	t.bufpool.Put(buf)
	return nil
}

// This file holds template attrs & methods

package mulate

import (
	"github.com/pkg/errors"
	"html/template"
	"net/http"
)

type Page struct {
	Title       string
	Status      int
	ContentType []string
	layout      string
	content     template.HTML
	JS          []string
	CSS         []string
	API         template.FuncMap
}

func (p *Page) SetLayout(name string) (string, error) {
	p.layout = name
	return "", nil
}

func (p *Page) SetTitle(name string) (string, error) {
	p.Title = name
	return "", nil
}

func (p *Page) SetError(status int, title, message string, abort bool) (string, error) {
	p.Status = status
	p.Title = title
	p.content = template.HTML(message)
	p.layout = "error"
	if abort {
		return "", errors.New(message)
	}
	return "", nil
}

func (p *Page) RedirectFound(uri string) (string, error) {
	p.Status = http.StatusFound
	p.Title = uri
	return "", errors.New("RedirectFound")
}

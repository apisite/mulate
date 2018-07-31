// This file holds template attrs & methods

package mulate

import (
	"github.com/pkg/errors"
	"html/template"
	"net/http"
)

// ErrRedirect is an error returned when page needs to be redirected
var ErrRedirect = errors.New("Abort with redirect")

// Page holds page attributes
type Page struct {
	Title       string
	Status      int
	ContentType []string
	Layout      string
	JS          []string
	CSS         []string
	Request     *http.Request
	content     template.HTML
	funcs       template.FuncMap
	errLayout   string
}

// SetLayout - set page layout
func (p *Page) SetLayout(name string) (string, error) {
	p.Layout = name
	return "", nil
}

// SetTitle - set page title
func (p *Page) SetTitle(name string) (string, error) {
	p.Title = name
	return "", nil
}

// Raise - abort template processing (if given) and raise error
func (p *Page) Raise(status int, title, message string, abort bool) (string, error) {
	p.Status = status
	p.Title = title
	p.content = template.HTML(message)
	p.Layout = p.errLayout
	if abort {
		return "", errors.New(message)
	}
	return "", nil
}

// RedirectFound - abort template processing and return redirect with StatusFound status
func (p *Page) RedirectFound(uri string) (string, error) {
	p.Status = http.StatusFound
	p.Title = uri
	return "", ErrRedirect // TODO: Is there a way to pass status & title via error?
}

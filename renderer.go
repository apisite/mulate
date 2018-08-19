// This file holds per-request rendering methods

package mulate

import (
	"net/http"
)

// Renderer holds per request rendering attributes
type Renderer struct {
	page *Page
	mlt  *Template
}

// NewRenderer creates new renderer object
func NewRenderer(mlt *Template, page *Page) *Renderer {
	return &Renderer{mlt: mlt, page: page}
}

// Render - render page and write it to w
func (r Renderer) Render(w http.ResponseWriter) error {
	return r.mlt.RenderLayout(w, r.page)
}

// WriteContentType writes page content type
func (r Renderer) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	// TODO: r.Page.ContentType
	//	if val := header["Content-Type"]; len(val) == 0 {
	header["Content-Type"] = []string{r.page.ContentType}
	//	}
}

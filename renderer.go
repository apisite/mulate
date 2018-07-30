// This file holds per-request rendering methods

package mulate

import (
	"net/http"
)

type Renderer struct {
	page *Page
	mlt  *Template
}

func NewRenderer(mlt *Template, page *Page) *Renderer {
	return &Renderer{mlt: mlt, page: page}
}

func (r Renderer) Render(w http.ResponseWriter) error {
	return r.mlt.RenderLayout(w, r.page)
}

func (r Renderer) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	// TODO: r.Page.ContentType
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = r.mlt.config.ContentType
	}
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/onrik/logrus/filename"
	//    "filename"

	"github.com/jessevdk/go-flags"

	mapper "github.com/birkirb/loggers-mapper-logrus"
	"gopkg.in/birkirb/loggers.v1"

	//	"github.com/Masterminds/sprig"

	"github.com/apisite/mulate"
)

// Config holds all config vars
type Config struct {
	Addr string `long:"http_addr" default:"localhost:8081"  description:"Http listen address"`

	Template mulate.Config `group:"Template Options"`
}

func main() {

	cfg := &Config{}
	p := flags.NewParser(cfg, flags.Default)

	_, err := p.Parse()
	if err != nil {
		if !strings.HasPrefix(err.Error(), "\nUsage") {
			fmt.Fprintf(os.Stderr, "error: %+v", err)
		}
		os.Exit(0)
	}

	l := logrus.New()

	if true /* gin.IsDebugging() */ {
		l.SetLevel(logrus.DebugLevel)
		l.AddHook(filename.NewHook())
	}
	log := mapper.NewLogger(l)

	mlt := mulate.New(cfg.Template)
	mlt.DisableCache(true)

	// extra template functions
	// allFuncs := sprig.FuncMap()

	// See also; github.com/htfy96/reformism
	allFuncs := make(template.FuncMap, 0)
	allFuncs["HTML"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	// Set blank at parse time
	// See https://stackoverflow.com/a/18302879
	allFuncs["data"] = func() interface{} {
		return nil
	}

	err = mlt.LoadTemplates(allFuncs)
	if err != nil {
		log.Fatal(err)
	}
	for _, uri := range mlt.Pages() {
		log.Debugf("Registering uri: %s", uri)
		http.HandleFunc("/"+uri, handleHTML(mlt, uri, log))
	}

	if err := http.ListenAndServe(cfg.Addr, nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

// handleHTML returns page handler
func handleHTML(mlt *mulate.Template, uri string, log loggers.Contextual) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Handling page (%s)", uri)
		reqFuncs := mlt.Funcs // per request copy
		reqFuncs["data"] = func() interface{} {
			return data
		}

		p, err := mlt.RenderPage(uri, reqFuncs, r)
		if err != nil {
			if p.Status == http.StatusMovedPermanently || p.Status == http.StatusFound {
				http.Redirect(w, r, p.Title, p.Status)
				return
			}
			log.Errorf("page error: (%+v)", err)
			if p.Status == http.StatusOK {
				p.Status = http.StatusInternalServerError
				p.Raise(p.Status, "Internal", err.Error(), false)
			}
		}
		renderer := mulate.NewRenderer(mlt, p)
		renderer.WriteContentType(w)
		log.Debugf("render layout (%s)", p.Layout)
		err = renderer.Render(w)
		if err != nil {
			log.Errorf("Error while handling uri (%s): %s", uri, err)
		}
	}
}

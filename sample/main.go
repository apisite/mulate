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

	"github.com/apisite/mulate"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var data = TodoPageData{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{Title: "Task 1", Done: false},
		{Title: "Task 2", Done: true},
		{Title: "Task 3", Done: true},
	},
}

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

	mlt, _ := mulate.New(cfg.Template, log)
	mlt.DisableCache(true)

	// See also; github.com/htfy96/reformism
	allFuncs := make(template.FuncMap, 0)
	// Set blank at parse time
	// See https://stackoverflow.com/a/18302879
	// TODO: use ginapi func types
	allFuncs["api"] = func(nsp string, method string, v ...interface{}) (interface{}, error) { return nil, nil }
	allFuncs["get"] = func(name string) (*string, error) { return nil, nil }

	err = mlt.LoadTemplates(allFuncs)
	if err != nil {
		log.Fatal(err)
	}
	for _, uri := range mlt.Pages() {
		log.Debugf("Registering uri: %s", uri)
		u := uri
		http.HandleFunc("/"+uri, func(w http.ResponseWriter, r *http.Request) {
			log.Debugf("Handling uri: %s", u)

			//			tmpl.Execute(w, data)
			allFuncs := make(template.FuncMap, 0)
			//allFuncs["api"] = API{}.APIFunc()
			//	allFuncs["get"] = e.API.APIMeta()
			p, err := mlt.RenderPage(u, allFuncs)
			if err != nil {
				if p.Status == http.StatusMovedPermanently || p.Status == http.StatusFound {
					http.Redirect(w, r, p.Title, p.Status)
					return
				}
				log.Debugf("page error: (%+v)", err)
				if p.Status == http.StatusOK {
					p.Status = http.StatusInternalServerError
					p.Raise(p.Status, "Internal", err.Error(), false)
				}
			}
			renderer := mulate.NewRenderer(mlt, p)
			renderer.WriteContentType(w)
			err = renderer.Render(w)
			if err != nil {
				log.Errorf("Error while handling uri (%s): %s", u, err)
			}
		})

	}

	if err := http.ListenAndServe(cfg.Addr, nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

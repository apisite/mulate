# mulate
> golang library for multiple layouts & templates

[![GoCard][gc1]][gc2]
 [![GitHub Release][gr1]][gr2]
 [![GitHub code size in bytes][sz]]()
 [![GitHub license][gl1]][gl2]

[gc1]: https://goreportcard.com/badge/apisite/mulate
[gc2]: https://goreportcard.com/report/github.com/apisite/mulate
[gr1]: https://img.shields.io/github/release/apisite/mulate.svg
[gr2]: https://github.com/apisite/mulate/releases
[sz]: https://img.shields.io/github/languages/code-size/apisite/mulate.svg
[gl1]: https://img.shields.io/github/license/apisite/mulate.svg
[gl2]: LICENSE

![Request processing flow](flow.png)

## Template structure

According to [sample](sample/), site templates tree might looks like: 

```
tmpl
├── inc
│   ├── footer.tmpl
│   ├── head.tmpl
│   └── menu.tmpl
├── layout
│   ├── base.tmpl
│   ├── error.tmpl
│   └── wide.tmpl
└── page
    ├── admin
    │   └── index.tmpl
    ├── err.tmpl
    ├── index.tmpl
    └── page.tmpl

```

## Usage

```go
    mlt, _ := mulate.New(cfg.Template, log)
    mlt.DisableCache(true)

    allFuncs := make(template.FuncMap, 0)

    err = mlt.LoadTemplates(allFuncs)
```

See also: [sample](sample/)

### Template methods

```
{{ .SetTitle "admin:index" -}}
```

```
{{ .SetLayout "wide" -}}
```

```
{{ .SetError 403 "Test error" "Error description" true }}
```

## TODO

* [ ] .Redirect
* [ ] gin-mulate
* [ ] pack renderering funcs
* [ ] move html/template to interface
* [ ] check another template engine
* [ ] test coverage
* [ ] docs

## Library name

**mulate** means **mu**ltiple **la**youts & **te**mplates.

## License

The MIT License (MIT), see [LICENSE](LICENSE).

Copyright (c) 2018 Aleksei Kovrizhkin <lekovr+apisite@gmail.com>


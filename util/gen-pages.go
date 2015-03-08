// This program generates a pages meta data needed by the frontend router.
// It is done automatically by gulp tasks but you can also run it manually:
// go run util/gen-pages.go > app/scripts/pages.js

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	// flags
	templatesRoot = flag.String("d", "app/templates", "templates dir")

	// metaTemplates defines which templates go into a page meta as string values.
	metaTemplates = []string{"title", "mastheadBgClass"}
)

type pageMeta map[string]interface{}

func main() {
	pages := make(map[string]pageMeta)

	err := filepath.Walk(*templatesRoot, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if p == *templatesRoot || fi.IsDir() {
			return nil
		}
		ext := filepath.Ext(p)
		if ext != ".html" || strings.HasPrefix(fi.Name(), "layout_") {
			return nil
		}
		name := p[len(*templatesRoot)+1 : len(p)-len(ext)]
		t, err := template.New("").Delims("{%", "%}").ParseFiles(p)
		if err != nil {
			return err
		}
		pages[name] = metaFromTemplate(t)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	var b []byte
	if b, err = json.MarshalIndent(pages, "", "  "); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "// auto-generated - do not modify\nIOWA.PAGES = %s;\n", b)
}

func metaFromTemplate(t *template.Template) pageMeta {
	m := make(pageMeta)
	m["hasBeenLoaded"] = false
	for _, n := range metaTemplates {
		b := new(bytes.Buffer)
		if err := t.ExecuteTemplate(b, n, nil); err != nil {
			continue
		}
		m[n] = b.String()
	}
	return m
}
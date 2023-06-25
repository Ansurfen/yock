// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestFuncs(t *testing.T) {
	literalSet := map[string]any{
		`{{.}}`:                    "Hello World!",
		`The {{ .name }} welcome!`: map[string]string{"name": "ansurfen"},
		`{{if eq .x .y}} x = y {{else if lt .x .y}} x < y {{else}} x > y {{end}}`:                                                          map[string]int{"x": 1, "y": 2},
		`{{define "T1"}}ONE{{end}}{{define "T2"}}TWO{{end}}{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}{{- template "T3" -}}`: nil,
		`{{with $x := "output"}}{{$x | printf "%q"}}{{end}}`:                                                                               nil,
		`{{range $x := . -}}{{println $x}}{{- end}}`:                                                                                       []int{1, 2, 3},
		`{{ block "T3" . }}T3{{ end }}`:                                                                                                    nil,
	}
	tmpl := NewTemplate()
	for k, v := range literalSet {
		res, err := tmpl.OnceParse(k, v)
		if err != nil {
			fmt.Printf("%s err: %v\n", k, err)
		} else {
			fmt.Printf("%s\n", res)
		}
	}
	fmt.Println("---------------------")
	gen := TemplateLiteralGenerator{}
	generatorSet := map[string]any{
		gen.Get(""): "你好，世界",
		fmt.Sprintf("The %s welcome!", gen.Get("name")): map[string]string{"name": "ansurfen"},
		gen.Batch(gen.Define("T1", "ONE"),
			gen.Define("T2", "TWO"),
			gen.Define("T3",
				gen.Batch(gen.Template("T1"), " ", gen.Template("T2"))),
			gen.Template("T3")): nil,
		gen.With(
			gen.Var("x", gen.Str("output")),
			gen.Scope(gen.Batch(gen.Ref("x"), gen.Pipe(), gen.Call("printf", gen.Str("%q"))))): nil,
		gen.If([]string{"eq .x .y", "lt .x .y"}, []string{"x = y", "x < y", "x > y"}): map[string]int{"x": 1, "y": 2},
	}
	for k, v := range generatorSet {
		res, err := tmpl.OnceParse(k, v)
		if err != nil {
			fmt.Printf("%s err: %v\n", k, err)
		} else {
			fmt.Printf("%s\n", res)
		}
	}
	linkerSet := map[string]any{
		`{{using "sprig"}}`:                         nil,
		`{{load "test.html"}}{{- template "T3" -}}`: nil,
	}
	for k, v := range linkerSet {
		res, err := tmpl.OnceParse(tmpl.Preprocess(k), v)
		if err != nil {
			fmt.Printf("%s err: %v\n", k, err)
		} else {
			fmt.Printf("%s\n", res)
		}
	}
	fmt.Println(tmpl.LazyLoadPkgs().LazyLoadFiles().OnceParse(`{{load "test.html"}}{{- lazytemplate "T3" -}}`, nil))
}

func TestLuaFunc(t *testing.T) {
	tmpl := NewTemplate().LoadPkgs("lua", "tmpl", "css")
	// script := `
	// {{using "lua"}}
	// `
	// tmpl.OnceParse(script, nil)
	// tmpl.LazyLoadPkgs()
	doc, _ := tmpl.OnceParse(`{{ SetVar "a" (Selectors (Str "Hello World!") (Str ".title")) | unescaped }} {{StrJoin "a" (Str "b") (Str "c") | unescaped}}`, nil)
	fmt.Println(tmpl.OnceParse(doc, nil))
}

func TestWorkflow(t *testing.T) {
	script := `
	{{ using "lua" "tmpl" "css" }}
	{{load "test.html"}} 
	{{- template "T3" -}}
	{{ LVar "a" (Selectors (Str "Hello World!") (Str ".title")) | unescaped }} {{StrJoin "a" (Str "b") (Str "c") | unescaped}}
	`
	tmpl := NewTemplate(TemplateOpt{
		epoch: 3,
		todos: []struct {
			preprocess bool
			loadfiles  []string
			loadpkgs   []string
		}{
			{preprocess: true, loadpkgs: []string{"lua", "tmpl", "css"}},
			{}, {},
		},
	})
	if doc, err := tmpl.Run(script); err != nil {
		panic(err)
	} else {
		fmt.Println(doc)
	}
}

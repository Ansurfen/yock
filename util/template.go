// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"strings"

	"github.com/Masterminds/sprig"
)

var (
	globalLuaGenerator LuaLiteralGenerator
	globalCSSGenerator CSSLiteralGenerator
)

var templatePkgs = map[string]template.FuncMap{
	"sprig": sprig.FuncMap(),
	"lua": {
		"LVar":    globalLuaGenerator.Var,
		"SetVar":  globalLuaGenerator.SetVar,
		"StrJoin": globalLuaGenerator.StrJoin,
		"Eq":      globalLuaGenerator.Eq,
		"Range":   globalLuaGenerator.Range,
		"IRange":  globalLuaGenerator.IRange,
	},
	"tmpl": {
		"Str": globalLuaGenerator.Str,
	},
	"css": {
		"Selectors": globalCSSGenerator.Selectors,
		"Style":     globalCSSGenerator.Style,
	},
}

func GetPkg(name string) template.FuncMap {
	if pkg, ok := templatePkgs[name]; ok {
		return pkg
	}
	return nil
}

type TemplateLinker struct {
	Pkgs  []string
	Files []string
}

func (ld *TemplateLinker) Using(pkgs ...string) string {
	ld.Pkgs = append(ld.Pkgs, pkgs...)
	return ""
}

func (ld *TemplateLinker) Load(files ...string) string {
	ld.Files = append(ld.Files, files...)
	return ""
}

type TemplateOpt struct {
	epoch int
	todos []struct {
		preprocess bool
		loadfiles  []string
		loadpkgs   []string
	}
}

type Template struct {
	*template.Template
	ld  *TemplateLinker
	opt TemplateOpt
}

func NewTemplate(opt ...TemplateOpt) *Template {
	var topt TemplateOpt
	if len(opt) > 0 {
		topt = opt[0]
	} else {
		topt = TemplateOpt{epoch: 1}
	}
	tmpl := &Template{
		Template: template.New("TEMPLATE"),
		ld:       &TemplateLinker{},
		opt:      topt,
	}
	tmpl.Funcs(template.FuncMap{
		"using":        tmpl.ld.Using,
		"load":         tmpl.ld.Load,
		"lazytemplate": lazytemplate,
		"unescaped": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	return tmpl
}

func (tmpl *Template) Run(doc string) (string, error) {
	var (
		err        error
		preprocess bool
	)
	for i := 0; i < tmpl.opt.epoch; i++ {
		if preprocess {
			tmpl.LazyLoadFiles().LazyLoadPkgs()
			preprocess = false
		}
		todo := tmpl.opt.todos[i]
		if todo.preprocess {
			doc = tmpl.Preprocess(doc)
			preprocess = true
		}
		if len(todo.loadfiles) > 0 {
			tmpl.LoadFiles(todo.loadfiles...)
		}
		if len(todo.loadpkgs) > 0 {
			tmpl.LoadPkgs(todo.loadpkgs...)
		}
		doc, err = tmpl.OnceParse(doc, nil)
		if err != nil {
			return "", err
		}
	}
	return doc, nil
}

func lazytemplate(names ...string) template.HTML {
	str := ""
	for _, name := range names {
		str += (fmt.Sprintf(`"%s" `, name))
	}
	return template.HTML(fmt.Sprintf(`{{ template %s }}`, str))
}

func (tmpl *Template) Preprocess(doc string) string {
	return strings.ReplaceAll(doc, "template", "lazytemplate")
}

func (tmpl *Template) LazyLoadFiles() *Template {
	tmpl.ParseFiles(tmpl.ld.Files...)
	return tmpl
}

func (tmpl *Template) LoadFiles(filename ...string) *Template {
	tmpl.ParseFiles(filename...)
	return tmpl
}

func (tmpl *Template) LoadPkgs(pkgs ...string) *Template {
	for _, pkgName := range pkgs {
		if pkg, ok := templatePkgs[pkgName]; ok {
			tmpl.Template = tmpl.Funcs(pkg)
		}
	}
	return tmpl
}

func (tmpl *Template) LazyLoadPkgs() *Template {
	for _, pkgName := range tmpl.ld.Pkgs {
		if pkg, ok := templatePkgs[pkgName]; ok {
			tmpl.Template = tmpl.Funcs(pkg)
		}
	}
	return tmpl
}

func (tmpl *Template) OnceParse(text string, data any) (string, error) {
	tmp, err := tmpl.Template.Clone()
	if err != nil {
		return "", err
	}
	tmp, err = tmp.Parse(text)
	if err != nil {
		return "", err
	}
	w := bytes.NewBufferString("")
	err = tmp.Execute(w, data)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

func (tmpl *Template) OnceParseFile(data any, filenames ...string) (string, error) {
	tmp, err := tmpl.Template.ParseFiles(filenames...)
	if err != nil {
		return "", err
	}
	tmp, err = tmp.ParseFiles(filenames...)
	if err != nil {
		return "", err
	}
	w := bytes.NewBufferString("")
	err = tmp.Execute(w, data)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

type TemplateLiteralGenerator struct{}

func (TemplateLiteralGenerator) Get(name string) string {
	return fmt.Sprintf(`{{ .%s }}`, name)
}

func (TemplateLiteralGenerator) Define(name, expression string) string {
	return fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, expression)
}

func (TemplateLiteralGenerator) Template(name string) string {
	return fmt.Sprintf(`{{template "%s"}}`, name)
}

func (TemplateLiteralGenerator) Block(name, scope, expression string) string {
	return fmt.Sprintf(`{{block "%s" %s}}%s{{end}}`, name, scope, expression)
}

func (TemplateLiteralGenerator) Batch(cmd ...string) string {
	return strings.Join(cmd, "")
}

func (TemplateLiteralGenerator) Str(str string) string {
	return fmt.Sprintf(`"%s"`, str)
}

func (TemplateLiteralGenerator) With(condition, expression string) string {
	return fmt.Sprintf(`{{with %s}}%s{{end}}`, condition, expression)
}

func (TemplateLiteralGenerator) Ref(name string) string {
	return fmt.Sprintf(`$%s`, name)
}

func (TemplateLiteralGenerator) Pipe() string {
	return " | "
}

func (TemplateLiteralGenerator) Call(name string, params ...string) string {
	return fmt.Sprintf(`%s %s`, name, strings.Join(params, " "))
}

func (TemplateLiteralGenerator) Var(name, value string) string {
	return fmt.Sprintf(`$%s := %s`, name, value)
}

func (TemplateLiteralGenerator) Scope(cmd string) string {
	return fmt.Sprintf(`{{%s}}`, cmd)
}

func (TemplateLiteralGenerator) Range(condition, expression string) string {
	return fmt.Sprintf(`{{range %s -}}%s{{- end}}`, condition, expression)
}

func (TemplateLiteralGenerator) If(conditions []string, expressions []string) string {
	condLen, expLen := len(conditions), len(expressions)
	if condLen != expLen && condLen != expLen-1 {
		return ""
	}
	stat := ""
	for idx, condition := range conditions {
		expression := expressions[idx]
		switch idx {
		case 0:
			stat += fmt.Sprintf(`{{if %s}} %s`, condition, expression)
		default:
			stat += fmt.Sprintf(`{{else if %s}}%s`, condition, expression)
		}
	}
	if condLen == expLen-1 {
		stat += fmt.Sprintf(`{{else}}%s`, expressions[expLen-1])
	}
	stat += "{{end}}"
	return stat
}

type LuaLiteralGenerator struct {
	TemplateLiteralGenerator
}

func (gen LuaLiteralGenerator) Var(name, value string) string {
	return fmt.Sprintf(`local %s = %s;`, name, value)
}

func (LuaLiteralGenerator) SetVar(name, value string) string {
	return fmt.Sprintf(`%s = %s;`, name, value)
}

func (LuaLiteralGenerator) GetVar(name string) string {
	return name
}

func (LuaLiteralGenerator) Range(kv, t string) string {
	kvSet := strings.Split(kv, ",")
	if len(kvSet) != 2 {
		panic(errors.New(""))
	}
	return fmt.Sprintf("%s, %s in pairs(%s)", kvSet[0], kvSet[1], t)
}

func (LuaLiteralGenerator) IRange(kv, t string) string {
	kvSet := strings.Split(kv, ",")
	if len(kvSet) != 2 {
		panic(errors.New(""))
	}
	return fmt.Sprintf("%s, %s in ipairs(%s)", kvSet[0], kvSet[1], t)
}

func (LuaLiteralGenerator) Eq(a, b string) string {
	return fmt.Sprintf("%s == %s", a, b)
}

func (LuaLiteralGenerator) StrJoin(name string, value ...string) string {
	return fmt.Sprintf(`%s = %s .. %s;`, name, name, strings.Join(value, " .. "))
}

type CSSLiteralGenerator struct{}

func (CSSLiteralGenerator) Selectors(name string, selectors ...string) string {
	return fmt.Sprintf(`{{ Style %s %s }}`, name, strings.Join(selectors, " "))
}

func (CSSLiteralGenerator) Style(name string, selectors ...string) string {
	return fmt.Sprintf(`style:Render(%s, %s)`, name, strings.Join(selectors, ","))
}

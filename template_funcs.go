package main

import (
	"log"
	"strings"
	"text/template"
	"unicode"
)

var knownAcronyms = []string{
	"ID",
	"API",
	"URI",
	"URL",
	"FQDN",
}

func sqlToGoType(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "tinyint", "smallint", "mediumint", "int", "integer", "serial":
		return "int"
	case "bigint":
		return "int64"
	case "float", "double", "decimal":
		return "float64"
	case "char", "varchar", "text", "mediumtext", "longtext":
		return "string"
	case "date", "datetime", "timestamp":
		return "time.Time"
	default:
		return "interface{}"
	}
}

func goToSqlType(goType string) string {
	switch goType {
	case "int", "int8", "int16", "int32", "uint",
		"uint8", "uint16", "uint32":
		return "INTEGER"
	case "int64", "uint64":
		return "BIGINT"
	case "float32", "float64":
		return "FLOAT"
	case "string":
		return "VARCHAR(255)"
	case "time.Time":
		return "TIMESTAMP"
	default:
		return "VARCHAR(255)"
	}
}
func pluralize(word string) string {
	var pw string
	lw := strings.ToLower(word)
	switch {
	case strings.HasSuffix(lw, "s") ||
		strings.HasSuffix(lw, "sh") ||
		strings.HasSuffix(lw, "ch"):
		pw = word + "es"
	case strings.HasSuffix(lw, "y"):
		if strings.Index("aeiou", string(lw[len(lw)-2])) == -1 {
			pw = word[:len(word)-1] + "ies"
		} else {
			pw = word + "s"
		}
	default:
		pw = word + "s"
	}
	if strings.ToUpper(word) == word {
		return strings.ToUpper(pw)
	}
	return pw
}

func singularize(word string) string {
	var sw string
	lw := strings.ToLower(word)
	switch {
	case strings.HasSuffix(lw, "ies"):
		sw = word[:len(word)-3] + "y"
	case strings.HasSuffix(lw, "es"):
		sw = word[:len(word)-2]
	case strings.HasSuffix(lw, "s"):
		sw = word[:len(word)-1]
	default:
		sw = word
	}
	if strings.ToUpper(word) == word {
		return strings.ToUpper(sw)
	}
	return sw
}

func snakeToCamel(s string) string {
	acronyms := func(as []string) map[string]string {
		am := make(map[string]string, len(as))
		for _, a := range as {
			am[strings.ToLower(a)] = a
		}
		return am
	}(knownAcronyms)
	var ok bool
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i], ok = acronyms[word]
		if !ok {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, "")
}

func camelToSnake(s string) string {
	acronyms := func(as []string) map[string]string {
		am := make(map[string]string, len(as))
		for _, a := range as {
			am[a] = strings.Title(strings.ToLower(a)) // e.g. "ID" -> "Id"
		}
		return am
	}(knownAcronyms)

	for k, v := range acronyms {
		s = strings.ReplaceAll(s, k, v)
	}
	var words []string
	start := 0
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			words = append(words, s[start:i])
			start = i
		}
	}
	words = append(words, s[start:])
	return strings.ToLower(strings.Join(words, "_"))
}

var tmplFuncs = template.FuncMap{
	"singularize":  singularize,
	"pluralize":    pluralize,
	"sqlToGoType":  sqlToGoType,
	"goToSqlType":  goToSqlType,
	"snakeToCamel": snakeToCamel,
	"camelToSnake": camelToSnake,
	"tolower":      strings.ToLower,
	"totablename":  func(s string) string { return pluralize(strings.ToLower(s)) },
	"tomodelname":  func(s string) string { return singularize(strings.ToUpper(s)) },
}

var tmpls *template.Template

func init() {
	var err error
	tmpls, err = template.New("").Funcs(tmplFuncs).ParseGlob("templates/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
}

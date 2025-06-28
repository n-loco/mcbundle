package projfiles

import "regexp"

var templateRegExp = regexp.MustCompile(`{{([_a-zA-Z][_a-zA-Z0-9]*)}}`)

var ssCodeRegExp = regexp.MustCompile(`§.`)

type variables struct {
	intern map[string]string
}

func (vars *variables) get(varName string) string {
	value, ok := vars.intern[varName]

	if ok {
		return value
	}

	return "null"
}

func (vars *variables) set(varName string, value string) {
	vars.intern[varName] = value
}

func (vars *variables) apply(s string) string {
	return templateRegExp.ReplaceAllStringFunc(s, func(p string) string {
		return vars.get(p[2 : len(p)-2])
	})
}

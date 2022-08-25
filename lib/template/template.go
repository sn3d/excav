package template

import (
	"html/template"
	"strings"
)

func Subst(body string, params map[string]interface{}) string {
	tmplt, err := template.New("t").Parse(body)
	if err != nil {
		return body
	}

	var b strings.Builder
	err = tmplt.Execute(&b, params)
	if err != nil {
		return body
	}

	return b.String()
}

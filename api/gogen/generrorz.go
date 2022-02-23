package gogen

import (
	"fmt"
)

const errorzTemplate = `package errorz

type Err struct {
	code    string
	message string
}

func New(code string, msgs ...string) error {
	e := &Err{
		code: code,
	}
	if len(msgs) > 0 {
		e.message = msgs[0]
	}
	return e
}

func (e *Err) Error() string {
	return e.code
}

func (e *Err) GetMessage() string {
	return e.message
}

`

func genErrorz(rootPkg string, params map[string]interface{}) error {

	var commonPath string
	if _, ok := params["commonPath"]; !ok || params["commonPath"] == "" {
		commonPath = "../"
	} else {
		commonPath = fmt.Sprintf("%s", params["commonPath"])
	}

	return genFile(fileGenConfig{
		dir:             commonPath,
		subdir:          "/common/errorz",
		filename:        "errorz.go",
		templateName:    "errorzTemplate",
		category:        category,
		templateFile:    "errorz.tpl",
		builtinTemplate: errorzTemplate,
		data:            map[string]interface{}{},
	})
}

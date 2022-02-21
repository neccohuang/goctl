package gogen

import (
	"fmt"
)

const stateTemplate = `package responsex

var (
	SUCCESS           = "0"     //"操作成功"
	FAIL              = "EX000" //"Fail"
	INVALID_PARAMETER = "EX001" //"参数不合法"
)
`

func genState(rootPkg string,params map[string]interface{}) error {

	var rootPath string
	if _,ok:= params["rootPath"]; !ok || params["rootPath"]==""{
		rootPath = "../"
	}else {
		rootPath = fmt.Sprintf("%s", params["rootPath"])
	}

	return genFile(fileGenConfig{
		dir:             rootPath,
		subdir:          "/common/responsex",
		filename:        "state.go",
		templateName:    "stateTemplate",
		category:        category,
		templateFile:    "state.tpl",
		builtinTemplate: stateTemplate,
		data: map[string]interface{}{
		},
	})
}

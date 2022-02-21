package gen

import (
	"github.com/neccohuang/goctl/model/sql/template"
	"github.com/neccohuang/goctl/util"
	"github.com/neccohuang/goctl/util/pathx"
)

func genImports(withCache, timeImport bool, table Table) (string, error) {
	if withCache {
		text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
			"time": timeImport,
			"data": table,
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	}

	text, err := pathx.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCache)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
		"time": timeImport,
		"data": table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

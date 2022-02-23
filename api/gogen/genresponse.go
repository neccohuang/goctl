package gogen

import (
	"fmt"
	"strings"
)

const responseTemplate = `package responsex

import (
	{{.errorz}}
	{{.locales}}
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gioco-play/easy-i18n/i18n"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/language"
	"net/http"
)

type Body struct {
	Code    string      ` + "`" + `json:"code"` + "`" + `
	Message string      ` + "`" + `json:"message"` + "`" + `
	Data    interface{} ` + "`" + `json:"data,omitempty"` + "`" + `
	Trace   string      ` + "`" + `json:"trace"` + "`" + `
}

func Json(w http.ResponseWriter, r *http.Request, code string, resp interface{}, err error) {
	var body Body

	span := trace.SpanFromContext(r.Context())

	i18n.SetLang(language.English)

	body.Code = code
	body.Message = i18n.Sprintf(code)
	if err != nil {
		var msg string
		if v, ok := err.(*errorz.Err); ok && v.Error() != "" {
		    msg = v.Error()
		} else {
            msg = body.Message
		}
		span.RecordError(errors.New(fmt.Sprintf("(%s)%s", code, msg)))
	} else {
		body.Data = resp
	}
	body.Trace = span.SpanContext().TraceID().String()

	if responseBytes, err := json.Marshal(body); err == nil {
		span.SetAttributes(attribute.KeyValue{
			Key:   "response",
			Value: attribute.StringValue(string(responseBytes)),
		})
	}

	httpx.OkJson(w, body)
}
`

func genResponse(rootPkg string, params map[string]interface{}) error {

	path := strings.Split(rootPkg, "/")

	var rootPath string
	if _, ok := params["rootPath"]; !ok || params["rootPath"] == "" {
		rootPath = "../"
	} else {
		rootPath = fmt.Sprintf("%s", params["rootPath"])
	}

	return genFile(fileGenConfig{
		dir:             rootPath,
		subdir:          "/common/responsex",
		filename:        "responsex.go",
		templateName:    "responseTemplate",
		category:        category,
		templateFile:    "response.tpl",
		builtinTemplate: responseTemplate,
		data: map[string]interface{}{
			"errorz":  fmt.Sprintf(`"%s/common/errorz"`, path[0]),
			"locales": fmt.Sprintf(`_ "%s/locales"`, path[0]),
		},
	})
}

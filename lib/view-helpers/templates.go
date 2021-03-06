package viewhelpers

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/golang/glog"
)

// RespondWithTemplate writes an http response to "w" by executing the given template with the params.
func RespondWithTemplate(w http.ResponseWriter, contentType string, tmpl *template.Template, name string, params interface{}) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, name, params); err != nil {
		glog.Errorf("Failed to execute template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if _, err := io.Copy(w, &buf); err != nil {
		glog.Errorf("Failed to send response: %v", err)
		return
	}
}

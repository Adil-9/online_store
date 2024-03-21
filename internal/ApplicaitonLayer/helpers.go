package applicationlayer

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (wh *WebHandler) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	wh.ErrorLog.Output(3, trace)
	// execute template
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (wh *WebHandler) clientError(w http.ResponseWriter, WebPage string, status int, err error) {
	if WebPage != "" && err != nil {
		Err := struct{ Error string }{Error: err.Error()}
		wh.InfoLog.Println(err)

		webErr := wh.templates.ExecuteTemplate(w, WebPage, Err)
		if webErr != nil {
			http.Error(w, http.StatusText(status), status)
		}
		return
	}
	http.Error(w, http.StatusText(status), status)
}

func (wh *WebHandler) notFound(w http.ResponseWriter) {
	//execute template
	wh.clientError(w, "", http.StatusNotFound, nil)
}

func (wh *WebHandler) render(w http.ResponseWriter, status int, webpage string, data any) {
	var buf bytes.Buffer
	err := wh.templates.ExecuteTemplate(&buf, webpage, data)
	if err == nil {
		w.WriteHeader(status)
		buf.WriteTo(w)
	} else {
		wh.ErrorLog.Println(err)
		wh.serverError(w, err)
	}
}

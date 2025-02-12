package controllers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := `
  <!DOCTYPE html>
  <html>
  <head>
    <title>Upload CSV</title>
  </head>
  <body>
    <h1>Service 2 is working fine!</h1>
  </body>
  </html>`
	w.Write([]byte(tmpl))
}

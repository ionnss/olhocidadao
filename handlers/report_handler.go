package handlers

import (
	"fmt"
	"net/http"
)

func SubmitReport(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would parse and validate the form data here.
	// For now, we'll just send a simple success message.

	fmt.Fprintf(w, "<div class=\"alert alert-success\" role=\"alert\">Denúncia enviada com sucesso! Obrigado por sua contribuição.</div>")
}

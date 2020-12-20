package mycustomhandler

import (
	"net/http"

	"GoSocket/userauthentication"
)

func Authentication(w http.ResponseWriter, r *http.Request) {

	var userID = r.Header.Get("UserID")

	if userauthentication.UserValid(userID) {
		userauthentication.StoreSession(userID)
	}
	w.Write([]byte("Secao inicianda"))
}


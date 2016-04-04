package mem

import (
	"errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func getID(res http.ResponseWriter, req *http.Request) (string, error) {
	ctx := appengine.NewContext(req)
	var id, origin string
	var cookie *http.Cookie
	origin = "COOKIE"
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		// try to get the id from the URL
		origin = "URL"
		id := req.FormValue("id")
		if id == "" {
			origin = "BRAND NEW VIA LOGOUT"
			log.Infof(ctx, "ID CAME FROM: %s", origin)
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return id, errors.New("ERROR: redirect to /logout because no session id accessible")
		}
		
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	log.Infof(ctx, "ID CAME FROM %s", origin)
	return id, nil
}
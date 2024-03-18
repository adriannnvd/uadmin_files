package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
		Err         string
		ErrExists   bool
		OTPRequired bool
		Username    string
		Password    string
	}

	c := Context{}

	if r.Method == "POST" {
		//This is a login request from user.
		username := r.PostFormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("password")
		otp := r.PostFormValue("otp")

		//Login2FA login using username, password and otp for users with OTPRequired = true.
		session := uadmin.Login2FA(r, username, password, otp)

		//Check whether the session returned is nil or the user is not active.
		if session == nil || !session.User.Active {

			c.ErrExists = true
			c.Err = "Invalid username/password or inactive user"
		} else {
			if session.PendingOTP {
				uadmin.Trail(uadmin.INFO, "User: %s", session.User.Username, session.User.GetOTP())
			}
			uadmin.Trail(uadmin.DEBUG, "Username and password is invalid")
		}

		uadmin.Trail(uadmin.DEBUG, "Session: %s", session)
		//Display the results here.
		uadmin.Trail(uadmin.DEBUG, "Username: %s", username)
		uadmin.Trail(uadmin.DEBUG, "Password: %s", password)
		uadmin.Trail(uadmin.DEBUG, "OTP: %s", otp)
	}

	uadmin.RenderHTML(w, r, "templates/login.html", c)

}

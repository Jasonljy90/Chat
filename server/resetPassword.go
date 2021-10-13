package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dchest/authcookie"
	"github.com/gorilla/mux"
)

// MinTokenLength is the minimum allowed length of token string.
// It is useful for avoiding DoS attacks with very long tokens: before passing
// a token to VerifyToken function, check that it has length less than [the
// maximum login length allowed in your application] + MinTokenLength.
var (
	MinTokenLength = authcookie.MinLength
)

var (
	ErrMalformedToken = errors.New("malformed token")
	ErrExpiredToken   = errors.New("token expired")
	ErrWrongSignature = errors.New("wrong token signature")
)

func generateSecretKey() []byte {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error Generating Key")
	}
	return key
}

func resetUserPassword(email string) {
	token, err := maker.CreateToken(email, time.Minute*30)
	if err != nil {
		fmt.Println("Error creating token")
		return
	}

	Link := "https://localhost:5221/usertoken/" + token
	sendEmail(email, Link)
}

func resetUserPasswordLinkClicked(res http.ResponseWriter, req *http.Request) {
	// Get token from user once user click the reset password link
	vars := mux.Vars(req)
	tokenStr := vars["token"]

	// Verify whether token is valid
	_, err := maker.VerifyToken(tokenStr)
	if err != nil {
		io.WriteString(res, `
			<html>
			 <meta http-equiv='refresh' content='5; url=/ '/>
			 Password reset link has expired! <br>
			 You will be redirected shortly in 5 seconds...<br>
			</html>
		`)
		return
	}

	// Redirect user to change password page
	http.Redirect(res, req, "/userresetchangepassword", http.StatusSeeOther)
}

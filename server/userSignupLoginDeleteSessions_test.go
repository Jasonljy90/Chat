package main

import (
	"net/http"
	"testing"
	"log"
)

/*
 resp, err := http.Get("https://golangcode.com")
    if err != nil {
        log.Fatal(err)
    }

    // Print the HTTP Status Code and Status Name
    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
    } else {
        fmt.Println("Argh! Broken")
    }
*/

// Working
func TestUserSignup(t *testing.T) {
	t.Run("returns Sign up Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/usersignup")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestUserLogin(t *testing.T) {
	t.Run("returns Login Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/userlogin")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("returns Delete User Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/deleteuser")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestUserChangePassword(t *testing.T) {
	t.Run("returns Change Password Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/userchangepassword")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestUserResetPassword(t *testing.T) {
	t.Run("returns Reset Password Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/userresetpassword")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestUserResetChangePassword(t *testing.T) {
	t.Run("returns Reset Change Password Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/userresetchangepassword")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestLogout(t *testing.T) {
	t.Run("returns Logout Result", func(t *testing.T) {
		response, err := http.Get("http://localhost:5221/logout")
		if err != nil {
		 	log.Fatal(err)
		}

		got := response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}


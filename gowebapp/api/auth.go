package api

import (
	"encoding/json"
	"fmt"
	"gowebapp/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/tago"
)

// AuthRouters are the collection of all URLs for the Auth App.
func AuthRouters(r *mux.Router) {
	r.HandleFunc("/api/v1/user/login", LoginUserEndpoint).Methods("POST")
}

// LoginUserEndpoint is to validate the user's login credential
func LoginUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		itrlog.Error(errBody)
		panic(errBody.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	userName := strings.TrimSpace(keyVal["username"])
	password := keyVal["password"]
	isSiteKeepMe, _ := strconv.ParseBool(keyVal["isSiteKeepMe"])

	fmt.Print("userName: ", userName)
	fmt.Print("password: ", password)
	fmt.Print("isSiteKeepMe: ", isSiteKeepMe)

	// Check if username is empty
	if len(strings.TrimSpace(userName)) == 0 {
		w.Write([]byte(`{ "isSuccess": "false", "alertTitle": "Username is Required BK", "alertMsg": "Please enter your username.", "alertType": "error" }`))
		return
	}

	// Check if password is empty
	if len(strings.TrimSpace(password)) == 0 {
		w.Write([]byte(`{ "isSuccess": "false", "alertTitle": "Password is Required BK", "alertMsg": "Please enter your password.", "alertType": "error" }`))
		return
	}

	// Set the cookie expiry in days.
	expDays := "1" // default to expire in 1 day.
	if isSiteKeepMe == true {
		expDays = config.UserCookieExp
	}

	// Encrypt the username value to store it from the user's cookie.
	encryptedUserName, err := tago.Encrypt(userName, config.MyEncryptDecryptSK)
	if err != nil {
		itrlog.Error(err)
	}

	w.Write([]byte(`{ "isSuccess": "true", "alertTitle": "Login Successful", "alertMsg": "Your account has been verified and it's successfully logged-in.",
		"alertType": "success", "redirectTo": "` + config.SiteBaseURL + `dashboard", "eUsr": "` + encryptedUserName + `", "expDays": "` + expDays + `" }`))
}

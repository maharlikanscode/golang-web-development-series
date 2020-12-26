package api

import (
	"encoding/json"
	"fmt"
	"gowebapp/config"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/tago"
)

// AuthRouters are the collection of all URLs for the Auth App.
func AuthRouters(r *mux.Router) {
	r.HandleFunc("/api/v1/user/login", LoginUserEndpoint).Methods("POST")
	r.HandleFunc("/login", Login).Methods("GET")
}

// Login function is to render the homepage page.
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(config.SiteRootTemplate+"front/login.html", config.SiteHeaderTemplate, config.SiteFooterTemplate))

	data := contextData{
		"PageTitle":    "Login - " + config.SiteShortName,
		"PageMetaDesc": config.SiteShortName + " account, sign in to access your account.",
		"CanonicalURL": r.RequestURI,
		"CsrfToken":    csrf.Token(r),
		"Settings":     config.SiteSettings,
	}
	tmpl.Execute(w, data)
}

type jsonResponse struct {
	IsSuccess  string `json:"isSuccess"`
	AlertTitle string `json:"alertTitle"`
	AlertMsg   string `json:"alertMsg"`
	AlertType  string `json:"alertType"`
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

	itrlog.Info("userName: ", userName)
	itrlog.Info("password: ", password)
	itrlog.Info("isSiteKeepMe: ", isSiteKeepMe)

	// Check if username is empty
	if len(strings.TrimSpace(userName)) == 0 {
		w.Write([]byte(`{ "IsSuccess": "false", "AlertTitle": "Username is Required BK", "AlertMsg": "Please enter your username.", "AlertType": "error" }`))
		return
	}

	// Check if password is empty
	if len(strings.TrimSpace(password)) == 0 {
		w.Write([]byte(`{ "IsSuccess": "false", "AlertTitle": "Password is Required BK", "AlertMsg": "Please enter your password.", "AlertType": "error" }`))
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

	// // Initialize the database connection
	// dbCon, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/carsier?parseTime=true&charset=utf8mb4,utf8")
	// if err != nil {
	// 	itrlog.Error(err)
	// }
	// defer dbCon.Close()

	// // Now, insert the new user's information here
	// ins, err := dbCon.Prepare("INSERT INTO yabi_user (username, password, email, first_name, " +
	// 	"middle_name, last_name, suffix, is_superuser, is_admin, date_joined, is_active) VALUES" +
	// 	"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	// if err != nil {
	// 	itrlog.Error(err)
	// }

	// // Pass on all the parameter values here
	// ins.Exec(userName, password, "politz@live.com", "P O L", "D.", "Peligro", "Jr.", 1, 0, time.Now(), 0)
	// defer ins.Close()

	w.Write([]byte(`{ "isSuccess": "true", "alertTitle": "Login Successful", "alertMsg": "Your account has been verified and it's successfully logged-in.",
		"alertType": "success", "redirectTo": "` + config.SiteBaseURL + `dashboard", "eUsr": "` + encryptedUserName + `", "expDays": "` + expDays + `" }`))
}

package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var API_LINK = "N/A"

func main() {

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	registerNewBasicHandler("templates/index.html", nil, "/", nil)
	registerNewBasicHandler("templates/forum.html", nil, "/forum", nil)
	registerNewBasicHandler("templates/hdv.html", nil, "/hdv", nil)
	registerNewBasicHandler("templates/leaderboard.html", nil, "/leaderboard", nil)
	registerNewBasicHandler("templates/server.html", nil, "/server", nil)
	registerNewBasicHandler("templates/shop.html", nil, "/shop", nil)
	registerNewBasicHandler("templates/wiki.html", nil, "/wiki", nil)
	registerNewBasicHandler("template/account.html", nil, "/account", nil)
	registerNewBasicHandler("templates/faction.html", nil, "/faction", nil)
	registerNewBasicHandler("templates/index.html", nil, "/login", nil)

	go InstallTailwindCSS()

	http.ListenAndServe(":8080", nil)
	println("Bye")
}

func InstallTailwindCSS() {
	println("Installing TailwindCSS...")
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "tailwind.bat")
	} else {
		cmd = exec.Command("/bin/bash", "tailwind.sh")
	}
	err := cmd.Run()
	if err != nil {
		println(err)
	}
	println("Tailwind installed and is built !")
}

func registerNewBasicHandler(page string, data any, link string, process func(http.ResponseWriter, *http.Request, *any)) {
	println("Registering " + link + " to " + page)
	http.HandleFunc(link, func(w http.ResponseWriter, r *http.Request) {

		if process != nil {
			process(w, r, &data)
		}

		tmpl, err := template.ParseFiles(page)

		if err != nil {
			json.NewEncoder(w).Encode("Internal error: " + err.Error())
		}

		tmpl.Execute(w, data)
	})

}

func processPostLogin(w http.ResponseWriter, r *http.Request, data *any) {
	if r.Method != "POST" {
		return
	}

	r.ParseForm()

	login := r.FormValue("input_loginusername")
	password := r.FormValue("input_loginpassword")

	authCooki := new(http.Cookie)
	authCooki.Name = "LOGIN"
	authCooki.Value = login

	if checkLogin(login, password) {
		http.SetCookie(w, authCooki)
	}

}

// BA LINK
func checkLogin(login string, password string) bool {
	db, _ := sql.Open("MYSQL", "root:@tcp(172.18.0.2:3306)/lowcalldata")
	defer db.Close()
	result, _ := db.Query("SELECT password FROM authme WHERE username=?", login)

	if result.Next() {
		pass := new(string)
		result.Scan(pass)

		println("check for " + login)
		return isValidPassword(password, *pass)
	}

	return false
}

func isValidPassword(password string, hash string) bool {
	// $SHA$salt$hash, where hash := sha256(sha256(password) . salt)
	parts := strings.Split(hash, "$")
	return len(parts) == 4 && parts[3] == fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%x", sha256.Sum256([]byte(password)))+parts[2])))
}

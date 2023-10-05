package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
)

var API_LINK = "N/A"

func main() {

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	registerNewBasicHandler("templates/index.html", nil, "/")
	registerNewBasicHandler("templates/forum.html", nil, "/forum")
	registerNewBasicHandler("templates/hdv.html", nil, "/hdv")
	registerNewBasicHandler("templates/leaderboard.html", nil, "/leaderboard")
	registerNewBasicHandler("templates/server.html", nil, "/server")
	registerNewBasicHandler("templates/shop.html", nil, "/shop")
	registerNewBasicHandler("templates/wiki.html", nil, "/wiki")

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

func registerNewBasicHandler(page string, data any, link string) {
	http.HandleFunc(link, func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(page)

		if err != nil {
			json.NewEncoder(w).Encode("Internal error: " + err.Error())
		}

		tmpl.Execute(w, data)
	})

}

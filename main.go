package main

import (
	"embed"
	"io"
	"net/http"
	"text/template"
	"time"

	"github-graph-drawer/log"
	"github-graph-drawer/utils/graphgen"
)

var (
	//go:embed resources/*
	res embed.FS
)

func main() {
	templates := template.Must(template.ParseGlob("./templates/html/*"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
			log.Errorln(err.Error())
			return
		}
	})

	http.HandleFunc("/contribution-graph", func(w http.ResponseWriter, r *http.Request) {
		msg, exists := r.URL.Query()["msg"]
		if !exists {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		font, _ := r.URL.Query()["font"]

		gg := graphgen.NewContributionsGraphGenerator(
			graphgen.HtmlGeneratorType,
			graphgen.ContributionsGraph{}.Init(time.Now().Year()),
		)

		switch font[0] {
		case "3x3":
		case "3x5":
			gg.SetFont(graphgen.Font3x5)
		}

		buf, err := gg.GetFinalForm(msg[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorln(err.Error())
			return
		}

		w.Header().Set("Content-Type", "text/html")
		_, _ = io.Copy(w, buf)
	})

	http.HandleFunc("/generate-script", func(w http.ResponseWriter, r *http.Request) {
		msg, exists := r.URL.Query()["msg"]
		if !exists {
			log.Warningln("someone sent a bad request...")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		font, _ := r.URL.Query()["font"]

		// TODO: add zis to ze generator interface
		commitCount, exists := r.URL.Query()["commit-count"]
		if !exists {
			log.Warningln("someone sent a bad request...")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = commitCount

		gg := graphgen.NewContributionsGraphGenerator(
			graphgen.CheatScriptGeneratorType,
			graphgen.ContributionsGraph{}.Init(time.Now().Year()),
		)

		switch font[0] {
		case "3x3":
		case "3x5":
			gg.SetFont(graphgen.Font3x5)
		}

		buf, err := gg.GetFinalForm(msg[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorln(err.Error())
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		_, _ = io.Copy(w, buf)
	})

	http.HandleFunc("/schedule-emails", func(w http.ResponseWriter, r *http.Request) {
		msg, exists := r.URL.Query()["msg"]
		if !exists {
			log.Warningln("someone sent a bad request...")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = msg

		// TODO: add zis to ze generator interface
		commitCount, exists := r.URL.Query()["commit-count"]
		if !exists {
			log.Warningln("someone sent a bad request...")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = commitCount

		email, exists := r.URL.Query()["email"]
		if !exists {
			log.Warningln("someone sent a bad request...")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = email
		log.Info("email: ", email)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<span style=\"color: white;\">NOT IMPLEMENTED :)</span>"))
	})

	http.Handle("/resources/", http.FileServer(http.FS(res)))

	log.Infoln("server started...")
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":8080", nil))
}

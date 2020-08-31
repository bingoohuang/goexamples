package main

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/markbates/pkger"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// dir := http.FileServer(pkger.Dir("/public"))
	// return http.ListenAndServe(addr, dir)
	addr := ":3000"
	fmt.Println("listening on", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			f, err := pkger.Open("/public/index.html")
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			defer f.Close()

			stat, _ := f.Stat()
			http.ServeContent(w, r, "", stat.ModTime(), f)
		case "/list":
			w := tabwriter.NewWriter(w, 0, 0, 0, ' ', tabwriter.Debug)

			pkger.Walk("/public", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					panic(err)
				}

				fmt.Fprintf(w,
					"%s \t %d \t %s \t %s \t\n",
					info.Name(),
					info.Size(),
					info.Mode(),
					info.ModTime().Format(time.RFC3339),
				)

				return nil
			})

			w.Flush()
		default:
			f, err := pkger.Open("/public" + r.URL.Path)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			defer f.Close()

			stat, _ := f.Stat()
			http.ServeContent(w, r, "", stat.ModTime(), f)
		}
	})

	return http.ListenAndServe(addr, mux)
}

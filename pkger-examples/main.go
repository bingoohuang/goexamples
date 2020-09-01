package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	addr := ":3000"
	fmt.Println("listening on", addr)

	dir := http.FileServer(pkger.Dir("/public"))
	// return http.ListenAndServe(addr, dir)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/":
			ServeStaticFolder(w, r, dir, "/")
		case r.URL.Path == "/static/list":
			// 静态资源列表
			ListStaticFiles(w)
		case strings.HasPrefix(r.URL.Path, "/static"):
			// 以/static开头的，去除/static前缀后，使用dir的FileServer响应
			ServeStaticFolder(w, r, dir, "/static")
		default:
			// 具体单个文件，直接查找静态文件，返回文件内容
			_ = ServeStaticFile(w, r)
		}
	})

	return http.ListenAndServe(addr, mux)
}

func ServeStaticFolder(w http.ResponseWriter, r *http.Request, dir http.Handler, prefix string) {
	switch prefix {
	case "", "/": // ignore
	default:
		dir = http.StripPrefix(prefix, dir)
	}

	dir.ServeHTTP(w, r)
}

func ServeStaticFile(w http.ResponseWriter, r *http.Request) error {
	f, err := pkger.Open("/public" + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}

	http.ServeContent(w, r, "", stat.ModTime(), f)
	return nil
}

func ListStaticFiles(w http.ResponseWriter) {
	tw := tabwriter.NewWriter(w, 0, 0, 0, ' ', tabwriter.Debug)

	pkger.Walk("/public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(tw, "%s \t %d \t %s \t %s \t\n",
			info.Name(), info.Size(), info.Mode(), info.ModTime().Format(time.RFC3339))

		return nil
	})

	tw.Flush()
}

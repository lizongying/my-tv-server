package main

import (
	"flag"
	"fmt"
	"github.com/lizongying/my-tv-server/internal/mytv"
	"github.com/lizongying/my-tv-server/statics"
	"io/fs"
	"log"
	"net/http"
	"strconv"
)

func main() {
	portPtr := flag.String("port", "8000", "")
	flag.Parse()
	port, _ := strconv.Atoi(*portPtr)

	log.Printf("http://127.0.0.1:%d/mytv/index.html\n", port)

	m, _ := mytv.NewMytv(port)

	files, _ := fs.Sub(statics.Mytv, "mytv")
	http.Handle("/mytv/", http.StripPrefix("/mytv/", http.FileServer(http.FS(files))))
	http.HandleFunc("/upload", m.Upload)
	http.HandleFunc("/0", m.Channels)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Println(err)
	}
}

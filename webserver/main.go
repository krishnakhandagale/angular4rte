package webserverangular

import (
	"net/http"
	"html/template"
	"log"

	c "newrte/config"
	//"newrte/externals/gorilla/mux"
	"time"

	"strings"
	"os"
	"strconv"
	"bufio"
)

var (
	CacheSince = time.Now().Format(http.TimeFormat)
	CacheUntil = time.Now().AddDate(0, 0, 1).Format(http.TimeFormat)
)
func Start() {



	//Route handlers for pages
	http.HandleFunc("/landing", func (wr http.ResponseWriter, req *http.Request) {
		log.Println("Landing page request")

		if req.Method == "GET" {
			t, err := template.ParseFiles("blog/dist/index.html")

			if err != nil {
				http.Error(wr, "Error parsing template, ["+err.Error()+"]", http.StatusInternalServerError)
				return
			}

			err = t.Execute(wr, nil)
			if err != nil {
				http.Error(wr, "Error parsing template, ["+err.Error()+"]", http.StatusInternalServerError)
			}
		}
	})

	http.HandleFunc("/", func (wr http.ResponseWriter, req *http.Request) {
	log.Println("In :" + req.URL.Path)
	path :=  "blog/dist" + req.URL.Path
	log.Println(path)
	var contentType string

	// identify the content type of the requested file
	if strings.HasSuffix(path, ".css") {
	contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
	contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
	contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".gif") {
	contentType = "image/gif"
	} else if strings.HasSuffix(path, ".svg") {
	contentType = "image/svg+xml"
	} else if strings.HasSuffix(path, ".js") {
	contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".swf") {
	contentType = "application/x-shockwave-flash"
	} else if strings.HasSuffix(path, ".mp4") {
	contentType = "video/mp4"
	} else if strings.HasSuffix(path, ".pdf") {
	contentType = "application/pdf"
	} else {
	contentType = "text/plain"
	}
	//https://stackoverflow.com/questions/499966/etag-vs-header-expires
	f, err := os.Open(path) // open the file
	if err == nil {
	fileInfo, _ := f.Stat()
	fileModTime := fileInfo.ModTime().String()
	fileSize := fileInfo.Size()
	eTag := fileModTime + strconv.Itoa(int(fileSize))
	defer f.Close()
	wr.Header().Add("Content-Type", contentType) // add the file content type to the http header
	//wr.Header().Set("Last-Modified", CacheSince)
	//wr.Header().Set("Expires", CacheUntil)
	wr.Header().Add("ETag", eTag)

	if match := req.Header.Get("If-None-Match"); match != "" {
	if strings.Contains(match, eTag) {
	wr.WriteHeader(http.StatusNotModified)
	return
	}
	}
	br := bufio.NewReader(f)                     // read the file to a buffer
	br.WriteTo(wr)                               // write the buffer to the client

	} else {
		log.Println(err)
	http.NotFound(wr, req) // return a 404 if there was an error Opening the file
	}

	//log.Println("Out :" + req.URL.Path)
	})


	//rtr.Handle("/node_modules/", http.StripPrefix("/node_modules", http.FileServer(http.Dir("./node_modules"))))
	//rtr.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir("./html"))))
	//rtr.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))
	//rtr.Handle("/ts/", http.StripPrefix("/ts", http.FileServer(http.Dir("./ts"))))
	//rtr.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	//root now holds all the above defined handlers\

	log.Printf("EvilCorp: %s\n", c.EVILCORP_DIR)
	log.Printf("Mediacluster: %s\n", c.MEDIACLUSTER_DIR)
	log.Println("Listening to port..." + c.WEBPORT)
	log.Fatal(http.ListenAndServe(c.WEBPORT, nil))

}

package main

// starts importing
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"

	"golang.org/x/net/context"
)

//stops importing

type data struct {
	Token    string `json : "token"`
	userName string `json:"user_name"`
	age      int    `json:"age"`
	gender   string `json: gender`
}

func init() {
	http.HandleFunc("/", handler)


}

type app struct {
	Router  *mux.Router
	ctx     context.Context
	client  *storage.Client
	storage *storage.Bucket
	w   io.Writer
	failed bool
}
func (a *app) errorf(format string, args ...interface{}) {
	a.failed = true
	fmt.Fprintln(a.w, fmt.Sprintf(format, args...))
	log.Errorf(a.ctx, format, args...)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
func main() {
	//let us declear an array

}

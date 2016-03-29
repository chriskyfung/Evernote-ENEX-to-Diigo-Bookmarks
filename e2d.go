package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
)

// EnExport is en-export tag
type EnExport struct {
	Note []Note `xml:"note"`
}

// Note is note tag
type Note struct {
	Title  string `xml:"title"`
	Source string `xml:"note-attributes>source-url"`
	Tags []string `xml:"tag"`
	Tag string
}

func convert(r io.Reader) error {
	e := EnExport{}
	if err := xml.NewDecoder(r).Decode(&e); err != nil {
		return err
	}

	// add a dummy url if source-url is empty
	// join tags with comma
	// call the post function with passing url, title, and tags
	dummy := "https://getpocket.com/"
	for i, v := range e.Note {
		if v.Source == "" {
			e.Note[i].Source = dummy + fmt.Sprint(i)
		}
		if len(v.Tags) > 1 {
			e.Note[i].Tag = "inEvernote," + strings.Join(v.Tags, ",")
		}
		post("&url="+e.Note[i].Source+"&title="+v.Title+"&tags="+e.Note[i].Tag) 
	}	
	
	return nil
}

func post(p string) string {
    client := &http.Client{}
	var key string = "<your Diigo API key>"
	var username string = "<your Diigo username>"
	var passwd string = "<your Diigo password>"
	
	p = strings.Replace(p, " ", "+", -1)
	
    URL := "https://secure.diigo.com/api/v2/bookmarks?key="+ key + p
	fmt.Println(p)
	
    //pass the values to the request's body	
    req, err := http.NewRequest("POST", URL, nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
		fmt.Println(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
	fmt.Println(s)
    return s
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid arguments")
	}

	in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()


	if err := convert(in); err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"flag"
	"text/template"

	"gopkg.in/yaml.v2"
)

var f = flag.String("template", "/input/template", "The go template to use.")
var d = flag.String("data", "/input/data", "The YAML file to read configuration from.")
var o = flag.String("output", "", "This is the file that the output will be writen to. If not set, will print to stdout")

// ExecuteTemplate asdf
func ExecuteTemplate(b []byte, o io.Writer, t ...string) error {
	tpl, err := template.ParseFiles(t...)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err)
	}

	var values map[string]interface{}
	err = yaml.Unmarshal(b, &values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}

	err = tpl.Execute(o, values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()

	var out *os.File
	var err error
	if *o == "" {
	  out = os.Stdout
	} else {
		out, err = os.Create(*o)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	d, err := ioutil.ReadFile(*d)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = ExecuteTemplate(d, out, *f)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

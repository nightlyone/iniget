package main

import "os"
import "io"
import "log"
import "flag"
import "code.google.com/p/goconf/conf"

var section = flag.String("section", "default", "section of ini file we look into")
var filename = flag.String("filename", "", "filename of ini file we look into")
var key = flag.String("key", "", "key of ini file we look into")

func main() {
	flag.Parse()
	if len(*key) == 0 || len(*filename) == 0 || len(*section) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	config, err := conf.ReadConfigFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	value, err := config.GetString(*section, *key)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(os.Stdout, value)
	if err != nil {
		log.Fatal(err)
	}
}

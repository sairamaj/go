package main

import(
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"flag"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){

	json_file_option := flag.String("json-file", "", "json file")
	path_option := flag.String("path", "", "path to extract")

	flag.Parse()
	json_file := *json_file_option
	path := *path_option

	if len(json_file) == 0 {
		flag.Usage()
		return
	}

	if len(path) == 0 {
		flag.Usage()
		return
	}

	dat, err := ioutil.ReadFile(json_file)
	check(err)

	value := gjson.Get(string(dat), path)
	log.Printf("%s ->|%s|", path, value)
}
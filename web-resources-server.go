package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

var config *yamlconfig

func main() {
	config = loadConfig("config.yaml")
	serveDirectory()
	log.Println(config.Directory, config.Port)
}

/*Loading the YAML config into the struct*/
type yamlconfig struct {
	Port      string `yaml:"port"`
	Directory string `yaml:"directory"`
	Path      string `yaml:"path"`
}

func loadConfig(path string) *yamlconfig {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln("cannot load config: " + err.Error())
	}

	var conf yamlconfig
	err = yaml.Unmarshal(b, &conf)

	if err != nil {
		log.Fatalln("cannot unmarshal config: " + err.Error())
	}

	return &conf
}

/*Serving the content of the directory*/
func serveDirectory() {

	log.Println(config.Directory)
	fsys := http.Dir(config.Directory)
	fs := http.FileServer(fsys)
	address := ":" + config.Port
	prefix := config.Path[0 : len(config.Path)-1]
	log.Println(prefix)

	http.Handle(config.Path, http.StripPrefix(prefix, fs))
	log.Fatalln(http.ListenAndServe(address, nil))
}

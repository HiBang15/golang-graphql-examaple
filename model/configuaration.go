package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

func (env *Environment) SetEnvironment() {
	_, filename, _, _ := runtime.Caller(1)

	//return path file config
	env.path = path.Join(path.Dir(filename), "environment/Connection.yml")

	//returns a FileInfo describing the named file.
	_, err := os.Stat(env.path)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (env *Environment) LoadConfig() {
	content, err := ioutil.ReadFile(env.path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(string(content)), env)
	if err != nil {
		log.Fatal(err)
	}

	if env.App.Debug == false {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Config load successfully!")
	return
}

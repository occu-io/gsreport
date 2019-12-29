package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

func IniParse(config_path string) *ini.File {
	log.Printf("Parsing file: %s", config_path)
	cfg, err := ini.Load(config_path)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

package config

import (
	"encoding/json"
	"fmt"
	"github.com/MathisBurger/yb-http/Var"
	"github.com/MathisBurger/yb-http/models"
	"os"
	"path/filepath"
)

func LoadConfigurations() {
	var files []string

	root := "./config/http/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	first := false
	for _, file := range files {
		if first {
			status, config := checkSyntax(file)
			if status {
				Var.AppendConfig(config)
			} else {
				fmt.Println("The config", file, "dies not have the correct sytax")
			}
		} else {
			first = true
		}
	}
}

func checkSyntax(filename string) (b bool, c *models.HttpConfig) {
	f, err := os.Open(filename)
	if err != nil {
		return false, nil
	}
	c = new(models.HttpConfig)
	err = json.NewDecoder(f).Decode(c)
	return true, c
}

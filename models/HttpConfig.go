package models

type HttpConfig struct {
	Server struct {
		Name         string `json:"name"`
		DocumentRoot string `json:"documentroot"`
		EntryPoint   string `json:"entrypoint"`
	} `json:"server"`
}

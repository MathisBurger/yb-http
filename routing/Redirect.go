package routing

import (
	"github.com/MathisBurger/yb-http/Var"
	"github.com/MathisBurger/yb-http/models"
	"github.com/MathisBurger/yb-http/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"strings"
)

func Redirect(c *fiber.Ctx) error {
	url := c.BaseURL()
	split := strings.Split(url, "://")
	cfg := Var.GetConfig(split[1])
	if cfg == nil {
		return c.SendString("Server is not configured correctly")
	}
	href := c.OriginalURL()
	root := autocomplete(href, cfg)
	return c.SendFile(root)
}

func autocomplete(url string, cfg *models.HttpConfig) string {
	requested := cfg.Server.DocumentRoot + url
	isDir, err := utils.IsDirectory(requested)
	if err != nil {
		return cfg.Server.DocumentRoot + "/" + cfg.Server.EntryPoint
	}
	if isDir {
		return scanDirForIndexFiles(cfg.Server.DocumentRoot+url, url)
	} else {
		return cfg.Server.DocumentRoot + url
	}
}

func scanDirForIndexFiles(root string, url string) string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	var subFiles []string
	for _, file := range files {
		if strings.Contains(file, root) {
			subFiles = append(subFiles, file)
		}
	}
	s := []rune(url)
	if string(s[len(s)-1:]) == "/" {
		return root + "index.html"
	} else {
		return root + "/index.html"
	}
}

package app_config

import (
	"log"
	"path"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
)

var Data = koanf.New(".")

func init() {
	basepath := "../../"
	err := Data.Load(file.Provider(path.Join(basepath, ".env")), dotenv.Parser())
	if err != nil {
		log.Fatalf("error loading config : %v", err)
	}
}

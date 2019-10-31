package codegen

import (
	"github.com/solo-io/autopilot/codegen/util"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Run(dir string, forceOverwrite bool) error {
	config := filepath.Join(dir, "autopilot.yaml")

	project, err := Load(config)
	if err != nil {
		return err
	}

	files, err := Generate(project)
	if err != nil {
		return err
	}

	if err := util.DeepcopyGen(project.TypesImportPath); err != nil {
		return err
	}

	for _, file := range files {
		name := filepath.Join(os.Getenv("GOPATH"), "src", file.OutPath)
		content := file.Content

		if !forceOverwrite && file.SkipOverwrite {
			if _, err := os.Stat(name); err == nil {
				log.Printf("skippinng file %v because it exists", name)
				continue
			}
		}

		if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
			return err
		}
		if err := ioutil.WriteFile(name, []byte(content), 0644); err != nil {
			return err
		}

		formatted, err := imports.Process(name, []byte(content), nil)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(name, []byte(formatted), 0644); err != nil {
			return err
		}
	}

	log.Printf("Finished generating %v", project.ApiVersion+"."+project.Kind)

	return nil
}

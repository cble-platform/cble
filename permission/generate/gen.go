//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

var templatesToExecute = map[string]string{
	"actions_gen.graphqls.tmpl": "actions/actions_gen.graphqls",
	"actions_gen.go.tmpl":       "actions/actions_gen.go",
	"permissions_gen.go.tmpl":   "permissions_gen.go",
}

// see Go source code:
// https://github.com/golang/go/blob/f57ebed35132d02e5cf016f324853217fb545e91/src/cmd/go/internal/modload/init.go#L1283
func findModuleRoot(dir string) (roots string) {
	if dir == "" {
		panic("dir not set")
	}
	dir = filepath.Clean(dir)

	// Look for enclosing go.mod.
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir
		}
		d := filepath.Dir(dir)
		if d == dir { // the parent of the root is itself, so we can go no further
			break
		}
		dir = d
	}
	return ""
}

var templateFuncs = template.FuncMap{
	"Title":   strcase.ToCamel,
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"TrimCommentPrefix": func(s string) string {
		return strings.TrimLeft(s, "# ")
	},
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Read in the actions.yml (from root dir)
	actionsFilePath := path.Join(findModuleRoot(cwd), "actions.yml")
	actionsFileBytes, err := os.ReadFile(actionsFilePath)
	if err != nil {
		panic(fmt.Errorf("cannot read actions.yml file: %v", err))
	}

	// Marshal the file as a map for the templates
	var actionsMap map[string]yaml.Node
	err = yaml.Unmarshal(actionsFileBytes, &actionsMap)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal actions.yml file: %v", err))
	}

	for templatePath, outputPath := range templatesToExecute {
		// Parse in the template file
		templateName := path.Base(templatePath)
		templateAbsPath := path.Join(cwd, "generate", templatePath)
		t, err := template.New(templateName).Funcs(templateFuncs).ParseFiles(templateAbsPath)
		if err != nil {
			panic(fmt.Errorf("failed to parse %s: %v", templatePath, err))
		}

		// Execute the template into memory
		var outputBuffer bytes.Buffer
		err = t.Execute(&outputBuffer, actionsMap)
		if err != nil {
			panic(fmt.Errorf("failed to execute template: %v", err))
		}

		var outputBytes []byte
		// If go file, format the output before writing to file
		if path.Ext(outputPath) == ".go" {
			outputBytes, err = format.Source(outputBuffer.Bytes())
			if err != nil {
				panic(fmt.Errorf("failed to execute gofmt on output: %v", err))
			}
		} else {
			outputBytes = outputBuffer.Bytes()
		}

		// Create the output file (or truncate it)
		outputAbsPath := path.Join(cwd, outputPath)
		outputFile, err := os.Create(outputAbsPath)
		if err != nil {
			panic(fmt.Errorf("failed to create output file %s: %v", outputPath, err))
		}

		// Write the output
		_, err = outputFile.Write(outputBytes)
		if err != nil {
			panic(fmt.Errorf("failed to write output to %s: %v", outputPath, err))
		}

		// Close the output file
		outputFile.Close()
	}
}

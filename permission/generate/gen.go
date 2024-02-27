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

	// Parse in the actions_gen template file
	actionsTemplatePath := path.Join(cwd, "generate", "actions_gen.go.tmpl")
	t, err := template.New("actions_gen.go.tmpl").Funcs(templateFuncs).ParseFiles(actionsTemplatePath)
	if err != nil {
		panic(fmt.Errorf("failed to parse generate/actions_gen.go.tmpl: %v", err))
	}

	// Execute the template into memory
	var actionsOutputBuffer bytes.Buffer
	err = t.Execute(&actionsOutputBuffer, actionsMap)
	if err != nil {
		panic(fmt.Errorf("failed to execute actions template: %v", err))
	}

	// Format the output before writing to file
	actionsOutputBytes, err := format.Source(actionsOutputBuffer.Bytes())
	if err != nil {
		panic(fmt.Errorf("failed to execute gofmt on output: %v", err))
	}

	// Create the output file (or truncate it)
	actionsOutputPath := path.Join(cwd, "actions", "actions_gen.go")
	actionsOutputFile, err := os.Create(actionsOutputPath)
	if err != nil {
		panic(fmt.Errorf("failed to create output file actions_gen.go: %v", err))
	}
	defer actionsOutputFile.Close()

	// Write the output
	_, err = actionsOutputFile.Write(actionsOutputBytes)
	if err != nil {
		panic(fmt.Errorf("failed to write output to actions_gen.go: %v", err))
	}

	// Parse in the permissions_gen template file
	permissionsTemplatePath := path.Join(cwd, "generate", "permissions_gen.go.tmpl")
	t, err = template.New("permissions_gen.go.tmpl").Funcs(templateFuncs).ParseFiles(permissionsTemplatePath)
	if err != nil {
		panic(fmt.Errorf("failed to parse generate/permissions_gen.go.tmpl: %v", err))
	}

	// Execute the template into memory
	var permissionsOutputBuffer bytes.Buffer
	err = t.Execute(&permissionsOutputBuffer, actionsMap)
	if err != nil {
		panic(fmt.Errorf("failed to execute permissions template: %v", err))
	}

	// Format the output before writing to file
	permissionsOutputBytes, err := format.Source(permissionsOutputBuffer.Bytes())
	if err != nil {
		panic(fmt.Errorf("failed to execute gofmt on output: %v", err))
	}

	// Create the output file (or truncate it)
	permissionsOutputPath := path.Join(cwd, "permissions_gen.go")
	permissionsOutputFile, err := os.Create(permissionsOutputPath)
	if err != nil {
		panic(fmt.Errorf("failed to create output file permissions_gen.go: %v", err))
	}
	defer actionsOutputFile.Close()

	// Write the output
	_, err = permissionsOutputFile.Write(permissionsOutputBytes)
	if err != nil {
		panic(fmt.Errorf("failed to write output to permissions_gen.go: %v", err))
	}
}

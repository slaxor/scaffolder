package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type Attribute struct {
	Name string
	Type string
}

type Target struct {
	Template string
	Filename string
}

type Model struct {
	Name       string
	ProjectDir string
	Attributes []Attribute
	Targets    []Target
}

func (m *Model) GenerateTargetFiles() error {
	for _, t := range m.Targets {
		var err error
		var buf bytes.Buffer
		err = tmpls.ExecuteTemplate(&buf, t.Template, m)
		if err != nil {
			return err
		}

		err = os.MkdirAll(path.Dir(t.Filename), 0755)
		if err != nil {
			return err
		}
		targetFile, err := os.Create(t.Filename)
		if err != nil {
			return err
		}
		defer targetFile.Close()
		fmt.Fprintf(targetFile, buf.String())
	}
	return nil
}

// func (m *Model) GenerateModelTestFile() error {
// return nil
// }

// func (m *Model) GenerateCRUDRoutes() error {
// Write the route and handler functions to create, read, update, and delete instances of the model
// ...
//
// return nil
// }

func (m *Model) GenerateViews() error {
	// Create the directory for the views
	err := os.MkdirAll(fmt.Sprintf("%s/views/%s", m.ProjectDir, strings.ToLower(m.Name)), 0755)
	if err != nil {
		return err
	}

	// Write the HTML templates for the create, update, and list views
	createFile, err := os.Create(fmt.Sprintf("%s/views/%s/create.html", m.ProjectDir, strings.ToLower(m.Name)))
	if err != nil {
		return err
	}
	defer createFile.Close()

	updateFile, err := os.Create(fmt.Sprintf("%s/views/%s/update.html", m.ProjectDir, strings.ToLower(m.Name)))
	if err != nil {
		return err
	}
	defer updateFile.Close()

	listFile, err := os.Create(fmt.Sprintf("%s/views/%s/list.html", m.ProjectDir, strings.ToLower(m.Name)))
	if err != nil {
		return err
	}
	defer listFile.Close()

	// Write the HTML templates using Go's template package
	// ...

	return nil
}

func generateScaffold(projectDir, modelName string, attributes []string) error {
	var err error
	var attrs []Attribute
	migVer := time.Now().Format("20060102150405")
	// Parse the attributes and create a Model struct
	for _, attr := range attributes {
		parts := strings.Split(attr, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid attribute: %s", attr)
		}
		attrs = append(attrs, Attribute{Name: parts[0], Type: parts[1]})
	}
	m := Model{
		Name:       modelName,
		ProjectDir: projectDir,
		Attributes: attrs,
		Targets: []Target{
			{
				Template: "migrate.down.sql.tmpl",
				Filename: fmt.Sprintf(
					"%s/migrations/%s_create_%s.down.sql",
					projectDir,
					migVer,
					strings.ToLower(modelName),
				),
			},
			{
				Template: "migrate.up.sql.tmpl",
				Filename: fmt.Sprintf(
					"%s/migrations/%s_create_%s.up.sql",
					projectDir,
					migVer,
					strings.ToLower(modelName),
				),
			},
			{
				Template: "model.go.tmpl",
				Filename: fmt.Sprintf("%s/models/%s.go", projectDir, strings.ToLower(modelName)),
			},
			{
				Template: "model_test.go.tmpl",
				Filename: fmt.Sprintf("%s/models/%s.go", projectDir, strings.ToLower(modelName)),
			},
			{
				Template: "render.go.tmpl",
				Filename: fmt.Sprintf("%s/views/render_%s.go", projectDir, strings.ToLower(modelName)),
			},
		},
	}

	err = os.MkdirAll(projectDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	err = m.GenerateTargetFiles()
	if err != nil {
		return err
	}

	// Generate the files and code for the model
	// err = m.GenerateModelFile()
	// if err != nil {
	// return err
	// }

	// err = m.GenerateMigrationFile()
	// if err != nil {
	// return err
	// }

	// err = m.GenerateCRUDRoutes()
	// if err != nil {
	// return err
	// }

	err = m.GenerateViews()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var err error
	// Define the command-line interface using the flag package
	var projectDir string
	var modelName string
	var attributes string
	flag.StringVar(&projectDir, "dir", ".", "The target folder")
	flag.StringVar(&modelName, "name", "", "The name of the model")
	flag.StringVar(&attributes, "attributes", "", "A comma-separated list of attributes in the format name:type")
	flag.Parse()

	// Define the command-line interface using the flag package
	// var modelName string
	// var attributes stringList
	// flag.StringVar(&modelName, "name", "", "The name of the model")
	// flag.Var(&attributes, "attributes", "A comma-separated list of attributes in the format name:type")
	// flag.Parse()

	// Generate the scaffold for the model
	// err := generateScaffold(modelName, attributes)
	err = generateScaffold(projectDir, modelName, strings.Split(attributes, ","))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// This code defines a `Model` struct and an `Attribute` struct to
// represent the data for the scaffold generator. It also defines
// functions to generate a Go file, a database migration file, CRUD
// routes and handlers, and HTML templates for the model. Finally, it
// defines a `generateScaffold` function to combine these functions into
// a single command to generate the scaffold for the model.

// To use this tool, you can run a command like this:
// go run main.go --name=Product --attributes=name:string,price:float,quantity:int

// This command will generate a Go file, a database migration file, CRUD
// routes and handlers, and HTML templates for a `Product` model with
// three attributes: `name` (a string), `price` (a float), and
// `quantity` (an integer). The generated files will be saved to the
// appropriate locations in the project directory.

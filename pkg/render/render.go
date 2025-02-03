package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/krishmagar/go-api/pkg/config"
	"github.com/krishmagar/go-api/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// This function sets the "var app *config.AppConfig"
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(tempData *models.TemplateData) *models.TemplateData {
	return tempData
}

// This function renders templates using html/template
func RenderTemplate(w http.ResponseWriter, html string, tempData *models.TemplateData) {
	var tempCache map[string]*template.Template

	if app.UseCache {
		tempCache = app.TemplateCache
	} else {
		tempCache, _ = CreateTemplateCache()
	}

	temp, ok := tempCache[html]
	if !ok {
		fmt.Println("Could not get template from template cache")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer) // Creates a new buffer variable

	tempData = AddDefaultData(tempData)

	err := temp.Execute(buf, tempData) // Execute the template and write output to the buffer
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w) // Write the buffer content to the response writer
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

	// // Template Parsing means to analyzing & identifying
	// // the placeholders in the data to be replaced.
	// parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	// // Execute the parsed template
	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error parsing template:", err)
	// 	return
	// }
}

// This function creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Find every file ending with .page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Find all layout templates
	layouts, err := filepath.Glob("./templates/*.layout.html")
	if err != nil {
		log.Println("Error finding layout templates:", err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Processing Template:", name) // about.page.html, home.page.html

		// Parse the page template
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// If layout templates exist, parse them
		if len(layouts) > 0 {
			// Add the parsed layout templates to the existing parsed page
			// template instance (ts); Fully parsed template (page + layouts)
			ts, err = ts.ParseFiles(layouts...)
			if err != nil {
				log.Println("Error parsing layout templates:", err)
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

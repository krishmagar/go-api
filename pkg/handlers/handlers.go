package handlers

import (
	"net/http"

	"github.com/krishmagar/go-api/pkg/config"
	"github.com/krishmagar/go-api/pkg/models"
	"github.com/krishmagar/go-api/pkg/render"
)

// It is the repo variable used by the handlers
var Repo *Repository

// This function sets the value of Repo variable
func NewHandlers(r *Repository) {
	Repo = r
}

// It is the repository structure
type Repository struct {
	App *config.AppConfig
}

// It creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Business Logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	//

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	values, err := divideValues(4, 0)
// 	if err != nil {
// 		fmt.Fprint(w, err)
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("The divide values are: %d", values))
// }
//
// func divideValues(x, y int) (int, error) {
// 	if y <= 0 {
// 		err := errors.New("Cannot divide value eq or less than 0")
// 		return 0, err
// 	}
// 	return x / y, nil
// }

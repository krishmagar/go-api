package handlers

import (
	"github.com/krishmagar/go-api/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
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

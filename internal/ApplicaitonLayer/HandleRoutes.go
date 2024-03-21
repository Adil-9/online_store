package applicationlayer

import (
	"errors"
	"net/http"

	servicelayer "github.com/Adil-9/online_store/internal/ServiceLayer"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (wh *WebHandler) HandleRoutes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./templates/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.Handler(http.MethodPost, "/sign-up", wh.sessionManager.LoadAndSave(http.HandlerFunc(wh.signUpPost)))
	router.Handler(http.MethodGet, "/sign-up", wh.sessionManager.LoadAndSave(http.HandlerFunc(wh.signUp)))
	router.Handler(http.MethodPost, "/login", wh.sessionManager.LoadAndSave(http.HandlerFunc(wh.loginPost)))
	router.Handler(http.MethodGet, "/login", wh.sessionManager.LoadAndSave(http.HandlerFunc(wh.login)))

	standard := alice.New(wh.recoverPanic, wh.logRequest, secureHeaders)

	return standard.Then(router)
}

func (wh *WebHandler) loginPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wh.clientError(w, "", http.StatusMethodNotAllowed, nil)
	}
	err := r.ParseForm()
	if err != nil {
		wh.serverError(w, err)
		return
	}
	name, password := r.PostForm.Get("username"), r.PostForm.Get("password")
	err = wh.services.CheckUserLogin(name, password)
	if err != nil {
		if errors.Is(err, servicelayer.ErrInvalidNameOrPassword) {
			wh.clientError(w, "Login", http.StatusBadRequest, err)
			return
		}
	}

	wh.sessionManager.Put(r.Context(), "flash", "User successfully logged in")

	flash := wh.sessionManager.PopString(r.Context(), "flash")

	wh.render(w, http.StatusOK, "Login", struct {
		Flash string
		Error error
	}{Flash: flash})
	//give tokens
	//redirect to home page

}

func (wh *WebHandler) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		wh.clientError(w, "", http.StatusMethodNotAllowed, nil)
	}
	wh.render(w, http.StatusOK, "Login", nil)
}

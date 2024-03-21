package applicationlayer

import (
	"errors"
	"net/http"

	servicelayer "github.com/Adil-9/online_store/internal/ServiceLayer"
	structures "github.com/Adil-9/online_store/internal/Structures"
)

func (wh *WebHandler) signUpPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wh.clientError(w, "", http.StatusMethodNotAllowed, nil)
	}
	err := r.ParseForm()
	if err != nil {
		wh.ErrorLog.Println(err)
		wh.serverError(w, err)
		return
	}
	user := structures.User{
		Name:       r.PostForm.Get("username"),
		Email:      r.PostForm.Get("email"),
		Password:   r.PostForm.Get("password"),
		RePassword: r.PostForm.Get("confirm-password"),
	}

	err = wh.services.CheckUserSignUp(user)
	if errors.Is(err, servicelayer.ErrUserExists) || errors.Is(err, servicelayer.ErrInvalidPassword) || errors.Is(err, servicelayer.ErrInvalidNameOrEmail) {
		wh.InfoLog.Println(err, r.RemoteAddr)
		user.Error = err.Error() //graceful error
		wh.render(w, http.StatusUnprocessableEntity, "Sign-Up", user)
		return
	} else if err != nil {
		wh.ErrorLog.Println(err, r.RemoteAddr)
		wh.clientError(w, "", http.StatusInternalServerError, nil)
		return
	}
	wh.render(w, http.StatusOK, "Sign-Up", nil)
}

func (wh *WebHandler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		wh.clientError(w, "", http.StatusMethodNotAllowed, nil)
	}
	wh.render(w, http.StatusOK, "Sign-Up", nil)
}

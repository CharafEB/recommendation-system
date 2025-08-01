package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (user *Control) Moul() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	//r.Mount("/", user.Admin(r))

	return r

}

func (app *Application) Run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.Address,
		Handler:      mux,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server is live in Port %s", app.Address)
		return err
	}
	return nil
}

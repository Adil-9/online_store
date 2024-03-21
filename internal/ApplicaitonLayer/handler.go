package applicationlayer

import (
	"database/sql"
	"log"
	"text/template"
	"time"

	dblayer "github.com/Adil-9/online_store/internal/DBlayer"
	servicelayer "github.com/Adil-9/online_store/internal/ServiceLayer"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type WebHandler struct {
	templates      *template.Template
	services       *servicelayer.ServiceHandler
	ErrorLog       *log.Logger
	InfoLog        *log.Logger
	sessionManager *scs.SessionManager
}

func (wh *WebHandler) MustLoadHadnler(db *sql.DB) {
	//initiate template
	wh.templates, _ = template.ParseGlob("templates/webpages/*.html")

	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	wh.sessionManager = sessionManager

	//service and database layer
	DBhandler := dblayer.NewDBH(db)
	services := servicelayer.NewSVH(DBhandler)

	wh.services = services
}

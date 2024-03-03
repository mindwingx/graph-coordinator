package driver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mindwingx/graph-coordinator/app/handler"
	"github.com/mindwingx/graph-coordinator/driver/abstractions"
	"github.com/mindwingx/graph-coordinator/driver/middleware"
	"net/http"
	"time"
)

type (
	Mux struct {
		router     *mux.Router
		handler    *http.Server
		conf       routerConfig
		workerPool chan struct{}
	}

	routerConfig struct {
		Host         string `mapstructure:"HOST"`
		Port         string `mapstructure:"PORT"`
		WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
		ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	}
)

func InitRouter(registry abstractions.RegAbstraction) abstractions.RouterAbstraction {
	m := new(Mux)
	registry.Parse(&m.conf)
	m.router = mux.NewRouter()

	return m
}

func (mux *Mux) Routes() {
	// API routes
	api := mux.router.PathPrefix("/api").Subrouter()
	api.Use(middleware.MsgMiddleware)
	api.HandleFunc("/send", func(rw http.ResponseWriter, r *http.Request) {
		handler.SendHandler(rw, r)
	}).Methods(http.MethodPost)

	mux.handler = &http.Server{
		Handler:      mux.router,
		Addr:         fmt.Sprintf("%s:%s", mux.conf.Host, mux.conf.Port),
		WriteTimeout: time.Duration(mux.conf.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(mux.conf.ReadTimeout) * time.Second,
	}
}

func (mux *Mux) Serve() error {
	return mux.handler.ListenAndServe()
}

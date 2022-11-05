package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/lucianthorr/goProject/db"
)

type Config struct {
	Hostname string        `yaml:"hostname"`
	Port     int           `yaml:"port"`
	Timeout  time.Duration `yaml:"timeout"`
}

type Client interface {
	// set public methods
	Run() error
}

func New(cfg *Config, dbCli db.Client) Client {
	c := &client{
		cfg:   cfg,
		dbCli: dbCli,
	}
	c.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	return c
}

type client struct {
	cfg    *Config
	dbCli  db.Client
	server *http.Server
}

func (c *client) Run() error {
	http.HandleFunc("/", makeGetThing(c.dbCli))
	return c.server.ListenAndServe()
}

// The following route just shows how the API uses its DB dependency and then does some post-processing on the result
// It expects something like "/xyz" where "xyz" is an ID that is looked up in the DB
func makeGetThing(dbCli db.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := strings.TrimPrefix(req.URL.Path, "/")
		tokens := strings.Split(path, "/")
		if len(tokens) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Missing Thing ID in path")
			return
		}
		if len(tokens) > 1 {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Expects one Thing ID as path parameter")
			return
		}
		thingID := tokens[0]
		result, err := dbCli.GetThing(thingID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}
		processedResult := result + "_processed"
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, processedResult)
	}
}

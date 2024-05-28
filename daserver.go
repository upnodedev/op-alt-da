package plasma

import (
	"fmt"
	"golang.org/x/exp/slog"
	"net/http"
	"plasma/config"
	"plasma/da"
)

type DAServer struct {
	logger *slog.Logger
	config config.App
	store  da.KVStore
}

func NewDAServer(cfgApp config.App, store da.KVStore, logger *slog.Logger) *DAServer {
	s := &DAServer{
		logger: logger,
		config: cfgApp,
		store:  store,
	}

	http.HandleFunc("/get", s.HandleGet)
	http.HandleFunc("/put", s.HandlePut)

	return s
}

func (d *DAServer) Start() {
	port := fmt.Sprintf(":%s", d.config.HttpPort)
	d.logger.Info("starting server", "port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		d.logger.Error("server start failed", "err", err)
		panic(err)
	}
}

func (d *DAServer) HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get\n"))
}

func (d *DAServer) HandlePut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put\n"))
}

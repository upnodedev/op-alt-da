package plasma

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"path"
	"plasma/common"
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

	http.HandleFunc("/get/", s.HandleGet)
	http.HandleFunc("/put/", s.HandlePut)

	return s
}

func (d *DAServer) Start() {
	port := fmt.Sprintf(":%d", d.config.Port)
	d.logger.Info("starting server", "port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		d.logger.Error("server start failed", "err", err)
		panic(err)
	}
}

func (d *DAServer) HandleGet(w http.ResponseWriter, r *http.Request) {
	route := path.Dir(r.URL.Path)
	if route != "/get" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := path.Base(r.URL.Path)
	comm, err := hexutil.Decode(key)
	if err != nil {
		d.logger.Error("Failed to decode commitment", "err", err, "key", key)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input, err := d.store.Get(r.Context(), comm)
	if err != nil && errors.Is(err, common.ErrNotFound) {
		d.logger.Error("Commitment not found", "key", key, "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		d.logger.Error("Failed to read commitment", "err", err, "key", key)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(input); err != nil {
		d.logger.Error("Failed to write pre-image", "err", err, "key", key)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *DAServer) HandlePut(w http.ResponseWriter, r *http.Request) {
	d.logger.Info("handling put request", "url", r.URL.Path)

	route := path.Dir(r.URL.Path)
	d.logger.Info("route", "route", route)
	if route != "/put" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input, err := io.ReadAll(r.Body)
	if err != nil {
		d.logger.Error("Failed to read request body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.URL.Path == "/put" || r.URL.Path == "/put/" { // without commitment
		comm := NewKeccak256Commitment(input).Encode()

		if err = d.store.Put(r.Context(), comm, input); err != nil {
			d.logger.Error("Failed to store commitment to the DA server", "err", err, "comm", comm)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		d.logger.Info("stored commitment", "key", hex.EncodeToString(comm), "input_len", len(input))

		if _, err := w.Write(comm); err != nil {
			d.logger.Error("Failed to write commitment request body", "err", err, "comm", comm)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		key := path.Base(r.URL.Path)
		comm, err := hexutil.Decode(key)
		if err != nil {
			d.logger.Error("Failed to decode commitment", "err", err, "key", key)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := d.store.Put(r.Context(), comm, input); err != nil {
			d.logger.Error("Failed to store commitment to the DA server", "err", err, "key", key)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

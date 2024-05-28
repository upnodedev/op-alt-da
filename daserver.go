package plasma

import (
	"github.com/ethereum/go-ethereum/log"
	"net"
	"net/http"
	"plasma-da/da"
	"strconv"
)

type DAServer struct {
	logger     log.Logger
	endpoint   string
	store      da.KVStore
	httpServer *http.Server
	listener   net.Listener
}

func NewDAServer(host string, port int, store da.KVStore, logger log.Logger) *DAServer {
	endpoint := net.JoinHostPort(host, strconv.Itoa(port))
	return &DAServer{
		logger:   logger,
		endpoint: endpoint,
		store:    store,
		httpServer: &http.Server{
			Addr: endpoint,
		},
	}
}

func (d *DAServer) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/get/", d.HandleGet)
	mux.HandleFunc("/put/", d.HandlePut)

	d.httpServer.Handler = mux

	listener, err := net.Listen("tcp", d.endpoint)
	if err != nil {
		return err
	}
	d.listener = listener

	go func() {
		if err := d.httpServer.Serve(d.listener); err != nil {
			d.logger.Error("failed to serve", "err", err)
		}
	}()

	return nil
}

func (d *DAServer) Stop() error {
	return d.httpServer.Close()
}

func (d *DAServer) HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get\n"))
}

func (d *DAServer) HandlePut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put\n"))
}

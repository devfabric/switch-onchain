package profiling

import (
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"

	"github.com/gorilla/mux"
)

func CheckAndEnableProfiling(pport string) error {
	if pport != "" {
		iport, err := strconv.Atoi(pport)
		if err != nil || iport < 0 {
			return fmt.Errorf("Profile port %s is not a valid, not enabling profiling", pport)
		} else {
			addr := net.JoinHostPort("0.0.0.0", pport)
			listener, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			go func() {
				r := mux.NewRouter()
				r.HandleFunc("/debug/pprof/", pprof.Index)
				r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
				r.HandleFunc("/debug/pprof/profile", pprof.Profile)
				r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
				r.HandleFunc("/debug/pprof/trace", pprof.Trace)
				err := http.Serve(listener, nil)
				panic(err)

			}()
		}
	}
	return nil
}

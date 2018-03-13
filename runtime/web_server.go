package runtime

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/Southclaws/sampctl/print"
	"github.com/Southclaws/sampctl/types"
)

// RunWithHTTP runs a server with a HTTP server running alongside to provide administration and data
// endpoints to interface with the server via a web frontend.
func RunWithHTTP(ctx1 context.Context, cfg types.Runtime, cacheDir string, recover bool, output io.Writer, input io.Reader) (err error) {
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(ctx1)
	defer cancel()

	go func() {
		binary := "./" + getServerBinary(cfg.Platform)
		fullPath := filepath.Join(cfg.WorkingDir, binary)
		print.Verb("starting", binary, "in", cfg.WorkingDir)

		errChan <- run(ctx, fullPath, cfg.Mode, recover, output, input)
	}()

	go func() {
		errChan <- initHTTP(cfg)
	}()

	err = <-errChan
	if err != nil {
		return errors.Wrap(err, "runtime encountered an error")
	}

	return
}

func initHTTP(cfg types.Runtime) (err error) {
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	addr := fmt.Sprintf("%s:%s", cfg.WebConfig.IP, cfg.WebConfig.Port)

	print.Verb("Starting HTTP server at", addr)

	return http.ListenAndServe(addr, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
}

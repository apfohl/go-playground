package modules

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
)

func nameHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Awesome Server")
}

func versionHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Awesome Server")
}

func NewServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run HTTP servers",
		Run:   runServer,
	}
}

func runServer(_ *cobra.Command, _ []string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	nameMux := http.NewServeMux()
	nameMux.HandleFunc("/name", nameHandler)
	nameServer := http.Server{Addr: ":8080", Handler: nameMux}

	versionMux := http.NewServeMux()
	versionMux.HandleFunc("/version", versionHandler)
	versionServer := http.Server{Addr: ":8081", Handler: versionMux}

	var waitGroup sync.WaitGroup

	waitGroup.Go(func() {
		go func() {
			_, _ = fmt.Fprintln(os.Stdout, "Name server is listening on port 8080...")
			if err := nameServer.ListenAndServe(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Name server error: %v\n", err)
			}
		}()

		<-ctx.Done()
		_, _ = fmt.Fprintln(os.Stdout, "Shutting down name server...")
		err := nameServer.Shutdown(ctx)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Name server shutdown error: %v\n", err)
		}
	})

	waitGroup.Go(func() {
		go func() {
			_, _ = fmt.Fprintln(os.Stdout, "Version server is listening on port 8080...")
			if err := versionServer.ListenAndServe(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Version server error: %v\n", err)
			}
		}()

		<-ctx.Done()
		_, _ = fmt.Fprintln(os.Stdout, "Shutting down version server...")
		err := versionServer.Shutdown(ctx)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Version server shutdown error: %v\n", err)
		}
	})

	waitGroup.Wait()
	_, _ = fmt.Fprintf(os.Stdout, "Servers stopped.")
}

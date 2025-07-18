package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/a-h/templ"
	"github.com/crispyfishtech/crispyfish-demo/views"
	"github.com/urfave/cli"
)

var (
	mux           = http.NewServeMux()
	sessionCookie = "session"
	waitGroup     = sync.WaitGroup{}
	started       = time.Now()
	requests      = 0
)

type (
	Ping struct {
		Instance  string `json:"instance"`
		RequestID string `json:"request_id,omitempty"`
	}

	Info struct {
		Hostname string `json:"hostname"`
		Uptime   string `json:"uptime"`
		Requests int    `json:"requests"`
	}
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return hostname
}

func getInfo() (*Info, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	uptime := time.Since(started)

	return &Info{
		Hostname: hostname,
		Uptime:   uptime.String(),
		Requests: requests,
	}, nil
}

func info(w http.ResponseWriter, r *http.Request) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	w.Header().Set("Connection", "close")

	i, err := getInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func load(w http.ResponseWriter, r *http.Request) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	// add a false delay
	time.Sleep(2 * time.Second)

	w.Header().Set("Connection", "close")

	i, err := getInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func fail(w http.ResponseWriter, r *http.Request) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	// add a false delay
	time.Sleep(2 * time.Second)

	w.Header().Set("Connection", "close")
	w.WriteHeader(http.StatusInternalServerError)
	views.InternalServerError().Render(r.Context(), w)
}

func missing(w http.ResponseWriter, r *http.Request) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	// add a false delay
	time.Sleep(2 * time.Second)

	w.Header().Set("Connection", "close")
	w.WriteHeader(http.StatusNotFound)
	views.NotFound().Render(r.Context(), w)
}

func ping(w http.ResponseWriter, r *http.Request) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	w.Header().Set("Connection", "close")

	hostname := getHostname()

	p := Ping{
		Instance: hostname,
	}

	requestID := r.Header.Get("X-Request-Id")
	if requestID != "" {
		p.RequestID = requestID
	}

	current, _ := r.Cookie(sessionCookie)
	if current == nil {
		current = &http.Cookie{
			Name:    sessionCookie,
			Value:   fmt.Sprintf("%d", time.Now().UnixNano()),
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, 1),
			MaxAge:  86400,
		}
	}
	fmt.Printf("cookie: %s\n", current.Value)

	http.SetCookie(w, current)

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func waitForDone(ctx context.Context) {
	waitGroup.Wait()
	ctx.Done()
}

func counter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		h.ServeHTTP(w, r)
	})
}

func main() {
	app := cli.NewApp()
	app.Name = "crispyfish-demo"
	app.Usage = "crispyfish kubernetes demo application with fish!"
	app.Version = "1.4.1"
	app.Author = "@crispyfishtech"
	app.Email = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "listen-addr, l",
			Usage: "listen address",
			Value: ":8080",
		},
		cli.StringFlag{
			Name:  "tls-cert, c",
			Usage: "tls certificate",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tls-key, k",
			Usage: "tls certificate key",
			Value: "",
		},
	}
	app.Action = func(c *cli.Context) error {
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
		mux.Handle("/demo", counter(http.HandlerFunc(ping)))
		mux.Handle("/info", counter(http.HandlerFunc(info)))
		mux.Handle("/load", counter(http.HandlerFunc(load)))
		mux.Handle("/fail", counter(http.HandlerFunc(fail)))
		mux.Handle("/404", counter(http.HandlerFunc(missing)))
		mux.Handle("/", templ.Handler(views.Index("Crispyfish Demo")))

		hostname := getHostname()
		listenAddr := c.String("listen-addr")
		tlsCert := c.String("tls-cert")
		tlsKey := c.String("tls-key")

		srv := &http.Server{
			Handler:      mux,
			Addr:         listenAddr,
			WriteTimeout: time.Second * 10,
			ReadTimeout:  time.Second * 10,
		}

		log.Printf("instance: %s\n", hostname)
		log.Printf("listening on %s\n", listenAddr)

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)

		go func() {
			<-ch
			log.Println("stopping")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			waitForDone(ctx)

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal(err)
			}
		}()

		var err error
		if tlsCert != "" && tlsKey != "" {
			err = srv.ListenAndServeTLS(tlsCert, tlsKey)
		} else {
			err = srv.ListenAndServe()
		}

		return err
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

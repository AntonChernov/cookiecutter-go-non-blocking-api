package log

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LogWriter struct {
}

func (writer LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [DEBUG] " + string(bytes))
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr

		log.Printf("%s - - [%s] %q %q %q",
			addr,
			time.Now().Format("2006-01-02T15:04:05Z07:00"),
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
			r.Referer(),
			r.UserAgent())
		// Do stuff here
		// log.Printf("Method: %v URI: %v", r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

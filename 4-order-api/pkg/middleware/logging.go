package middleware

import (
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type responseRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	// если handler не вызвал WriteHeader — статус будет 200
	if r.status == 0 {
		r.status = http.StatusOK
	}
	n, err := r.ResponseWriter.Write(b)
	r.bytes += n
	return n, err
}

func LoggingJSON(l *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()

			rr := &responseRecorder{ResponseWriter: w}

			next.ServeHTTP(rr, req)

			latency := time.Since(start)

			l.WithFields(logrus.Fields{
				"ts":         time.Now().UTC().Format(time.RFC3339Nano),
				"method":     req.Method,
				"path":       req.URL.Path,
				"query":      req.URL.RawQuery,
				"status":     rr.status,
				"durationMs": latency.Milliseconds(),
				"bytes":      rr.bytes,
				"ip":         clientIP(req),
				"userAgent":  req.UserAgent(),
			}).Info("http_request")
		})
	}
}

func clientIP(r *http.Request) string {
	// сначала X-Forwarded-For (если есть прокси/балансер)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// берём первое значение
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' {
				return xff[:i]
			}
		}
		return xff
	}

	// иначе RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

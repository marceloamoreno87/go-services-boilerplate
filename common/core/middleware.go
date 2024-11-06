package core

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"sendzap-checkout/common/helpers"
	"sendzap-checkout/common/infra/jwt"
	"github.com/prometheus/client_golang/prometheus"
)

func SetDefaultMiddleware(mux *chi.Mux) {
	metricsMiddleware := MetricsMiddleware{}

	mux.Use(middleware.Logger)

	mux.Use(middleware.Recoverer)

	mux.Use(middleware.RequestID)

	mux.Use(middleware.RealIP)

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Use(metricsMiddleware.NewMetrics())

	mux.Use(metricsMiddleware.NewPatternMetrics())
}

type AuthMiddleware struct{}

func (c AuthMiddleware) CheckAuth() func(next http.Handler) http.Handler {
	// Define a função que será usada como middleware
	fn := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := OPTL.Start(r.Context(), "Middleware.CheckAuth")
			defer span.End()

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				helpers.LogError(span, helpers.REQUEST_BODY, "authorization header is required")
				helpers.Unauthorized(w, errors.New("error.authorization_header_is_required"), span.SpanContext().TraceID().String())
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				helpers.LogError(span, helpers.REQUEST_BODY, "bearer token is required")
				helpers.Unauthorized(w, errors.New("error.bearer_token_is_required"), span.SpanContext().TraceID().String())
				return
			}

			jwtInstance := jwt.JWT{JwtSecret: JWT_SECRET}
			claims, err := jwtInstance.VerifyToken(tokenString)
			if err != nil {
				helpers.LogError(span, helpers.REQUEST_BODY, "invalid token")
				helpers.Unauthorized(w, errors.New("error.invalid_token"), span.SpanContext().TraceID().String())
				return
			}

			ctx = context.WithValue(ctx, "userID", claims.UserID)
			ctx = context.WithValue(ctx, "jwtExp", claims.Exp)
			ctx = context.WithValue(ctx, "planExpirationDate", claims.PlanExpirationDate)
			ctx = context.WithValue(ctx, "role", claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	return fn
}

type AdminMiddleware struct{}

func (c AdminMiddleware) CheckAdmin() func(next http.Handler) http.Handler {
	fn := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := OPTL.Start(r.Context(), "Middleware.CheckAdmin")
			defer span.End()

			role := helpers.GetRoleFromContext(ctx)
			if role != "admin" {
				helpers.LogError(span, helpers.REQUEST_BODY, "user is not an admin")
				helpers.Unauthorized(w, errors.New("error.user_is_not_an_admin"), span.SpanContext().TraceID().String())
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
	return fn
}

var (
	dflBuckets = []float64{300, 1200, 5000}
)

const (
	reqsName           = "chi_requests_total"
	latencyName        = "chi_request_duration_milliseconds"
	patternReqsName    = "chi_pattern_requests_total"
	patternLatencyName = "chi_pattern_request_duration_milliseconds"
)

type MetricsMiddleware struct {
	reqs    *prometheus.CounterVec
	latency *prometheus.HistogramVec
}

func (c MetricsMiddleware) NewMetrics(buckets ...float64) func(next http.Handler) http.Handler {
	var m MetricsMiddleware
	m.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        reqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"service": os.Getenv("PROJECT_NAME")},
		},
		[]string{"code", "method", "path"},
	)
	prometheus.MustRegister(m.reqs)

	if len(buckets) == 0 {
		buckets = dflBuckets
	}
	m.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        latencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
		ConstLabels: prometheus.Labels{"service": os.Getenv("PROJECT_NAME")},
		Buckets:     buckets,
	},
		[]string{"code", "method", "path"},
	)
	prometheus.MustRegister(m.latency)
	return m.handler
}

func (c MetricsMiddleware) handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		c.reqs.WithLabelValues(http.StatusText(ww.Status()), r.Method, r.URL.Path).Inc()
		c.latency.WithLabelValues(http.StatusText(ww.Status()), r.Method, r.URL.Path).Observe(float64(time.Since(start).Nanoseconds()) / 1000000)
	}
	return http.HandlerFunc(fn)
}

func (c MetricsMiddleware) NewPatternMetrics(buckets ...float64) func(next http.Handler) http.Handler {
	var m MetricsMiddleware
	m.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        patternReqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path (with patterns).",
			ConstLabels: prometheus.Labels{"service": os.Getenv("PROJECT_NAME")},
		},
		[]string{"code", "method", "path"},
	)
	prometheus.MustRegister(m.reqs)

	if len(buckets) == 0 {
		buckets = dflBuckets
	}
	m.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        patternLatencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path (with patterns).",
		ConstLabels: prometheus.Labels{"service": os.Getenv("PROJECT_NAME")},
		Buckets:     buckets,
	},
		[]string{"code", "method", "path"},
	)
	prometheus.MustRegister(m.latency)
	return m.patternHandler
}

func (c MetricsMiddleware) patternHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		rctx := chi.RouteContext(r.Context())
		routePattern := strings.Join(rctx.RoutePatterns, "")
		routePattern = strings.Replace(routePattern, "/*/", "/", -1)

		c.reqs.WithLabelValues(http.StatusText(ww.Status()), r.Method, routePattern).Inc()
		c.latency.WithLabelValues(http.StatusText(ww.Status()), r.Method, routePattern).Observe(float64(time.Since(start).Nanoseconds()) / 1000000)
	}
	return http.HandlerFunc(fn)
}

type CacheMiddleware struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (c *CacheMiddleware) WriteHeader(statusCode int) {
	c.statusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
}

func (c *CacheMiddleware) Write(b []byte) (int, error) {
	c.body.Write(b)
	return c.ResponseWriter.Write(b)
}

func (m *CacheMiddleware) Cache() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := OPTL.Start(r.Context(), "Middleware.CacheMiddleware")
			defer span.End()

			if r.Method != http.MethodGet {
				next.ServeHTTP(w, r)
				return
			}

			cacheKey := r.URL.RequestURI() + "_" + helpers.GetUserIDFromContextStr(ctx) + helpers.GetExpFromContext(ctx)
			var data map[string]interface{}
			err := REDISCONN.HGetAll(ctx, cacheKey).Scan(&data)
			if err == nil {
				helpers.ReturnSuccess(ctx, w, data)
				return
			}

			rec := &CacheMiddleware{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
				body:           bytes.NewBufferString(""),
			}

			next.ServeHTTP(rec, r)

			if rec.statusCode == http.StatusOK {
				REDISCONN.Set(ctx, cacheKey, rec.body.String(), 1*time.Minute)
			}
		})
	}
}

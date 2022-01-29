package pkg

import (
	"context"
	"github.com/ardikapras/owldetect/internal/constant"
	"net"
	"net/http"
	"net/url"
	gpath "path"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type IRouter interface {
	Group(path string, m ...Middleware) *Router
	Use(m ...Middleware) *Router
	GET(path string, handler httprouter.Handle)
	DELETE(path string, handler httprouter.Handle)
	HEAD(path string, handler httprouter.Handle)
	OPTIONS(path string, handler httprouter.Handle)
	PATCH(path string, handler httprouter.Handle)
	POST(path string, handler httprouter.Handle)
	PUT(path string, handler httprouter.Handle)
	Handle(method, path string, handler httprouter.Handle)
	ServeFiles(path string, root http.FileSystem)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// A Middleware chains http.Handlers.
type Middleware func(httprouter.Handle) httprouter.Handle

// A Router is a http.Handler which supports routing and middlewares.
type Router struct {
	router      *httprouter.Router
	middlewares []Middleware
	path        string
}

// New creates a new Router.
func New() IRouter {
	r := httprouter.New()
	return &Router{router: r, path: "/"}
}

// Group returns a new Router with given path and middlewares.
// It should be used for handlers which have same path prefix or
// common middlewares.
func (r *Router) Group(path string, m ...Middleware) *Router {
	return &Router{
		middlewares: append(m, r.middlewares...),
		path:        gpath.Join(r.path, path),
		router:      r.router,
	}
}

// Use appends new middlewares to current Router.
func (r *Router) Use(m ...Middleware) *Router {
	r.middlewares = append(m, r.middlewares...)
	return r
}

// Handle registers a new request handler combined with middlewares.
func (r *Router) Handle(method, path string, handler httprouter.Handle) {
	for _, v := range r.middlewares {
		handler = v(handler)
	}
	r.router.Handle(method, gpath.Join(r.path, path), handler)
}

// GET -
func (r *Router) GET(path string, handler httprouter.Handle) {
	r.Handle("GET", path, r.handle(handler))
}

// DELETE -
func (r *Router) DELETE(path string, handler httprouter.Handle) {
	r.Handle("DELETE", path, r.handle(handler))
}

// HEAD -
func (r *Router) HEAD(path string, handler httprouter.Handle) {
	r.Handle("HEAD", path, r.handle(handler))
}

// OPTIONS -
func (r *Router) OPTIONS(path string, handler httprouter.Handle) {
	r.Handle("OPTIONS", path, r.handle(handler))
}

// PATCH -
func (r *Router) PATCH(path string, handler httprouter.Handle) {
	r.Handle("PATCH", path, r.handle(handler))
}

// POST -
func (r *Router) POST(path string, handler httprouter.Handle) {
	r.Handle("POST", path, r.handle(handler))
}

// PUT -
func (r *Router) PUT(path string, handler httprouter.Handle) {
	r.Handle("PUT", path, r.handle(handler))
}

// ServeFiles -
func (r *Router) ServeFiles(path string, root http.FileSystem) {
	r.router.ServeFiles(path, root)
}

// ServeHTTP -
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) handle(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		t := time.Now()

		values := url.Values{}
		for _, v := range p {
			values.Add(v.Key, v.Value)
		}

		ctx := req.Context()
		ctx = context.WithValue(ctx, constant.ContextUUID, uuid.New().String())
		ctx = context.WithValue(ctx, constant.ContextReferenceUUID, req.Header.Get(string(constant.ContextReferenceUUID)))
		ctx = context.WithValue(ctx, constant.ContextUserAgent, req.Header.Get("User-Agent"))
		ctx = context.WithValue(ctx, constant.ContextUserIP, getUserIP(req))
		ctx = context.WithValue(ctx, constant.ContextDeviceID, req.Header.Get("X-DPL-Device-ID"))
		ctx = context.WithValue(ctx, constant.ContextRouterParam, values)
		ctx = context.WithValue(ctx, constant.ContextBirthTime, t)
		req = req.WithContext(ctx)
		handler(w, req, p)
	}
}

func getUserIP(req *http.Request) string {
	// Get from Header X-Forwarded-For, if not exist get from RemoteAddr
	ipAddress := req.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = req.RemoteAddr
	}

	// The response may be comma separated IP, split them and get the first one
	ipAddress = strings.Replace(ipAddress, " ", "", -1)
	ipAddressArr := strings.Split(ipAddress, ",")
	if len(ipAddressArr) > 0 {
		ipAddress = ipAddressArr[0]
	}

	// Separate the port string (if exist)
	ipAddressArr = strings.Split(ipAddress, ":")
	if len(ipAddressArr) > 0 {
		ipAddress = ipAddressArr[0]
	}

	// Make sure that the IP Address is valid
	addr := net.ParseIP(ipAddress)
	if addr == nil {
		ipAddress = ""
	}

	return ipAddress
}

package responserecorder

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

// used to produce a bitmap indicating
// additional interfaces a ResponseRecorder must
// implement to wrap the original http.ResposeWriter correctly.
const (
	closeNotifier = 1 << iota
	flusher
	hijacker
	readerFrom
	pusher
)

// ResponseRecorder is an interface which wraps
// the http.ResponseWriter and provides additional
// methods to retrieve characteristics of the HTTP response.
type ResponseRecorder interface {
	http.ResponseWriter
	// the content-length of the response payload
	ContentLength() int
	// the returned status code
	StatusCode() int
}

func NewResponseRecorder(rw http.ResponseWriter) ResponseRecorder {
	bitmap := 0
	if _, ok := rw.(http.CloseNotifier); ok {
		bitmap = bitmap | closeNotifier
	}
	if _, ok := rw.(http.Flusher); ok {
		bitmap = bitmap | flusher
	}
	if _, ok := rw.(http.Hijacker); ok {
		bitmap = bitmap | hijacker
	}
	if _, ok := rw.(io.ReaderFrom); ok {
		bitmap = bitmap | readerFrom
	}
	if _, ok := rw.(http.Pusher); ok {
		bitmap = bitmap | pusher
	}

	return factory[bitmap](rw)
}

// responseRecorder implements the ResponseRecorder
// interface by proxying the embedded http.ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	length int
	code   int
}

// the content-length of the response payload
func (r *responseRecorder) ContentLength() int {
	return r.length
}

// the returned status code
func (r *responseRecorder) StatusCode() int {
	return r.code
}

// Header() proxies directly to the ResponseWriter
func (r *responseRecorder) Header() http.Header {
	return r.ResponseWriter.Header()
}

// Write proxies to the ResponseWriter and then records
// the number of bytes wrote.
func (r *responseRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	if err != nil {
		return n, err
	}
	r.length = n
	return n, err
}

// WriteHeader records the provided statusCode and then
// proxies to the ResponseWriter
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.code = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// closeNotifierWrap wraps a ResponseWriter which smuggles
// the http.CloseNotifier interface and proxies to it.
type closeNotifierWrap struct {
	http.ResponseWriter
}

func (w *closeNotifierWrap) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

// flusherWrap wraps a ResponseWriter which smuggles
// the http.Flusher interface and proxies to it.
type flusherWrap struct {
	http.ResponseWriter
}

func (w *flusherWrap) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

// hijackerWrap wraps a ResponseWriter which smuggles
// the http.Hijacker interface and proxies to it.
type hijackerWrap struct {
	http.ResponseWriter
}

func (h *hijackerWrap) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return h.ResponseWriter.(http.Hijacker).Hijack()
}

// readerFromWrap wraps a ResponseWriter which smuggles
// the io.ReaderFrom interface and proxies to it.
type readerFromWrap struct {
	http.ResponseWriter
}

func (r *readerFromWrap) ReadFrom(reader io.Reader) (n int64, err error) {
	return r.ResponseWriter.(io.ReaderFrom).ReadFrom(reader)
}

// pusherWrap wraps a ResponseWriter which smuggles
// the http.Pusher interface and proxies to it.
type pusherWrap struct {
	http.ResponseWriter
}

func (p *pusherWrap) Push(target string, opts *http.PushOptions) error {
	return p.ResponseWriter.(http.Pusher).Push(target, opts)
}

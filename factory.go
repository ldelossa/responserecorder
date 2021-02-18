package responserecorder

import (
	"io"
	"net/http"
)

// factory indexes ResponseRecorder constructors.
// The ResponseRecorder created by the indexed constructor
// will smuggle additional interfaces in accordance with the bitmap.
var factory = [32]func(rw http.ResponseWriter) ResponseRecorder{
	// [0,0,0,0,0] = 0
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
		}{&responseRecorder{rw, 0, 200}}
	},
	// [0,0,0,0,1] = 1
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier)}
	},
	// [0,0,0,1,0] = 2
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher)}
	},
	// [0,0,0,1,1] = 3
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher)}
	},
	// [0,0,1,0,0] = 4
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Hijacker
		}{&responseRecorder{rw, 0, 200}, rw.(http.Hijacker)}
	},
	// [0,0,1,0,1] = 5
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Hijacker
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Hijacker)}
	},
	// [0,0,1,1,0] = 6
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			http.Hijacker
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(http.Hijacker)}
	},
	// [0,0,1,1,1] = 7
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			http.Hijacker
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(http.Hijacker)}
	},
	// [0,1,0,0,0] = 8
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(io.ReaderFrom)}
	},
	// [0,1,0,0,1] = 9
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(io.ReaderFrom)}
	},
	// [0,1,0,1,0] = 10
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(io.ReaderFrom)}
	},
	// [0,1,0,1,1] = 11
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(io.ReaderFrom)}
	},
	// [0,1,1,0,0] = 12
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Hijacker
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.Hijacker), rw.(io.ReaderFrom)}
	},
	// [0,1,1,0,1] = 13
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Hijacker), rw.(io.ReaderFrom)}
	},
	// [0,1,1,1,0] = 14
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(http.Hijacker), rw.(io.ReaderFrom)}
	},
	// [0,1,1,1,1] = 15
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(http.Hijacker), rw.(io.ReaderFrom)}
	},
	// [1,0,0,0,0] = 16
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Pusher)}
	},
	// [1,0,0,0,1] = 17
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Pusher)}
	},
	// [1,0,0,1,0] = 18
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(http.Pusher)}
	},
	// [1,0,0,1,1] = 19
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(http.Pusher)}
	},
	// [1,0,1,0,0] = 20
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Hijacker
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Hijacker), rw.(http.Pusher)}
	},
	// [1,0,1,0,1] = 21
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Hijacker), rw.(http.Pusher)}
	},
	// [1,0,1,1,0] = 22
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			http.Hijacker
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(http.Hijacker), rw.(http.Pusher)}
	},
	// [1,0,1,1,1] = 23
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			http.Hijacker
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(http.Hijacker), rw.(http.Pusher)}
	},
	// [1,1,0,0,0] = 24
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,0,0,1] = 25
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,0,1,0] = 26
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,0,1,1] = 27
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,1,0,0] = 28
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Hijacker), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,1,0,1] = 29
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Hijacker), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,1,1,0] = 30
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.Flusher), rw.(http.Hijacker), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
	// [1,1,1,1,1] = 31
	func(rw http.ResponseWriter) ResponseRecorder {
		return &struct {
			ResponseRecorder
			http.CloseNotifier
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{&responseRecorder{rw, 0, 200}, rw.(http.CloseNotifier), rw.(http.Flusher), rw.(http.Hijacker), rw.(io.ReaderFrom), rw.(http.Pusher)}
	},
}

# ResponseRecorder

ResponseRecorder allows the extraction of HTTP response information.
Typically used in a middleware context.

## Usage

```golang
func RecorderHandler(next http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rr := NewResponseRecorder(w)
        next.ServeHTTP(rr, r)

        log.Printf("content-length: %v", rr.ContentLength())
        log.Printf("status code: %v", rr.StatusCode())
    } 
}
```

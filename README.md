Log5Go Negroni: A Negroni middleware adaptor for Log5Go
===

Usage: 

```go

import l5g "gihub.com/neocortical/log5go"
import l5g-negroni "github.com/neocortical/log5go-negroni"

log := l5g.Logger(l5g.LogAll)

l5g-negroni.New(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  start := time.Now()
  log.Info("Before request: $s", r.RequestURI)

  next(rw, r)

  log.Info("After request. Request took %d ms", time.Since(start))
})

```

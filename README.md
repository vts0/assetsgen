# AssetsGen - Quick Start

Embed static files into Go code and serve them via HTTP in a single binary.  

---

## 1. Prepare assets

Place your files in a folder, e.g., `assets/`:

```bash
assets/
 ├─ index.html
 ├─ style.css
 └─ script.js
 ```

## 2. Generate Go code

Run the generator:

```bash
go run generate.go -s assets -o assets_gen.go
```

## 3. Use in your Go server

Import the generated package and register the handler:

```go
package main

import (
    "log"
    "net/http"
    "path/to/generated/assets"
)

func main() {
    http.HandleFunc("/", assets.HandleAssets)
    log.Println("Serving on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
````

Now all your assets are embedded, served with proper MIME types, ETag caching, gzip compression, and Range support automatically.

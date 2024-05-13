# KVDB: A key value store

kvdb is a simple key value store written in golang. this is a work in progress and is not ready for production use.
It is designed with the db storage principles outlined in the book "
Designing Data-Intensive Applications" by Martin Kleppmann.

## Project Context

This project is a hobby project, developed in my spare time to explore and understand the principles of designing
databases. It is not intended for production use.

While the project aims to implement key features of a key-value database, it is a simplified version and does not
include all the complexities and optimizations found in production-ready databases. Feedback, suggestions, and
contributions are welcome as they help improve the learning process.

## Usage

```go
package main

import (
	"fmt"
	"github.com/bsushmith/kvdb"
)

func main() {
	db := kvdb.NewDB()
	db.Set("key", "value")
	value, _ := db.Get("key")
	fmt.Println(value)
}
```

Available methods:

- Get
- Set
- Delete
- Exists
- GetAllKeys

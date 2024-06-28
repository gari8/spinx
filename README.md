## SPINX

spinx is a loading package built in Go that runs on the CLI.

### How to install
```bash
go get github.com/gari8/spinx
```

### How to use
standard
```go
package main

import "github.com/gari8/spinx"

func main() {
	// select rune or string
	s := spinx.NewSpinner[rune]()
	s.Spin()
	defer s.Stop()
	//...
}
```

use option
```go
package main

import (
	"github.com/gari8/spinx"
	"time"
)

func main() {
    t := 100 * time.Millisecond
    s := spinx.NewSpinner(spinx.Option[string]{
        SpinSpeed:     &t,	
    })
    s.Spin()
    defer s.Stop()
    //...
}
```

use custom Chars
```go
package main

import (
	"github.com/gari8/spinx"
	"time"
)

func main() {
    t := 500 * time.Millisecond
    s := spinx.NewSpinner(spinx.Option[rune]{
        Chars:     []rune{'😀', '😃'},
        SpinSpeed: &t,
    })
    s.Spin()
    defer s.Stop()
    //...
}
```

gear shift
```go
package main

import (
	"github.com/gari8/spinx"
	"time"
)

func main() {
    t := 500 * time.Millisecond
    s := spinx.NewSpinner(spinx.Option[rune]{
        Chars:     []rune{'😀', '😃'},
        SpinSpeed: &t,
    })
    s.Spin()
    defer s.Stop()
    //...
	s.Shift(t / 2)
}
```
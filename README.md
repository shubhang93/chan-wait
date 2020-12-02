## chan-wait

### Wait on multiple channels with ease

Usage

```shell script
go get -v -u github.com/shubhang93/chan-wait/wc
```

```go
package main

import (
	"context"
	"github.com/shubhang93/chan-wait/wc"
	"time"
)

func asyncTask() chan interface{} {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	return done
}

func main() {
	taskChans := []chan interface{}{}
	for i := 0; i < 5; i++ {
		taskChans = append(taskChans, asyncTask())
	}
	<-wc.Wait(context.Background(), taskChans...)
}
```


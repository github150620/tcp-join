# tcp-join
Read from connection A and write to connection B, at the same time read from connection B and write to connection A.

## Install
```
go get github.com/github150620/tcpjoin
```

## Usage
```
import (
	"net"
	"github.com/github150620/tcp-join"
)

func main() {
	conn1, err := net.Dial("tcp", "192.168.1.99:80")
	if err != nil {
		return
	}
  
	conn2, err := net.Dial("tcp", "192.168.1.100:80")
	if err != nil {
		return
	}

	join = TCPJoin.New(conn1, conn2)
	join.Run()
}
```

## Futures
* If either connection A or connection B close, then the other will be closed.
* Connection idle time is 5 minutes.

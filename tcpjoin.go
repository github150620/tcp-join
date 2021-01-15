package tcpjoin

import (
	"net"
	"time"
)

type TCPJoin struct {
	done chan struct{}

	conn1 net.Conn
	conn2 net.Conn
}

func New(conn1, conn2 net.Conn) *TCPJoin {
	return &TCPJoin{
		done:  make(chan struct{}, 2),
		conn1: conn1,
		conn2: conn2,
	}
}

func (join *TCPJoin) Run() {
	go c.readAndWriteServe(join.conn1, join.conn2)
	go c.readAndWriteServe(join.conn2, join.conn1)
	<-join.done
	join.conn1.Close()
	join.conn2.Close()
	<-join.done
}

func (join *TCPJoin) readAndWriteServe(rc, wc net.Conn) {
	defer func() {
		join.done <- struct{}{}
	}()

	buf := make([]byte, 1024)
	for {
		rc.SetReadDeadline(time.Now().Add(time.Second * 300))
		n, err := rc.Read(buf)
		if err != nil {
			break
		}

		wc.SetWriteDeadline(time.Now().Add(time.Second * 300))
		_, err = wc.Write(buf[:n])
		if err != nil {
			break
		}
	}
}

package snellapi

import (
	"context"
	"io"
	"net"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
)

// this file refers to github.com/Trojan-Qt5/go-shadowsocks/ssapi/
type TrafficMeter interface {
	Count(sent uint64, recv uint64)
	Query() (sent uint64, recv uint64)
	io.Closer
}

type MemoryTrafficMeter struct {
	TrafficMeter
	sent uint64
	recv uint64
}

func (m *MemoryTrafficMeter) Count(sent, recv uint64) {
	atomic.AddUint64(&m.sent, uint64(sent))
	atomic.AddUint64(&m.recv, uint64(recv))
}

func (m *MemoryTrafficMeter) Query() (uint64, uint64) {
	return m.sent, m.recv
}

type RecordConn struct {
	net.Conn
	meter TrafficMeter
}

func NewRecordConn(conn net.Conn, meter TrafficMeter) *RecordConn {
	return &RecordConn{
		Conn:  conn,
		meter: meter,
	}
}

func (c *RecordConn) Read(b []byte) (int, error) {
	n, err := c.Conn.Read(b)
	c.meter.Count(uint64(n), 0)
	return n, err
}

func (c *RecordConn) Write(b []byte) (int, error) {
	n, err := c.Conn.Write(b)
	c.meter.Count(0, uint64(n))
	return n, err
}

type ClientAPIService struct {
	SnellServiceServer
	meter         TrafficMeter
	uploadSpeed   uint64
	downloadSpeed uint64
	lastSent      uint64
	lastRecv      uint64
	ctx           context.Context
}

func (s *ClientAPIService) QueryStats(ctx context.Context, req *StatsRequest) (*StatsReply, error) {
	sent, recv := s.meter.Query()
	reply := &StatsReply{
		UploadTraffic:   sent,
		DownloadTraffic: recv,
		UploadSpeed:     s.uploadSpeed,
		DownloadSpeed:   s.downloadSpeed,
	}
	return reply, nil
}

func (s *ClientAPIService) calcSpeed() {
	for {
		select {
		case <-time.After(time.Second):
			sent, recv := s.meter.Query()
			s.uploadSpeed = sent - s.lastSent
			s.downloadSpeed = recv - s.lastRecv
			s.lastSent = sent
			s.lastRecv = recv
		case <-s.ctx.Done():
			return
		}
	}
}

func RunClientAPIService(ctx context.Context, APIAddress string, meter TrafficMeter) error {
	server := grpc.NewServer()
	service := &ClientAPIService{
		meter: meter,
		ctx:   ctx,
	}
	go service.calcSpeed()
	RegisterSnellServiceServer(server, service)
	listener, err := net.Listen("tcp", APIAddress)
	if err != nil {
		return err
	}
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Serve(listener)
	}()
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		server.Stop()
		return nil
	}
}

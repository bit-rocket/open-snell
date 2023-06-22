/*
 * This file is part of open-snell.
 * open-snell is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * open-snell is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU General Public License
 * along with open-snell.  If not, see <https://www.gnu.org/licenses/>.
 */

package socks5

import (
	"io"
	"io/ioutil"
	"net"

	snellapi "github.com/bit-rocket/open-snell/snell-api"
	log "github.com/golang/glog"
)

type SocksCallback func(net.Conn, Addr)

type SockListener struct {
	net.Listener
	address  string
	closed   bool
	callback SocksCallback
}

func NewSocksProxyV2(addr string, cb SocksCallback, meter snellapi.TrafficMeter) (*SockListener, error) {

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	sl := &SockListener{l, addr, false, cb}
	go func() {
		log.Infof("SOCKS proxy listening at: %s\n", addr)
		for {
			c, err := l.Accept()
			if err != nil {
				if sl.closed {
					break
				}
				continue
			}
			recordConn := c
			if meter != nil {
				recordConn = snellapi.NewRecordConn(c, meter)
			}
			go handleSocks(recordConn, sl.callback)
		}
	}()

	return sl, nil
}

func NewSocksProxy(addr string, cb SocksCallback) (*SockListener, error) {
	return NewSocksProxyV2(addr, cb, nil)
}

func (l *SockListener) Close() {
	l.closed = true
	l.Listener.Close()
}

func (l *SockListener) Address() string {
	return l.address
}

func handleSocks(conn net.Conn, cb SocksCallback) {
	target, command, err := ServerHandshake(conn)
	if err != nil {
		conn.Close()
		return
	}
	if c, ok := conn.(*net.TCPConn); ok {
		c.SetKeepAlive(true)
	}
	if command == CmdUDPAssociate {
		defer conn.Close()
		io.Copy(ioutil.Discard, conn)
		return
	}
	cb(conn, target)
}

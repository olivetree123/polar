package server

import (
	"context"
	"encoding/json"
	"errors"

	"net"

	"github.com/olivetree123/polar/response"

	"github.com/olivetree123/polar/commands"
	"github.com/olivetree123/polar/log"
)

type PushHandler func(data response.Response)
type ApiHandler func(ctx context.Context, h PushHandler, command commands.Command) *response.Response

type Server struct {
	address     string
	listener    net.Listener
	funcs       map[int]ApiHandler
	connections map[string]*net.Conn
}

func NewServer(address string) *Server {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Logger.Error("Fail to start server ", err)
		return nil
	}
	return &Server{
		address:     address,
		listener:    listener,
		funcs:       make(map[int]ApiHandler),
		connections: make(map[string]*net.Conn),
	}
}

func (server *Server) Register(commandCode int, f ApiHandler) {
	server.funcs[commandCode] = f
}

func (server *Server) Listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := server.listener.Accept()
			if err != nil {
				log.Logger.Error("Fail to connect ", err)
				continue
			}
			log.Logger.Info("发现新的连接...")
			go func() {
				if err = server.controller(ctx, conn); err != nil {
					log.Logger.Error(err)
				}
			}()
		}
	}
}

func (server *Server) controller(ctx context.Context, conn net.Conn) error {
	if conn == nil {
		return errors.New("conn is nil")
	}
	defer conn.Close()
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	for {
		select {
		case <-ctx.Done():
			goto finishLabel
		default:
			var command commands.Command
			if err := decoder.Decode(&command); err != nil {
				return err
			}
			log.Logger.Info("发现新的请求: code = ", command.Code)
			if command.ClientID == "" {
				return errors.New("ClientID is nil")
			}
			server.connections[command.ClientID] = &conn
			var r *response.Response
			if f, found := server.funcs[command.Code]; found {
				r = f(ctx, server.push, command)
			} else {
				return errors.New("invalid command code")
			}
			if err := encoder.Encode(r); err != nil {
				return err
			}
		}
	}
finishLabel:
	return nil
}

func (server *Server) push(data response.Response) {
	content, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	// 这里有问题 ?
	conn, found := server.connections["Gate"]
	if !found {
		log.Logger.Error("connection for clientID = %s not found", data.UserID)
		return
	}
	if n, err := (*conn).Write(content); err != nil {
		log.Logger.Error(err)
	} else {
		log.Logger.Info("推送成功，推送字节数：", n)
	}
}

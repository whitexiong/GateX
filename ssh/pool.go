package ssh

import (
	"bytes"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"runtime"
	"sync"
)

var Pool *ConnectionPool

type ConnectionPool struct {
	config  *ssh.ClientConfig
	addr    string
	pool    chan *ssh.Client
	maxConn int
	mu      sync.Mutex
	closed  bool
}

func NewSSHConnectionPool(user, password, addr string, maxConn int) *ConnectionPool {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	pool := &ConnectionPool{
		config:  config,
		addr:    addr,
		pool:    make(chan *ssh.Client, maxConn),
		maxConn: maxConn,
	}
	runtime.SetFinalizer(pool, func(p *ConnectionPool) {
		p.safeClose()
	})
	return pool
}

func (s *ConnectionPool) GetConnection() (*ssh.Client, error) {
	select {
	case conn := <-s.pool:
		return conn, nil
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		if len(s.pool) < s.maxConn {
			conn, err := ssh.Dial("tcp", s.addr, s.config)
			if err != nil {
				return nil, err
			}
			return conn, nil
		}
		return <-s.pool, nil
	}
}

func (s *ConnectionPool) ReturnConnection(conn *ssh.Client) {
	select {
	case s.pool <- conn:
	default:
		conn.Close()
	}
}

func (s *ConnectionPool) Close() {
	close(s.pool)
	for conn := range s.pool {
		conn.Close()
	}
}

func (s *ConnectionPool) safeClose() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.closed {
		s.closed = true
		close(s.pool)
		for conn := range s.pool {
			conn.Close()
		}
	}
}

func InitializeSSHConnectionPool() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := os.Getenv("SSH_USER")
	password := os.Getenv("SSH_PASSWORD")
	addr := os.Getenv("SSH_HOST") + ":" + os.Getenv("SSH_PORT")
	maxConn := 10

	Pool = NewSSHConnectionPool(user, password, addr, maxConn)
}

func ExecuteCustomSSHCommand(command string) (string, error) {
	conn, err := Pool.GetConnection()
	if err != nil {
		return "", err
	}
	defer Pool.ReturnConnection(conn)

	session, err := conn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run(command); err != nil {
		return "", err
	}

	return b.String(), nil
}

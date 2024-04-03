package server

import (
	"fmt"
	"log"

	"github.com/emersion/go-imap/client"
	"github.com/rodrigocam/fin/pkg/db"
)

type Server struct {
	ProviderHost string
	ProviderPort int
	Database *db.Database
}

func Default() *Server {
	return &Server{
		ProviderHost: "imap.gmail.com",
		ProviderPort: 993,
		Database: &db.Database{
			Path: "fin.db",
		},
	}
}

func (s *Server) Start(email, password string) error {
	log.Println("Starting watcher server")

	host := fmt.Sprintf("%s:%d", s.ProviderHost, s.ProviderPort)

	c, err := client.DialTLS(host, nil)
	if err != nil {
		return err
	}

	defer c.Logout()

	log.Printf("Successfully connected to the `%s` imap server\n", host)

	if err := c.Login(email, password); err != nil {
		return err
	}
	
	log.Printf("Successfully logged into `%s`\n", email)

	return nil
}

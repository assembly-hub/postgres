package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
)

type Config struct {
	User            string
	Password        string
	Host            string
	Port            int
	Database        string
	SSLMode         string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
	ConnMaxIdleTime int
}

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) *Client {
	c := new(Client)
	c.cfg = cfg
	return c
}

func (c *Client) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=%s",
		c.cfg.User, c.cfg.Password, c.cfg.Host, c.cfg.Port, c.cfg.Database, c.cfg.SSLMode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Duration(c.cfg.ConnMaxLifeTime) * time.Millisecond)
	db.SetConnMaxIdleTime(time.Duration(c.cfg.ConnMaxIdleTime) * time.Millisecond)
	db.SetMaxOpenConns(c.cfg.MaxOpenConn)
	db.SetMaxIdleConns(c.cfg.MaxIdleConn)
	return db, err
}

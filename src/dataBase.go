package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	_ "sort"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type TablesNames struct {
	ExecutedScripts string
}

//go:embed migrations/init.sql
var initScript string

func NewDatabase(cfg DBConfig, names TablesNames) (*DataBase, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(initScript); err != nil {
		return nil, err
	}
	return &DataBase{db, names}, nil
}

type LocalStorage interface {
	findScript(script string) (string, error)
	addScript(script string) error
	checkUser(name, password string) (bool, error)
	addUser(name, password string) error
}

type DataBase struct {
	DB    *sql.DB
	Names TablesNames
}

func (d DataBase) findScript(script string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d DataBase) addScript(script string) error {
	//TODO implement me
	panic("implement me")
}

func (d DataBase) checkUser(name, password string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (d DataBase) addUser(name, password string) error {
	//TODO implement me
	panic("implement me")
}

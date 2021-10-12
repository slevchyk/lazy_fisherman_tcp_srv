package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"

	_ "github.com/lib/pq"
	"github.com/slevchyk/lazy_fisherman_tcp_srv/core"
	"github.com/slevchyk/lazy_fisherman_tcp_srv/database"
	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
)

var cfg models.Config
var db *sql.DB

func init() {

	var err error

	cfg, err = loadConfiguration("config.json")
	if err != nil {
		log.Fatal("Can't load configuration file config.json", err.Error())
	}

	db, err = database.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Can't connect to database", err.Error())
	}
	database.InitDb(db)
}

func main() {

	l, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		for {
			bs := make([]byte, 1024)
			n, err := conn.Read(bs)
			if err != nil {
				break
			}

			request := string(bs[:n])
			request = strings.TrimSuffix(request, "\n")

			cmds := strings.Split(request, "?")

			if len(cmds) > 0 {

				if cmds[0] == "close" {
					break
				}

				if cmds[0] == "-c" && len(cmds) == 2 {

					sa := core.GetCoordinates(db, cmds[1])

					response := fmt.Sprintf("s:%v;%v\n", sa.Status, sa.Response)
					io.WriteString(conn, response)
					continue
				}

				if cmds[0] == "-rb" && len(cmds) == 2 {

					sa := core.GetCoordinates(db, cmds[1])

					response := fmt.Sprintf("s:%v;%v\n", sa.Status, sa.Response)
					io.WriteString(conn, response)
					continue
				}
			}

			answer := request + "-x\n"
			io.WriteString(conn, answer)
		}

		conn.Close()
	}
}

func loadConfiguration(file string) (models.Config, error) {
	var config models.Config

	cfgFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return config, err
	}

	err = json.Unmarshal(cfgFile, &config)
	if err != nil {
		log.Println(err)
		return config, err
	}

	return config, nil
}

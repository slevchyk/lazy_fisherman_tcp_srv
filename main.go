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

			cmds := strings.Split(request, ";")

			if len(cmds) > 0 {

				if cmds[0] == "close" {
					conn.Close()
					break
				}

				if cmds[0] == "coordinates" {
					if len(cmds) == 2 {
						answer, err := getCoordinates(cmds[1])
						if err != nil {
							log.Println(err.Error())
							continue
						}
						io.WriteString(conn, answer)
						continue
					}
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

func getCoordinates(boardId string) (string, error) {

	var answer string

	rows, err := database.SelectMapMarkerByBoardId(db, boardId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var mm models.MapMarker

		database.ScanMapMarker(rows, &mm)
		answer += fmt.Sprintf("%f", mm.Lng)
		answer += ":"
		answer += fmt.Sprintf("%f", mm.Lat)
		answer += ";\n"

	}

	return answer, nil
}

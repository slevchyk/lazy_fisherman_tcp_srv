package core

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/slevchyk/lazy_fisherman_tcp_srv/database"
	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
	"github.com/slevchyk/lazy_fisherman_tcp_srv/status"
)

func GetCoordinates(db *sql.DB, r string) models.ServerAnswer {

	var sa models.ServerAnswer
	var boardId string
	var waterId string

	prms := strings.Split(r, "&")

	for _, v := range prms {
		ss := strings.Split(v, ":")

		if len(ss) != 2 {
			sa.Status = status.CoordinatesCorruptedParameters
			return sa
		}

		switch ss[0] {
		case "id":
			boardId = ss[1]

		case "wid":
			waterId = ss[1]

		default:
			sa.Status = status.CoordinatesUnknownParameter
			return sa
		}
	}

	if boardId == "" {
		sa.Status = status.CoordinatesEmptyKeyParameter
		return sa
	}

	if waterId != "" {
		return getCoordinatesByBoardIdWaterId(db, boardId, waterId)
	} else {
		return getCoordinatesByBoardId(db, boardId)
	}
}

func getCoordinatesByBoardIdWaterId(db *sql.DB, boardId, waterId string) models.ServerAnswer {

	var sa models.ServerAnswer
	var response string
	var bid, wid int64
	var err error

	bid, err = strconv.ParseInt(boardId, 10, 64)
	if err != nil {
		sa.Status = status.CoordinatesIncorrectFormatBoardId
		return sa
	}

	wid, err = strconv.ParseInt(waterId, 10, 64)
	if err != nil {
		sa.Status = status.CoordinatesIncorrectFormatWaterId
		return sa
	}

	rows, err := database.SelectMapMarkerByBoardIdWaterId(db, bid, wid)
	if err != nil {
		sa.Status = status.InternalServerError
		return sa
	}
	defer rows.Close()

	for rows.Next() {
		var mm models.MapMarker

		database.ScanMapMarker(rows, &mm)
		response += fmt.Sprintf("%f:%f;", mm.Lng, mm.Lat)
	}

	sa.Response = response
	return sa
}

func getCoordinatesByBoardId(db *sql.DB, boardId string) models.ServerAnswer {

	var sa models.ServerAnswer
	var response string
	var bid int64
	var err error

	bid, err = strconv.ParseInt(boardId, 10, 64)
	if err != nil {
		sa.Status = status.CoordinatesIncorrectFormatBoardId
		return sa
	}

	rows, err := database.SelectMapMarkerByBoardId(db, bid)
	if err != nil {
		sa.Status = status.InternalServerError
		return sa
	}
	defer rows.Close()

	for rows.Next() {
		var mm models.MapMarker

		database.ScanMapMarker(rows, &mm)
		response += fmt.Sprintf("%f:%f;", mm.Lng, mm.Lat)
	}

	sa.Response = response
	return sa
}

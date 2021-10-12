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

func registerBoard(db *sql.DB, r string) models.ServerAnswer {

	var sa models.ServerAnswer
	var board models.Board
	var name, currentFirmware, latestFirmware, serialNumber, appUserIdString, boardVersionIdString string
	var appUserId, boardVersionId int64
	var err error

	prms := strings.Split(r, "&")

	for _, v := range prms {
		ss := strings.Split(v, ":")

		if len(ss) != 2 {
			sa.Status = status.CoordinatesCorruptedParameters
			return sa
		}

		switch ss[0] {
		case "n":
			name = ss[1]

		case "cf":
			currentFirmware = ss[1]

		case "lf":
			latestFirmware = ss[1]

		case "sn":
			serialNumber = ss[1]

		case "au":
			appUserIdString = ss[1]

		case "bv":
			boardVersionIdString = ss[1]

		default:
			sa.Status = status.CoordinatesUnknownParameter
			return sa
		}
	}

	if name == "" {
		sa.Status = status.BoardsEmptyNameParametr
		return sa
	}

	if currentFirmware == "" {
		sa.Status = status.BoardsEmptyCurrentFirmwareParametr
		return sa
	}

	if latestFirmware == "" {
		sa.Status = status.BoardsEmptyLatestFirmwareParametr
		return sa
	}

	if serialNumber == "" {
		sa.Status = status.BoardsEmptySerialNumberParametr
		return sa
	}

	if appUserIdString == "" {
		sa.Status = status.BoardsEmptySerialNumberParametr
		return sa
	} else {
		appUserId, err = strconv.ParseInt(appUserIdString, 10, 64)
		if err != nil {
			sa.Status = status.BoardsIncorrectFormatAppUserId
			return sa
		}

		rows, err := database.SelectAppUserById(db, appUserId)
		if err != nil {
			sa.Status = status.InternalServerError
			return sa
		}

		if !rows.Next() {
			sa.Status = status.DbAppUserNotFound
			return sa
		}

		rows.Close()
	}

	if boardVersionIdString == "" {
		sa.Status = status.BoardsEmptySerialNumberParametr
		return sa
	} else {
		boardVersionId, err = strconv.ParseInt(boardVersionIdString, 10, 64)
		if err != nil {
			sa.Status = status.BoardsIncorrectFormatAppUserId
			return sa
		}

		rows, err := database.SelectBoardVersionById(db, boardVersionId)
		if err != nil {
			sa.Status = status.InternalServerError
			return sa
		}

		if !rows.Next() {
			sa.Status = status.DbBoardVersionNotFound
			return sa
		}

		rows.Close()
	}

	board = models.Board{
		Name:            name,
		CurrentFirmware: currentFirmware,
		LatestFirmware:  latestFirmware,
		SerialNumber:    serialNumber,
		AppUserId:       appUserId,
		BoardVersionId:  boardVersionId,
	}

	boardId, err := database.InsertBoard(db, board)
	if err != nil {
		sa.Status = status.InternalServerError
	}

	sa.Status = status.Ok
	sa.Response = fmt.Sprint(boardId)

	return sa
}

package status

const (
	NotQualifiedError = 0
	Ok                = 1

	InternalServerError = 10001

	//DB
	DbBoardNotFound        = 10101
	DbAppUserNotFound      = 10102
	DbBoardVersionNotFound = 10103

	//Coordinates
	CoordinatesCorruptedParameters    = 11001
	CoordinatesUnknownParameter       = 11002
	CoordinatesEmptyKeyParameter      = 11003
	CoordinatesIncorrectFormatBoardId = 11003
	CoordinatesIncorrectFormatWaterId = 11005

	//Boards
	BoardsEmptyNameParametr             = 12001
	BoardsEmptyCurrentFirmwareParametr  = 12002
	BoardsEmptyLatestFirmwareParametr   = 12003
	BoardsEmptySerialNumberParametr     = 12004
	BoardsEmptyAppUserIdParametr        = 12005
	BoardsEmptyBoardVersionIdParametr   = 12006
	BoardsIncorrectFormatAppUserId      = 12007
	BoardsIncorrectFormatBoardVersionId = 12008
)

package models

type Board struct {
	Id              int64  `json:"id"`
	AppUserId       int64  `json:"app_user_id"`
	BoardVersionId  int64  `json:"board_version_id"`
	Name            string `json:"name"`
	CurrentFirmware string `json:"current_firmware"`
	LatestFirmware  string `json:"latest_firmware"`
	SerialNumber    string `json:"serial_number"`
	BoardVersion    BoardVersion
}

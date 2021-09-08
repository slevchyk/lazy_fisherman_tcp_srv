package models

type Board struct {
	Id              string `json:"id"`
	OwnerId         string `json:"owner_id"`
	BoardVersionId  string `json:"board_version_id"`
	Name            string `json:"name"`
	CurrentFirmware string `json:"current_firmware"`
	LatestFirmware  string `json:"latest_firmware"`
	BoardVersion    BoardVersion
}

package models

type Config struct {
	AppAuth authConfig `json:"app_auth"`
	AdminAuth authConfig `json:"admin_auth"`
	DB   dBConfig `json:"db"`
	WinService winService `json:"win_service"`
}

type authConfig struct {
	User string `json:"user"`
	Password string `json:"password"`
}

type dBConfig struct {	
	User string `json:"user"`
	Password string `json:"password"`
	Server string `json:"server"`
	Name string `json:"name"`
}

type winService struct {
	Name        string `json:"name"`
	LongName    string `json:"long_name"`
	Description string `json:"description"`
}
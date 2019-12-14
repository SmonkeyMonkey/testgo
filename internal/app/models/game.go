package models

type Game struct {
	PointsGained string `json:"points_gained"`
	WinStatus    string `json:"win_status"`
	GameType     string `json:"game_type"`
	Created      string `json:"created"`
}

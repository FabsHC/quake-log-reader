package model

type GameInfo struct {
	TotalKills   int            `json:"total_kills"`
	Players      []string       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}

func CreateGameInfo(totalKills int, players []string, kills map[string]int) *GameInfo {
	return &GameInfo{
		TotalKills: totalKills,
		Players:    players,
		Kills:      kills,
	}
}

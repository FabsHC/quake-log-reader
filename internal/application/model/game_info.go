package model

import "slices"

type GameInfo struct {
	TotalKills   int            `json:"total_kills"`
	Players      []string       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}

func CreateGameInfo() *GameInfo {
	return &GameInfo{
		Players:      make([]string, 0),
		Kills:        make(map[string]int),
		KillsByMeans: make(map[string]int),
	}
}

func (g *GameInfo) AddTotalPlayers(playerName string) {
	if !slices.Contains(g.Players, playerName) {
		g.Players = append(g.Players, playerName)
	}
	if _, ok := g.Kills[playerName]; !ok {
		g.Kills[playerName] = 0
	}
}

func (g *GameInfo) ProcessKills(playerName string, point int) {
	totalPoints := g.Kills[playerName]
	totalPoints += point
	g.Kills[playerName] = totalPoints
}

func (g *GameInfo) ProcessKillsByMean(mean string) {
	totalPoints := g.KillsByMeans[mean]
	totalPoints += 1
	g.KillsByMeans[mean] = totalPoints
}

package usecase_test

import (
	"quake-log-reader/internal/aplication/usecase"
	"slices"
	"testing"
)

func TestProcessEmptyLogMessage(t *testing.T) {
	pe := usecase.NewProcessEventUseCase()
	pe.Execute("")
	if len(pe.GetAllGamesResult()) > 0 {
		t.Errorf("game results should be empty")
	}
}

func TestInitGameLogMessage(t *testing.T) {
	pe := usecase.NewProcessEventUseCase()
	pe.Execute(" 20:37 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\bot_minplayers\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0\n")
	if pe.GetActualGame() == nil {
		t.Errorf("fail to start game report")
	}
}

func TestKillGameLogMessageWithoutInitGameShouldFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("use case should have failed")
		}
	}()
	pe := usecase.NewProcessEventUseCase()
	pe.Execute(" 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT\n")
}

func TestProcessWorldKillGameLogMessageAndFinishGame(t *testing.T) {
	pe := usecase.NewProcessEventUseCase()
	pe.Execute(" 20:37 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\bot_minplayers\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0\n")
	if pe.GetActualGame() == nil {
		t.Errorf("fail to start game report")
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("log message should have been processed")
		}
	}()
	pe.Execute(" 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT\n")
	actualGame := pe.GetActualGame()
	if !slices.Contains(actualGame.Players, "Isgalamido") {
		t.Errorf("game should contain Isgalamido as a player")
	}
	if actualGame.TotalKills != 1 {
		t.Errorf("game should contain 1 kill")
	}
	if actualGame.Kills["Isgalamido"] != -1 {
		t.Errorf("player Isgalamido should have -1 kill points")
	}
	if actualGame.KillsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Errorf("game should contain MOD_TRIGGER_HURT with 1 kill count")
	}
	pe.FinishOpenGames()
	if len(pe.GetAllGamesResult()) != 1 {
		t.Errorf("game report should have 1 game result")
	}
}

func TestStartAndFinishTwoGames(t *testing.T) {
	pe := usecase.NewProcessEventUseCase()
	// GAME 1
	pe.Execute(" 20:37 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\bot_minplayers\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0\n")
	if pe.GetActualGame() == nil {
		t.Errorf("fail to start game report")
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("log message should have been processed")
		}
	}()
	pe.Execute(" 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT\n")
	actualGame := pe.GetActualGame()
	if !slices.Contains(actualGame.Players, "Isgalamido") {
		t.Errorf("game should contain Isgalamido as a player")
	}
	if actualGame.TotalKills != 1 {
		t.Errorf("game should contain 1 kill")
	}
	if actualGame.Kills["Isgalamido"] != -1 {
		t.Errorf("player Isgalamido should have -1 kill points")
	}
	if actualGame.KillsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Errorf("game should contain MOD_TRIGGER_HURT with 1 kill count")
	}
	// GAME 2
	pe.Execute(" 20:37 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\bot_minplayers\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0\n")
	if pe.GetActualGame() == nil {
		t.Errorf("fail to start game report")
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("log message should have been processed")
		}
	}()
	pe.Execute("  2:11 Kill: 2 4 6: Dono da Bola killed Zeh by MOD_ROCKET\n")
	actualGame = pe.GetActualGame()
	if !slices.Contains(actualGame.Players, "Dono da Bola") {
		t.Errorf("game should contain Dono da Bola as a player")
	}
	if !slices.Contains(actualGame.Players, "Zeh") {
		t.Errorf("game should contain Zeh as a player")
	}
	if actualGame.TotalKills != 1 {
		t.Errorf("game should contain 1 kill")
	}
	if actualGame.Kills["Zeh"] != 0 {
		t.Errorf("player Zeh should have 0 kill points")
	}
	if actualGame.Kills["Dono da Bola"] != 1 {
		t.Errorf("player Dono da Bola should have 1 kill points")
	}
	if actualGame.KillsByMeans["MOD_ROCKET"] != 1 {
		t.Errorf("game should contain MOD_ROCKET with 1 kill count")
	}
	pe.FinishOpenGames()
	if len(pe.GetAllGamesResult()) != 2 {
		t.Errorf("game report should have 2 games result")
	}
}

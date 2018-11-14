package fake

import (
	mock "github.com/stretchr/testify/mock"

	dem "github.com/markus-wa/demoinfocs-golang"
	common "github.com/markus-wa/demoinfocs-golang/common"
	st "github.com/markus-wa/demoinfocs-golang/sendtables"
)

// GameState is a mock for of demoinfocs.IGameState.
type GameState struct {
	mock.Mock

	TeamTerroristsVal        common.TeamState
	TeamCounterTerroristsVal common.TeamState
	ParticipantsMock         Participants
	GrenadeProjectilesVal    map[int]*common.GrenadeProjectile
	InfernosVal              map[int]*common.Inferno
	EntitiesVal              map[int]*st.Entity
	BombVal                  common.Bomb
}

// IngameTick is a mock-implementation of IGameState.IngameTick().
//
// Returns the mocked value.
func (gs *GameState) IngameTick() int {
	return gs.Called().Int(0)
}

// TeamCounterTerrorists is a mock-implementation of IGameState.TeamCounterTerrorists().
//
// Returns GameState.MockTeamTerrorists.
func (gs *GameState) TeamCounterTerrorists() *common.TeamState {
	return &gs.TeamCounterTerroristsVal
}

// TeamTerrorists is a mock-implementation of IGameState.TeamTerrorists().
//
// Returns GameState.MockTeamTerrorists.
func (gs *GameState) TeamTerrorists() *common.TeamState {
	return &gs.TeamTerroristsVal
}

// Participants is a mock-implementation of IGameState.Participants().
//
// Returns GameState.ParticioantsMock.
func (gs *GameState) Participants() dem.IParticipants {
	return &gs.ParticipantsMock
}

// GrenadeProjectiles is a mock-implementation of IGameState.GrenadeProjectiles().
//
// Returns GameState.GrenadeProjectilesVal.
func (gs *GameState) GrenadeProjectiles() map[int]*common.GrenadeProjectile {
	return gs.GrenadeProjectilesVal
}

// Infernos is a mock-implementation of IGameState.Infernos().
//
// Returns GameState.InfernosVal.
func (gs *GameState) Infernos() map[int]*common.Inferno {
	return gs.InfernosVal
}

// Entities is a mock-implementation of IGameState.Entities().
//
// Returns GameState.EntitiesVal.
func (gs *GameState) Entities() map[int]*st.Entity {
	return gs.EntitiesVal
}

// Bomb is a mock-implementation of IGameState.Bomb().
//
// Returns GameState.BombVal.
func (gs *GameState) Bomb() *common.Bomb {
	return &gs.BombVal
}

// TotalRoundsPlayed is a mock-implementation of IGameState.TotalRoundsPlayed().
//
// Returns the mocked value.
func (gs *GameState) TotalRoundsPlayed() int {
	return gs.Called().Int(0)
}

// IsWarmupPeriod is a mock-implementation of IGameState.IsWarmupPeriod().
//
// Returns the mocked value.
func (gs *GameState) IsWarmupPeriod() bool {
	return gs.Called().Bool(0)
}

// IsMatchStarted is a mock-implementation of IGameState.IsMatchStarted().
//
// Returns the mocked value.
func (gs *GameState) IsMatchStarted() bool {
	return gs.Called().Bool(0)
}

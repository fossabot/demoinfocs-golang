package fake

import (
	mock "github.com/stretchr/testify/mock"

	common "github.com/markus-wa/demoinfocs-golang/common"
)

// Participants is a mock for of demoinfocs.IParticipants.
type Participants struct {
	mock.Mock

	ByUserIDVal      map[int]*common.Player
	ByEntityIDVal    map[int]*common.Player
	AllVal           []*common.Player
	PlayingVal       []*common.Player
	TeamMembersMock  func(common.Team) []*common.Player
	FindByHandleMock func(int) *common.Player
}

// ByUserID is a mock-implementation of IParticipants.ByUserID().
//
// Returns Participants.ByUserIDVal.
func (ptcp *Participants) ByUserID() map[int]*common.Player {
	return ptcp.ByUserIDVal
}

// ByEntityID is a mock-implementation of IParticipants.ByEntityID().
//
// Returns Participants.ByEntityIDVal.
func (ptcp *Participants) ByEntityID() map[int]*common.Player {
	return ptcp.ByEntityIDVal
}

// All is a mock-implementation of IParticipants.All().
//
// Returns Participants.AllVal.
func (ptcp *Participants) All() []*common.Player {
	return ptcp.AllVal
}

// Playing is a mock-implementation of IParticipants.Playing().
//
// Returns Participants.PlayingVal.
func (ptcp *Participants) Playing() []*common.Player {
	return ptcp.PlayingVal
}

// TeamMembers is a mock-implementation of IParticipants.TeamMembers().
//
// Returns Participants.TeamMembersMock(team).
func (ptcp *Participants) TeamMembers(team common.Team) []*common.Player {
	if ptcp.TeamMembersMock == nil {
		return nil
	}

	return ptcp.TeamMembersMock(team)
}

// FindByHandle is a mock-implementation of IParticipants.FindByHandle().
//
// Returns Participants.FindByHandleMock(handle).
func (ptcp *Participants) FindByHandle(handle int) *common.Player {
	if ptcp.FindByHandleMock == nil {
		return nil
	}

	return ptcp.FindByHandleMock(handle)
}

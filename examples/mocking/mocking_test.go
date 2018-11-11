package mocking

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/assert"

	common "github.com/markus-wa/demoinfocs-golang/common"
	events "github.com/markus-wa/demoinfocs-golang/events"
	mock "github.com/markus-wa/demoinfocs-golang/mock"
)

func TestCollectKills(t *testing.T) {
	parser := mock.NewParser()
	kill1 := kill(common.EqAK47)
	kill2 := kill(common.EqScout)
	parser.Events[0] = kill1
	parser.Events[1] = kill2
	parser.On("ParseToEnd").Return(nil) // Return no error

	actual, err := collectKills(parser)

	assert.Nil(t, err)
	expected := []events.Kill{kill1, kill2}
	assert.Equal(t, expected, actual)
}

func kill(wep common.EquipmentElement) events.Kill {
	eq := common.NewEquipment(wep)
	return events.Kill{
		Killer: new(common.Player),
		Weapon: &eq,
		Victim: new(common.Player),
	}
}

func TestCollectKillsError(t *testing.T) {
	parser := mock.NewParser()
	expectedErr := errors.New("Test error")
	parser.On("ParseToEnd").Return(expectedErr)

	kills, actualErr := collectKills(parser)

	assert.Nil(t, kills)
	assert.Equal(t, expectedErr, actualErr)
}

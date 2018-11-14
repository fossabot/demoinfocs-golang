// Package fake provides basic mocks for IParser, IGameState and IParticipants.
// See examples/mocking (https://github.com/markus-wa/demoinfocs-golang/tree/master/examples/mocking).
package fake

import (
	"time"

	dp "github.com/markus-wa/godispatch"
	mock "github.com/stretchr/testify/mock"

	dem "github.com/markus-wa/demoinfocs-golang"
	common "github.com/markus-wa/demoinfocs-golang/common"
	st "github.com/markus-wa/demoinfocs-golang/sendtables"
)

// Parser is a mock for of demoinfocs.IParser.
type Parser struct {
	mock.Mock

	HeaderVal     common.DemoHeader
	Events        map[int]interface{}
	NetMessages   map[int]interface{}
	GameStateMock GameState

	eventDispatcher dp.Dispatcher
	msgDispatcher   dp.Dispatcher
	currentFrame    int
}

// NewParser returns a new parser mock with pre-initialized Events and NetMessages.
func NewParser() *Parser {
	return &Parser{
		Events:      make(map[int]interface{}),
		NetMessages: make(map[int]interface{}),
	}
}

// ServerClasses is a mock-implementation of IParser.ServerClasses().
//
// Not yet implemented.
func (p *Parser) ServerClasses() st.ServerClasses {
	panic("ServerClasses not implemented")
}

// Header is a mock-implementation of IParser.Header().
//
// Returns Parser.HeaderVal.
func (p *Parser) Header() common.DemoHeader {
	return p.HeaderVal
}

// GameState is a mock-implementation of IParser.GameState().
//
// Returns Parser.GameStateMock.
func (p *Parser) GameState() dem.IGameState {
	return &p.GameStateMock
}

// CurrentFrame is a mock-implementation of IParser.CurrentFrame().
func (p *Parser) CurrentFrame() int {
	return p.Called().Int(0)
}

/*
CurrentTime is a mock-implementation of IParser.CurrentTime().

Returns the mocked value as nano-seconds.

	oneMinInNanoSeconds := 1000 * 1000 * 1000 * 60
	mock.On("CurrentTime").Return(oneMinInNanoSeconds)
	mock.CurrentTime() // 1 min
*/
func (p *Parser) CurrentTime() time.Duration {
	return time.Duration(p.Called().Int(0))
}

/*
Progress is a mock-implementation of IParser.Progress().

Returns the mocked value divided by 100.

	mock.On("Progress").Return(55)
	mock.Progress() // 0.55
*/
func (p *Parser) Progress() float32 {
	return float32(p.Called().Int(0)) / 100
}

// RegisterEventHandler is a mock-implementation of IParser.RegisterEventHandler().
func (p *Parser) RegisterEventHandler(handler interface{}) dp.HandlerIdentifier {
	return p.eventDispatcher.RegisterHandler(handler)
}

// UnregisterEventHandler is a mock-implementation of IParser.UnregisterEventHandler().
func (p *Parser) UnregisterEventHandler(identifier dp.HandlerIdentifier) {
	p.eventDispatcher.UnregisterHandler(identifier)
}

// RegisterNetMessageHandler is a mock-implementation of IParser.RegisterNetMessageHandler().
func (p *Parser) RegisterNetMessageHandler(handler interface{}) dp.HandlerIdentifier {
	return p.msgDispatcher.RegisterHandler(handler)
}

// UnregisterNetMessageHandler is a mock-implementation of IParser.UnregisterNetMessageHandler().
func (p *Parser) UnregisterNetMessageHandler(identifier dp.HandlerIdentifier) {
	p.msgDispatcher.UnregisterHandler(identifier)
}

// ParseHeader is a mock-implementation of IParser.ParseHeader().
//
// Returns Parser.HeaderVal.
func (p *Parser) ParseHeader() (common.DemoHeader, error) {
	return p.HeaderVal, nil
}

// ParseToEnd is a mock-implementation of IParser.ParseToEnd().
//
// Dispatches Parser.Events and Parser.NetMessages in the specified order.
//
// Returns the mocked error value.
func (p *Parser) ParseToEnd() (err error) {
	maxFrame := max(p.Events)
	maxNetMessageFrame := max(p.NetMessages)
	if maxFrame < maxNetMessageFrame {
		maxFrame = maxNetMessageFrame
	}

	for ; p.currentFrame <= maxFrame; p.currentFrame++ {
		event, ok := p.Events[p.currentFrame]
		if ok {
			p.eventDispatcher.Dispatch(event)
		}

		msg, ok := p.NetMessages[p.currentFrame]
		if ok {
			p.msgDispatcher.Dispatch(msg)
		}
	}

	return p.Called().Error(0)
}

func max(numbers map[int]interface{}) (maxNumber int) {
	for maxNumber = range numbers {
		break
	}
	for n := range numbers {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return
}

// Cancel is a mock-implementation of IParser.Cancel().
//
// Not yet implemented.
func (p *Parser) Cancel() {
	panic("Cancel not implemented")
}

// ParseNextFrame is a mock-implementation of IParser.ParseNextFrame().
//
// Not yet implemented.
func (p *Parser) ParseNextFrame() (b bool, err error) {
	panic("ParseNextFrame not implemented")
}

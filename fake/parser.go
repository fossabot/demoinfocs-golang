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

	Events      map[int][]interface{}
	NetMessages map[int][]interface{}

	eventDispatcher dp.Dispatcher
	msgDispatcher   dp.Dispatcher
	currentFrame    int
}

// NewParser returns a new parser mock with pre-initialized Events and NetMessages.
func NewParser() *Parser {
	p := &Parser{
		Events:      make(map[int][]interface{}),
		NetMessages: make(map[int][]interface{}),
	}

	return p
}

// ServerClasses is a mock-implementation of IParser.ServerClasses().
//
// Unfortunately sendtables.ServerClasses currently isn't mockable.
func (p *Parser) ServerClasses() st.ServerClasses {
	return p.Called().Get(0).(st.ServerClasses)
}

// Header is a mock-implementation of IParser.Header().
func (p *Parser) Header() common.DemoHeader {
	return p.Called().Get(0).(common.DemoHeader)
}

// GameState is a mock-implementation of IParser.GameState().
func (p *Parser) GameState() dem.IGameState {
	return p.Called().Get(0).(dem.IGameState)
}

// CurrentFrame is a mock-implementation of IParser.CurrentFrame().
func (p *Parser) CurrentFrame() int {
	return p.Called().Int(0)
}

// CurrentTime is a mock-implementation of IParser.CurrentTime().
func (p *Parser) CurrentTime() time.Duration {
	return p.Called().Get(0).(time.Duration)
}

// Progress is a mock-implementation of IParser.Progress().
func (p *Parser) Progress() float32 {
	return p.Called().Get(0).(float32)
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
func (p *Parser) ParseHeader() (common.DemoHeader, error) {
	args := p.Called()
	return args.Get(0).(common.DemoHeader), args.Error(1)
}

// ParseToEnd is a mock-implementation of IParser.ParseToEnd().
//
// Dispatches Parser.Events and Parser.NetMessages in the specified order.
//
// Returns the mocked error value.
func (p *Parser) ParseToEnd() (err error) {
	args := p.Called()

	maxFrame := max(p.Events)
	maxNetMessageFrame := max(p.NetMessages)
	if maxFrame < maxNetMessageFrame {
		maxFrame = maxNetMessageFrame
	}

	for p.currentFrame <= maxFrame {
		p.parseNextFrame()
	}

	return args.Error(0)
}

func (p *Parser) parseNextFrame() {
	events, ok := p.Events[p.currentFrame]
	if ok {
		for _, e := range events {
			p.eventDispatcher.Dispatch(e)
		}
	}

	messages, ok := p.NetMessages[p.currentFrame]
	if ok {
		for _, msg := range messages {
			p.msgDispatcher.Dispatch(msg)
		}
	}

	p.currentFrame++
}

// ParseNextFrame is a mock-implementation of IParser.ParseNextFrame().
//
// Dispatches Parser.Events and Parser.NetMessages in the specified order.
//
// Returns the mocked bool and error values.
func (p *Parser) ParseNextFrame() (b bool, err error) {
	args := p.Called()

	p.parseNextFrame()

	return args.Bool(0), args.Error(1)
}

func max(numbers map[int][]interface{}) (maxNumber int) {
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

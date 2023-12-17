// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Sequencer is an autogenerated mock type for the Sequencer type
type Sequencer struct {
	mock.Mock
}

type Sequencer_GetCur struct {
	*mock.Call
}

func (_m Sequencer_GetCur) Return(_a0 uint64) *Sequencer_GetCur {
	return &Sequencer_GetCur{Call: _m.Call.Return(_a0)}
}

func (_m *Sequencer) OnGetCur() *Sequencer_GetCur {
	c_call := _m.On("GetCur")
	return &Sequencer_GetCur{Call: c_call}
}

func (_m *Sequencer) OnGetCurMatch(matchers ...interface{}) *Sequencer_GetCur {
	c_call := _m.On("GetCur", matchers...)
	return &Sequencer_GetCur{Call: c_call}
}

// GetCur provides a mock function with given fields:
func (_m *Sequencer) GetCur() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

type Sequencer_GetNext struct {
	*mock.Call
}

func (_m Sequencer_GetNext) Return(_a0 uint64) *Sequencer_GetNext {
	return &Sequencer_GetNext{Call: _m.Call.Return(_a0)}
}

func (_m *Sequencer) OnGetNext() *Sequencer_GetNext {
	c_call := _m.On("GetNext")
	return &Sequencer_GetNext{Call: c_call}
}

func (_m *Sequencer) OnGetNextMatch(matchers ...interface{}) *Sequencer_GetNext {
	c_call := _m.On("GetNext", matchers...)
	return &Sequencer_GetNext{Call: c_call}
}

// GetNext provides a mock function with given fields:
func (_m *Sequencer) GetNext() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

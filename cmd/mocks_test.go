// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cmd_test

import (
	"github.com/legaard/uuid/cmd"
	"sync"
)

// Ensure, that WriterMock does implement cmd.Writer.
// If this is not the case, regenerate this file with moq.
var _ cmd.Writer = &WriterMock{}

// WriterMock is a mock implementation of cmd.Writer.
//
//	func TestSomethingThatUsesWriter(t *testing.T) {
//
//		// make and configure a mocked cmd.Writer
//		mockedWriter := &WriterMock{
//			WriteFunc: func(p []byte) (int, error) {
//				panic("mock out the Write method")
//			},
//		}
//
//		// use mockedWriter in code that requires cmd.Writer
//		// and then make assertions.
//
//	}
type WriterMock struct {
	// WriteFunc mocks the Write method.
	WriteFunc func(p []byte) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// Write holds details about calls to the Write method.
		Write []struct {
			// P is the p argument value.
			P []byte
		}
	}
	lockWrite sync.RWMutex
}

// Write calls WriteFunc.
func (mock *WriterMock) Write(p []byte) (int, error) {
	callInfo := struct {
		P []byte
	}{
		P: p,
	}
	mock.lockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	mock.lockWrite.Unlock()
	if mock.WriteFunc == nil {
		var (
			nOut   int
			errOut error
		)
		return nOut, errOut
	}
	return mock.WriteFunc(p)
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//
//	len(mockedWriter.WriteCalls())
func (mock *WriterMock) WriteCalls() []struct {
	P []byte
} {
	var calls []struct {
		P []byte
	}
	mock.lockWrite.RLock()
	calls = mock.calls.Write
	mock.lockWrite.RUnlock()
	return calls
}

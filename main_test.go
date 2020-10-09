package main

import (
	"testing"

	"github.com/nealwolff/mockery/mocks"

	"github.com/stretchr/testify/mock"
)

func TestTestInterfaceWithMock(t *testing.T) {
	theMock := &mocks.TestInterface{}
	TestImplementation = theMock
	passString := "passing data using mock implementation"
	expectedValue := "MOCK IMPLEMENTATION CALLED, STRING MOCKED: " + passString

	theMock.On("TestFunc", mock.Anything).Return(expectedValue).Once()

	//now create an object that uses a function that calls the test implementation
	theType := TestType{}

	theType.AssignData(passString)

	if theType.GetData() != expectedValue {
		t.Fatal("invalid data obtained, got: " + theType.GetData())
	}

}

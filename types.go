package main

//TestInterface is the basic test interface
type TestInterface interface {
	TestFunc(string) string
}

//implementation for testInterface
type testImpl struct {
}

//testFunc implementation for testImpl
func (t testImpl) TestFunc(the string) string {
	return "IMPLEMENTED FUNCTION CALLED, string value passed in: " + the
}

//variable implementation for testInterface
var (
	TestImplementation TestInterface = testImpl{}
)

//TestType random data struct that will use testImplementation to call testFunc
type TestType struct {
	theData string
}

//AssignType assigns the data
func (t *TestType) AssignData(data string) {
	t.theData = TestImplementation.TestFunc(data)
}

//GetData returns private data field theData
func (t *TestType) GetData() string {
	return t.theData
}

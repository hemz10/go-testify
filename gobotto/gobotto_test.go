package gobotto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

/* This is our suite
To use suite please make sure to import suite package
*/

type RobotsURLSuite struct {
	suite.Suite
}

// Testing RobotsURL function
func TestRobotsURLSuccessfull(t *testing.T) {
	// Giving Expected and Actual value to compare
	expected := "http://my-cool-domain.com/robots.txt"
	actual, _ := RobotsURL("http://my-cool-domain.com/blog-post/1")

	// Without Testify we can't use assertion but we can use t.Fatal to capture Error
	// if actual != expected {
	// 	t.Fatal("expected: ", expected, " but got: ", actual)
	// }

	// With Testify we can use assertion and easily test expected and actual values. Infact the error will be displayed more precisely with the diff
	assert.Equal(t, expected, actual, "")
}

// Adding another test to check if err is nil
func TestRobotsURLError(t *testing.T) {
	/* Tests in Go run sequentially unless marked as parallel.
	To mark a test as parallel you just gotta call the Parallel function on the provided *testing.T argument, just like this:
	*/
	t.Parallel()
	_, err := RobotsURL("http://my-cool-domain.com/blog-post/1")

	assert.Nil(t, err)
}

/*
	One of our test using suite

Now we can add whatever test methods we want to that suite just by adding methods that start with Test and that have this suite as the receiver.
You should also notice that now we can call the assertion methods straight from the testing suite.
*/

// Please make sure to not add (t *testing.T) as arguments here.
func (suite *RobotsURLSuite) TestRobotsURLSuccess() {
	expectedURL := "http://my-cool-domain.com/robots.txt"
	result, _ := RobotsURL("http://my-cool-domain.com/blog-post/1")

	// Notice we are now using `suite` to call the assertion methods
	suite.Equal(expectedURL, result)
}

// Adding one more test to the suite to demostrate SetupTest life cycle.
func (suite *RobotsURLSuite) TestRobotsURLError() {
	_, err := RobotsURL("http://my-cool-domain.com/blog-post/1")
	suite.Nil(err)
}

/*
the above test will not run by go test.
We still need to call the suite.Run method within a standard testing method (one that starts with Test so it can get run automatically by go test).
*/
func TestSuite(t *testing.T) {
	suite.Run(t, new(RobotsURLSuite))
}

// In order to add lifecycle hooks all you’ve gotta do is implement them with their corresponding names and make your test suite the receiver for these methods.
/* Below are the life Cycle hooks
BeforeTest(suiteName, testName string) - Runs right before the test starts
AfterTest(suiteName, testName string) - Runs right after the test finishes
SetupSuite() - Runs before the tests in the suite
SetupTest() - Runs before each test in the suite
TearDownTest() - Runs after each test in the suite
TearDownSuite() - Runs after all the tests in the suite have been run
*/

// Example of how to use these hooks

func (suite *RobotsURLSuite) SetupTest() {
	fmt.Printf("Before starting test print before SetupTest")
}

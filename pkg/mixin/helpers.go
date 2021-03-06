package mixin

import (
	"testing"

	"github.com/deislabs/porter/pkg/context"
)

type TestRunner struct {
	*Runner
	TestContext *context.TestContext
}

// NewTestRunner initializes a mixin test runner, with the output buffered, and an in-memory file system.
func NewTestRunner(t *testing.T, mixin string, runtime bool) *TestRunner {
	c := context.NewTestContext(t)
	mixinDir := "/root/.porter/mixins/exec"
	r := &TestRunner{
		Runner:      NewRunner(mixin, mixinDir, runtime),
		TestContext: c,
	}
	r.Context = c.Context

	// Setup Mixin Home
	c.AddTestDirectory("../../bin/mixins", "/root/.porter/mixins")

	return r
}

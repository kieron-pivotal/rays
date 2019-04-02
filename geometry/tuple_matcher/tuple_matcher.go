package tuple_matcher

import (
	"fmt"

	"github.com/kieron-pivotal/rays/geometry"
	"github.com/onsi/gomega/types"
)

type tupleMatcher struct {
	expected geometry.Tuple
}

func Equal(expected geometry.Tuple) types.GomegaMatcher {
	return &tupleMatcher{
		expected: expected,
	}
}

func (m *tupleMatcher) Match(actual interface{}) (success bool, err error) {
	actualTuple, ok := actual.(geometry.Tuple)
	if !ok {
		return false, fmt.Errorf("EqualsTuple matcher expects a geometry.Tuple")
	}

	return actualTuple.Equals(m.expected), nil
}

func (m *tupleMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nto equal\n\t%v", actual, m.expected)
}

func (m *tupleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nnot to equal\n\t%v", actual, m.expected)
}

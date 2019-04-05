package color

import (
	"fmt"

	"github.com/onsi/gomega/types"
)

type colorMatcher struct {
	expected Color
}

func Equal(expected Color) types.GomegaMatcher {
	return &colorMatcher{
		expected: expected,
	}
}

func (m *colorMatcher) Match(actual interface{}) (success bool, err error) {
	actualColor, ok := actual.(Color)
	if !ok {
		return false, fmt.Errorf("color.Equal matcher expects a Color")
	}

	return actualColor.Equals(m.expected), nil
}

func (m *colorMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nto equal\n\t%v", actual, m.expected)
}

func (m *colorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nnot to equal\n\t%v", actual, m.expected)
}

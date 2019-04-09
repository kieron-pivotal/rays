package matrix

import (
	"fmt"

	"github.com/onsi/gomega/types"
)

type matrixMatcher struct {
	expected *Matrix
}

func Equal(expected *Matrix) types.GomegaMatcher {
	return &matrixMatcher{
		expected: expected,
	}
}

func (m *matrixMatcher) Match(actual interface{}) (success bool, err error) {
	actualMatrix, ok := actual.(*Matrix)
	if !ok {
		return false, fmt.Errorf("matrix.Equal matcher expects a Matrix")
	}

	return actualMatrix.Equals(m.expected), nil
}

func (m *matrixMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nto equal\n\t%v", actual, m.expected)
}

func (m *matrixMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nnot to equal\n\t%v", actual, m.expected)
}

package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateStep(t *testing.T) {
	testCases := []struct {
		desc     string
		prevStep Event
		nextStep Event
		expected bool
	}{
		{
			desc:     "Left != Right",
			prevStep: Left,
			nextStep: Right,
			expected: false,
		},
		{
			desc:     "Left == Left",
			prevStep: Left,
			nextStep: Left,
			expected: true,
		},
		{
			desc:     "Left == Up",
			prevStep: Left,
			nextStep: Up,
			expected: true,
		},
		{
			desc:     "Left == Down",
			prevStep: Left,
			nextStep: Down,
			expected: true,
		},
		{
			desc:     "Up != Down",
			prevStep: Up,
			nextStep: Down,
			expected: false,
		},
		{
			desc:     "Up == Left",
			prevStep: Up,
			nextStep: Left,
			expected: true,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tC.expected, validateStep(tC.prevStep, tC.nextStep), tC.desc)
		})
	}
}

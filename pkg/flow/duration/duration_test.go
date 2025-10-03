package duration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDurationAll(t *testing.T) {
	t.Parallel()

	for w := range time.Duration(2) {
		for d := range time.Duration(2) {
			for h := range time.Duration(2) {
				for m := range time.Duration(2) {
					for s := range time.Duration(2) {
						for ms := range time.Duration(2) {
							input := fmt.Sprintf("%dw %dd %dh %dm %ds %dms", w, d, h, m, s, ms)
							t.Run(input, func(t *testing.T) {
								output, err := ParseDuration(input)
								require.NoError(t, err)
								assert.Equal(t, w*7*24*time.Hour+d*24*time.Hour+h*time.Hour+m*time.Minute+s*time.Second+ms*time.Millisecond, output)
							})
						}
					}
				}
			}
		}
	}
}

func TestParseDurationNoZero(t *testing.T) {
	t.Parallel()

	for w := range time.Duration(2) {
		for d := range time.Duration(2) {
			for h := range time.Duration(2) {
				for m := range time.Duration(2) {
					for s := range time.Duration(2) {
						for ms := range time.Duration(2) {
							input := ""
							if w > 0 {
								input += "1w"
							}
							if d > 0 {
								input += "1d"
							}
							if h > 0 {
								input += "1h"
							}
							if m > 0 {
								input += "1m"
							}
							if s > 0 {
								input += "1s"
							}
							if ms > 0 {
								input += "1ms"
							}
							if w == 0 && d == 0 && h == 0 && m == 0 && ms == 0 { // This is invalid
								continue
							}
							t.Run("no-zero-"+input, func(t *testing.T) {
								output, err := ParseDuration(input)
								require.NoError(t, err)
								assert.Equal(t, w*7*24*time.Hour+d*24*time.Hour+h*time.Hour+m*time.Minute+s*time.Second+ms*time.Millisecond, output)
							})
						}
					}
				}
			}
		}
	}
}

func TestParseDurationErrors(t *testing.T) {
	for _, ts := range []struct {
		in  string
		err string
	}{
		{"1h m", "unit hm not recognized in 1hm"},
		{"m", "duration without value: m"},
		{"1m 1h", "unit h is higher than max unit (s)"},
		{"", "empty duration"},
	} {
		t.Run(ts.in, func(t *testing.T) {
			d, err := ParseDuration(ts.in)
			assert.Zero(t, d)
			require.ErrorContains(t, err, ts.err)
		})
	}
}

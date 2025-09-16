package flow

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
								output, err := parseDuration(input)
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
							t.Run("no-zero-"+input, func(t *testing.T) {
								output, err := parseDuration(input)
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
		{"1h m", "format specifier (m) without value"},
		{"m", "format specifier (m) without value"},
		{"1m 1h", "format specifier (hour) is higher than max unit (second)"},
	} {
		t.Run(ts.in, func(t *testing.T) {
			d, err := parseDuration(ts.in)
			assert.Zero(t, d)
			require.ErrorContains(t, err, ts.err)
		})
	}
}

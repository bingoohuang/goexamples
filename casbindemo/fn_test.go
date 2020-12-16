package casbindemo_test

import (
	"casbindemo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWildcardMatch(t *testing.T) {
	assert.True(t, casbindemo.WildcardMatch("abc", "ab*"))
	assert.True(t, casbindemo.WildcardMatch("adc", "ab*/ad*"))
	assert.False(t, casbindemo.WildcardMatch("axc", "ab*/ad*"))
	assert.False(t, casbindemo.WildcardMatch("adc", "ab*"))
	assert.True(t, casbindemo.WildcardMatch("adc", "a?c"))
}

func TestTimeAllow(t *testing.T) {
	// ignore/deny
	assert.True(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "-"))
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "x"))

	// until absolute datetime
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "2020-12-16 18:00:24"))
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "2020-12-16 18:00:25"))
	assert.True(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "2020-12-16 18:00:26"))

	// until duration
	casbindemo.CasbinStartTime, _ = time.Parse(casbindemo.CasbinTimeLayout, "2020-12-16 18:00:00")
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "10s"))
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "25s"))
	assert.True(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "26s"))

	// until range
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:24", "2020-12-16 18:00:25/2020-12-16 18:00:26"))
	assert.True(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "2020-12-16 18:00:25/2020-12-16 18:00:26"))
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:26", "2020-12-16 18:00:25/2020-12-16 18:00:26"))
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:27", "2020-12-16 18:00:25/2020-12-16 18:00:26"))

	// until range
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:24", "2020-12-16 18:00:25/10s"))
	assert.True(t, casbindemo.TimeAllow("2020-12-16 18:00:25", "2020-12-16 18:00:25/10s"))
	assert.True(t, casbindemo.TimeAllow("2020-12-16 18:00:26", "2020-12-16 18:00:25/10s"))
	assert.False(t, casbindemo.TimeAllow("2020-12-16 18:00:35", "2020-12-16 18:00:25/10s"))
}

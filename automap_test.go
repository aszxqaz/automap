package automap_test

import (
	"testing"

	"github.com/aszxqaz/automap"
	"github.com/stretchr/testify/assert"
)

func TestKeysValuesLen(t *testing.T) {
	var m automap.Map[int, string]

	m.Set(1, "one")
	m.Set(2, "two")

	keys := m.Keys()
	assert.ElementsMatch(t, []int{1, 2}, keys)

	values := m.Values()
	assert.ElementsMatch(t, []string{"one", "two"}, values)

	assert.Equal(t, 2, m.Len())
}

func TestGet(t *testing.T) {
	var m1 automap.Map[int, *string]

	s1, ok := m1.Get(1)
	assert.False(t, ok)
	assert.Nil(t, s1)

	var m2 automap.Map[int, string]

	s2, ok := m2.Get(1)
	assert.False(t, ok)
	assert.Equal(t, "", s2)
}

func TestSet(t *testing.T) {
	var m automap.Map[int, string]

	m.Set(1, "one")

	s, ok := m.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "one", s)
}

func TestDelete(t *testing.T) {
	var m automap.Map[int, string]

	m.Set(1, "one")

	ok := m.Delete(1)
	assert.True(t, ok)

	ok = m.Delete(2)
	assert.False(t, ok)
}

func TestWhere(t *testing.T) {
	var m automap.Map[int, string]

	m.Set(1, "one")
	m.Set(2, "two")

	s, ok := m.FirstWhere(func(k int, v string) bool { return v == "one" })
	assert.True(t, ok)
	assert.Equal(t, "one", s)

	s, ok = m.FirstWhere(func(k int, v string) bool { return v == "three" })
	assert.False(t, ok)
	assert.Equal(t, "", s)
}

func TestUpdate(t *testing.T) {
	var m automap.Map[int, string]

	m.Set(1, "one")
	m.Set(2, "two")

	ok := m.Update(1, func(k int, v string) string { return "ONE" })
	assert.True(t, ok)
	s, ok := m.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "ONE", s)

	ok = m.Update(3, func(k int, v string) string { panic("unreachable") })
	assert.False(t, ok)
}

func TestUpdateWhere(t *testing.T) {
	var m automap.Map[int, string]

	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")
	m.Set(4, "four")

	ok := m.UpdateWhere(
		func(k int, v string) bool { return k > 2 },
		func(k int, v string) string { return "too much" },
	)

	assert.True(t, ok)
	s, ok := m.Get(3)
	assert.True(t, ok)
	assert.Equal(t, "too much", s)

	s, ok = m.Get(4)
	assert.True(t, ok)
	assert.Equal(t, "too much", s)

	ok = m.UpdateWhere(
		func(k int, v string) bool { return len(v) < 3 },
		func(k int, v string) string { panic("unreachable") },
	)
	assert.False(t, ok)
}

func TestReduce(t *testing.T) {
	var m automap.Map[string, int]

	m.Set("1", 1)
	m.Set("2", 2)
	m.Set("3", 3)

	sum := m.Reduce(0, func(k string, v int, r any) any {
		return r.(int) + v
	})
	assert.Equal(t, 6, sum)
}

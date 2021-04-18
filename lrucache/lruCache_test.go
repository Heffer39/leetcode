package lrucache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type pair struct {
	key, value int
}

var testCases = []struct {
	name     string
	capacity int
	input    []pair
	expected []pair
}{
	{name: "standard input",
		capacity: 2,
		input: []pair{
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4}},
		expected: []pair{
			{1, -1},
			{2, -1},
			{3, 3},
			{4, 4}}},
	{name: "overwrite value",
		capacity: 3,
		input: []pair{
			{1, 1},
			{2, 1},
			{3, 3},
			{4, 4}},
		expected: []pair{
			{1, -1},
			{2, 1},
			{3, 3},
			{4, 4}}},
	{name: "capacity 1",
		capacity: 1,
		input: []pair{
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4}},
		expected: []pair{
			{1, -1},
			{2, -1},
			{3, -1},
			{4, 4}}},
	{name: "capacity 2 overwrite value",
		capacity: 2,
		input: []pair{
			{2, 1},
			{1, 1},
			{2, 3},
			{4, 1}},
		expected: []pair{
			{1, -1},
			{2, 3},
			{4, 1}}},
	{name: "capacity 2",
		capacity: 2,
		input: []pair{
			{2, 1},
			{3, 1},
			{4, 3}},
		expected: []pair{
			{2, -1},
			{3, 1},
			{4, 3}}},
}

func TestLRUCache(t *testing.T) {
	for _, test := range testCases {
		t.Logf("Running test: %q", test.name)
		obj := Constructor(test.capacity)
		for _, input := range test.input {
			obj.Put(input.key, input.value)
		}
		for _, output := range test.expected {
			assert.Equal(t, output.value, obj.Get(output.key))
		}
	}
}

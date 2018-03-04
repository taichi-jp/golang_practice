package main

import (
	"reflect"
	"testing"
)

var anscestorTests = []struct {
	descendant Category
	anscestor  Category
	out        bool
}{
	{categories[0], categories[0], true},
	{categories[3], categories[0], true},
	{categories[1], categories[0], false},
	{categories[5], categories[0], true},
}

func TestAnscestorIs(t *testing.T) {
	for _, test := range anscestorTests {
		actual := test.descendant.AnscestorIs(test.anscestor, categories)
		expected := test.out
		if actual != expected {
			t.Errorf("descendant: %v\nanscestor: %v\ngot: %v\nwant: %v", test.descendant, test.anscestor, actual, expected)
		}
	}
}

func BenchmarkAnscestorIs(b *testing.B) { // O(n^3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range anscestorTests {
			test.descendant.AnscestorIs(test.anscestor, categories)
		}
	}
}

// 普通にやると通らない(値渡しなので)。ポインタをとって比較するアプローチで正しい？？
var descendantsTests = []struct {
	target *Category
	out    []*Category
}{
	{&categories[0], []*Category{&categories[0], &categories[2], &categories[3], &categories[5]}},
	{&categories[1], []*Category{&categories[1], &categories[4]}},
	{&categories[2], []*Category{&categories[2], &categories[5]}},
}

func TestDescendantsOf(t *testing.T) {
	for _, test := range descendantsTests {
		actual := DescendantsOf(test.target, categories)
		expected := test.out
		if reflect.DeepEqual(actual, expected) {
			t.Errorf("target: %v\ngot: %v\nwant: %v", test.target, actual, expected)
		}
	}
}

func BenchmarkDescendantsOf(b *testing.B) { // O(n^2) + append
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range descendantsTests {
			DescendantsOf(test.target, categories)
		}
	}
}

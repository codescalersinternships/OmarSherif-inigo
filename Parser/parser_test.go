package main

import (
	"reflect"
	"testing"
)

//Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the []
func TestParser(t *testing.T) {

	t.Run("Get the Name Of the Section", func(t *testing.T) {
		got := section.Name()
		want := "section1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
	t.Run("Get the Key list of a Section", func(t *testing.T) {
		got := section.GetKeyList()
		want := []string{"key1", "key2"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Get value of a key of a Section", func(t *testing.T) {
		got, err := section.GetValue("key1")
		want := "value1"
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

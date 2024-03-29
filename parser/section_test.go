package main

import (
	"reflect"
	"testing"
)

//Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the []
func TestSection(t *testing.T) {
	sectionName := "section1"

	var sectionDictionary Dictionary = map[string]string{"key1": "value1", "key2": "value2"}

	section := Section{sectionName, sectionDictionary}

	t.Run("Get the Name Of the Section", func(t *testing.T) {
		got := section.Name()
		want := "section1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("Get value of a key of a Section", func(t *testing.T) {
		got, err := section.GetValue("key1")
		want := "value1"
		if err != nil {
			t.Errorf("Error %q", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

package main

import "testing"

//Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the []
func TestName(t *testing.T) {

	section := Section{"Doctor", string[]{}, Dictionary{}, string[]{}}

	t.Run("Get the Name Of the Section", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		t.Errorf("got %q want %q", got, want)

	})

}

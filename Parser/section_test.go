package main

import "testing"

//Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the []
func TestName(t *testing.T) {
	sectionName := "section1"
	var commentList []string
	var keyList []string
	var sectionDictionary Dictionary
	var sectionDictionaryMap = make(map[string]string)
	sectionDictionary = Dictionary(sectionDictionaryMap)

	section := Section{sectionName, commentList, keyList, sectionDictionary}

	t.Run("Get the Name Of the Section", func(t *testing.T) {
		got := section.Name()
		want := "section1"
		t.Errorf("got %q want %q", got, want)

	})

}

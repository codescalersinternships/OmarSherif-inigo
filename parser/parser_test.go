package main

import (
	"reflect"
	"testing"
)

//Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the []
func TestParser(t *testing.T) {
	const input = `[owner]
	name = John Doe
	organization = Acme Widgets Inc.

	[database]
	; use IP address in case network name resolution is not working
	server = 192.0.2.62
	port = 143
	file = "payroll.dat"`

	parser := NewParser()
	parser.LoadFromString(input)

	t.Run("get the names of all the sections", func(t *testing.T) {
		got, err := parser.GetSectionNames()
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		want := []string{"owner", "database"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("get a value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("owner", "name")
		want := "John Doe"
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns the section dictionary", func(t *testing.T) {
		got, err := parser.GetSections()
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		want := sectionDictionary{map[string]Section{"owner": {"owner",
			map[string]string{"name": "John Doe", "organization": "Acme Widgets Inc."}},
			"database": {"database", map[string]string{"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns the section", func(t *testing.T) {
		got, err := parser.GetSection("owner")
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		want := Section{"owner", map[string]string{"name": "John Doe", "organization": "Acme Widgets Inc."}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns the value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("owner", "name")
		want := "John Doe"
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns the value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("owner", "organization")
		want := "Acme Widgets Inc."
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns the value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("database", "server")
		want := "192.0.2.62"
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns the value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("database", "port")
		want := "143"
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns the value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("database", "file")
		want := "\"payroll.dat\""
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns the value of a key in a section", func(t *testing.T) {
		got, err := parser.Get("database", "unknown")
		want := ""
		if err != ErrKeyNotFound {
			t.Errorf("Error: %v", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

package main

import (
	"strings"
	"testing"
)

func TestModelView(t *testing.T) {
	m := model{}
	view := m.View()
	if !strings.Contains(view, "Metsaman!") {
		t.Errorf("View() = %q; want it to contain \"Metsaman!\"", view)
	}
}

package provider

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		id   string
		want string
	}{
		{"google", "Gemini"},
		{"openai", "ChatGPT"},
		{"anthropic", "Claude"},
		{"unknown", "Generic"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			p := Get(tt.id)
			if p.Name() != tt.want {
				t.Errorf("Get(%q).Name() = %q, want %q", tt.id, p.Name(), tt.want)
			}
		})
	}
}

func TestDetect(t *testing.T) {
	tests := []struct {
		name  string
		tools []string
		want  string
	}{
		{"google match", []string{"run_shell_command"}, "google"},
		{"openai match", []string{"code_interpreter"}, "openai"},
		{"anthropic match", []string{"computer"}, "anthropic"},
		{"no match", []string{"ls", "cat"}, "generic"},
		{"case insensitive match", []string{"RUN_SHELL_COMMAND"}, "google"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Detect(tt.tools)
			if got != tt.want {
				t.Errorf("Detect(%v) = %q, want %q", tt.tools, got, tt.want)
			}
		})
	}
}

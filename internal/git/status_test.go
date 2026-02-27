package git

import (
	"reflect"
	"testing"
)

func TestParsePorcelainStatus(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []StagedFile
	}{
		{
			name:  "added file",
			input: "A  main.go\n",
			expected: []StagedFile{
				{Path: "main.go", Status: Added},
			},
		},
		{
			name:  "modified file",
			input: "M  README.md\n",
			expected: []StagedFile{
				{Path: "README.md", Status: Modified},
			},
		},
		{
			name:  "deleted file",
			input: "D  old.txt\n",
			expected: []StagedFile{
				{Path: "old.txt", Status: Deleted},
			},
		},
		{
			name:  "mixed files",
			input: "A  new.go\nM  existing.go\n?? untracked.go\n",
			expected: []StagedFile{
				{Path: "new.go", Status: Added},
				{Path: "existing.go", Status: Modified},
			},
		},
		{
			name:  "renamed file",
			input: "R  old.go -> new.go\n",
			expected: []StagedFile{
				{Path: "new.go", Status: Renamed},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParsePorcelainStatus(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParsePorcelainStatus() = %v, want %v", got, tt.expected)
			}
		})
	}
}

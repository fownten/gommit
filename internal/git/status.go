package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type FileStatus string

const (
	Added    FileStatus = "Added"
	Modified FileStatus = "Modified"
	Deleted  FileStatus = "Deleted"
	Unknown  FileStatus = "Unknown"
)

type StagedFile struct {
	Path   string
	Status FileStatus
}

func GetStagedFiles() ([]StagedFile, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run git status: %w", err)
	}

	return ParsePorcelainStatus(string(output)), nil
}

func ParsePorcelainStatus(output string) []StagedFile {
	lines := strings.Split(output, "\n")
	var staged []StagedFile

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}

		// In porcelain mode, the first column is the index status
		indexStatus := line[0]
		path := strings.TrimSpace(line[3:])

		var status FileStatus
		switch indexStatus {
		case 'A':
			status = Added
		case 'M':
			status = Modified
		case 'D':
			status = Deleted
		default:
			// If it's not in the index (staged), we skip it for this tool
			continue
		}

		staged = append(staged, StagedFile{
			Path:   path,
			Status: status,
		})
	}

	return staged
}

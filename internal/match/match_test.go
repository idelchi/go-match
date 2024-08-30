package match_test

import (
	"reflect"
	"testing"

	"code.swisscom.com/swisscom/scsa-shared-tools/go-match/internal/match"
)

func TestMatch(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name     string
		paths    []string
		includes []string
		excludes []string
		want     []string
		wantErr  bool
	}{
		{
			name:     "Match all with no patterns",
			paths:    []string{"file1.txt", "file2.go", "dir/file3.js"},
			includes: []string{},
			excludes: []string{},
			want:     []string{"file1.txt", "file2.go", "dir/file3.js"},
			wantErr:  false,
		},
		{
			name:     "Include *.txt files",
			paths:    []string{"file1.txt", "file2.go", "dir/file3.txt"},
			includes: []string{"*.txt"},
			excludes: []string{},
			want:     []string{"file1.txt"},
			wantErr:  false,
		},
		{
			name:     "Include **/*.txt files",
			paths:    []string{"file1.txt", "file2.go", "dir/file3.txt", "dir/subdir/file4.txt"},
			includes: []string{"**/*.txt"},
			excludes: []string{},
			want:     []string{"file1.txt", "dir/file3.txt", "dir/subdir/file4.txt"},
			wantErr:  false,
		},
		{
			name:     "Exclude *.go files",
			paths:    []string{"file1.txt", "file2.go", "dir/file3.js"},
			includes: []string{},
			excludes: []string{"*.go"},
			want:     []string{"file1.txt", "dir/file3.js"},
			wantErr:  false,
		},
		{
			name:     "Exclude **/*.go files",
			paths:    []string{"file1.txt", "file2.go", "dir/file3.js", "dir/subdir/file.go"},
			includes: []string{},
			excludes: []string{"**/*.go"},
			want:     []string{"file1.txt", "dir/file3.js"},
			wantErr:  false,
		},
		{
			name:     "Include *.txt and exclude dir/*",
			paths:    []string{"file1.txt", "file2.go", "dir/file3.txt"},
			includes: []string{"*.txt"},
			excludes: []string{"dir/*"},
			want:     []string{"file1.txt"},
			wantErr:  false,
		},
		{
			name:     "Include **/*.txt (recursive)",
			paths:    []string{"file1.txt", "dir/file2.txt", "dir/subdir/file3.txt"},
			includes: []string{"**/*.txt"},
			excludes: []string{},
			want:     []string{"file1.txt", "dir/file2.txt", "dir/subdir/file3.txt"},
			wantErr:  false,
		},
		{
			name:     "Exclude **/subdir/* (recursive)",
			paths:    []string{"file1.txt", "dir/file2.txt", "dir/subdir/file3.txt"},
			includes: []string{},
			excludes: []string{"**/subdir/*"},
			want:     []string{"file1.txt", "dir/file2.txt"},
			wantErr:  false,
		},
		{
			name:     "Invalid include pattern",
			paths:    []string{"file1.txt"},
			includes: []string{"[invalid"},
			excludes: []string{},
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "Invalid exclude pattern",
			paths:    []string{"file1.txt"},
			includes: []string{},
			excludes: []string{"[invalid"},
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := match.Match(tc.paths, tc.includes, tc.excludes)
			if (err != nil) != tc.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tc.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Test %q: Match() = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

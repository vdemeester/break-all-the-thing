package strings

import (
	"fmt"
	"testing"
)

// mergePaths takes in two string slices and returns the combination of them
// without any duplicate entries.
func mergePaths(a, b []string) []string {
	newPaths := make([]string, 0, len(a)+len(b))
	newPaths = append(newPaths, a...)
	for p := 0; p < len(b); p++ {
		if !containsString(newPaths, b[p]) {
			newPaths = append(newPaths, b[p])
		}
	}
	return newPaths
}

// containsString takes in a string slice and looks for the provided string
// within the slice.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func mergePathsWithMap(a, b []string) []string {
	newPaths := make([]string, 0, len(a)+len(b))
	pathsMap := make(map[string]bool)
	for p := 0; p < len(a); p++ {
		pathsMap[a[p]] = true
	}
	for p := 0; p < len(b); p++ {
		pathsMap[b[p]] = true
	}
	for k, _ := range pathsMap {
		newPaths = append(newPaths, k)
	}
	return newPaths
}

func BenchmarkSliceContaining(b *testing.B) {
	for n := 1; n <= 16; n *= 2 {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				t1 := []string{"A", "B", "C"}
				t2 := []string{"B", "D", "E"}
				c := mergePaths(t1, t2)
				if len(c) != 5 {
					b.Error(".")
				}
			}
		})
	}
}

func BenchmarkWithMap(b *testing.B) {
	for n := 1; n <= 16; n *= 2 {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				t1 := []string{"A", "B", "C"}
				t2 := []string{"B", "D", "E"}
				c := mergePathsWithMap(t1, t2)
				if len(c) != 5 {
					b.Error(".")
				}
			}
		})
	}
}

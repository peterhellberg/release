// Package release is used to parse scene release names
package release

import (
	"regexp"
	"strconv"
	"strings"
)

// Release represents a scene release
type Release struct {
	Title      string
	Year       int
	Type       string
	Resolution string
	Format     string
	Group      string
}

var (
	yearPattern       = regexp.MustCompile(`[1,2]\d{3}`)
	groupPattern      = regexp.MustCompile(`-[[:alnum:]]+(\z|\s|\[)`)
	groupNamePattern  = regexp.MustCompile(`[[:alnum:]]+`)
	resolutionPattern = regexp.MustCompile(`(720|1080)p`)
	typePattern       = regexp.MustCompile(`\.(Unrated|UNRATED|INTERNAL|LiMiTED|LIMITED|Remastered\.DC|DC|MULTiSubs|PROPER|EXTENDED|Extended|IMAX\.EDITION)\.`)
	formatPattern     = regexp.MustCompile(`[\.\s][xXh]264[\-\.\s]`)
)

// Parse parses a release name
func Parse(s string) (*Release, error) {
	r := Release{}

	// Find the group
	if loc := groupPattern.FindStringIndex(s); loc != nil {
		r.Group = groupNamePattern.FindString(s[loc[0]+1:])
	}

	// Find the format
	if loc := formatPattern.FindStringIndex(s); loc != nil {
		r.Format = strings.ToLower(s[loc[0]+1:])
	}

	// Find the resolution
	if loc := resolutionPattern.FindStringIndex(s); loc != nil {
		r.Resolution = s[loc[0]:loc[1]]
		r.Title = cleanString(s[:loc[0]])
	}

	// Find the type
	if loc := typePattern.FindStringIndex(s); loc != nil {
		r.Type = strings.Replace(strings.ToUpper(s[loc[0]+1:loc[1]-1]), ".", " ", -1)
		r.Title = cleanString(s[:loc[0]])
	}

	// Try to find the year
	if loc := yearPattern.FindStringIndex(s); loc != nil {
		year, err := strconv.Atoi(s[loc[0]:loc[1]])
		if err == nil {
			r.Year = year
			r.Title = cleanString(s[:loc[0]])
		}
	}

	// Try to find the group again
	if r.Group == "" {
		if loc := formatPattern.FindStringIndex(s); loc != nil {
			r.Group = s[loc[1]:]
		}
	}

	return &r, nil
}

func cleanString(s string) string {
	return strings.Trim(strings.Replace(s, ".", " ", -1), " (")
}

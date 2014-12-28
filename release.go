// Package release is used to parse scene release names
package release

import (
	"regexp"
	"strconv"
	"strings"
)

// Release represents a scene release
type Release struct {
	Input      string
	Title      string
	Year       int
	Type       string
	Category   string
	Resolution string
	Format     string
	Audio      string
	Group      string
}

var (
	yearPattern       = regexp.MustCompile(`[1,2]\d{3}`)
	groupPattern      = regexp.MustCompile(`-[[:alnum:]]+(\z|\s|\[)`)
	groupNamePattern  = regexp.MustCompile(`[[:alnum:]]+`)
	resolutionPattern = regexp.MustCompile(`(720|1080)p`)
	typePattern       = regexp.MustCompile(`\.(Unrated|UNRATED|INTERNAL|LiMiTED|LIMITED|Remastered\.DC|DC|MULTiSubs|PROPER|EXTENDED|Extended|IMAX\.EDITION)\.`)
	bluRayPattern     = regexp.MustCompile(`[\.\s](BluRay|Bluray|Blu\.Ray|BRRip|BDRIP|BDRip)[\.\s]`)
	categoryPattern   = regexp.MustCompile(`[\.\s](HDRip|WEB-DL|Webrip)[\.\s]`)
	formatPattern     = regexp.MustCompile(`[\.\s]([xXhH]264|XviD)[\-\.\s]`)
	audioPattern      = regexp.MustCompile(`(AAC|AC3|DTS|DD5\.1)`)
)

// Parse parses a release name
func Parse(s string) (*Release, error) {
	r := Release{Input: s}

	// Find the group
	if loc := groupPattern.FindStringIndex(s); loc != nil {
		r.Group = groupNamePattern.FindString(s[loc[0]+1:])
	}

	// Find the format
	if loc := formatPattern.FindStringIndex(s); loc != nil {
		r.Format = strings.ToLower(s[loc[0]+1 : loc[1]-1])
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

	// Find the category, try different versions of BluRay first
	if loc := bluRayPattern.FindStringIndex(s); loc != nil {
		r.Category = "BluRay"
	} else if loc := categoryPattern.FindStringIndex(s); loc != nil {
		r.Category = s[loc[0]+1 : loc[1]-1]
	}

	// Try to find the year
	if loc := yearPattern.FindStringIndex(s); loc != nil {
		year, err := strconv.Atoi(s[loc[0]:loc[1]])
		if err == nil {
			r.Year = year
			r.Title = cleanString(s[:loc[0]])
		}
	}

	// Find the audio
	if loc := audioPattern.FindStringIndex(s); loc != nil {
		r.Audio = s[loc[0]:loc[1]]
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

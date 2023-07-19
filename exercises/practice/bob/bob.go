// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob


import (
	"strings"
	"unicode"
)


// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.TrimSpace(remark)
	
	if isSilence(remark) {
		return "Fine. Be that way!"
	}

	switch {
	case isQuestion(remark) && isYell(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isYell(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYell(remark string) bool {
	if strings.IndexFunc(remark, unicode.IsLetter) == -1 {
		return false
	}

	return strings.ToUpper(remark) == remark
}

func isSilence(remark string) bool {
	return remark == ""
}
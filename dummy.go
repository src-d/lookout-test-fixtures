// just a file for test

package fixtures

// NoopYes is a function that does absolutely nothing but returns true. I need to make this comment long enough so it would be catched by any linter that doesn't like long lines.
func NoopYes() bool {
	// new very long comment to trigger analyzer moar moar moar text moar moar moar text
	return true
	// another new very long comment to trigger analyzer moar moar moar text moar moar moar text
}

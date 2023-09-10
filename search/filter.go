package search

import "regexp"

var jsonRegex = regexp.MustCompile(`.+\.json$`)
var dbgJsonRegex = regexp.MustCompile(`.+\.dbg\.json$`)

func IsArtifactFile(name string) bool {
	return jsonRegex.MatchString(name) && !dbgJsonRegex.MatchString(name)
}

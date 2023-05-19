package regutil

import (
	"regexp"
	"strings"
)

const RegBase58 = "^[A-HJ-NP-Za-km-z1-9]+$"
const metabloxDidPrefix = "did:metablox:"

var regBase58 = regexp.MustCompile(RegBase58)

func IsBase58(s string) bool {
	return regBase58.MatchString(s)
}

func IsMetabloxDid(did string) bool {
	if len(did) == len(metabloxDidPrefix)+44 {
		return strings.HasPrefix(did, metabloxDidPrefix) && regBase58.MatchString(did[len(metabloxDidPrefix):])
	}
	return false
}

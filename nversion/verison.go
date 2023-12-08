package nversion

import (
	"errors"
	"fmt"
	"strings"
)

func CheckVersionBool(version string) bool {
	return len(version) > 0 && version[0:1] == "v" &&
		len(strings.Split(strings.TrimPrefix(version, "v"), ".")) == 3
}

func CheckVersion(version string) error {
	if len(version) == 0 || version[0:1] != "v" { // make sure have a v at the start v0.1.1
		return errors.New(fmt.Sprintf("incorrect provided: %s version number try: v1.2.3", version))
	}
	p := strings.Split(strings.TrimPrefix(version, "v"), ".")
	if len(p) != 3 {
		return errors.New(fmt.Sprintf("incorrect length provided: %s version number try: v1.2.3", version))
	}
	return nil
}

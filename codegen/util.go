package codegen

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type goversion struct {
	major int
	minor int
	patch int
}

func goVersion() (goversion, error) {
	v := runtime.Version()
	errmsg := fmt.Errorf("can't parse go version string \"%v\"", v)
	if !strings.HasPrefix(v, "go") {
		return goversion{}, errmsg

	}
	v = strings.TrimPrefix(v, "go")
	values := strings.Split(v, ".")
	if len(values) != 3 {
		return goversion{}, errmsg
	}
	major, e1 := strconv.Atoi(values[0])
	minor, e2 := strconv.Atoi(values[1])
	patch, e3 := strconv.Atoi(values[2])
	if e1 != nil || e2 != nil || e3 != nil {
		return goversion{}, errmsg
	}
	return goversion{major, minor, patch}, nil
}

// cmpGoVersion compares v1 and v2
// if v1 > v2  returns  1
// if v1 == v2 returns  0
// if v1 < v2  returns -1
func cmpGoVersion(v1, v2 goversion) int {
	if v1.major > v2.major {
		return 1
	} else if v1.major < v2.major {
		return -1
	} else if v1.minor > v2.minor {
		return 1
	} else if v1.minor < v2.minor {
		return -1
	} else if v1.patch > v2.patch {
		return 1
	} else if v1.patch < v2.patch {
		return -1
	} else {
		return 0
	}

}

// ice (internal compiler error) calls panic with "Internal error " + msg.
func ice(msg string) string {
	panic(fmt.Sprintf("Internal error, \"%v\"", msg))
}

func addIndent(assembly, indent string) string {
	lines := strings.Split(assembly, "\n")
	indented := ""
	for _, line := range lines {
		// skip debug comments
		indented += indent + line + "\n"
	}
	return indented
}

func stripDebug(assembly, indent string) string {
	lines := strings.Split(assembly, "\n")
	stripped := ""
	begin := indent + "// BEGIN"
	end := indent + "// END"
	for _, line := range lines {
		// skip debug comments
		if strings.HasPrefix(line, begin) || strings.HasPrefix(line, end) {
			continue
		}
		stripped += line + "\n"
	}
	return stripped
}

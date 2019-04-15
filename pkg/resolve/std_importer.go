package resolve

import (
	"path/filepath"
	"strings"

	"github.com/jkcfg/jk/pkg/std"
)

const (
	stdPrefix = "@jkcfg/std"
)

// StdImporter is the standard library importer.
type StdImporter struct {
	// PublicModules are modules users are allowed to import.
	PublicModules []string
	Source        []byte
}

func isStdModule(name string) bool {
	return strings.HasPrefix(name, stdPrefix)
}

func (i *StdImporter) isPublicModule(path string) bool {
	for _, name := range i.PublicModules {
		if name == path {
			return true
		}
	}
	return false
}

// Import implements importer.
func (i *StdImporter) Import(basePath, specifier, referrer string) ([]byte, string, []Candidate) {
	candidate := []Candidate{{specifier, staticRule}}

	// Short circuit the lookup when:
	//  - we're not trying to load a @jkcfg/std.* module
	//  - we're not inside the std library resolution
	if !isStdModule(specifier) && !strings.HasPrefix(basePath, stdPrefix) {
		return nil, "", candidate
	}

	path := specifier
	if isStdModule(path) {
		path = specifier[len(stdPrefix):]
	}
	if path == "" {
		path = "std.js"
	}
	if !strings.HasSuffix(path, ".js") {
		path += ".js"
	}

	// fmt.Printf("import %s from %s (basePath=%s, path=%s)\n", specifier, referrer, basePath, path)

	// Ensure we only allow users to import PublicModules. Modules from the std lib
	// itself are allowed to import internal private modules.
	if !isStdModule(referrer) && !i.isPublicModule(path) {
		return nil, "", candidate
	}

	module := std.Module(path)
	if len(module) == 0 {
		return nil, "", candidate
	}

	return module, filepath.Join(stdPrefix, path), candidate
}

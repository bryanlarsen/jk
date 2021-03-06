package std

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"sync"
)

// ModuleResources keeps track of the base paths for modules, as well
// as generating the magic modules when they are imported.
type ModuleResources struct {
	mu sync.RWMutex
	// module hash -> basePath for resource reads
	modules map[string]string
	salt    []byte
}

// NewModuleResources initialises a new ModuleResources
func NewModuleResources() *ModuleResources {
	r := &ModuleResources{
		modules: map[string]string{},
	}
	r.salt = make([]byte, 32)
	rand.Read(r.salt)
	return r
}

// ResourceBase provides the module base path given the hash.
func (r *ModuleResources) ResourceBase(hash string) (string, bool) {
	r.mu.RLock()
	path, ok := r.modules[hash]
	r.mu.RUnlock()
	return path, ok
}

// MakeModule generates resource module code (and path) given the
// importing module's base path.
func (r *ModuleResources) MakeModule(basePath string) ([]byte, string) {
	hash := sha256.New()
	hash.Write([]byte(basePath))
	hash.Write(r.salt)
	moduleHash := fmt.Sprintf("%x", hash.Sum(nil))
	r.mu.Lock()
	r.modules[moduleHash] = basePath
	r.mu.Unlock()

	code := `
import std from '@jkcfg/std';

const module = %q;

function read(path, {...rest} = {}) {
  return std.read(path, {...rest, module});
}

function dir(path) {
  return std.dir(path, { module });
}

function info(path) {
  return std.info(path, { module });
}

export { read, dir, info };
`
	return []byte(fmt.Sprintf(code, moduleHash)), "resource:" + basePath
}

package godi

import (
	"fmt"
	"runtime/debug"
)

// Def is a dependency definition
type Def struct {
	// Name is a dependency name
	Name string

	// Build builds dependency object
	Build BuildFn

	// Validate validates dependency definition on add
	Validate ValidateFn

	// Close finalizes dependency object
	Close CloseFn

	// Lazy is a flag. If true, Build will be executed only on Container.Get() call.
	Lazy bool

	obj   any
	built bool
}

// build builds dependency's object
func (d *Def) build(ctn *Container) error {
	if d.built {
		return nil
	}

	if d.Build == nil {
		return fmt.Errorf("%s: %w", d.Name, ErrBuildFunctionMissing)
	}

	var buildErr error

	func() {
		defer func() {
			if r := recover(); r != nil {
				stack := string(debug.Stack())

				buildErr = fmt.Errorf("build paniced: %v; stack: %s", r, stack)
			}
		}()

		d.obj, buildErr = d.Build(ctn)
	}()

	d.built = true

	if buildErr != nil {
		d.obj = nil

		return buildErr
	}

	return nil
}

// definitions is a dependencies definitions map
type definitions map[string]Def

// ValidateFn is a dependency validation function
type ValidateFn func(ctn *Container) (err error)

// BuildFn is a dependency build function
type BuildFn func(ctn *Container) (obj any, err error)

// CloseFn is a dependency close function
type CloseFn func(obj any) (err error)

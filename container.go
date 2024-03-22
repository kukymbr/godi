package godi

import (
	"errors"
	"fmt"
)

// Container is a dependency container
type Container struct {
	defs definitions
}

// Has checks if dependency is registered in Container
func (c *Container) Has(name string) bool {
	_, ok := c.defs[name]

	return ok
}

// Get returns built dependency. Panics on error.
func (c *Container) Get(name string) (obj any) {
	obj, err := c.SafeGet(name)
	if err != nil {
		panic(err.Error())
	}

	return obj
}

// SafeGet returns built dependency
func (c *Container) SafeGet(name string) (obj any, err error) {
	def, ok := c.defs[name]
	if !ok {
		return nil, fmt.Errorf("%s: %w", name, ErrDefinitionNotFound)
	}

	if def.Lazy {
		err = def.build(c)
		if err != nil {
			return nil, err
		}

		c.defs[name] = def
	}

	return def.obj, nil
}

// Len returns count of definitions in the Container
func (c *Container) Len() int {
	return len(c.defs)
}

// Close finalizes dependencies
func (c *Container) Close() (err error) {
	for _, def := range c.defs {
		if !def.built {
			continue
		}

		if def.Close == nil {
			continue
		}

		defErr := def.Close(def.obj)
		if defErr != nil {
			err = errors.Join(err, fmt.Errorf("%s: %w", def.Name, defErr))
		}
	}

	return err
}

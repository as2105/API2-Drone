// Copyright © 2018 Optum
package persisteth

type ResourceStore interface {
	Create(interface{}) (string, string, error)
	Read(uuid string, version string) (interface{}, error)
	Update(uuid string) (interface{}, error)
	Delete(uuid string) error
}

package harmonyclient

import "errors"

// ContainerEnv holds a container's environment variable
type ContainerEnv struct {
	ID          string `jsonapi:"name=id"`
	ContainerID string `jsonapi:"name=container_id"`
	Name        string `jsonapi:"name=name"`
	Value       string `jsonapi:"name=value"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerEnv) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerEnv) SetID(id string) error {
	c.ID = id
	return nil
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *ContainerEnv) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		// we can skip this, becase MachineID should already be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

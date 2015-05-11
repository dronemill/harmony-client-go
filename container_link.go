package harmonyclient

import "errors"

// ContainerLink holds a container's link resource
type ContainerLink struct {
	ID              string `jsonapi:"name=id"`
	ContainerID     string `jsonapi:"name=container_id"`
	ContainerFromID string `jsonapi:"name=container_from_id"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerLink) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerLink) SetID(id string) error {
	c.ID = id
	return nil
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *ContainerLink) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		// we can skip this, becase MachineID should already be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

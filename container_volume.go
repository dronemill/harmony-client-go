package harmonyclient

import "errors"

// ContainerVolume holds a container's volume
type ContainerVolume struct {
	ID            string `jsonapi:"name=id"`
	ContainerID   string `jsonapi:"name=container_id"`
	PathHost      string `jsonapi:"name=path_host"`
	PathContainer string `jsonapi:"name=path_container"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerVolume) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerVolume) SetID(id string) error {
	c.ID = id
	return nil
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *ContainerVolume) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		// we can skip this, becase MachineID should already be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

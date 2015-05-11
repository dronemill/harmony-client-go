package harmonyclient

import "errors"

// ContainerExpose holds a container's expose resource
type ContainerExpose struct {
	ID          string `jsonapi:"name=id"`
	ContainerID string `jsonapi:"name=container_id"`
	RangeStart  string `jsonapi:"name=range_start"`
	RangeEnd    string `jsonapi:"name=range_end"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerExpose) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerExpose) SetID(id string) error {
	c.ID = id
	return nil
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *ContainerExpose) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		// we can skip this, becase MachineID should already be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

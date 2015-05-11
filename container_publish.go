package harmonyclient

import "errors"

// ContainerPublish holds a container's publish resource
type ContainerPublish struct {
	ID            string `jsonapi:"name=id"`
	ContainerID   string `jsonapi:"name=container_id"`
	ContainerPort string `jsonapi:"name=container_port"`
	HostPort      string `jsonapi:"name=host_port"`
	IP            string `jsonapi:"name=ip"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerPublish) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerPublish) SetID(id string) error {
	c.ID = id
	return nil
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *ContainerPublish) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		// we can skip this, becase MachineID should already be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

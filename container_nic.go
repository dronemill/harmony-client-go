package harmonyclient

import "errors"

// ContainerNic holds a container's nic
type ContainerNic struct {
	ID           string `jsonapi:"name=id"`
	ContainerID  string `jsonapi:"name=container_id"`
	BridgeDev    string `jsonapi:"name=bridge_dev"`
	ContainerDev string `jsonapi:"name=container_dev"`
	IP           string `jsonapi:"name=ip"`
	Netmask      string `jsonapi:"name=netmask"`
	Gateway      string `jsonapi:"name=gateway"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerNic) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerNic) SetID(id string) error {
	c.ID = id
	return nil
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *ContainerNic) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		// we can skip this, becase MachineID should already be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

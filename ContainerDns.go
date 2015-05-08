package harmonyclient

// ContainerDns holds a container's dns resource
type ContainerDns struct {
	ID          string `jsonapi:"name=id"`
	ContainerID string `jsonapi:"name=container_id"`
	Nameserver  string `jsonapi:"name=nameserver"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ContainerDns) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ContainerDns) SetID(id string) error {
	c.ID = id
	return nil
}

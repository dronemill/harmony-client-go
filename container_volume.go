package harmonyclient

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

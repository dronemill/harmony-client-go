package harmonyclient

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

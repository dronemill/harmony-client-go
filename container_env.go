package harmonyclient

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

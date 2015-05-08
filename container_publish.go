package harmonyclient

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

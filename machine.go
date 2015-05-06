package harmonyclient

// Machine holds a Harmony Machine
type Machine struct {
	ID       string `jsonapi:"name=id"`
	Name     string `jsonapi:"name=name"`
	Hostname string `jsonapi:"name=hostname"`
	IP       string `jsonapi:"name=ip"`

	ContainerIDs []string `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Machine) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Machine) SetID(id string) error {
	c.ID = id
	return nil
}

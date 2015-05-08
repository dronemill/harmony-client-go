package harmonyclient

import (
	"errors"

	"github.com/univedo/api2go/jsonapi"
)

// Container is a generic database user
type Container struct {
	ID         string `jsonapi:"name=id"`
	MachineID  string `jsonapi:"name=machine_id"`
	CID        string `jsonapi:"name=cid"`
	Name       string `jsonapi:"name=name"`
	Hostname   string `jsonapi:"name=hostname"`
	Restart    string `jsonapi:"name=restart"`
	Image      string `jsonapi:"name=image"`
	EntryPoint string `jsonapi:"name=entry_point"`
	Enabled    bool   `jsonapi:"name=enabled"`
	CreatedAt  string `jsonapi:"name=created_at"`
	UpdatedAt  string `jsonapi:"name=updated_at"`

	ContainerDns          []ContainerDns     `json:"-"`
	ContainerDnsIDs       []string           `json:"-"`
	ContainerEnvs         []ContainerEnv     `json:"-"`
	ContainerEnvsIDs      []string           `json:"-"`
	ContainerExposes      []ContainerExpose  `json:"-"`
	ContainerExposesIDs   []string           `json:"-"`
	ContainerLinks        []ContainerLink    `json:"-"`
	ContainerLinksIDs     []string           `json:"-"`
	ContainerNics         []ContainerNic     `json:"-"`
	ContainerNicsIDs      []string           `json:"-"`
	ContainerPublishes    []ContainerPublish `json:"-"`
	ContainerPublishesIDs []string           `json:"-"`
	ContainerVolumes      []ContainerVolume  `json:"-"`
	ContainerVolumesIDs   []string           `json:"-"`
	Machine               Machine            `json:"-"`
}

// "container_envs": { … },
// "container_volumes": { … },
// "container_nics": { … },
// "container_dns": { … },
// "container_exposes": { … },
// "container_links": { … },
// "container_publishs": { … },

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Container) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Container) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (c Container) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "container_dns",
			Name: "container_dns",
		},
		{
			Type: "container_envs",
			Name: "container_envs",
		},
		{
			Type: "container_exposes",
			Name: "container_exposes",
		},
		{
			Type: "container_links",
			Name: "container_links",
		},
		{
			Type: "container_nics",
			Name: "container_nics",
		},
		{
			Type: "container_publishes",
			Name: "container_publishes",
		},
		{
			Type: "container_volumes",
			Name: "container_volumes",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (c Container) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	// handle ContainerDns
	for _, containerDns := range c.ContainerDns {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerDns.ID,
			Type: "container_dns",
			Name: "container_dns",
		})
	}

	// handle ContainerEnvs
	for _, containerEnv := range c.ContainerEnvs {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerEnv.ID,
			Type: "container_envs",
			Name: "container_envs",
		})
	}

	// handle ContainerExposes
	for _, containerExpose := range c.ContainerExposes {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerExpose.ID,
			Type: "container_exposes",
			Name: "container_exposes",
		})
	}

	// handle ContainerLinks
	for _, containerLink := range c.ContainerLinks {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerLink.ID,
			Type: "container_links",
			Name: "container_links",
		})
	}

	// handle ContainerNics
	for _, containerNic := range c.ContainerNics {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerNic.ID,
			Type: "container_nics",
			Name: "container_nics",
		})
	}

	// handle ContainerPublishes
	for _, containerPublish := range c.ContainerPublishes {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerPublish.ID,
			Type: "container_publishes",
			Name: "container_publishes",
		})
	}

	// handle ContainerVolumes
	for _, containerVolume := range c.ContainerVolumes {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerVolume.ID,
			Type: "container_volumes",
			Name: "container_volumes",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (c Container) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	// handle ContainerDns
	for key := range c.ContainerDns {
		result = append(result, c.ContainerDns[key])
	}

	// handle ContainerEnvs
	for key := range c.ContainerEnvs {
		result = append(result, c.ContainerEnvs[key])
	}

	// handle ContainerExposes
	for key := range c.ContainerExposes {
		result = append(result, c.ContainerExposes[key])
	}

	// handle ContainerLinks
	for key := range c.ContainerLinks {
		result = append(result, c.ContainerLinks[key])
	}

	// handle ContainerNics
	for key := range c.ContainerNics {
		result = append(result, c.ContainerNics[key])
	}

	// handle ContainerPublishes
	for key := range c.ContainerPublishes {
		result = append(result, c.ContainerPublishes[key])
	}

	// handle ContainerVolumes
	for key := range c.ContainerVolumes {
		result = append(result, c.ContainerVolumes[key])
	}

	return result
}

// 	SetToOneReferenceID sets the reference ID and satifices jsonapi.UnmarshalToOneRelations interface
func (c *Container) SetToOneReferenceID(name, ID string) error {
	if name == "machine" {
		// we can skip this, becase MachineID should alreadt be populated
	} else {
		return errors.New("There is no to-one relationship with name " + name)
	}

	return nil
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (c *Container) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "container_dns" {
		c.ContainerDnsIDs = IDs
	} else if name == "container_envs" {
		c.ContainerEnvsIDs = IDs
	} else if name == "container_exposes" {
		c.ContainerExposesIDs = IDs
	} else if name == "container_links" {
		c.ContainerLinksIDs = IDs
	} else if name == "container_nics" {
		c.ContainerNicsIDs = IDs
	} else if name == "container_publishes" {
		c.ContainerPublishesIDs = IDs
	} else if name == "container_volumes" {
		c.ContainerVolumesIDs = IDs
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

// AddToManyIDs adds some new sweets that a users loves so much
func (c *Container) AddToManyIDs(name string, IDs []string) error {
	if name == "container_dns" {
		c.ContainerDnsIDs = append(c.ContainerDnsIDs, IDs...)
	} else if name == "container_envs" {
		c.ContainerEnvsIDs = append(c.ContainerEnvsIDs, IDs...)
	} else if name == "container_exposes" {
		c.ContainerExposesIDs = append(c.ContainerExposesIDs, IDs...)
	} else if name == "container_links" {
		c.ContainerLinksIDs = append(c.ContainerLinksIDs, IDs...)
	} else if name == "container_nics" {
		c.ContainerNicsIDs = append(c.ContainerNicsIDs, IDs...)
	} else if name == "container_publishes" {
		c.ContainerPublishesIDs = append(c.ContainerPublishesIDs, IDs...)
	} else if name == "container_volumes" {
		c.ContainerVolumesIDs = append(c.ContainerVolumesIDs, IDs...)
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (c *Container) DeleteToManyIDs(name string, IDs []string) error {
	if name == "container_dns" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerDnsIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerDns = append(c.ContainerDns[:pos], c.ContainerDns[pos+1:]...)
				}
			}
		}
	} else if name == "container_envs" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerEnvsIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerEnvs = append(c.ContainerEnvs[:pos], c.ContainerEnvs[pos+1:]...)
				}
			}
		}
	} else if name == "container_exposes" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerExposesIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerExposes = append(c.ContainerExposes[:pos], c.ContainerExposes[pos+1:]...)
				}
			}
		}
	} else if name == "container_links" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerLinksIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerLinks = append(c.ContainerLinks[:pos], c.ContainerLinks[pos+1:]...)
				}
			}
		}
	} else if name == "container_nics" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerNicsIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerNics = append(c.ContainerNics[:pos], c.ContainerNics[pos+1:]...)
				}
			}
		}
	} else if name == "container_publishes" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerPublishesIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerPublishes = append(c.ContainerPublishes[:pos], c.ContainerPublishes[pos+1:]...)
				}
			}
		}
	} else if name == "container_volumes" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerVolumesIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerVolumes = append(c.ContainerVolumes[:pos], c.ContainerVolumes[pos+1:]...)
				}
			}
		}
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

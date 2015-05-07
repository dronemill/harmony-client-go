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

	ContainerEnvs       []ContainerEnv    `json:"-"`
	ContainerEnvsIDs    []string          `json:"-"`
	ContainerVolumes    []ContainerVolume `json:"-"`
	ContainerVolumesIDs []string          `json:"-"`
	Machine             Machine           `json:"-"`
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
			Type: "container_envs",
			Name: "container_envs",
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

	// handle ContainerEnvs
	for _, containerEnv := range c.ContainerEnvs {
		result = append(result, jsonapi.ReferenceID{
			ID:   containerEnv.ID,
			Type: "container_envs",
			Name: "container_envs",
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

	// handle ContainerEnvs
	for key := range c.ContainerEnvs {
		result = append(result, c.ContainerEnvs[key])
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
	if name == "container_envs" {
		c.ContainerEnvsIDs = IDs
	} else if name == "container_volumes" {
		c.ContainerVolumesIDs = IDs
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

// AddToManyIDs adds some new sweets that a users loves so much
func (c *Container) AddToManyIDs(name string, IDs []string) error {
	if name == "container_envs" {
		c.ContainerEnvsIDs = append(c.ContainerEnvsIDs, IDs...)
	} else if name == "container_volumes" {
		c.ContainerVolumesIDs = append(c.ContainerVolumesIDs, IDs...)
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (c *Container) DeleteToManyIDs(name string, IDs []string) error {
	if name == "container_envs" {
		for _, ID := range IDs {
			for pos, oldID := range c.ContainerEnvsIDs {
				if ID == oldID {
					// match, this ID must be removed
					c.ContainerEnvs = append(c.ContainerEnvs[:pos], c.ContainerEnvs[pos+1:]...)
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

package harmonyclient

import (
	"errors"

	"github.com/manyminds/api2go/jsonapi"
)

// Machine holds a Harmony Machine
type Machine struct {
	ID           string `jsonapi:"name=id"`
	Name         string `jsonapi:"name=name"`
	Hostname     string `jsonapi:"name=hostname"`
	IP           string `jsonapi:"name=ip"`
	Es_client_id string `jsonapi:"name=es_client_id"`
	CreatedAt    string `jsonapi:"name=created_at"`
	UpdatedAt    string `jsonapi:"name=updated_at"`

	Containers   []Container `json:"-"`
	ContainerIDs []string    `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (m Machine) GetID() string {
	return m.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (m *Machine) SetID(id string) error {
	m.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (m Machine) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "containers",
			Name: "containers",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (m Machine) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, container := range m.Containers {
		result = append(result, jsonapi.ReferenceID{
			ID:   container.ID,
			Type: "containers",
			Name: "containers",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (m Machine) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range m.Containers {
		result = append(result, m.Containers[key])
	}

	return result
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (m *Machine) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "containers" {
		m.ContainerIDs = IDs
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

// AddToManyIDs adds some new sweets that a users loves so much
func (m *Machine) AddToManyIDs(name string, IDs []string) error {
	if name == "containers" {
		m.ContainerIDs = append(m.ContainerIDs, IDs...)
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (m *Machine) DeleteToManyIDs(name string, IDs []string) error {
	if name == "containers" {
		for _, ID := range IDs {
			for pos, oldID := range m.ContainerIDs {
				if ID == oldID {
					// match, this ID must be removed
					m.Containers = append(m.Containers[:pos], m.Containers[pos+1:]...)
				}
			}
		}
	} else {
		return errors.New("There is no to-many relationship with the name " + name)
	}

	return nil
}

package main

import (
	"errors"
	"testing"

        "github.com/google/go-cmp/cmp"
)

type EntitiesStub struct {
    getUser func(id string) (User, error)
    getPets func(userID string) ([]Pet, error)
    saveUser func(user User) error
}

func (e EntitiesStub) GetPets(userID string) ([]Pet, error) {
    return e.getPets(userID)
}

func (e EntitiesStub) GetUser(id string) (User, error) {
    return e.getUser(id)
}

func (e EntitiesStub) SaveUser(user User) error {
    return e.saveUser(user)
}


func TestGetPetNames(t *testing.T) {
    tdata := []struct {
        name string
        getPets func(userID string) ([]Pet, error)
        UserID string
        petNames []string
        errMesg string
    }{
        {
            "case1",
            func(userID string) ([]Pet, error) {
                return []Pet{
                    {Name: "pet1"},
                    {Name: "pet2"},
                }, 
                nil
            }, 
            "user1", 
            []string{"pet1", "pet2"}, 
            "",
        },
        {
            "case2",
            func(userID string) ([]Pet, error) {
                return nil, errors.New("invalid id: user3")
            }, 
            "user3", 
            nil,
            "invalid id: user3",
        },
    }

    l := Logic{}
    for _, d := range tdata {
        t.Run(d.name, func(t *testing.T) {
            l.Entities = EntitiesStub{
                getPets: d.getPets,
            }
            petNames, err := l.GetPetNames(d.UserID)
            if diff := cmp.Diff(petNames, d.petNames); diff != "" {
                t.Errorf("petNames mismatch (-want +got):\n%s", diff)
            }
            if err != nil && err.Error() != d.errMesg {
                t.Errorf("errMesg mismatch: want \"%s\", got \"%s\"", d.errMesg, err.Error())
            }
        })
    }
}

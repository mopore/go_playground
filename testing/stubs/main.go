package main

type User struct {
    Name string
}

type Pet struct {
    Name string
}

type Entities interface {
    GetUser(id string) (User, error)
    GetPets(userID string) ([]Pet, error)
    SaveUser(user User) error
}

type Logic struct {
    Entities Entities
}

// We want to test this function
func (l Logic) GetPetNames(userID string) ([]string, error) {
    pets, err := l.Entities.GetPets(userID)
    if err != nil {
        return nil, err
    }

    var petNames []string
    for _, pet := range pets {
        petNames = append(petNames, pet.Name)
    }

    return petNames, nil
}


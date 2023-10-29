package main

import (
	"context"
	"log"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type YasmDatabase struct {
    Driver neo4j.DriverWithContext
    Context context.Context
}

func NewYasmDatabase(dbUri, dbUser, dbPass string) (YasmDatabase, error) {
    ctx := context.Background()
    driver, err := neo4j.NewDriverWithContext(
        dbUri, 
        neo4j.BasicAuth(dbUser, dbPass, ""))
    if err != nil {
        log.Fatal(err)
    }
    return YasmDatabase{
        Driver: driver,
        Context: ctx,
    }, nil
}

func (y YasmDatabase) QueryFiveSkills( skillName string) {
    searchParam := strings.ToLower(skillName)
    result, err := neo4j.ExecuteQuery(y.Context, y.Driver, 
        "MATCH (s:Skill) WHERE tolower(s.name) CONTAINS $skillname RETURN s.name, s.id LIMIT 5",
        map[string]any{"skillname": searchParam},
        neo4j.EagerResultTransformer,
        // neo4j.ExecuteQueryWithDatabase("neo4j"),
    )

    if err != nil {
        log.Fatal(err)
    }
    for _, record := range result.Records {
        log.Printf("Skill: %s, ID: %s", record.Values[0].(string), record.Values[1].(string))
    }
}

func (y YasmDatabase) CreateTimestamp(name string) {
    result, err := neo4j.ExecuteQuery(y.Context, y.Driver, 
        "CREATE (t:Timestamp {name: $name, time: timestamp()}) RETURN t.name, t.time",
        map[string]any{"name": name},
        neo4j.EagerResultTransformer,
    )

    if err != nil {
        log.Fatal(err)
    }
    for _, record := range result.Records {
        log.Printf("Timestamp: %s, Time: %d", record.Values[0].(string), record.Values[1].(int64))
    }
}

func (y YasmDatabase) Close() error {
    return y.Driver.Close(y.Context)
}

func main() {
    log.Println("Starting Neo4j driver example")
    
    dbUri := "bolt://localhost:7687"
    dbUser := "neo4j"
    dbPass := "password"
    y, err := NewYasmDatabase(dbUri, dbUser, dbPass)
    if err != nil {
        log.Fatal(err)
    }
    defer y.Close()

    y.QueryFiveSkills("Jakarta")
    y.CreateTimestamp("test")
}

package main

import (
    "log"
    "os"
    "encoding/json"
    //"github.com/kubeless/kubeless/pkg/functions"
    f "github.com/fauna/faunadb-go/v3/faunadb"
)


type Pet struct {
     Name string `fauna:"name", json:"name, omitempty`
     Age  int    `fauna:"age" , json:"age,  omitempty`
}

func main() {
    log.Println("ok")
    secret := os.Getenv("SECRET")
    dbUrl  := os.Getenv("FAUNA_URL") //http://localhost:8443

    log.Println("secret ", secret)
    log.Println("dbUrl ", dbUrl)

    client := f.NewFaunaClient(secret, f.Endpoint(dbUrl))
    log.Println("client ", client)    

   // log.Println(event.Data)
   inputPet := `{"name":"Fuffy", "age":4}`

    var pet Pet   
    json.Unmarshal([]byte(inputPet), &pet)

    client.Query(
        f.Create(
            f.Class("pets"),
            f.Obj{"data": pet},
        ),
    )

    res, _ := client.Query(
           f.Get(
            f.MatchTerm(
                f.Index("pets_by_name"),
                "Fuffy",
            ),
        ),
    )

     var pet2 Pet

     res.At(f.ObjKey("data")).Get(&pet2)
     log.Println(pet2)

     data, _ := json.Marshal(pet2)

    log.Println("final result: ", string(data))

}
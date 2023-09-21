package main

import (
	"encoding/json"
	"fmt"
)

type User struct  {
  Name string
  Age int32
}


func main() {
  reqUser := User{
    Name: "test",
    Age: 5,
  }

  s := make([]User, 0, 2)
  s = append(s, reqUser, reqUser)

  fmt.Println(s)

  j, err := json.Marshal(s)

  if err != nil {
    panic(err)
  }

  fmt.Println(string(j))

  var users [2]User

  err = json.Unmarshal(j,&users)

  for _, user := range users {
    fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
  }
}


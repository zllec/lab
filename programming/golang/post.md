# Common way of creating POST Request

1. Encode the data as json
2. Create a new POST request, set the request headers
3. Create a new HTTP Client and make the request
4. Decode the JSON data from the response

```go
package main

import (
 "bytes"
 "encoding/json"
 "net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
 // encode the data as json
 jsonData, err := json.Marshal(data)
 if err != nil {
  return User{}, err
 }

 // create a new POST request
 req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
 if err != nil {
  return User{}, err
 }

 // set request headers
 req.Header.Set("Content-Type", "application/json")
 req.Header.Set("X-API-Key", apiKey)

 // create a new client and make the request
 client := &http.Client{}
 res, err := client.Do(req)
 if err != nil {
  return User{}, err
 }
 defer res.Body.Close()

 // decode the json data from the response
 // into a new User struct
 var user User
 decoder := json.NewDecoder(res.Body)
 if err := decoder.Decode(&user); err != nil {
  return User{}, err
 }

 return user, err
}
```

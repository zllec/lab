# Common way of consuming a GET API

- Send GET request
- use json.NewDecoder(res.Body) to create a decoder to "parse" JSON
- call Decode to into a slice of struct
  - `decoder.Decode(&issues)`
- don't forget to close the body

### Decode

- From JSON payload to Struct using:
  - `json.Decoder`
  - `json.Unmarshal`
- `Decode` method from `json.Decoder` streams data from an io.reader into a Go struct
- Unmarshal works with data that's already in []byte format
- When working on a large JSON data, `Decoder` is ideal since it doesnt load all the data into memory
- When working with a smaller JSON data, `Unmarshal` is preferred

#### Sample Decoder

```go
package main

import (
 "fmt"
 "net/http"
 "encoding/json"
)

func getIssues(url string) ([]Issue, error) {
 res, err := http.Get(url)
 if err != nil {
  return nil, fmt.Errorf("error creating request: %w", err)
 }
 defer res.Body.Close()

 var issues []Issue
 decoder := json.NewDecoder(res.Body)
 // Decode method streams JSON into a Go struct
 if err := decoder.Decode(&issues); err != nil {
  fmt.Println("error decoding response body")
  return nil, err
 }
 return issues, nil
}
```

#### Sample Unmarshal

```go
package main

import (
 "encoding/json"
 "fmt"
 "io"
 "net/http"
)

func getIssues(url string) ([]Issue, error) {
 res, err := http.Get(url)
 if err != nil {
  return nil, fmt.Errorf("error creating request: %w", err)
 }
 defer res.Body.Close()

 // io.ReadAll returns a []byte which works well with Unmarshal
 data, err := io.ReadAll(res.Body)
 if err != nil {
  return nil, err
 }

 var issues [] Issue
 if err := json.Unmarshal(data, &issues); err != nil {
  return nil, err
 }
 return issues, nil
}
```

### For

- `for idx, item := range items`
  - range returns two items - index, value

### DNS

- domain name/ host name to IP address
- managed by ICANN

package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
)

type UserResponse struct {
  Github_health string `json:"Github_health"`
  Updated_at string `json:"Updated_at"`
}

type Response struct {
  Status IStatus
  Page IPage
}

type IStatus struct {
  Indicator string `json:"indicator"`
  Description string `json:"description"`
}

type IPage struct {
  Id string `json:"id"`
  Name string `json:"name"`
  url string `json:"url"`
  TimeZone string `json:"time_zone"`
  UpdatedAt string `json:"updated_at"`
}
func main() {
  fmt.Println("Hello World")

  // Get the url endpoint
  url := "https://kctbh9vrtdwd.statuspage.io/api/v2/status.json"

  // Make Http get request
  res, err := http.Get(url)

  // Check if error first
  // Then process it
  if err != nil {
    log.Fatal("Error : ", err)
    return
  }

  body, err2 := ioutil.ReadAll(res.Body)

  // Check if error first
  // Then process it
  if err2 != nil {
    log.Fatal("Error : ", err2)
    return
  }

  jsonBody := Response{}

  err3 := json.Unmarshal(body, &jsonBody)
  
  // Check if error first
  // Then process it
  if err3 != nil {
    log.Fatal("Error : ", err3)
    return
  }  
  updatedAt := jsonBody.Page.UpdatedAt
  githubHealth := jsonBody.Status.Indicator
  fmt.Println(githubHealth)

  responseObject := map[string] interface{} {
    "Github_health" : githubHealth,
    "Updated_at" : updatedAt,
  }

  jsonRes, errMarshal := json.Marshal(responseObject)
  if errMarshal != nil {
    log.Fatal("Error : ", errMarshal)
    return
  }  

  http.HandleFunc("/checkStatus", func(w http.ResponseWriter, r *http.Request){	
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
  })

  fmt.Println("Starting server at port 8080")
  if err4 := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err4)
  }

}
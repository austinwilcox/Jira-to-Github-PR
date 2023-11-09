package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
  jiraUrl := "arbinger.atlassian.net"
  ticketId := "DEV-1937"
  username := "awilcox@arbinger.com"
  output, err := exec.Command("./credentialHelper.sh").Output(); if err != nil {
    log.Fatal(err)
  }
  password := output

  response := getJiraTicketDescription(jiraUrl, ticketId, username, password)
  for _, resp := range response {
    fmt.Println(resp)
  }
}

func getJiraTicketDescription(jiraUrl string, ticketId string, username string, password []byte) []string {
  auth := username + ":" + string(password)
  authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

  req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/rest/api/3/issue/%s", jiraUrl, ticketId), nil); if err != nil {
    log.Fatal(err)
  }

  req.Header.Add("Authorization", authHeader)
  client := &http.Client{}
  response, err := client.Do(req); if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  responseData, err := io.ReadAll(response.Body); if err != nil {
    log.Fatal(err)
  }
  myData := JiraIssue{}
  json.Unmarshal(responseData, &myData)
  fmt.Println(myData)
  var myDescriptionText []string
  for _, content := range myData.Fields.Description.Content {
    for _, text := range content.Content {
      if text.Type == "text" {
        myDescriptionText = append(myDescriptionText, text.Text)
      }
    }
  }

  return myDescriptionText
}

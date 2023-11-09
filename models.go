package main

type JiraIssue struct {
    Fields struct {
      Description struct {
      Type    string `json:"type"`
      Version int    `json:"version"`
      Content []struct {
          Type    string `json:"type"`
          Content []struct {
              Type string `json:"type"`
              Text string `json:"text"`
          } `json:"content"`
          Attrs struct {
              Layout string `json:"layout"`
          } `json:"attrs"`
      } `json:"content"`
    } `json:"description"`
  } `json:"fields"`
}

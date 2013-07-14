package main

import (
  "github.com/NOX73/go-twitter-stream-api"
  "launchpad.net/goyaml"
  "io/ioutil"
  "fmt"
)

type Config struct {
  Credentials struct { Oauth_consumer_key, Oauth_token, Oauth_consumer_secret, Oauth_token_secret string }
  Params map[string]string
}

func main() {

  fmt.Println("Reading config ... ")

  config_yaml, _ := ioutil.ReadFile("config.yml")
  config := Config{}

  err := goyaml.Unmarshal(config_yaml, &config)
  if err != nil{ panic(err) }

  fmt.Println("Config: ", config)

  fmt.Println("Connecting to twitter api ... ")

  credentials := twitter_api.NewCredentials(
    config.Credentials.Oauth_consumer_key,
    config.Credentials.Oauth_token,
    config.Credentials.Oauth_consumer_secret,
    config.Credentials.Oauth_token_secret,
  ) 

  ch := make(chan twitter_api.Message, 10)

  go twitter_api.TwitterStream(ch, credentials, config.Params)

  for {
    message := <- ch

    if message.Error != nil {
      fmt.Printf("Error: %s \n", message.Error)
      fmt.Printf("Response: %s \n", message.Response)
      break
    }

    fmt.Printf("Channel length: %s \n", len(ch))
  }


}



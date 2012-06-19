// Package httputil implements some http utility functions.
package httputil

import "io/ioutil"
import "net/http"
import "strings"

// client is the default client used by httputil requests.
var client = http.DefaultClient

// SetClient sets the default client used by httputil requests.
func SetClient(c *http.Client) {
   client = c
}

// Post issues a POST to the specified URL.
func Post(rawUrl, bodyType, data string) (s string, err error) {
   res, err := client.Post(rawUrl, bodyType, strings.NewReader(data))
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   buf, err := ioutil.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   return string(buf), nil
}

// Get issues a GET to the specified URL.
func Get(rawUrl string) (s string, err error) {
   res, err := client.Get(rawUrl)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   buf, err := ioutil.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   return string(buf), nil
}
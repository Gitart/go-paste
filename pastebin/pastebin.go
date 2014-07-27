// Package pastebin wraps the basic functions of the Pastebin API and exposes a
// Go API.
package pastebin

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const pastebinDevKey = "d06a9df64b29123b8eeda23f53d6535d"

var PastebinPutError = errors.New("Pastebin Put Failed!")

// Function Put uploads text to Pastebin with optional title returning the ID or
// an error.
func Put(text, title string) (id string, err error) {
	data := url.Values{}
	// Required values.
	data.Set("api_dev_key", pastebinDevKey)
	data.Set("api_option", "paste") // Create a paste.
	data.Set("api_paste_code", text)
	// Optional values.
	data.Set("api_paste_name", title)      // The paste should have title "title".
	data.Set("api_paste_private", "0")     // Create a public paste.
	data.Set("api_paste_expire_date", "N") // The paste should never expire.

	// Parse and URLEncode the values ready to pass to Pastebin.
	body := bytes.NewBufferString(data.Encode())

	resp, err := http.Post("http://pastebin.com/api/api_post.php", "application/x-www-form-urlencoded", body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", PastebinPutError
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(respBody), "http://pastebin.com/", "", -1), nil
}
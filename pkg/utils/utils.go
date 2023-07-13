// unmarshal the json data to struct
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody parses the body of the request and unmarshals it to the struct passed in the function argument
func ParseBody(r *http.Request, x interface{}) {
 if body, err := ioutil.ReadAll(r.Body); err == nil {
	if err := json.Unmarshal([]byte(body), &x); err != nil {
		return 
	}
}
}


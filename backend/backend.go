package backend

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/groups", getGroups)
}

const response = `
    {
        "Groups": [{
            "Name": "GoSV",
            "URL": "http://www.meetup.com/golangsv",
            "Members": 194,
            "City": "San Mateo",
            "Country": "US"
        }, {
            "Name": "GoSF",
            "URL": "http://www.meetup.com/golangsf",
            "Members": 1393,
            "City": "San Francisco",
            "Country": "US"
        }],
        "Errors": [
            "something bad happened"
        ]
    }
`

func getGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, response)
}

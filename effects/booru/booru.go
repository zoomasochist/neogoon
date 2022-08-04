package booru

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Booru interface {
	Next() Image
}

type Image struct {
	Ext   string
	Bytes []byte
}

func FromString(s string) (Booru, error) {
	switch s {
	case "e621":
		return E621{}, nil
	default:
		return E621{}, errors.New(fmt.Sprintln("No booru named", s))
	}
}

func MakeHttpRequest(uri string) ([]byte, error) {
	// TODO: I think we can reuse the client
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("User-Agent", "Goonware")

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

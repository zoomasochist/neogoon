package booru

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type E621Response struct {
	Posts []E621Post `json:"posts"`
}

type E621Post struct {
	File E621File `json:"file"`
}

type E621File struct {
	Url string `json:"url"`
}

type E621 struct {
	Tags         []string
	MinimumScore int
}

// I'd like this to be part to be part of the above struct, but for whatever reason I can't
// get it to maintain the state of these variables. It keeps emptying them - I think it's something
// to do with passing by value in Next() or something. But this works, so here it will stay.
var imageUrlCache []string
var imageDataCache []Image

func (e E621) Next() Image {
	if len(imageDataCache) > 0 {
		image := imageDataCache[0]
		imageDataCache = imageDataCache[1:]
		return image
	}

	if len(imageUrlCache) > 0 {
		e.fillDataCache()
		return e.Next()
	}

	e.fillUrlCache()
	return e.Next()
}

func (e *E621) fillDataCache() error {
	imageUrls := imageUrlCache[:5]
	imageUrlCache = imageUrlCache[5:]

	for _, image := range imageUrls {
		imageBytes, err := MakeHttpRequest(image)
		if err != nil {
			return err
		}

		components := strings.Split(image, ".")
		imageDataCache = append(imageDataCache, Image{
			Ext:   components[len(components)-1],
			Bytes: imageBytes,
		})

		// E621's hard rate limit is 1 request/500ms.
		time.Sleep(700 * time.Millisecond)
	}

	return nil
}

func (e *E621) fillUrlCache() error {
	var tagSelection string
	if len(e.Tags) == 0 {
		tagSelection = ""
	} else {
		tagSelection = e.Tags[rand.Intn(len(e.Tags))]
	}

	url := fmt.Sprintf("https://e621.net/posts.json?limit=100&"+
		"tags=order:random+rating:e+score:>%d+%s", e.MinimumScore, tagSelection)

	resp, err := MakeHttpRequest(url)
	if err != nil {
		return err
	}

	var marshalledResponse E621Response
	if err = json.Unmarshal(resp, &marshalledResponse); err != nil {
		return err
	}

	for _, response := range marshalledResponse.Posts {
		if response.File.Url != "" {
			imageUrlCache = append(imageUrlCache, response.File.Url)
		}
	}

	return nil
}

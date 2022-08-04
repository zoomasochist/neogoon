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

	imageUrlCache  []string
	imageDataCache []Image
}

func (e E621) Next() Image {
	if len(e.imageDataCache) > 0 {
		image := e.imageDataCache[0]
		e.imageDataCache = e.imageDataCache[1:]
		return image
	}

	if len(e.imageUrlCache) > 0 {
		e.fillDataCache()
		return e.Next()
	}

	e.fillUrlCache()
	return e.Next()
}

func (e E621) fillDataCache() error {
	imageUrls := e.imageUrlCache[:5]
	e.imageDataCache = e.imageDataCache[5:]

	for _, image := range imageUrls {
		imageBytes, err := MakeHttpRequest(image)
		if err != nil {
			return err
		}

		components := strings.Split(image, ".")
		e.imageDataCache = append(e.imageDataCache, Image{
			Ext:   components[len(components)-1],
			Bytes: imageBytes,
		})

		// Double E621's hard rate limit.
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (e E621) fillUrlCache() error {
	var tagSelection string
	if len(e.Tags) == 0 {
		tagSelection = ""
	} else {
		tagSelection = e.Tags[rand.Intn(len(e.Tags))]
	}

	url := fmt.Sprintf("https://e621.net/posts.json?limit=100&"+
		"tags=rating:e+score:>%d+%s", e.MinimumScore, tagSelection)

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
			e.imageUrlCache = append(e.imageUrlCache, response.File.Url)
		}
	}

	return nil
}

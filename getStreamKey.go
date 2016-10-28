package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// StreamKey is a key for authenticate streaming programm.
type StreamKey struct {
	StreamKey string `json:"streamKey"`
}

// GetStreamKey get the stream key for channel.
func GetStreamKey(channel string, authToken AuthToken) (StreamKey, error) {
	var args fasthttp.Args
	args.Add("authToken", authToken.Token)
	requestURL := fmt.Sprintf("%s/mediakey/%s?%s", API, channel, args.String())
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return StreamKey{}, err
	}
	var obj StreamKey
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return StreamKey{}, err
	}
	return obj, nil
}

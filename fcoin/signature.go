package fcoin

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

// Sign request with client secret using HMAC-SHA1
// args should be ordered URI format
func Sign(method, uri, ts, args, key string) string {
	prep := method + uri + ts + args
	b64prep := []byte(base64.StdEncoding.EncodeToString([]byte(prep)))
	mac := hmac.New(sha1.New, []byte(key))

	mac.Write(b64prep)
	hmac_prep := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmac_prep)
}

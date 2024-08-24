package utils

import (
	"encoding/base64"
	"fmt"
	"time"
)

/*
This shortening technique is simple but not the best as it can produce collisions:
We get the current timestamp in seconds and then use base64 encoding to shorten it and return the first 8 characters.
*/
func GetShortCode() string {
	fmt.Println("Shortening URL")
	// Generate a random key, in a real world scennario, this would also need to be unique
	ts := time.Now().UnixNano()
	fmt.Println("Timestamp: ", ts)
	// We convert the timestamp to byte slice and then encode it to base64 string
	ts_bytes := []byte(fmt.Sprintf("%d", ts))
	key := base64.StdEncoding.EncodeToString(ts_bytes)
	fmt.Println("Key: ", key)

	// We remove the last two chars since they are usuall always equal signs (==)
	key = key[:len(key)-2]

	// We return the last chars after 16 chars, these are almost always different
	return key[16:]
}

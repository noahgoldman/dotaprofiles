package upload

import (
	"fmt"
)

// The bucket.URL function doesn't produce any errors, so neither
// does this function (for now)
func GetURL(path string) string {
	return fmt.Sprintf("http://%s.s3.amazonaws.com/%s", BUCKET_NAME, path)
}

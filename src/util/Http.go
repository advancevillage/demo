//author: richard
package util

import (
	"net/http"
)

func HttpQueryParamsService(r *http.Request, key string) (value string) {
	qs := r.URL.Query()
	if values, exist := qs[key]; exist {
		value = values[0]
	}
	return
}

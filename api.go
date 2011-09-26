/* Package go140 implements interaction with the Twitter API */
package go140

import (
	"github.com/alloy-d/goauth"
	"fmt"
	"http"
	"io/ioutil"
	"json"
	"os"
	"url"
)

type API struct {
	Root string
	oauth.OAuth
}

// TODO: this is SILLY.
func addQueryParams(resource string, params map[string]string) string {
	str := resource

	first := true
	for k, v := range params {
		if first {
			str += "?"
			first = false
		} else {
			str += "&"
		}

		rawv, err := url.QueryUnescape(v)
		if err == nil {
			v = rawv
		}
		str += k + "=" + url.QueryEscape(v)
	}
	return str
}

func (api *API) Get(resource string, params map[string]string) (*http.Response, os.Error) {
	if api.Authorized() {
		return api.OAuth.Get(resource, params)
	}

	fullURL := addQueryParams(resource, params)
	return http.Get(fullURL)
}

type Status struct {
	Date string "created_at"
	Text string
	Location string "place"
}

type User struct {
	ScreenName string "screen_name"
	Name string
	Location string
	Description string
	Status *Status
}

func (api *API) UserByID(id uint) (*User, os.Error) {
	return api.user(map[string]string{
		"id": fmt.Sprintf("%d", id),
	})
}

func (api *API) User(screen_name string) (*User, os.Error) {
	return api.user(map[string]string{
		"screen_name": screen_name,
	})
}

func (api *API) user(params map[string]string) (*User, os.Error) {
	url := api.Root + "/1/users/show.json"

	resp, err := api.Get(url, params)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (api *API) Status() (*Status, os.Error) {
	user, err := api.UserByID(api.UserID())
	if err != nil {
		return nil, err
	}

	return user.Status, nil
}

func (api *API) Update(s string) (string, os.Error) {
	if len(s) > 140 {
		return "", tweetError{"Tweet too long!"}
	}

	url := api.Root + "/1/statuses/update.json"
	params := map[string]string{"status": s}

	_, err := api.Post(url, params)
	if err != nil {
		return "", err
	}
	return "<not yet implemented>", nil
}

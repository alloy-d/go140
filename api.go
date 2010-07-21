package go140

import (
    "github.com/alloy-d/goauth"
    "fmt"
    //"http"
    "io/ioutil"
    "json"
    "os"
)

type API struct {
    Root string
    oauth.OAuth
}

func (api *API) User(id string) (interface{}, os.Error) {
    url := api.Root + "/1/users/show.json"
    params := map[string]string{
        "id": id,
    }

    resp, err := api.Get(url, params)
    if err != nil {
        return nil, err
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var obj interface{} = nil
    err = json.Unmarshal(data, &obj)
    if err != nil {
        return nil, err
    }
    return obj, nil
}

func (api *API) Status() (string, os.Error) {
    user, err := api.User(fmt.Sprintf("%d", api.UserID()))
    if err != nil {
        return "", err
    }

    userMap := user.(map[string]interface{})
    if status, ok := userMap["status"]; ok {
        statusMap := status.(map[string]interface{})
        if text, ok := statusMap["text"]; ok {
            return text.(string), nil
        }
    }
    return "", nil
}

func (api *API) Update(s string) (string, os.Error) {
    if len(s) > 140 {
        return "", tweetError{"Tweet too long!"}
    }

    url := api.Root + "/1/statuses/update.json"
    params := map[string]string{"status": s,}

    _, err := api.Post(url, params)
    if err != nil {
        return "", err
    }
    return "<not yet implemented>", nil
}



package main

import (
    "github.com/alloy-d/goauth"
    "github.com/alloy-d/go140"
    "fmt"
    "os"
)

var home string

func init() {
    home = os.Getenv("HOME")
}

func Error(err os.Error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func authorize(api *go140.API) (err os.Error) {
    api.ConsumerKey = "24eNIyd8EskIpi32cYFEbg"
    api.ConsumerSecret = "heXe5AP9tT4jMgWHyF6y17NWJRkJeRX7S7I7W78VmCc"
    api.SignatureMethod = oauth.HMAC_SHA1

    api.RequestTokenURL = "https://api.twitter.com/oauth/request_token"
    api.OwnerAuthURL = "https://api.twitter.com/oauth/authorize"
    api.AccessTokenURL = "https://api.twitter.com/oauth/access_token"
    api.Callback = "oob"

    api.Root = "http://api.twitter.com"

    err = api.Load(home + "/.go140.oauth")
    if err != nil {
        err = api.GetRequestToken()
        Error(err)
        url, err := api.AuthorizationURL()
        Error(err)
        fmt.Printf("Please visit the following URL for authorization:\n%s\n", url)

        var verifier string
        fmt.Printf("PIN: ")
        fmt.Scanf("%s", &verifier)
        err = api.GetAccessToken(verifier)
        Error(err)

        err = api.Save(home + "/.go140.oauth")
        if err != nil {
            fmt.Fprintf(os.Stderr, "Couldn't save authorization information: %s", err)
        }
    }

    return nil
}

func main() {
    api := new(go140.API)
    authorize(api)

    status, err := api.Status()
    Error(err)
    fmt.Println(status)

    status, err = api.Update("Twitter workflow: edit random test file, change string in API call, recompile, run program. Über sophistication.")

    Error(err)
    fmt.Println(status)

    status, err = api.Status()
    Error(err)
    fmt.Println(status)
}


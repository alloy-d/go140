package go140

type tweetError struct {
    what string
}

func (te tweetError) String() string { return te.what }


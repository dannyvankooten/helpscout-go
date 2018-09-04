HelpScout Go
=============
[![GoDoc](https://godoc.org/github.com/dannyvankooten/helpscout-go?status.svg)](https://godoc.org/github.com/dannyvankooten/helpscout-go)
 [![Build Status](https://travis-ci.org/dannyvankooten/helpscout-go.png?branch=master)](https://travis-ci.org/dannyvankooten/helpscout-go)

An unofficial Go client library for [HelpScout Custom Apps](https://developer.helpscout.com/custom-apps/dynamic/). 

## Usage

```go
import (
    "net/http"
    "io/ioutil"
    "github.com/dannyvankooten/helpscout"
)

func handler(w http.ResponseWriter, r *http.Request) {
    helpscout.SecretKey = "your-40-char-secret-key"
    signature := r.Header.Get("X-HelpScout-Signature")

	// check helpscout signature
	payload, _ := ioutil.ReadAll(r.Body)
	if !helpscout.VerifySignature(payload, signature) {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	input, _ := helpscout.Decode(payload)

    // You can now access input.Customer.Email etc.. 
    // TODO: Respond to request with JSON
    // TODO: Handle errors
}
```

## License

MIT Licensed. See the [LICENSE](LICENSE) file for details.

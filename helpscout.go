// MIT License

// Copyright (c) 2017 Danny van Kooten

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package helpscout

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"os"
)

// Response is the response we return to the HelpScout request.
type Response struct {
	HTML string `json:"html"`
}

// Ticket represents a HelpScout ticket (conversation)
type Ticket struct {
	ID      string `json:"id"`
	Number  string `json:"number"`
	Subject string `json:"subject"`
}

// Customer represents a HelpScout customer
type Customer struct {
	ID     string   `json:"id"`
	Fname  string   `json:"fname"`
	Lname  string   `json:"lname"`
	Email  string   `json:"email"`
	Emails []string `json:"emails"`
}

// Input wraps all request data for a Custom Integration
type Input struct {
	Customer *Customer `json:"customer,omitempty"`
}

// SecretKey is used to verify request signatures.
// It must match the secret key in your HelpScout Custom App settings.
var SecretKey = os.Getenv("HELPSCOUT_SECRET_KEY")

//VerifySignature checks if the request signature matches the expected signature (given the secret key)
func VerifySignature(payload []byte, signature string) bool {
	mac := hmac.New(sha1.New, []byte(SecretKey))
	mac.Write(payload)
	hash := mac.Sum(nil)

	// compare with given signature (which is base64 encoded)
	expected := make([]byte, base64.StdEncoding.EncodedLen(len(hash)))
	base64.StdEncoding.Encode(expected, hash)
	return hmac.Equal([]byte(signature), expected)
}

// Decode data into struct
func Decode(payload []byte) (*Input, error) {
	input := &Input{}
	err := json.Unmarshal(payload, input)
	return input, err
}

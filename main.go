package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// request identity
func requestIdentity(w http.ResponseWriter, r *http.Request) {

	jsonBody := []byte(`{
        "didMetadata": {
            "method": "polygonid",
            "blockchain": "polygon",
            "network": "mumbai",
            "type": "BJJ"
        }
    }`)

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := "http://localhost:3001/v1/identities"

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userName := "user-issuer"
	pwd := "password-issuer"

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Basic dXNlci1pc3N1ZXI6cGFzc3dvcmQtaXNzdWVy")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(userName, pwd)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Response: %v\n", res.Status)

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Print the response body
	fmt.Printf("%s\n", responseBody)

}

// get identities
func getIdentities(w http.ResponseWriter, r *http.Request) {

	jsonBody := []byte(`{

    }`)

	bodyReader := bytes.NewReader(jsonBody)
	requestURL := "http://localhost:3001/v1/identities"

	req, err := http.NewRequest(http.MethodGet, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userName := "user-issuer"
	pwd := "password-issuer"

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Basic dXNlci1pc3N1ZXI6cGFzc3dvcmQtaXNzdWVy")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(userName, pwd)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Response: %v\n", res.Status)

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Print the response body
	fmt.Printf("%s\n", responseBody)
}

// Publish state on-chain
func publishStateOnChain(w http.ResponseWriter, r *http.Request) {

	jsonBody := []byte(`{

	}`)

	bodyReader := bytes.NewReader(jsonBody)
	requestURL := "http://localhost:3001/v1/did:polygonid:polygon:mumbai:2qM7PNe8nLjMKUHCaNMuCwXfgxhZKCS1g7bjKyKGNW/state/publish"

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userName := "user-issuer"
	pwd := "password-issuer"

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Basic dXNlci1pc3N1ZXI6cGFzc3dvcmQtaXNzdWVy")
	req.SetBasicAuth(userName, pwd)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Response: %v\n", res.Status)

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Print the response body
	fmt.Printf("%s\n", responseBody)

}

// Create claim
func createClaim(w http.ResponseWriter, r *http.Request) {
	jsonBody := []byte(`{
		"credentialSchema":"https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json/KYCAgeCredential-v3.json",
		"type":"KYCAgeCredential",
		"credentialSubject":
			{
				"id":"did:polygonid:polygon:mumbai:2qHSH78PVMnajgnvZZkxspH8vqBSzkas4VcJT1wsCR",
				"birthday":19960424,
				"documentType":2
			},
		"expiration":1903357766
		}`)

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := "http://localhost:3001/v1/did:polygonid:polygon:mumbai:2qHMY14FP4KYfHf2cZLUcYCpso19NgGPReK4bt3cRq/claims"

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userName := "user-issuer"
	pwd := "password-issuer"

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Basic dXNlci1pc3N1ZXI6cGFzc3dvcmQtaXNzdWVy")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(userName, pwd)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Response: %v\n", res.Status)

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Print the response body
	fmt.Printf("%s\n", responseBody)
}

// get claim
func getClaim(w http.ResponseWriter, r *http.Request) {
	jsonBody := []byte(`{}`)

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := "http://localhost:3001/v1/did%3Apolygonid%3Apolygon%3Amumbai%3A2qHMY14FP4KYfHf2cZLUcYCpso19NgGPReK4bt3cRq/claims/dbe30a0d-a585-11ee-bd8f-0242ac170006"

	req, err := http.NewRequest(http.MethodGet, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userName := "user-issuer"
	pwd := "password-issuer"

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Basic dXNlci1pc3N1ZXI6cGFzc3dvcmQtaXNzdWVy")

	req.SetBasicAuth(userName, pwd)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Response: %v\n", res.Status)

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Print the response body
	fmt.Printf("%s\n", responseBody)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/requestIdentity", requestIdentity)
	mux.HandleFunc("/getIdentities", getIdentities)
	mux.HandleFunc("/publishOnState", publishStateOnChain)
	mux.HandleFunc("/createClaim", createClaim)
	mux.HandleFunc("/getClaim", getClaim)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// Credentials Subject
// type CredentialSchema struct {
// 	id       string
// 	birthdat string
// }

// // Claims structure
// type Claims struct {
// 	credentialsSchema string
// 	schemaType        string
// }

// Identity structure
// type Identity struct {
// 	method     string
// 	blockchain string
// 	network    string
// 	idtype     string
// }

package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/assert"
)

func TestPostalCode(t *testing.T) {
    address := "Омск, п. Верхние пупки, д. 6"
    postalCode := "100000"

    var actualAddress string
    var actualPostalCode string

    server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        actualAddress = req.URL.Query().Get("address")

        rw.Write([]byte(postalCode))
    }))

    defer server.Close()

    postalCodeClient := PostalCodeClient{server.Client(), server.URL}
    actualPostalCode, err := postalCodeClient.PostalCode(address)
    if err != nil {
        t.Fatalf("failed to get postal code: %v", err)
    }

    if actualPostalCode != postalCode {
        t.Errorf("expected postal code %s, got %s", postalCode, actualPostalCode)
    }

    if actualAddress != address {
        t.Errorf("expected address %s, got %s", address, actualAddress)
    }
}

func TestPostalCodeWithAssert(t *testing.T) {
    address := "Омск, п. Верхние пупки, д. 6"
    postalCode := "100000"

    var actualAddress string
    var actualPostalCode string

    server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        actualAddress = req.URL.Query().Get("address")

        rw.Write([]byte(postalCode))
    }))

    defer server.Close()

    postalCodeClient := PostalCodeClient{server.Client(), server.URL}
    actualPostalCode, err := postalCodeClient.PostalCode(address)

    require.NoError(t, err)
    assert.Equal(t, postalCode, actualPostalCode)
    assert.Equal(t, address, actualAddress)
}
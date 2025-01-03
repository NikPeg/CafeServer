package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
)


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?city=moscow&count=" + strconv.Itoa(totalCount + 1), nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
    }

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    if len(list) != totalCount {
        t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))
    }
}
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
)

func TestGenerateShortURL(t *testing.T) {
    url1 := generateShortURL()
    if len(url1) != 6 {
        t.Errorf("La URL corta debería tener 6 caracteres, pero tiene %d", len(url1))
    }

    url2 := generateShortURL()
    if url1 == url2 {
        t.Error("Las URLs cortas generadas deberían ser diferentes")
    }
}

func TestIsValidURL(t *testing.T) {
    validURL := "http://example.com"
    if !isValidURL(validURL) {
        t.Errorf("La URL %s debería ser válida", validURL)
    }

    invalidURL := "example.com"
    if isValidURL(invalidURL) {
        t.Errorf("La URL %s debería ser inválida", invalidURL)
    }
}

func TestShortenAndRedirect(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/shorten", func(c *gin.Context) {
        longURL := c.PostForm("url")
        // Simula la lógica de tu servidor
        urlMap["abcde1"] = longURL
        c.JSON(http.StatusOK, gin.H{"shortenedURL": "http://localhost:8080/abcde1"})
    })
    router.GET("/:shortURL", func(c *gin.Context) {
        shortURL := c.Param("shortURL")
        longURL, found := urlMap[shortURL]
        if found {
            c.Redirect(http.StatusMovedPermanently, longURL)
        } else {
            c.String(http.StatusNotFound, "URL no encontrada")
        }
    })

    // Simular una petición POST para acortar una URL
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/shorten", bytes.NewBufferString("url=http://example.com"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Se esperaba un código de estado 200, pero se obtuvo %d", w.Code)
    }

    var response map[string]string
    json.Unmarshal(w.Body.Bytes(), &response)

    shortenedURL := response["shortenedURL"]

    // Simular una petición GET para la redirección
    wRedirect := httptest.NewRecorder()
    reqRedirect, _ := http.NewRequest("GET", shortenedURL, nil)
    router.ServeHTTP(wRedirect, reqRedirect)

    if wRedirect.Code != http.StatusMovedPermanently {
        t.Errorf("Se esperaba un código de estado 301, pero se obtuvo %d", wRedirect.Code)
    }
}
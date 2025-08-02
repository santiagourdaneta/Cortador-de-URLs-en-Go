package main

import (
    "log"
    "net/http"
    "math/rand"
    "time"
    "net/url" 
    "github.com/JGLTechnologies/gin-rate-limit"
    "github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)

func generateShortURL() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    shortURL := make([]byte, 6)
    for i := range shortURL {
        shortURL[i] = charset[rand.Intn(len(charset))]
    }
    return string(shortURL)
}

// New function to validate URLs
func isValidURL(toTest string) bool {
    _, err := url.ParseRequestURI(toTest)
    if err != nil {
        return false
    }
    u, err := url.Parse(toTest)
    return err == nil && u.Scheme != "" && u.Host != ""
}

func keyFunc(c *gin.Context) string {
    return c.ClientIP() // Limita por dirección IP del cliente
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
    c.JSON(http.StatusTooManyRequests, gin.H{"error": "Demasiadas peticiones. Por favor, espera un momento."})
}

func main() {

	 // ------------------ Inicialización de la semilla ------------------
    rand.Seed(time.Now().UnixNano())
    // ------------------------------------------------------------------
    
	router := gin.Default()

	// Create an in-memory store for rate limiting with the correct options.
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		// Set the rate and limit directly on the store's options.
		Rate:  time.Second, // The duration of the rate-limiting window
		Limit: 5,           // The number of requests allowed per duration
	})

	// Create the middleware with the store, key function, and error handler.
	middleware := ratelimit.RateLimiter(store, &ratelimit.Options{
		KeyFunc:      keyFunc,
		ErrorHandler: errorHandler,
	})

	// Apply the middleware to the router
	router.Use(middleware)
	// ---------------------------------------------------------------------------

	router.StaticFile("/", "templates/index.html")

	router.POST("/shorten", func(c *gin.Context) {
		longURL := c.PostForm("url")

		if longURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The URL cannot be empty."})
			return
		}

		if !isValidURL(longURL) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The URL format is invalid."})
			return
		}

		shortURL := generateShortURL()
		urlMap[shortURL] = longURL

		c.JSON(http.StatusOK, gin.H{"shortenedURL": "http://localhost:8080/" + shortURL})
	})

	router.GET("/:shortURL", func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		longURL, found := urlMap[shortURL]

		if found {
			c.Redirect(http.StatusMovedPermanently, longURL)
			return
		}

		c.String(http.StatusNotFound, "URL not found")
	})

	log.Println("Server started on http://localhost:8080")
	router.Run(":8080")
}
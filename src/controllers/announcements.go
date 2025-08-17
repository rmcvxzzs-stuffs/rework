// ... (imports and type definitions unchanged)

const adminToken = "your_secret_here" // Change this to a strong secret!

// ... (global variables unchanged)

func main() {
    r := gin.Default()

    // Middleware for simple admin token auth
    adminAuth := func(c *gin.Context) {
        token := c.GetHeader("X-Admin-Token")
        if token != adminToken {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }
        c.Next()
    }

    // POST endpoint to add announcements (admin only)
    r.POST("/announcements", adminAuth, func(c *gin.Context) {
        var input Announcement
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        annMutex.Lock()
        input.ID = nextID
        nextID++
        if input.CreatedAt.IsZero() {
            input.CreatedAt = time.Now()
        }
        announcements = append(announcements, input)
        annMutex.Unlock()
        c.JSON(http.StatusOK, input)
    })

    // GET endpoint unchanged (or add adminAuth for GET too if needed)
    r.GET("/announcements.xml", func(c *gin.Context) {
        // ... unchanged ...
    })

    r.Run(":8080")
}

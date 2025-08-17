package main

import (
    "encoding/xml"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Announcement struct {
    ID            uint      `gorm:"primaryKey" xml:"id"`
    Platform      string    `xml:"platform"`
    LanguageCode  string    `xml:"language_code"`
    CreatedAt     time.Time `xml:"created_at"`
    Subject       string    `xml:"subject"`
    Text          string    `xml:"text"`
}

type Announcements struct {
    Total            int            `xml:"total"`
    AnnouncementList []Announcement `xml:"AnnouncementList>Announcement"`
}

type ResponseStatus struct {
    ID      int    `xml:"id"`
    Message string `xml:"message"`
}

type Response struct {
    XMLName  xml.Name        `xml:"Response"`
    Status   ResponseStatus  `xml:"status"`
    Response []Announcements `xml:"response"`
}

// Replace this with your real admin token(s) or authentication logic!
const adminToken = "supersecrettoken"

// Middleware to check admin authentication via Authorization header: Bearer <token>
func AdminRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if !strings.HasPrefix(auth, "Bearer ") || strings.TrimPrefix(auth, "Bearer ") != adminToken {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "admin authorization required"})
            return
        }
        c.Next()
    }
}

func main() {
    db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    db.AutoMigrate(&Announcement{})

    r := gin.Default()

    // POST endpoint to add announcements (admin only)
    r.POST("/announcements", AdminRequired(), func(c *gin.Context) {
        var input Announcement
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if input.CreatedAt.IsZero() {
            input.CreatedAt = time.Now()
        }
        db.Create(&input)
        c.JSON(http.StatusOK, input)
    })

    // GET endpoint to fetch announcements as XML
    r.GET("/announcements.xml", func(c *gin.Context) {
        platform := c.Query("platform")
        var announcements []Announcement
        db.Where("platform = ?", platform).Find(&announcements)

        response := Response{
            Status: ResponseStatus{ID: 0, Message: "Successful completion"},
            Response: []Announcements{
                {
                    Total:            len(announcements),
                    AnnouncementList: announcements,
                },
            },
        }
        c.XML(http.StatusOK, response)
    })

    r.Run(":8080")
}

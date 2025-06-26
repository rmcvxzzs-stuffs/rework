package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
    app := fiber.New()

    // Static file serving - equivalent to router.use('/resources', static(__dirname + '/v1/static'))
    app.Static("/resources", "./v1/static")

    // Route handlers for XML endpoints
    app.Use("/preferences.xml", preferencesHandler)
    app.Use("/achievements.xml", achievementsHandler)  
    app.Use("/player_avatars", playerAvatarsHandler)
    app.Use("/servers/select.xml", serverSelectHandler)
    app.Use("/announcement.list.xml", announcementHandler)
    app.Use("/profanity_filters.xml", profanityFilterHandler)
    app.Use("/tags.xml", tagsHandler)
    app.Use("/single_player_games/create_finish_and_post_stats.xml", singlePlayerStatsHandler)
    app.Use("/sub_leaderboards", leaderboardHandler)

    // GET route for policies
    app.Get("/policies/view.xml", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile(filepath.Join(".", "policy.xml"))
    })

    // POST route for session presence
    app.Post("/session/set_presence.xml", func(c *fiber.Ctx) error {
        body := string(c.Body())
        presence := strings.Replace(body, "presence=", "", -1)
        fmt.Println("Presence: " + presence)
        
        c.Set("X-Status-ID", "0")
        c.Set("X-Status-Message", "Successful completion") 
        c.Set("Status", "200")
        c.Set("Content-Type", "application/xml;charset=utf-8")
        
        return c.Status(200).SendFile(filepath.Join(".", "success.xml"))
    })

    // POST route for session ping
    app.Post("/session/ping.xml", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile(filepath.Join(".", "success.xml"))
    })

    // GET route for player creation comments
    app.Get("/player_creation_comments.xml", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile(filepath.Join(".", "misc/example_comment_list.xml"))
    })

    // GET route for player profile
    app.Get("/player_profile/view.xml", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile(filepath.Join(".", "misc/example_profile.xml"))
    })

    // POST route for player creations verify
    app.Post("/player_creations/verify.xml", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile(filepath.Join(".", "player_creations/verify.xml"))
    })

    // POST route for player creation comments
    app.Post("/player_creation_comments.xml", func(c *fiber.Ctx) error {
        body := string(c.Body())
        decoded, _ := url.QueryUnescape(body)
        response := strings.Split(decoded, "&")
        if len(response) > 0 {
            commentParts := strings.Split(response[0], "=")
            if len(commentParts) > 1 {
                commentBody := commentParts[1]
                fmt.Println("Comment:", commentBody)
            }
        }
        return c.SendFile(filepath.Join(".", "success.xml"))
    })

    log.Fatal(app.Listen(":3000"))
}

// Placeholder handler functions - you'll need to implement these based on your v1 modules
func preferencesHandler(c *fiber.Ctx) error {
    return c.Next()
}

func achievementsHandler(c *fiber.Ctx) error {
    return c.Next()
}

func playerAvatarsHandler(c *fiber.Ctx) error {
    return c.Next()
}

func serverSelectHandler(c *fiber.Ctx) error {
    return c.Next()
}

func announcementHandler(c *fiber.Ctx) error {
    return c.Next()
}

func profanityFilterHandler(c *fiber.Ctx) error {
    return c.Next()
}

func tagsHandler(c *fiber.Ctx) error {
    return c.Next()
}

func singlePlayerStatsHandler(c *fiber.Ctx) error {
    return c.Next()
}

func leaderboardHandler(c *fiber.Ctx) error {
    return c.Next()
}
import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Fiber instance
    app := fiber.New(fiber.Config{
	CaseSensitive:  true
	ServerHeader:   "Fiber",	
	AppName:   "Horizon" // app name btw
})
    

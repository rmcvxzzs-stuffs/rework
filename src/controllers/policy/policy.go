package main

import (
    "encoding/xml"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

type Policy struct {
    ID         int    `xml:"id,attr"`
    IsAccepted bool   `xml:"is_accepted,attr"`
    Name       string `xml:"name,attr"`
    Text       string `xml:",chardata"`
}

type PolicyResponse struct {
    XMLName xml.Name `xml:"response"`
    Policy  Policy   `xml:"policy"`
}

func main() {
    r := gin.Default()

    // Example EULA policy in-memory
    eulaText := `
TESTING
`

    r.GET("/policies/view.xml", func(c *gin.Context) {
        policyType := c.DefaultQuery("policy_type", "EULA")
        // Only EULA policy is supported in this example
        var policy Policy
        switch strings.ToUpper(policyType) {
        case "EULA":
            policy = Policy{
                ID:         1,
                IsAccepted: false, // Set to true if user has accepted (implement as needed)
                Name:       "EULA",
                Text:       eulaText,
            }
        default:
            c.XML(http.StatusNotFound, gin.H{"error": "policy not found"})
            return
        }

        c.XML(http.StatusOK, PolicyResponse{
            Policy: policy,
        })
    })

    r.Run(":8080")
}

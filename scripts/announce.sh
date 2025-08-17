# run 'chmod +x announce.sh' first!
#!/bin/sh

# Prompt user for input
read -sp "Enter ADMIN_SECRET_TOKEN: " ADMIN_SECRET_TOKEN
echo
read -p "Enter announcement SUBJECT: " ANNOUNCEMENT_SUBJECT
read -p "Enter announcement TEXT: " ANNOUNCEMENT_TEXT

# CHANGE SERVER ENDPOINT TO YOUR LIKING
SERVER_URL="http://localhost:8080/announcements"

# Send POST request with curl
curl -X POST "$SERVER_URL" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_SECRET_TOKEN" \
  -d "{\"subject\": \"$ANNOUNCEMENT_SUBJECT\", \"text\": \"$ANNOUNCEMENT_TEXT\"}"

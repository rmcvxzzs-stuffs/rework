# dont make an announcement yet
curl -X POST http://localhost:8080/announcements \
     -H "Content-Type: application/json" \
     -H "X-Admin-Token: [replace]" \
     -d '{"platform":"web","language_code":"en","subject":"[replace]","text":"[replace]"}'

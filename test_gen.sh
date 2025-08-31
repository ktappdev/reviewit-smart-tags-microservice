#!/bin/bash

# Test script for /gen endpoint
echo "Testing /gen endpoint..."

# Test with a sample description
curl -X POST http://localhost:3003/gen \
  -H "Content-Type: application/json" \
  -d '{"description": "A modern coffee shop that serves artisanal coffee, pastries, and light meals. Features free WiFi and cozy atmosphere for remote work."}' \
  2>/dev/null | python3 -m json.tool || echo "Response received (may not be valid JSON if API key is invalid)"

echo -e "\n\nTest completed. If you see tags in JSON format above, the /gen endpoint is working correctly."
echo "Note: You need a valid OPEN_ROUTER_API_KEY in your .env file for the AI to work."

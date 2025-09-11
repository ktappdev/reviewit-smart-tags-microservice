# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a lightweight Go microservice that generates AI-powered tags from business descriptions using the OpenRouter API. The service is stateless with no database dependencies and provides a RESTful API endpoint for tag generation.

## Architecture

The service is built with Go 1.23.1 and uses the Fiber web framework. The main components are:

- **main.go**: Entry point with Fiber app setup, CORS configuration, and route handling
- **ai.go**: Core AI integration with OpenRouter API, including model fallback logic
- **direction.go**: Contains the AI prompt template for tag generation (30 tags per business description)
- **structs.go**: Data structures for API requests/responses and AI responses
- **tagsFromLiveDescription.go**: Main request handler that coordinates tag generation

## Key Dependencies

- `github.com/gofiber/fiber/v2`: Web framework
- `github.com/joho/godotenv`: Environment variable management

## Development Commands

### Running the Application
```bash
# Run with hot reload (recommended for development)
air

# Run directly
go run .

# Build the application
go build -o reviewit-smarttags-go .
```

### Testing
```bash
# Test the /gen endpoint
./test_gen.sh

# Manual test with curl
curl -X POST http://localhost:3003/gen \
  -H "Content-Type: application/json" \
  -d '{"description": "A modern coffee shop that serves artisanal coffee, pastries, and light meals."}'
```

### Dependencies
```bash
# Install dependencies
go mod tidy

# Update dependencies
go get -u
```

## Configuration

The service requires a `.env` file with:
- `OPEN_ROUTER_API_KEY`: OpenRouter API key (required)
- `PORT`: Server port (default: 3003)

Copy `.env.example` to `.env` and configure your API key.

## API Endpoint

- `POST /gen`: Accepts JSON with `description` field, returns JSON with `tags` array containing exactly 30 tags

## AI Integration

The service uses OpenRouter API with fallback models:
1. Primary: `meta-llama/llama-3.2-1b-instruct`
2. Fallback: `google/gemma-3-4b-it`

The AI is prompted to generate exactly 30 tags that are relevant to business descriptions, commonly used in job search, and follow specific formatting rules (lowercase, no duplicates).

## CORS Configuration

CORS is configured to allow requests from:
- `https://reviewit.gy`
- `http://localhost:3000`
- `http://localhost:3001`
- `http://127.0.0.1:3000`

## Build Output

The compiled binary is output as `reviewit-smarttags-go` in the project root.
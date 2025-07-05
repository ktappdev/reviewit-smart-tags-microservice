# reviewit-smart-tags-microservice

A Go microservice for generating AI-powered tags for business reviews using OpenRouter API.

## Features

- Generate tags from business descriptions using AI
- Regenerate tags for existing reviews
- RESTful API endpoints
- PostgreSQL database integration
- CORS support for web applications

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and fill in your configuration:
   ```bash
   cp .env.example .env
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run .
   ```

## API Endpoints

- `POST /gettags` - Get tags for a review by ID
- `POST /regen` - Regenerate tags for an existing review
- `POST /gen` - Generate tags from a description

## Environment Variables

- `DATABASE_URL` - PostgreSQL connection string
- `APP_API` - External API endpoint for review data
- `OPEN_ROUTER_API_KEY` - OpenRouter API key for AI requests
- `PORT` - Server port (default: 3003)

## Development

For development with hot reload, use Air:
```bash
air
```

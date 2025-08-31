# reviewit-smart-tags-microservice

A lightweight Go microservice for generating AI-powered tags from business descriptions using OpenRouter API.

## Features

- Generate tags from business descriptions using AI
- No database dependencies - stateless operation
- RESTful API endpoint
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

- `POST /gen` - Generate tags from a description

## Environment Variables

- `OPEN_ROUTER_API_KEY` - OpenRouter API key for AI requests (required)
- `PORT` - Server port (default: 3003)

## Development

For development with hot reload, use Air:
```bash
air
```

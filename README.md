# prajwal_portfolio_bot backend (Go)

Minimal backend for `telegram_bot` with static command responses.

## Features
- No database
- Static command-based responses
- Supports both polling and webhook mode

## Commands
- `/start`
- `/help`
- `/about`
- `/skills`
- `/experience`
- `/projects`
- `/contact`
- `/resume`
- `/website`
- `/github`
- `/linkedin`

## Setup
1. Copy env file:
   ```bash
   cp .env.example .env
   ```
2. Set `TELEGRAM_BOT_TOKEN` in `.env`.
   - Optional: set `RESUME_URL` to a public PDF URL for `/resume`.
3. Export env vars:
   ```bash
   export $(grep -v '^#' .env | xargs)
   ```
4. Run:
   ```bash
   go run .
   ```

By default it runs in polling mode.

## Webhook mode (optional)
Set:
- `APP_MODE=webhook`
- `PORT=8080`
- `WEBHOOK_SECRET=some_random_secret`
- `RESUME_URL=https://<public-file-url>/resume.pdf` (optional)

Then set webhook on Telegram:
```bash
curl -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/setWebhook" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://<your-domain>/webhook",
    "secret_token": "'$WEBHOOK_SECRET'"
  }'
```

Health check endpoint:
- `GET /health`

## Deploy (Render)
This repo includes `render.yaml` for quick deployment.

1. Push this repo to GitHub.
2. In Render, create a new Blueprint and select this repo.
3. Set `TELEGRAM_BOT_TOKEN` in Render environment variables.
   - Optional: set `RESUME_URL` to your public resume PDF URL.
4. Deploy. After deploy, note your service URL:
   - Example: `https://portfolio-bot.onrender.com`
5. Set Telegram webhook to your public URL:
   ```bash
   curl -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/setWebhook" \
     -H "Content-Type: application/json" \
     -d '{
       "url": "https://<your-render-domain>/webhook",
       "secret_token": "'$WEBHOOK_SECRET'"
     }'
   ```
6. Verify webhook:
   ```bash
   curl "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getWebhookInfo"
   ```

Notes:
- Ensure Render env has `APP_MODE=webhook`.
- Render provides `PORT` automatically; the app already supports it.

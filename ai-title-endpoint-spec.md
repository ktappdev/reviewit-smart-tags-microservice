# AI Title Generation Endpoint

**Purpose:** Auto-generate a review title from the review body text.

## Endpoint

```
POST {NEXT_PUBLIC_AI_SERVER}/api/ai/generate-title
```

`NEXT_PUBLIC_AI_SERVER` is already set in `.env` as `https://ai.reviewit.gy`.

## Request

```json
{
  "body": "The camera on this phone is incredible. Low light photos look professional and the zoom is sharp even at 10x.",
  "rating": 4,
  "productName": "iPhone 15 Pro"
}
```

| Field | Type | Required | Notes |
|-------|------|----------|-------|
| `body` | string | yes | Plain text extracted from rich text editor (HTML tags stripped). First ~200 chars sent. |
| `rating` | number | yes | 1–5 star rating user selected |
| `productName` | string | no | Product name for context. May be empty if product not loaded yet. |

## Response

### Success (200)

```json
{
  "title": "Incredible camera, amazing low light photos"
}
```

| Field | Type | Notes |
|-------|------|-------|
| `title` | string | Generated title, max 80 characters recommended. Should reference specifics from the body, not generic ("Great product"). |

### Error

Any non-200 response is treated as failure. The client will fall back to manual title entry — no error message shown to user.

## Client Behavior

| Trigger | Action |
|---------|--------|
| User types 4+ words in body | Wait 800ms debounce, then call |
| Response received | Title field auto-fills with animation |
| User manually edits title | Manual title respected. Regenerate button appears nearby |
| User clicks "Regenerate" | Re-calls endpoint with latest body text. No debounce — immediate call. |
| Network/server error | Title field remains as-is, subtle error indicator on regenerate button |

### Regenerate UX

- Small icon button (↻) or text link ("Regenerate title") next to title field
- Only visible after initial title is generated (auto or manual)
- On click: shows brief loading spinner on the button, then replaces title
n- Works on both mobile and desktop — touch-friendly hit area (min 44px)

## Expected AI Behavior

- Read the body text and extract the key sentiment or standout feature
- Generate a short, natural-sounding review title (not robotic)
- Incorporate specific details from the body when present
- Match tone to rating (positive for 4–5, constructive for 1–3)
- Examples:
  - Body: "battery lasts two full days, screen is gorgeous" → "Amazing battery life and gorgeous screen"
  - Body: "shipping was late and the box was damaged" → "Late shipping, arrived damaged"
  - Body: "does what it says, no complaints" → "Solid product, works as expected"

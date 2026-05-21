package main

var directionTitle = `You generate short, natural-sounding titles for product reviews.

Rules:
- Read the review body and extract the key sentiment
- Reference specific details from the body — never use generic phrases like "Great product"
- Match tone to rating: enthusiastic for 4-5 stars, critical/constructive for 1-3 stars  
- Maximum 80 characters
- Sound like a real person wrote it, not a robot
- If the review is negative, the title MUST be negative. If positive, title MUST be positive.
- Return ONLY a JSON object: {"title": "..."} — no markdown, no backticks, no explanation.

Examples:
Body: "battery lasts two full days, screen is gorgeous" | Rating: 5 stars → {"title": "Amazing battery life and gorgeous screen"}
Body: "shipping was late and the box was damaged" | Rating: 2 stars → {"title": "Late shipping, arrived damaged"}
Body: "this is the worst place on the planet, I would never eat here" | Rating: 1 star → {"title": "Worst place on the planet, never eating here again"}
Body: "does what it says, no complaints" | Rating: 3 stars → {"title": "Solid product, works as expected"}
`

var direction = `
Generate job sector tags based on the provided business description. Always return exactly 30 tags, following these guidelines:

1. Start with the most relevant tags directly related to the business's primary activities.
2. If 30 highly relevant tags cannot be generated, expand to closely related industries or skills.
3. If still short of 30, include broader sector categories that encompass the business.
4. If necessary, add complementary or adjacent industry tags to reach 30.

The tags should be:
1. Relevant to the business's activities, with the most pertinent tags listed first
2. Commonly used in job search or industry classification
3. A mix of specific and broader sector terms
4. All in lowercase, even for proper nouns
5. Without duplicates or extremely similar terms

Return the result as a JSON object with a single key "tags" containing an array of exactly 30 string values. Return ONLY the JSON object, with no markdown, backticks, explanations, or additional text.

Example output format:
{"tags": ["technology", "e-commerce", "digital marketing", "software development", "cloud computing", "data analytics", "user experience", "mobile apps", "artificial intelligence", "cybersecurity", "fintech", "saas", "big data", "blockchain", "iot", "machine learning", "web development", "digital transformation", "startup", "b2b", "it services", "product management", "devops", "agile methodology", "customer success", "business intelligence", "data science", "network security", "software engineering", "innovation"]}
`

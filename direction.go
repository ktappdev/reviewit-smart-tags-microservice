package main

var direction = `
Generate job sector tags based on the provided business description. Always return exactly 20 tags, following these guidelines:

1. Start with the most relevant tags directly related to the business's primary activities.
2. If 20 highly relevant tags cannot be generated, expand to closely related industries or skills.
3. If still short of 20, include broader sector categories that encompass the business.
4. If necessary, add complementary or adjacent industry tags to reach 20.

The tags should be:
1. Relevant to the business's activities, with the most pertinent tags listed first
2. Commonly used in job search or industry classification
3. A mix of specific and broader sector terms
4. All in lowercase, even for proper nouns
5. Without duplicates or extremely similar terms

Return the result as a JSON object with a single key "tags" containing an array of exactly 20 string values. Do not include any explanations or additional text.

Example output format:
{
  "tags": ["technology", "e-commerce", "digital marketing", "software development", "cloud computing", "data analytics", "user experience", "mobile apps", "artificial intelligence", "cybersecurity", "fintech", "saas", "big data", "blockchain", "iot", "machine learning", "web development", "digital transformation", "startup", "b2b"]
}
`

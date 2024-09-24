package main

var direction = `
Generate job sector tags based on the provided business description. Prioritize the number of tags as follows:
1. Aim for 5 tags if possible
2. If 5 relevant tags cannot be generated, provide 4 tags
3. If 4 relevant tags cannot be generated, provide 3 tags

The tags should be:
1. Highly relevant to the business's primary activities
2. Commonly used in job search or industry classification
3. Broad enough to capture the main sector, not overly specific
4. All tags must be in common case, meaning they should be written in lowercase letters, even if it is a proper noun.

Return the result as a JSON object with a single key "tags" containing an array of string values. The array should contain 5, 4, or 3 tags based on the prioritization above. Do not include any explanations or additional text.

Example output format:
{
  "tags": ["Technology", "E-commerce", "Digital Marketing", "Software Development", "Cloud Computing"]
}

or

{
  "tags": ["Technology", "E-commerce", "Digital Marketing", "Software Development"]
}

or

{
  "tags": ["Technology", "E-commerce", "Digital Marketing"]
}
`

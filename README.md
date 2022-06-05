# go-exercises-track2

Write an endpoint `GET /posts`

that returns the following JSON

```
[
	{
		"id": "abcde",
		"title": "I'm here!",
		"body": "What's occurring? What did I miss?",
		"publishedAt": "2022-05-16T15:51:09.000Z",
		"author": {
			"id": "dhirenb94",
			"name": "Dhiren B"
		}
	},
	{
		"id": "rtyui",
		"title": "What's for lunch?",
		"body": "I miss team lunches - what do people eat around here?",
		"publishedAt": "2022-05-16T12:21:09.000Z",
		"author": {
			"id": "aceakash",
			"name": "Akash K"
		}
	},
	{
		"id": "gjhgk",
		"title": "Bored and falling asleep",
		"body": "This place sucks, there's hardly anything to do here...",
		"publishedAt": "2022-05-12T11:11:00.000Z",
		"author": {
			"id": "vvirgitti",
			"name": "Vanessa V"
		}
	}
]

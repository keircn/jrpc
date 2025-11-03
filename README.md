# jrpc

Simple Discord RPC client written in Go configurable via JSON.

```json
{
  "clientId": "YOUR_CLIENT_ID",
  "state": "using arch btw",
  "details": "bla bla bla"
  "largeImage": "sylphie",
  "largeText": "tooltip for the large image"
  "smallImage": "arch",
  "smallText": "tooltip for the small image"
  "showTimestamp": true,
  "buttons": [
    {
      "label": "My AniList",
      "url": "https://anilist.co/user/keiran"
    }
  ]
}
```

- Get the client ID from the [Discord Developer Portal](https://discord.com/developers/applications)
- Buttons are limited to 2 by Discord
- Add images in the `Rich Presence` section of the developer portal


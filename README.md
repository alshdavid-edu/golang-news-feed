# Go Newsfeed API

## Getting Started

```bash
make dev
```

## Viewing Items:

Navigate to [http://localhost:8080/newsfeed](http://localhost:8080/newsfeed)

## Adding Items

Open the JavaScript console and enter:

```javascript
fetch('/newsfeed', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ 
    message: 'Hello World',
  })
})
```

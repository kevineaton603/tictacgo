# Tic Tac Go

Server Implementation of Tic Tac Toe using Echo, Templ, HTMX and Tailwind.css

## Development

Run the Project

```bash
go run main.go
```

Watch for Changes to Templ Files

```bash
templ generate --watch
```

Watch for Tailwind CSS Changes (NPM Required)

```bash
npx tailwindcss -i ./assets/css/main.css -o ./assets/css/output.css --watch
```

***Note: The project execution is not watched so you will need to restart the server when changes are made to Templ Files***

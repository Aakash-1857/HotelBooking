# 🏨 Hotel Room Booking Backend (Go)

A lightweight backend service for managing hotel room bookings, built using Go. This project demonstrates a clean architecture, RESTful routing, session management, and modular design — ideal for scalable web services.

---

## 📌 Features

- RESTful APIs for room booking and availability
- Modular folder structure with `cmd/`, `pkg/`, and `internal/`
- Secure session management (via `github.com/alexedwards/scs/v2`)
- Routing with `go-chi/chi`
- HTML template rendering for basic views
- Form validation and error handling

---

## 🧱 Tech Stack

- **Language:** Go (Golang)
- **Routing:** chi (`github.com/go-chi/chi/v5`)
- **Sessions:** SCS (`github.com/alexedwards/scs/v2`)
- **Templates:** Go's built-in `html/template`
- **Storage:** In-memory (can be extended to DB)

---

## 🗂️ Project Structure

```bash
.
├── cmd/             # Application entry point (main.go)
├── pkg/             # App-level config, handlers, routes
├── internal/        # Utilities, helpers, middlewares
├── templates/       # HTML templates
├── static/          # CSS, JS, etc.
└── go.mod

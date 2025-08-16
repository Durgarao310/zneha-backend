│── cmd/ # Application entry points
│ └── server/ # Main server entrypoint
│ └── main.go
│
│── configs/ # Configuration files (YAML, ENV, JSON)
│
│── internal/ # Core application logic (private to this repo)
│ ├── api/ # Gin handlers & middleware
│ ├── service/ # Business logic
│ ├── repository/ # Data access layer
│ ├── model/ # Domain models
│ ├── routes/ # Gin routes setup
│ ├── database/ # DB connection & migrations
│ └── config/ # Config loader
│
│── pkg/ # Reusable utility packages
│
│── test/ # Unit & integration tests
│
│── scripts/ # Automation scripts
│
│── .env # Environment variables
│── go.mod # Go modules
│── go.sum
│── README.md # Project overview
│── PROJECT_STRUCTURE.md # This file (explains structure)
# Simple Bank - Go Banking Application

![Go](https://img.shields.io/badge/Go-1.21+-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-blue)
![gRPC](https://img.shields.io/badge/gRPC-1.0+-blue)

A complete banking application built with Go featuring:

- REST & gRPC APIs
- JWT/PASETO authentication
- PostgreSQL database
- Kubernetes deployment

## 🚀 Features

### Core Functionality

- User registration & authentication
- Bank account management
- Money transfers between accounts
- Transaction history

### Technical Features

- ✅ REST API (Gin framework)
- ✅ SQLC for type-safe SQL queries
- ✅ Database migrations
- ✅ JWT & PASETO token support
- ✅ Kubernetes deployment (EKS)
- ✅ Unit tests

## 📂 Project Structure

simple-bank/
├── api/ # REST API implementation
│ ├── account.go # Account CRUD handlers
│ ├── account_test.go # Account tests
│ ├── middleware.go # Authentication middleware
│ ├── middleware_test.go # Middleware tests
│ ├── server.go # HTTP server setup
│ ├── transfer.go # Transfer handlers
│ ├── user.go # User authentication handlers
│ ├── user_test.go # User tests
│ └── validator.go # Request validation
│
├── db/ # Database operations
│ ├── migration/ # Database migrations
│ ├── mock/ # Mock database for testing
│ │ ├── store.go # Mock store interface
│ │ |── query/ # Mock queries
│ │ ├── account.sql
│ │ ├── entry.sql
│ │ ├── transfer.sql
│ │ └── user.sql
│ │
│ └── sqlc/ # SQLC generated code
│ ├── account.sql.go # Account queries
│ ├── entry.sql.go # Entry queries
│ ├── transfer.sql.go # Transfer queries
│ ├── user.sql.go # User queries
│ ├── db.go # Database connection
│ └── models.go # Data models
│
├── gapi/ # gRPC API implementation
│ └── server.go # gRPC server setup
│
├── proto/ # Protocol buffer definitions
│ ├── rpc_create_user.proto # CreateUser RPC
│ ├── rpc_login_user.proto # LoginUser RPC
│ ├── service_simple_bank.proto # Service definition
│ └── user.proto # User message definitions
│
├── token/ # Token management
│ ├── jwt_maker.go # JWT implementation
│ ├── jwt_maker_test.go # JWT tests
│ ├── paseto_maker.go # PASETO implementation
│ ├── paseto_maker_test.go # PASETO tests
│ ├── maker.go # Token maker interface
│ └── payload.go # Token payload structure
│
├── util/ # Utility functions
│
└── eks/ # Kubernetes deployment
├── aws-auth.yml # AWS IAM auth config
├── deployment.yml # Deployment config
└── service.yml # Service config

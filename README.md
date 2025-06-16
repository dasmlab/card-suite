# DASMLAB Card Suite: Cribbage & Beyond

> Modern, cross-platform suite for classic card games.  
> **Phase 1:** Cribbage (multiplayer, bots, slick UX, Go backend)

---

## Architecture Overview

```mermaid
flowchart LR
    %% Backend
    S1["games.dasmlab.org --- Gin Backend --- Go, OAuth, Prometheus"]

    %% Clients
    C1["Client Device 1 --- Android / iOS / Web"]
    C2["Client Device 2 --- Android / iOS / Web"]

    %% Stores
    GP["Google Play Store"]
    AS["Apple App Store"]

    %% App install/update flows
    GP -->|Install/Update| C1
    AS -->|Install/Update| C2

    %% Main flows
    S1 -->|Login/API/WebSocket| C1
    S1 -->|Login/API/WebSocket| C2

    C1 <--> |Multiplayer Game Events| C2

    %% Optional: Backend to stores for update links (can omit these if not needed)
    S1 -- App Update Metadata --> GP
    S1 -- App Update Metadata --> AS
```
# Features

Cribbage: All modes (1v1, 3-way, 2v2, 3 teams)

Bots: Solo/bot-mixed games

Modern UI: Animated board, avatars, touch/tap friendly

OAuth: Keycloak for secure login/session

Cross-platform: Native mobile + browser

Future-proof: Easy to add more games

# Development Quickstart
(Detailed steps per milestone—see PROJECT.MD for full roadmap)

1. Backend
Go 1.22+, Gin, Logrus, Prometheus

cd backend && go run main.go

2. Client
(After framework selection)

cd client && <build commands for target platform>

3. Game Logic
Pure Go, re-used in backend (and exported for bot/client logic if needed)

# Contributing
Open issues/PRs for ideas or improvements!

Game logic contributions especially welcome

# License
© DASMLAB, 202-5



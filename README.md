# DASMLAB Card Suite: Cribbage & Beyond

> Modern, cross-platform suite for classic card games.  
> **Phase 1:** Cribbage (multiplayer, bots, slick UX, Go backend)

---

## Architecture Overview

```mermaid
flowchart LR
    S1["games.dasmlab.org<br/><b>(GIN ENGINE)</b>"]

    C1["<b>client device 1</b>"]
    C2["<b>client device 2</b>"]

    GP["Google Play Store"]
    AS["Apple App Store"]

    %% Backend to clients
    S1 -- "Login/API/SSE/WEBSOCK" --> C1
    S1 -- "Login/API/SSE/WEBSOCK" --> C2

    %% Multiplayer events
    C1 <--> |"MultiPlayer Events"| C2

    %% App Stores to clients
    GP -- "Install Client" --> C1
    AS -- "Install Client" --> C2

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



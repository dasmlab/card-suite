# DASMLAB Card Suite: Cribbage & Beyond

> Modern, cross-platform suite for classic card games.  
> **Phase 1:** Cribbage (multiplayer, bots, slick UX, Go backend)

---

## Architecture Overview

```mermaid
flowchart TD
    subgraph Stores
        A1[Google Play Store]
        A2[Apple App Store]
    end

    C1[Client Device 1<br/>(Android/iOS/Web)]
    C2[Client Device 2<br/>(Android/iOS/Web)]

    S1[games.dasmlab.org<br/>Gin Backend<br/>(Go, OAuth, Prometheus)]

    %% Arrows
    A1-- Install App -->C1
    A2-- Install App -->C2
    C1-- Secure Login/API/WebSocket -->S1
    C2-- Secure Login/API/WebSocket -->S1
    C1<-- Multiplayer Game Events -->C2
    C2<-- Multiplayer Game Events -->C1
    S1-- App Updates/Store Links -->A1
    S1-- App Updates/Store Links -->A2
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



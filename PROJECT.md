# DASMLAB Card Suite — Project Charter

## Status
**Phase:** Planning & Scaffolding  
**Version:** 0.1.0  
**Lead:** Daniel Smith (dasm)  
**Start:** 2025-06-15  
**Last updated:** 2025-06-15

## Vision
Build an extensible, visually-rich cross-platform suite of card games, starting with Cribbage, supporting native and web deployments with a modern, secure backend. Enable multi-device, multi-user play with bots, social features, and tight branding control.

---

## Milestones

### Milestone 1: Project Setup & Framework Selection
- [ ] Research/select cross-platform UI toolkit (Flutter, React Native, Unity, Godot, or Qt; web: Flutter Web, React, or Vue/Quasar)
- [ ] Scaffold mono-repo (backend, shared core logic, client(s))
- [ ] Set up backend (Gin + Logrus/Prometheus, OAuth w/ Keycloak, scalable session mgmt)
- [ ] Define game protocol (WebSocket? GRPC-web? Evaluate for multiplayer)

### Milestone 2: Cribbage Core Engine
- [ ] Aggregate cribbage rules (min. 10 sources)
- [ ] Implement pure game logic in Go (standalone, testable, no UI)
- [ ] Expose core logic via Go API for server and (optionally) client consumption
- [ ] Build test suite: deal, play, score, peg, validate

### Milestone 3: Backend Services
- [ ] Implement lobby/game room mgmt (create/join/find)
- [ ] Integrate OAuth2 (Keycloak) and user mgmt
- [ ] Support bot logic as a first-class player
- [ ] Multiplayer comms (choose between client-server and peer relay, document tradeoffs)

### Milestone 4: Client UX/Framework
- [ ] Build "base client" with:
    - Start, Score, Exit, Branded Theme
    - Table layout: avatars, dynamic seating (2, 3, 4, 6)
    - Card/deck animation hooks
    - Server connection, lobby join, play
- [ ] Deploy/test on:
    - Android, iOS, Web
    - (Plan for Mac/Windows/Linux if framework supports)
- [ ] Integrate with App Store/Play Store requirements

### Milestone 5: Multi-Game Suite Preparation
- [ ] Abstract game logic API for new games (Poker, Euchre, Hearts, etc.)
- [ ] Document patterns for board/animation reuse

### Ongoing
- [ ] Automated CI/CD for backend and all client builds
- [ ] Security review (backend, client, comms)
- [ ] Branding/UX polish

---

## Project Structure (proposed)

```yaml
dasmlab-card-suite/
├── backend/ # Gin server, core logic, bots, API
├── cribbage-core/ # Pure Go cribbage game logic (reusable/testable)
├── client/
│ ├── mobile/ # Shared code (Flutter/ReactNative/Qt/Unity) for Android/iOS
│ └── web/ # Web build (if not fully shared)
├── infra/ # Deploy scripts, Docker, K8s, HAProxy configs
├── docs/
├── resources/
```


---

## Key Tech Decisions (Open Questions)
- **Client Framework:** Flutter (best for code sharing, performant UI, native/web), or React Native + Expo (web + mobile), or Qt (if C++/Go binding preferred), or Unity/Godot (for more “game” feel).
- **Client/Server Protocol:** WebSocket (WS) is the most common, with support for GRPC-web as an option.
- **Animation:** Most cross-platform toolkits have animation libraries, but Flutter is the most seamless for “table/card” visuals.
- **Private game rooms/networking:** Stick to server relay for now; evaluate direct peer networking later.

---

## Notes
- Future games can “plug in” to this framework with new logic/board layouts.
- All rules, scoring, and randomization are enforced on server for fairness.
- Client UX should be touch/gesture friendly but clean and simple for rapid iteration.

---

# END PROJECT.MD

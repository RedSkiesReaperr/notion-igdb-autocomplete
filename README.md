# Notion IGDB autocomplete

Rails web app that automates the completion of video game information in a Notion database.
It watches a Notion database for entries whose title is wrapped in `{{...}}`, looks them up on IGDB, and writes the result back (title, platforms, genres, franchises, release date, cover image).

Tasks are persisted in SQLite, processed asynchronously through Solid Queue, and surfaced live in the UI via Turbo Streams (Solid Cable as the backend).

## Table of contents

- [Requirements](#requirements)
- [Local setup](#local-setup)
- [Docker setup](#docker-setup)
- [Environment variables](#environment-variables)
- [Notion database schema](#notion-database-schema)
- [Web interface](#web-interface)
- [Architecture](#architecture)
- [Migrating from v1](#migrating-from-v1)

## Requirements

- Ruby 4.0.4 (`.tool-versions` and `.ruby-version` both pin it; use [asdf](https://asdf-vm.com) or your favorite manager)
- SQLite 3.8+
- Docker / Docker Compose (only for the containerized setup)

## Local setup

```bash
cp .env.example .env   # fill in the 5 keys, see below
bundle install
bin/rails db:prepare
```

Run the app in two terminals:

```bash
# terminal 1: web server + tailwind watcher + jobs worker (via foreman)
bin/dev
```

Then start the Notion watcher loop once. It will reschedule itself every `REFRESH_DELAY` seconds:

```bash
bin/rails runner "NotionWatcherJob.perform_later"
```

`bin/dev` (via `Procfile.dev`) starts the Rails server, the Tailwind CSS watcher, and the Solid Queue worker (`bin/jobs`). The web UI is at <http://localhost:3000>.

## Docker setup

End users do **not** build the image — they pull it from Docker Hub.

```bash
cp .env.example .env   # fill in the 5 keys
# download docker-compose.yml from the repo, then:
docker compose up
```

The `web` container runs `db:prepare` and schedules `NotionWatcherJob` on boot (only if no in-flight watcher job is already enqueued). The `worker` container runs Solid Queue.

Maintainers build and publish the image with:

```bash
docker build -t redskiesreaperr/notion-igdb-autocomplete:latest .
docker push redskiesreaperr/notion-igdb-autocomplete:latest
```

## Environment variables

| Key | Description |
|-----|-------------|
| `NOTION_API_SECRET` | Internal-integration token created at <https://www.notion.so/my-integrations>. The integration must be connected to the target database. |
| `NOTION_PAGE_ID` | ID of the Notion database (the long hex string in the database URL). |
| `IGDB_CLIENT_ID` | Twitch Developer application client ID. IGDB authenticates through Twitch OAuth. |
| `IGDB_SECRET` | Twitch Developer application secret. |
| `REFRESH_DELAY` | Seconds between two `NotionWatcherJob` runs. Defaults to `2`. |

`.env` is git-ignored. `.env.example` is the template.

## Notion database schema

The Notion database must expose the same properties as v1:

| Property | Type |
|----------|------|
| `Title` | Title |
| `Platforms` | Multi-select |
| `Genres` | Multi-select |
| `Franchises` | Multi-select |
| `Release date` | Date |
| `Time to complete (Main Story)` | Text |
| `Time to complete (Main + Sides)` | Text |
| `Time to complete (Completionist)` | Text |

The watcher picks up any row whose `Title` starts with `{{` **and** ends with `}}`. The markers are stripped before the IGDB search.

The three "Time to complete" properties are kept for forward compatibility with `TimeToBeatTask` ([see below](#migrating-from-v1)) but are not written to by the current job.

## Web interface

<http://localhost:3000> shows the 100 most recent tasks with title, type, status (waiting / running / finished / failed), and the error message if any. Rows update in place over Turbo Streams as jobs progress. Failed tasks expose a `Retry` button that creates a fresh task with the same `notion_id` and re-enqueues it.

## Architecture

| Layer | Technology |
|-------|------------|
| Framework | Rails 8.1.3 |
| Background jobs | Solid Queue (SQLite-backed, separate database) |
| Real-time UI | Turbo Streams over Solid Cable (SQLite-backed) |
| Database | SQLite, three logical databases (`primary`, `queue`, `cable`) |
| CSS | Tailwind CSS 4 |
| Service layer | [`service_actor`](https://github.com/sunny/actor) actors in `app/actors/` |

### Source layout

```
app/
├── actors/             # single-purpose service objects (search, pick, update)
├── controllers/        # TasksController#index / #retry
├── jobs/               # NotionWatcherJob (self-rescheduling), GameInfosJob
├── models/
│   ├── task.rb         # AR record with status / task_type enums
│   └── igdb/game.rb    # Notion::Writeable IGDB game representation
├── services/
│   ├── choose.rb       # closest_game via Levenshtein (amatch)
│   ├── igdb/client.rb  # IGDB REST + Twitch OAuth token cache
│   └── notion/         # Notion API client + Writeable module
└── views/tasks/        # ERB + Turbo Streams
```

### Job lifecycle

1. `NotionWatcherJob` queries Notion for `Title` rows wrapped in `{{...}}`, creates a `Task` per new entry, enqueues a `GameInfosJob`, and reschedules itself after `REFRESH_DELAY`.
2. `GameInfosJob` runs three actors: `SearchGame` (IGDB), `PickClosestGame` (Levenshtein), `UpdateNotionPage`. It updates the `Task` status and broadcasts a Turbo replace to `tasks` on every transition.
3. On failure, the task moves to `failed`, the error is shown in the UI, and a `Retry` button re-enqueues a fresh job.

The `Notion::Writeable` module mirrors v1's Go `PageUpdateRequester` interface: any object that implements `notion_update_request` can be passed to `Notion::Client.update_page`.

---

## Migrating from v1

> The v1 Go source code is preserved in [`legacy/`](./legacy) for reference. The CI workflow for the Go version is under `legacy/.github/workflows/ci.yml`.

### What changes

| V1 (Go) | V2 (Rails) |
|---------|------------|
| CLI / TUI application (`./bin/app`) | Web interface at <http://localhost:3000> |
| `-headless` mode via Docker | `docker compose up` with two services (`web` + `worker`) |
| Single Go binary | Rails server + Solid Queue worker |
| Viper-backed `.env` (hot-reload) | `dotenv-rails` (no hot-reload — restart to apply changes) |
| TUI dashboard for tasks | Web dashboard with Turbo Streams |
| In-memory `core.Task` state machine | Persistent `Task` ActiveRecord rows |

### What stays the same

- The `.env` file and its five keys are identical (`NOTION_API_SECRET`, `NOTION_PAGE_ID`, `IGDB_CLIENT_ID`, `IGDB_SECRET`, `REFRESH_DELAY`).
- The Notion database schema (property names and types).
- The Twitch / IGDB credentials.

### Migration steps

1. Download the new `docker-compose.yml` from the repository.
2. Your existing `.env` file is unchanged — no edits needed.
3. `docker compose pull && docker compose up`.
4. Confirm tasks appear at <http://localhost:3000>.

### Feature not yet ported

`TimeToBeatTask` (HowLongToBeat playtime lookup) is **not** available in v2. Notion entries that would trigger it (rows with no `{{...}}` markers but empty "Time to complete" properties) will not be processed until phase 2.

## License

MIT. See `LICENSE`.

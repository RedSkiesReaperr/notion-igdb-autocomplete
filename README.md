# Notion IGDB autocomplete

![demo](https://github.com/RedSkiesReaperr/notion-igdb-autocomplete/assets/64477486/02de6e81-974f-4ed1-948a-e261cbd29eba)

## How does it works
Watch the database you set up. Waits for a page with a title that matchs the following pattern: `{{GAME_TITLE}}`. Then it asks for IGDB informations and it fills the Notion page.

## Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)

## Setup
1. Follow the [getting started](https://developers.notion.com/docs/create-a-notion-integration#create-your-integration-in-notion) to get your `NOTION_API_SECRET` & `NOTION_PAGE_ID`.
2. Follow the [getting started](https://api-docs.igdb.com/#getting-started) IGDB API to get your `IGDB_CLIENT_ID` & `IGDB_SECRET`
3. Put these in your `.env` file
4. Complete `.env` following the `.env.example` file
5. Run `docker-compose up` command
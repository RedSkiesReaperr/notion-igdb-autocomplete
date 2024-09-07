
# Notion IGDB autocomplete [![CircleCI](https://dl.circleci.com/status-badge/img/gh/RedSkiesReaperr/notion-igdb-autocomplete/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/RedSkiesReaperr/notion-igdb-autocomplete/tree/main) [![reddit thread](https://img.shields.io/badge/Reddit-FF4500?logo=reddit&logoColor=white)](https://www.reddit.com/r/Notion/comments/17dw8js/created_integration_to_automatically_fill_in/?utm_source=share&utm_medium=web2x&context=3) [![Notion template](https://img.shields.io/badge/Notion-%23000000.svg?logo=notion&logoColor=white)](https://plant-pantry-77c.notion.site/Automated-video-games-library-c833cb560feb4b82935a310e508d34c2) [![DockerHub](https://img.shields.io/badge/DockerHub-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/redskiesreaperr/notion-igdb-autocomplete)

This Go project aims to automate the completion of video game information in a Notion database. It simplifies the task of collecting and updating details of video games, including titles, franchises, genres, platforms, and release dates, by connecting to the Notion database.

![demo](https://github.com/RedSkiesReaperr/notion-igdb-autocomplete/assets/64477486/02de6e81-974f-4ed1-948a-e261cbd29eba)

## Key Features
- Notion Integration: Connects to your Notion database to extract and update video game information.
- Automated Search: Performs online searches to retrieve up-to-date information about video games.
- Database Updates: Updates video game entries in Notion with the latest details, such as titles, franchises, genres, platforms, and release dates.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Running the Project](#running-the-project)
	- [Running on Local Machine](#running-on-local-machine)
	- [Running with Docker](#running-with-docker)
5. [Dependencies](#dependencies)
6. [Contributing](#contributing)
7. [License](#license)

## Prerequisites
- Go v1.21.1+ (if running on local machine)
- Make (if running on local machine)
- [Docker](https://www.docker.com/products/docker-desktop/) (if running on docker)
- Access to a Notion database and a Notion integration key (see Configuration section for details)
- Access to IGDB API by a client id and client secret (see Configuration section for details)

## Installation
> This section is mandatory only if you want to download the source & run the code on your local machine. Either, you can directly go to the [Configuration section]()

To get started, you need to install Go, download the source code and install all the dependencies it needs to work properly.

### Clone the Repository
```bash
$ git  clone  https://github.com/RedSkiesReaperr/notion-igdb-autocomplete
$ cd  notion-igdb-autocomplete
```

### Install Dependencies
```bash
$ go mod download
```

## Configuration
This section will help you to get all the mandatory variables from third party apps and create the base application configuration. 

### 1. Create environment configuration file: 
Create a `.env` file to save your configuration informations:
```shell
$ cp ./.env.example .env
```

#### 2. Create a Notion database
>**If you start from scratch**: Duplicate the [Notion template](https://plant-pantry-77c.notion.site/Automated-video-games-library-c833cb560feb4b82935a310e508d34c2?pvs=4) and go to step 2

>**If you already have a Notion database**: Continue to step 1

1. Create a Notion database (or configure an already existing one) with the following properties:
	- Title: (**type: Title**)
	- Platforms: (**type: Multi-select**)
	- Genres: (**type: Multi-select**)
	- Franchises: (**type: Multi-select**)
	- Release date: (**type: Date**)
	- Time to complete (Main Story): (**type: Text**)
	- Time to complete (Main + Sides): (**type: Text**)
	- Time to complete (Completionist):  (**type: Text**)
  
2. Create a private Notion integration on your account by following the [getting started page](https://developers.notion.com/docs/create-a-notion-integration#create-your-integration-in-notion) before the "Setting up the demo locally" step.
3. Put your integration API secret as value of the `NOTION_API_SECRET` in your `.env` file.
4. Go on your Notion databse, click on the "..." button and on the "copy link" option. As mentionned in [environment variables section](https://developers.notion.com/docs/create-a-notion-integration#environment-variables), Get the ID of your database and put it as value of the `NOTION_PAGE_ID` in your `.env` file.

> ***At this point you should have a Notion database, with all mandatory properties. You should have created a private Notion integration connected with you database. You should have a `.env` file in your cloned project directory who have two values filled: `NOTION_API_SECRET` and `NOTION_PAGE_ID`***

5. Create a Twitch Developer application (needed to use IGDB API) by following the "Account Creation" of [getting started page](https://api-docs.igdb.com/#getting-started). This will give you steps to get your `IGDB_CLIENT_ID` & `IGDB_SECRET`. Afterward fill `IGDB_CLIENT_ID` & `IGDB_SECRET` in your `.env` file. ***If you need more detailed explanations, follow the 'More details about IGDB configuration' section below***.

🎉 **Congrats, configuration done!** 🎉

### More details about IGDB configuration
Once you are on the Twitch developers portal:
1. On the left menu go in the "Applications" section
2. Then click on the "Register your application" purple button
3. In the registration form:
1. Name field: you fill whatever you want.
2. Second field (about OAuth things): `http://localhost`.
3. Click on create button
4. Once created it takes you to the applications listing, then click on "Manage" button for your newly created app.
5. Under the captcha, you have a field "client identifier" (or something like that). This value is your `IGDB_CLIENT_ID`
6. Click on the "New secret" button. It gives you the `IGDB_SECRET`

## Running the Project

### Running on Local Machine

1. Compile the project:
	```bash
	$ make build
	```
2. Run the generated binary:
	```bash
	$ ./bin/app
	```

By running the app this way you will get into a nicer and clearer TUI (text-based user interface) than the more conventional way to run. You can navigate through it by using your keyboard. It will also help you to have a better understanding of whats going on. Here is a sneak peak of this TUI:
![tui-demo](https://github.com/user-attachments/assets/3a0dca3a-7d67-42ff-8c35-016b7f55abd7)

### Running with Docker

1. Pull the docker image:
	```bash
	$ docker pull redskiesreaperr/notion-igdb-autocomplete:latest
	```
2. Create and run the container, send the config values as environment variables to the container:
```bash
$ docker run \
	-e NOTION_PAGE_ID=your_notion_page_id \
	-e IGDB_CLIENT_ID=your_igdb_client_id \
	-e IGDB_SECRET=your_igdb_secret \
	-e NOTION_API_SECRET=your_notion_api_secret \
	-e REFRESH_DELAY=5 \
	redskiesreaperr/notion-igdb-autocomplete:latest
```

## Build a release image
```bash
$ docker build \
	--platform linux/386,linux/amd64,linux/arm,linux/arm64 \
	-t image:tag \
	.
```

## Dependencies
Thanks to all the authors who created and maintains the following packages:
- [agnivade/levenshtein](https://github.com/agnivade/levenshtein)
- [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles)
- [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)
- [corpix/uarand](https://github.com/corpix/uarand)
- [fsnotify/fsnotify](https://github.com/fsnotify/fsnotify)
- [google/uuid](https://github.com/google/uuid)
- [jomei/notionapi](https://github.com/jomei/notionapi)
- [spf13/viper](https://github.com/spf13/viper)
- [RedSkiesReaperr/howlongtobeat](https://github.com/RedSkiesReaperr/howlongtobeat)

## Contributing
If you wish to contribute to this project, please follow these steps:
1. Fork this repository.
2. Create a branch for your feature: git checkout -b feature/feature-name
3. Commit your changes: git commit -m 'Added a new feature'
4. Push your branch: git push origin feature/feature-name
5. Open a Pull Request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
Feel free to open issues or submit feature requests if you have ideas to enhance this project.

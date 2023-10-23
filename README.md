# Notion IGDB autocomplete [![reddit thread](https://img.shields.io/badge/Reddit-FF4500?logo=reddit&logoColor=white)](https://www.reddit.com/r/Notion/comments/17dw8js/created_integration_to_automatically_fill_in/?utm_source=share&utm_medium=web2x&context=3) [![Notion template](https://img.shields.io/badge/Notion-%23000000.svg?logo=notion&logoColor=white)](https://plant-pantry-77c.notion.site/Automated-video-games-library-c833cb560feb4b82935a310e508d34c2) 

This Go project aims to automate the completion of video game information in a Notion database. It simplifies the task of collecting and updating details of video games, including titles, franchises, genres, platforms, and release dates, by connecting to the Notion database.

![demo](https://github.com/RedSkiesReaperr/notion-igdb-autocomplete/assets/64477486/02de6e81-974f-4ed1-948a-e261cbd29eba)

## Key Features
- Notion Integration: Connects to your Notion database to extract and update video game information.
- Automated Search: Performs online searches to retrieve up-to-date information about video games.
- Database Updates: Updates video game entries in Notion with the latest details, such as titles, franchises, genres, platforms, and release dates.

## System Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)
- Access to a Notion database and a Notion integration key (see Configuration section for details)
- Access to IGDB API by a client id and client secret (see Configuration section for details)

## Installation

1. Clone this repository to your local machine:
    ```sh
    git clone https://github.com/RedSkiesReaperr/notion-igdb-autocomplete
    cd notion-igdb-autocomplete
    ```

## Configuration

1. Create environment configuration file:
    ```sh
    cp ./example.env .env
    ```

2.
    >**If you start from scratch**: Duplicate the [Notion template](https://plant-pantry-77c.notion.site/Automated-video-games-library-c833cb560feb4b82935a310e508d34c2?pvs=4) and go to step 4
    >
    >**If you already have a Notion database**: Continue to step 3

3. Create a Notion database (or configure an already existing one) with the following properties:
    - Title:
        - Name: Title
        - Type: Title
    - Platforms:
        - Name: Platforms
        - Type: Multi-select
    - Genres:
        - Name: Genres
        - Type: Multi-select
    - Franchises:
        - Name: Franchises
        - Type: Multi-select
    - Release date:
        - Name: Release date
        - Type: Date

3. Create a private Notion integration on your account by following the [getting started page](https://developers.notion.com/docs/create-a-notion-integration#create-your-integration-in-notion) before the "Setting up the demo locally" step.

4. Put your integration API secret as value of the `NOTION_API_SECRET` in your `.env` file.

5. Go on your Notion databse, click on the "..." button and on the "copy link" option. As mentionned in [environment variables section](https://developers.notion.com/docs/create-a-notion-integration#environment-variables), Get the ID of your database and put it as value of the `NOTION_PAGE_ID` in your `.env` file.

> ***At this point you should have a Notion database, with all mandatory properties. You should have created a private Notion integration connected with you database. You should have a `.env` file in your cloned project directory who have two values filled: `NOTION_API_SECRET` and `NOTION_PAGE_ID`***

6. Create a Twitch Developer application (needed to use IGDB API) by following the "Account Creation" of [getting started page](https://api-docs.igdb.com/#getting-started). This will give you steps to get your `IGDB_CLIENT_ID` & `IGDB_SECRET`. Afterward fill `IGDB_CLIENT_ID` & `IGDB_SECRET` in your `.env` file. ***If you need more detailed explanations, follow the 'More details about IGDB configuration' section below***.

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

## Usage
Install & configure the application following the steps above.

1. Run the application:
    ```sh
    docker-compose up
    ```
2. In you Notion database create a new entry with a title matching the pattern `{{YOUR_DESIRED_GAME_NAME}}`

The application will connect to your Notion database, IGDB API and search for video game information.
Video game details, including titles, franchises, genres, platforms, and release dates, will be updated in your Notion database.

## Any question ?
The answer might be [there](https://www.reddit.com/r/Notion/comments/17dw8js/created_integration_to_automatically_fill_in/?utm_source=share&utm_medium=web2x&context=3)

## Contributing
If you wish to contribute to this project, please follow these steps:

1. Fork this repository.
2. Create a branch for your feature: git checkout -b feature/feature-name
3. Commit your changes: git commit -m 'Added a new feature'
4. Push your branch: git push origin feature/feature-name
5. Submit a Pull Request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Feel free to open issues or submit feature requests if you have ideas to enhance this project.

## Contributors

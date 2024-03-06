# Recipe API Backend Server Documentation

This documentation provides a guide on how to set up and use the Recipe API backend server. The server allows users to perform CRUD (Create, Read, Update, Delete) operations on recipes stored in a database.

## Table of Contents

- [Recipe API Backend Server Documentation](#recipe-api-backend-server-documentation)
  - [Table of Contents](#table-of-contents)
  - [Overview ](#overview-)
  - [Setup ](#setup-)
    - [Prerequisites ](#prerequisites-)
    - [Installation ](#installation-)
    - [Database Setup ](#database-setup-)
  - [Endpoints ](#endpoints-)
    - [Create Recipe ](#create-recipe-)
    - [List Recipes ](#list-recipes-)
    - [Get Recipe ](#get-recipe-)
    - [Update Recipe ](#update-recipe-)
    - [Delete Recipe ](#delete-recipe-)
  - [Conclusion](#conclusion)

## Overview <a name="overview"></a>

The Recipe API backend server provides HTTP endpoints for managing recipes. It supports creating, listing, getting, updating, and deleting recipes. The server is built using Go programming language and utilizes the GORM library for database operations.

## Setup <a name="setup"></a>

### Prerequisites <a name="prerequisites"></a>

Before setting up the server, ensure you have the following prerequisites installed:

-   Go programming language (version 1.14 or later)
-   SQLite database (or any other supported by GORM)
-   Git (optional, for cloning the repository)

### Installation <a name="installation"></a>

1. Clone the repository or download the source code:

    ```bash
    git clone https://github.com/StaphoneWizzoh/Recipes_StdLib.git
    ```

2. Navigate to the project directory:

    ```bash
    cd <project_directory>
    ```

3. Install dependencies:
    ```bash
    go mod tidy
    ```

### Database Setup <a name="database-setup"></a>

1. The server uses SQLite by default. If you prefer a different database, modify the database driver in `store.go`.

2. Create a SQLite database file:

    ```bash
    touch test.db
    ```

3. Perform database migrations to create necessary tables:
    ```bash
    go run main.go migrate
    ```

## Endpoints <a name="endpoints"></a>

The following endpoints are provided by the Recipe API backend server:

### Create Recipe <a name="create-recipe"></a>

-   **URL:** `/recipes`
-   **Method:** POST
-   **Description:** Creates a new recipe.
-   **Request Body:**
    ```json
    {
        "name": "Recipe Name",
        "ingredients": [{ "name": "Ingredient 1" }, { "name": "Ingredient 2" }]
    }
    ```
-   **Response:** HTTP status code 200 if successful.

### List Recipes <a name="list-recipes"></a>

-   **URL:** `/recipes`
-   **Method:** GET
-   **Description:** Retrieves a list of all recipes.
-   **Response:** JSON array containing all recipes.

### Get Recipe <a name="get-recipe"></a>

-   **URL:** `/recipes/{recipe_id}`
-   **Method:** GET
-   **Description:** Retrieves a specific recipe by ID.
-   **Response:** JSON object representing the recipe.

### Update Recipe <a name="update-recipe"></a>

-   **URL:** `/recipes/{recipe_id}`
-   **Method:** PUT
-   **Description:** Updates an existing recipe.
-   **Request Body:** Same as the request body for creating a recipe.
-   **Response:** HTTP status code 200 if successful.

### Delete Recipe <a name="delete-recipe"></a>

-   **URL:** `/recipes/{recipe_id}`
-   **Method:** DELETE
-   **Description:** Deletes a specific recipe by ID.
-   **Response:** HTTP status code 200 if successful.

## Conclusion

The Recipe API backend server provides a simple and efficient way to manage recipes via HTTP endpoints. With easy setup instructions and clear documentation on endpoints, users can quickly integrate and utilize the server for their recipe management needs.

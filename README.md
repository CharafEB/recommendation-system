# Recommendation System

This project is a recommendation system that provides movie recommendations to users based on user clustering and collaborative filtering. It consists of a Go backend and two Python services for clustering and recommendation.

## Architecture

The system has the following components:

*   **Go Backend**: A web server that connects to a PostgreSQL database, and periodically generates CSV files from the `users` and `ratings` tables.
*   **Clustering Service**: A Python service that clusters users into groups based on their features using KMeans clustering.
*   **Recommending Service**: A Python service that recommends movies to users based on user similarity and collaborative filtering. It uses Redis to cache recommendations.

## Machine Learning Algorithms

This project uses the following machine learning algorithms:

*   **K-Means Clustering**: To group users into clusters based on their features.
*   **Cosine Similarity**: To calculate the similarity between users.
*   **Collaborative Filtering (CF)**: As the main approach for generating movie recommendations.

## Services

### Clustering Service

The clustering service groups users into three clusters based on their features.

*   **File**: `clustring service/ClastringUsers.py`
*   **Functionality**:
    *   Reads user data from `tempf/Users.csv`.
    *   Performs KMeans clustering with 3 clusters.
    *   Saves a plot of the clusters to `clustring service/kmeansWUser2.png`.
    *   Updates the `cluster_id` for each user in the database.

Here is an example of the user clustering:

![User Clustering](clustring%20service/kmeansWUser2.png)

*   **How to run**: The service is designed to run as a cron job every 23 hours. You can run it manually using:
    ```bash
    python "clustring service/router.py"
    ```

### Recommending Service

The recommending service generates movie recommendations for each user.

*   **File**: `recomanding service/recommende.py`
*   **Functionality**:
    *   Reads user ratings from `tempf/UsersRating.csv`.
    *   Calculates user similarity using cosine similarity.
    *   Recommends movies from similar users.
    *   Caches the recommendations in Redis for 24 hours.
*   **How to run**: The service is designed to run as a cron job every 24 hours. You can run it manually using:
    ```bash
    python "recomanding service/reouter.py"
    ```

## Backend

The Go backend is the main application that serves the API and manages data.

*   **Main file**: `main.go`
*   **Functionality**:
    *   Connects to a PostgreSQL database.
    *   Periodically generates `Users.csv` and `UsersRating.csv` from the `users` and `ratings` tables every 10 seconds.
    *   Starts a web server.
*   **How to run**:
    ```bash
    go run main.go
    ```

## Database

The system uses a PostgreSQL database with at least two tables:

*   `users`: Stores user information, including `user_id`, `xfeature`, `yfeature`, and `cluster_id`.
*   `ratings`: Stores user ratings for movies, including `UserName`, `FilmName`, and `Rating`.

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
POSTGERS_API_LINE=your_database_url_here
PORT=your_port_here
REDISHOST=your_redis_host_here
REDISUSERNAME=your_redis_username_here
REDISPASSWORD=your_redis_password_here
```

## Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/recommendation-system.git
    cd recommendation-system
    ```
2.  **Install Go dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Install Python dependencies**:
    ```bash
    pip install -r requirements.txt
    ```
    *(Note: A `requirements.txt` file is not provided. You will need to create one based on the imports in the Python files: `scikit-learn`, `matplotlib`, `pandas`, `psycopg2-binary`, `python-dotenv`, `redis`, `schedule`)*

4.  **Set up the database**:
    *   Create a PostgreSQL database.
    *   Create the `users` and `ratings` tables.
    *   Update the `.env` file with your database credentials.

## Usage

1.  **Start the Go backend**:
    ```bash
    go run main.go
    ```
2.  **Run the clustering service**:
    ```bash
    python "clustring service/router.py"
    ```
3.  **Run the recommending service**:
    ```bash
    python "recomanding service/reouter.py"
    ```
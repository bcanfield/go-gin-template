# CRUD App Template with GORM and Gin

This is a CRUD (Create, Read, Update, Delete) app template built with GORM (a Go ORM) and Gin (a web framework). The template demonstrates the usage of GORM and Gin to create a simple app with three models: User, Video, and Comment.

## Prerequisites

Before getting started, ensure that you have the following installed on your system:

- Go https://go.dev/doc/install
- Docker https://docs.docker.com/get-docker/
- Task https://taskfile.dev/installation/

## Getting Started

To use this template, follow these steps:

1. Clone the repository:

    ```shell
    git clone https://github.com/bcanfield/go-gin-template
    cd go-gin-template
    ```

2. Install the dependencies:

    ```shell
    task install
    ```

3. Start your local Database via Docker:

    ```shell
    task db
    ```
4. Start your server:

    ```shell
    task dev
    ```
    
You should now have the app running locally at `http://localhost:3335`. You can use tools like cURL or Postman to interact with the API.

# Models and Relationships
The app consists of the following models:

### User
    ID: The unique identifier of the user.
    Name: The name of the user.
    Email: The email address of the user.
    Password: The password of the user.
    Videos: A list of videos associated with the user.
    Comments: A list of comments made by the user.
    
### Video
    ID: The unique identifier of the video.
    Title: The title of the video.
    Description: The description of the video.
    Views: The number of views the video has.
    Likes: The number of likes the video has.
    User: The user who uploaded the video.
    Comments: A list of comments associated with the video.

### Comment
    ID: The unique identifier of the comment.
    Content: The content of the comment.
    Likes: The number of likes the comment has.
    Dislikes: The number of dislikes the comment has.
    User: The user who posted the comment.
    Video: The video to which the comment belongs.

The relationships between the models are as follows:

### One-to-Many:
    User has a one-to-many relationship with Video (User has many Videos, Video belongs to a User).
    User has a one-to-many relationship with Comment (User has many Comments, Comment belongs to a User).
    Video has a one-to-many relationship with Comment (Video has many Comments, Comment belongs to a Video).
    Please note that this template provides basic CRUD functionality and can be extended as per your specific requirements.

### License
This project is licensed under the MIT License.

    Feel free to copy the above content into your README.md file and modify it to suit your specific project requirements.





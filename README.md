# HelloGo - A Blueprint for a Full-stack Go Web Application

This repo contains all of the example code mentioned in a companion blog post which lays the foundation for a full-stack Go web application. It sports many features out of the box:

* Connection to a Postgres database backend
* Full NPM-powered embedded front-end project
* Server-side templating using Go's `html/template` package
* API middleware pattern
* A Dev container to get to coding quick, even if you haven't installed Go before (or NPM, or Postgres, or...)

## Quick start
1. Install Docker, VS Code, and the Dev Containers extension in VS Code
    * Make sure your Docker engine is running
    * If on Windows, clone the code within WSL2
2. Open your code folder in VSCode, click the `><` in the bottom-left corner -> `Reopen in Container`
3. After everything finishes downloading (first time only will take a few minutes), configure your local PostgresDB instance
```sql
--connect to your Postgres server with an admin role; this is in your /.devcontainer/.env file
CREATE ROLE hellogo LOGIN PASSWORD 'hellogo';
CREATE DATABASE hellogo;
GRANT CONNECT ON DATABASE hellogo TO hellogo;
--connect to your new hellogo database with an admin role; this is in your /.devcontainer/.env file
CREATE SCHEMA hellogo;
GRANT CREATE, USAGE ON SCHEMA hellogo to hellogo;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA hellogo TO hellogo;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA hellogo TO hellogo;
```
4. Head to the `Run and Debug` tab in VSCode and click on the triangle button to run your server.
5. [https://localhost:8080/hello](https://localhost:8080/hello)
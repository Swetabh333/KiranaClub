# KiranaClub Task

This project is created as a part of selection process for KiranaClub. It implements a program for doctors and receptionists and helps perform the following tasks:

- Submit  a job for image processing.

- Get status of any jobs submitted.

## Testing out the project

The project is deployed at [https://kiranaclub.onrender.com](https://kiranaclub.onrender.com) .

[Here](https://app.swaggerhub.com/apis/SWETABHSHREYAM333/KiraanaClub/1.0.0) is the api documentation.

You can change and put the deployment link instead of `localhost:8080` if you do not wish to try the project locally.

**Note** - Since the project uses the free tier of render it might spin down due to inactivity , so while using the deployed link the first response might take a while , please have patience and retry in a minute. Thank you for your understanding.

The project exposes the following endponts :

### Job Submit Route

- **`/api/submit`**
  - **Method:** POST
  - **Description:** Receives a request from stores, checks if the store exists and creates a job to process the images.

### Job Status Route

- **`/api/status`**

  - **Method:** GET
  - **Description:** Checks and returns the status of a job taking jobid as a query param.


## Setting up the project Locally

To set up the project paste the follwing commands in your terminal:

```bash
git clone https://github.com/Swetabh333/KiranaClub.git
cd KiranaClub
go mod tidy
```

This will install all the required dependencies for the project.

Next you have to set the environment for the project.In your root directory in your terminal you have to paste the following information:

```bash
export DSN_STRING="<your_postgres_connection_string>"
```
Now you'll have to build the project with the following command in the root directory

```bash
go build -o bin/exe app/main/main.go
```

This will create an executable in your bin folder, which you can run using

**NOTE** : make sure no other process is running on port 8080

```bash
./bin/exe
```

**Your backend is now listening at port `8080`**.

## Setting up using docker

You can use docker compose to set this up as well

```bash
git clone https://github.com/Swetabh333/KiranaClub.git
cd KiranaClub
```

then run the docker build command

```bash
docker compose up --build
```

**Your backend is now listening at port `8080`**.


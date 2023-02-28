# B-zone
SEP-2023/Q3 Bumbal zone management


### Setting up Frontend
##### Installation
**Node.js** can be installed from the following link: https://nodejs.org/en/download/

##### Set up
In the terminal run `cd frontend`, to access the frontend app. To start the web app run `npm start`. The web app should open a tab automatically, but if this is not the case it is located at http://localhost:3000/. To stop the app, press `Ctrl + c` in the terminal and answer `y` to the prompt (if there is one).   

### Setting up DataBase: MongoDB + Docker
We will use Docker to host a MongoDB instance locally.

- [ ] Install [Docker Desktop](https://www.docker.com/products/docker-desktop/). Run it (update WSL2 if needed).

- [ ] Download [MongoDB Shell](https://www.mongodb.com/try/download/shell).

- [ ] Download [MongoDB Compass](https://www.mongodb.com/try/download/compass).

- [ ] Open Command Prompt (cmd). 

- [ ] Enter ```docker pull mongo```.

- [ ] Enter ```docker images```. You should see something similar:

```
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
mongo        latest    a440572ac3c1   3 weeks ago   639MB
```

- [ ] Create a folder "mongodb" in your root directory e.g. C:\Users\20193308.

- [ ] Enter `cd C:\yourdirectory\mongodb`. e.g. C:\Users\20193308\mongodb.

- [ ] Enter `docker run -d -p 27017:27017 -v ~/mongodb:/data/db --name Bumbal mongo:latest`. The entire command line should look like this: 

```
C:\Users\20193308\mongodb> docker run -d -p 27017:27017 -v ~/mongodb:/data/db --name Bumbal mongo:latest
```

- [ ]  Enter `docker ps`. You should see something similar:

```
C:\Users\20193308\mongodb> docker ps

CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                      NAMES
a56ffa709baa   mongo:latest   "docker-entrypoint.s…"   7 seconds ago    Up 7 seconds    0.0.0.0:27017->27017/tcp    Bumbal
```
- [ ]  Open the Docker Desktop. You should see Container "Bumbal" with status "Running".

- [ ] Press on the Bumbal Container and go to "Terminal". _(Or you can continue using cmd by entering: "docker exec -it Bumbal /bin/bash" in the cmd)_

- [ ]  Enter `mongosh`. You should see something similar:

```
C:\Users\20193308\mongodb> mongosh

Current Mongosh Log ID: 63fdcfa758222f13f077d16e
Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.2
Using MongoDB:          6.0.4
Using Mongosh:          1.6.2
```

- [ ]  Enter `show dbs`. You should see something similar:
```
C:\Users\20193308\mongodb> show dbs

admin   40.00 KiBur products, anonymous usage data is collected and sent to MongoDB periodically (https://www.mongodb.co
config  60.00 KiBolicy).
local   40.00 KiBy running the disableTelemetry() command.
```
Commands that can be used in the shell: `use admin`,`use local`, `use test`, `show collections`, etc... https://www.mongodb.com/docs/mongodb-shell/

- [ ]  Run `MongoDB Compass`. Enter `mongodb://localhost:27017`.


![](https://www.prisma.io/dataguide/content/mongodb/setting-up-a-local-database/windows/16_compass_running.png)

9 minutes Tutorial: https://www.youtube.com/watch?v=xBbSR7xU2Yw

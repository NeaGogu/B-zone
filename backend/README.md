
# How to run

The following steps are going to be the same for both dev teams until one point.

1. First make sure you have docker engine installed on your system ( it should be installed if you installed docker desktop). Follow these instructions for your platform https://docs.docker.com/get-docker/ 
2. In your terminal, go to `./B-zone/backend`. Run `docker compose up -d`. You should see some nicely coloured text like this
![image](https://user-images.githubusercontent.com/70640237/224291503-25dc7053-e831-4057-bc54-e18b16e9f87c.png)
3. If you look into docker desktop you'll also see something like this 
![image](https://user-images.githubusercontent.com/70640237/224291715-f23d961c-8d5e-4918-962d-27acdb4d5851.png)
4. Nice. Now if you lick on `b-zone-container` and you see these lines you're good to go
```bash
INFO    2023/03/10 10:25:16 Pinged your deployment. You successfully connected to MongoDB!
INFO    2023/03/10 10:25:16 Starting server on :4000
```

You can access the backend services at `localhost:4000`.


## For backend devs

Backend devs will want to test their changes made to backend, obviously. Instead of rebuilding the image every time, you can do the following.

1. Keep the `mongodb-container` running and stop the other one (and everything that runs on port 4000).

2. From the `backend` folder, run the following command in the terminal:

```go
go run ./cmd/web
```

You should see a line being printed in the terminal that looks like this:
`INFO	2023/02/28 11:44:50 Starting server on :4000`

By default, the server starts running on address `localhost:4000` and it's trying to connect to a mongodb instance accessible at `mongodb://localhost:27017`. You can pass a different port or mongo address as an environment variable like this:

> NOTE: this example only works on Unix systems...for Windows users idfk

```go
ADDR=":3400" MONGO_URL="some_mongo_address" go run ./cmd/web
```

This will run the server on port `3400` and you should see it reflected in the terminal.

## For frontend devs 

Normally you should not bother with anything else, you can just run the containers from Docker desktop. However, there might be changes made to backend that you need. In that case you can re-build the backend image with the following command:
1. go to the `backend` folder in your terminal
2. run `docker compose up --build --force-recreate b-zone`. **Dont forget the last argument. Otherwise your database instance will be rebuilt and you'll lose all your data**.
3. Nice 👍

![image](https://user-images.githubusercontent.com/70640237/224294005-238c103f-add2-49f2-9d5d-f7658ad4df92.png)

# API Endpoint Guide for our frontend friends
> !! All API calls should contain the JWT token as a Bearer Authorization header!!

![grogu](https://user-images.githubusercontent.com/53708808/226573731-2376ac58-cb0b-40a5-bbc8-b0590f503d43.png)

Here is an overview of the backend endpoints, together with instructions on how to access them:

1. Endpoint for running the `k-means algorithm`:
    - the `url` for this `PUT` request is: http://localhost:4000/bumbal/algorithm/kmeans
    - the request should have a `json` body which looks like this: 
    - ![image](https://user-images.githubusercontent.com/53708808/226574702-fe906675-0c77-47b5-b618-059c1d525e1f.png)
    - for convenience: `{"number_of_clusters":2,"number_of_candidate_clusters":1}`
    - the two fields are integer numbers, set them as needed, it's not mandatory to use the ones from the examples
    - please make sure to also include the users's `authentication token` in the request!
    - the body of the response is in `json` format
    - use the field named `result`, which contains a `list` of type ZoneModel
    - here is an example:
    - ![image](https://user-images.githubusercontent.com/53708808/226575577-102ed80b-5171-455d-a218-2b15a8908f58.png)
    - wow now you can finally say `SSIIIUUUUUU`

2. Endpoint for running the `genetic algorithm`:
    - the `url` for this `PUT` request is: http://localhost:4000/bumbal/algorithm/genetic
    - the request should have a `json` body which looks like this: 
    - ![image](https://user-images.githubusercontent.com/53708808/227586728-3f763ba9-748d-4b86-88ba-86563471c8e0.png)
    - for convenience: `{"number_of_zones":5,"number_of_generations":10000,"maximum_runtime":10}`
    - the three fields are integer numbers, set them as needed, it's not mandatory to use the ones from the examples. There is no need to have more than 10000 generations. The `maximum_runtime` field should include the maximum amount of minutes for which a user is willing to wait. In the example, the algorithm will return what it generated after 10 minutes even if it hasn't finished doing its calculations (in case of a very large input).
    - please make sure to also include the users's `authentication token` in the request!
    - the body of the response is in `json` format
    - use the field named `result`, which contains a `list` of type ZoneModel
    - the `json` structure present in the body of the response should be `similar` to the one from the `k-means` algorithm
    - `VAAAMOOSS`

- `[PUT] /plot/sync` -> call this right after logging in with bumbal; this ensures that the latest zones saved in bumbal are synced with our backend ( **OLD PLOTS WITH ORIGIN "bumbal" ARE REMOVED WHEN CALLING THIS**)
- `[GET] /user/plots` -> get a list of (plot id, plot name) that the user has saved on our backend. The response body looks like this:
```json
[
	{
		"user_plot_id": "47ca8b57-f43b-49b5-b33c-0041e56b1444",
		"user_plot_name": "Bumbal zone - 2023-03-21 21:08:16"
	},
	{
		"user_plot_id": "a8d28475-6e40-4fd6-8e0a-e7ae5673a260",
		"user_plot_name": "another plot"
	}
]
```
- `[GET] /plot/<PLOT_ID>` -> get a PlotModel object that corresponds with the ID provided in the path. The response body looks like this:
```json
{
	"plot_id": "0be65878-32c6-4d30-9686-a483acb906e7",
	"plot_name": "a new plot hi",
	"plot_zones": [
		{
			"zone_id": "",
			"zone_ranges": [
				{
					"zipcode_from": 1234,
					"zipcode_to": 4312
				},
				{
					"zipcode_from": 1234,
					"zipcode_to": 4312
				}
			],
			"zone_fuel_cost": 0,
			"zone_driving_time": 0
		}
	],
	"plot_created_at": "2023-03-21T20:19:04.518Z",
	"plot_saved_at": "2023-03-21T20:19:04.518Z",
	"origin": "bzone"
}
```
- `[POST] /plot/save` -> used to save a new plot configuration. THIS DOES NOT UPDATE EXISTING PLOTS. The request body should have this structure:
```json
{
	"plot_id": "blabla", <- not taken into account; can be ommited
	"plot_name": "a new plot hi", <- requirerd!
	"plot_zones": [     <- required!
		{
			"zone_ranges": [ <- required!
				{
					"zipcode_from": 1234, <- required!
					"zipcode_to": 4312 <- required!
				},
				{
					"zipcode_from": 1234,
					"zipcode_to": 4312
				}
			]
		}
	]
}
```


# Project structure and organization

The project structure looks something like:

- bzone/backend
    - **root** -> any config/non Go specific code 
    - cmd
        - web
        - another_application_which_does_not_exist_yet

    - internal
        - models, etc
    


The `cmd` directory will contain the application-specific code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the `cmd/web` directory.

The `internal` directory will contain the ancillary non-application-specific code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the database/Bumbal API models for the project.

## Why?

There are two big benefits:

1. It gives a clean separation between Go and non-Go assets. All the Go code we write will live exclusively under the `cmd` and `internal` directories, leaving the project root free to hold non-Go assets like UI files, makefiles and module definitions (including `go.mod` file). This can make things easier to manage when it comes to building and deploying the app later.

2. It scales really nicely if we want to add another executable application to our project. For example, we might want to add a CLI to automate some administrative tasks in the future. With this structure, we could create this CLI application under `cmd/cli` and it will be able to import and reuse all the code you’ve written under the internal directory.



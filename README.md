# B-zone
SEP-2023/Q3 Bumbal zone management

In this README file it is detailed which files should be checked or not be checked in the code quality assessment. This is separated into distinct backend and frontend parts, since the backend uses Go, which cannot be checked with the suggested tools and thus in discussion with the course coordinator different tools were chosen.

## Backend
*Should not be checked:*

Automatically generated:
- None

Compiled but unchanged third party resources:
- None

Data files/classes:
- bumbal_swag_models.go
- bzone_models.go
- bZonePlot.go

Test files:
- geneticAlgorithm_test.go
- population_test.go
- pos_test.go
- route_test.go
- sliceFunctions_test.go
- solution_test.go
- clusterToZipcode_test.go
- kMeans_test.go
- handlers_test.go
- helpers_test.go
- plotids_handler_test.go
- routes_test.go
- zip_handler_test.go
- bumbal_zones_test.go
- handler_helpers_test.go
- bumbal_handler_models_test.go
- bZonePlot_test.go
- plotids_test.go
- test_env_models_test.go
- users_test.go
- zipcodes_test.go
- script_test.go

*Should be checked:*

All of the other files in the backend folder.

## Frontend
*Should not be checked:*

Automatically generated:
- node_modules/.
- manifest.json
- robots.txt
- reporWebVitals.js
- setupTests.js
- package-lock.json
- package.json

Compiled but unchanged third party resources:
- None

Data files/classes:
- NL_Boundary.json

Test files:
- src/test/.

*Should be checked:*

Created/changed files, all in frontend/src:

- headerComponent.js
- heatmapComponent.js
- mapComponent.js
- polygonComponents.js
- sliderComponent.js
- textComponent.js
- getActivities.js
- getAllActivities.js
- querryDatabase.js
- pages/home/index.css
- pages/home/index.js
- pages/login/index.css
- pages/login/index.js
- App.css
- App.js

- index.html
- index.js




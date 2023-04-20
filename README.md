# B-zone
SEP-2023/Q3 Bumbal zone management

In this README file it is detailed which files should be checked or not be checked in the code quality assessment. This is separated into distinct backend and frontend parts, since the backend uses Go, which cannot be checked with the suggested tools and thus in discussion with the course coordinator different tools were chosen.

## Root
Files in the root should not be checked, all of B-Zone's source files are in either the backend or the frontend folder.

## Backend
All files listed below are in backend/.

*Should not be checked:*

Automatically generated:
- None

Compiled but unchanged third party resources:
- None

Data files/classes:
- internal/bumbal/bumbal_swag_models.go
- internal/models/bzone_models.go
- internal/models/bZonePlot.go

Test files:
- cmd/genetic/geneticAlgorithm_test.go
- cmd/genetic/population_test.go
- cmd/genetic/pos_test.go
- cmd/genetic/route_test.go
- cmd/genetic/sliceFunctions_test.go
- cmd/genetic/solution_test.go
- cmd/k-means/clusterToZipcode_test.go
- cmd/k-means/kMeans_test.go
- cmd/web/handlers_test.go
- cmd/web/helpers_test.go
- cmd/web/plotids_handler_test.go
- cmd/web/routes_test.go
- cmd/web/zip_handler_test.go
- internal/bumbal/bumbal_zones_test.go
- internal/bumbal/handler_helpers_test.go
- internal/models/bumbal_handler_models_test.go
- internal/models/bZonePlot_test.go
- internal/models/plotids_test.go
- internal/models/test_env_models_test.go
- internal/models/users_test.go
- internal/models/zipcodes_test.go
- scripts/script_test.go

*Should be checked:*

- cmd/genetic/geneticAlgorithm.go
- cmd/genetic/population.go
- cmd/genetic/pos.go
- cmd/genetic/route.go
- cmd/genetic/sliceFunctions.go
- cmd/genetic/solution.go
- cmd/k-means/clusterToZipcode.go
- cmd/k-means/kMeans.go
- cmd/web/bumbal_handler.go
- cmd/web/helpers.go
- cmd/web/main.go
- cmd/web/plotids_handler.go
- cmd/web/routes.go
- cmd/web/zip_handler.go
- internal/bumbal/bumbal_swag_models.go
- internal/bumbal/bumbal_zones_helper.go
- internal/bumbal/handler_helpers.go
- internal/bumbal/runGenetic_handler.go
- internal/bumbal/runKMeans_handler.go
- internal/models/bumbal_handler_models.go
- internal/models/bzone_models.go
- internal/models/bZonePlot.go
- internal/models/errors.go
- internal/models/plotids.go
- internal/models/users.go
- internal/models/zipcodes.go
- internal/test/testing_env.go
- scripts/populate_database.go
- scripts/populate_zones_dummy.go
- scripts/scrips_main.go

## Frontend
All files listed below are in frontend/.

*Should not be checked:*

Automatically generated:
- node_modules/.
- public/manifest.json
- public/robots.txt
- src/reporWebVitals.js
- src/setupTests.js
- package-lock.json
- package.json

Compiled but unchanged third party resources:
- None

Data files/classes:
- src/pages/home/Data/NL_Boundary.json

Test files:
- src/test/.

*Should be checked:*

Created/changed files, all in frontend/src/.:

- pages/home/components/headerComponent.js
- pages/home/components/heatmapComponent.js
- pages/home/components/mapComponent.js
- pages/home/components/polygonComponents.js
- pages/home/components/sliderComponent.js
- pages/home/components/textComponent.js
- pages/home/functions/getActivities.js
- pages/home/functions/getAllActivities.js
- pages/home/functions/querryDatabase.js
- pages/home/index.css
- pages/home/index.js
- pages/login/index.css
- pages/login/index.js
- App.css
- App.js
- index.css
- index.js

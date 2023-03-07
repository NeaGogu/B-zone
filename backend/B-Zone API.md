
# B-Zone API Endpoints documentation




| Type | Method              | First Header      | Second Header                              |
|------|---------------------|-------------------|--------------------------------------------|
| GET  | RetrieveZonesPlot  | Get/zones-plot/{timestamp}   | Retrieve a zones' plot           |
| GET  | CalculateZonesPlot | Get/zones-plot               |        Calculate a zones' plot   |
| GET  | RetrieveZone       | Get/zone/{zoneId}            | Retrieve zone                    |
| GET  | RetrieveZoneRange  | Get/zone-range/{zoneId}      | Retrieve zone range              |
| POST | SaveZonesPlot      | Post/zones-plot/{timestamp}  | Save the current map in the database|


![image](https://user-images.githubusercontent.com/64141509/223172499-ed7ff3bb-a0e0-4a2a-844e-0fcf568be519.png)

## UsersModel

| Name          | Type                  | Description                            | Notes                 |
|---------------|-----------------------|----------------------------------------|-----------------------|
| **Id**        | Pointer to **int64**  |                                        | [optional] [readonly] |
| **Uuid**      | Pointer to **string** | unique per user                        | [optional] [readonly] |



## Zones-PlotModel
| Name          | Type                              | Description                            | Notes      |
|---------------|-----------------------------------|----------------------------------------|------------|
| **Id**        | Pointer to **int64**              | Unique ID                              | [optional] |
| **Name**           | Pointer to **string**                                | ZonesPlot Name            | [optional]            |
| **Zones**     | Pointer to [**[]ZoneModel**]      | Zones that are in the ZonesPlot                           | [optional] |
| **ZonesPlotCreatedAt** | Pointer to **time.Time**           | Timestamp when the Zones Plot was created      | [optional] |
| **ZonesPlotSavedAt** | Pointer to **time.Time**           | Timestamp when the Zones Plot was saved      | [optional] |


## ZoneModel

| Name               | Type                                                 | Description          | Notes                 |
|--------------------|------------------------------------------------------|----------------------|-----------------------|
| **Id**             | Pointer to **int64**                                 | Unique ID            | [optional]            |
| **ZoneRanges**     | Pointer to [**[]ZoneRangeModel**]                    |                      | [optional]            |
| **ZoneFuelCost**       | Pointer to **float64**                               | Total fuel cost      | [optional]            |
| **ZoneDrivingTime**    | Pointer to **int64**                                 | Total driving time   | [optional]            |

## ZoneRangeModel

| Name               | Type                                                 | Description          | Notes                 |
|--------------------|------------------------------------------------------|----------------------|-----------------------|
| **Id**             | Pointer to **int64**                                 | Unique Zone typeID   | [optional]            |
| **zipcode_from**   | Pointer to **int64**                                 | Zipcode range start  | [optional]            |
| **zipcode_to**     | Pointer to **int64**                                 | Zipcode range start  | [optional]            |
| **IsoCountry**     | Pointer to **string**                                 | Iso Country | [optional]            |





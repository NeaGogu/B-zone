
# B-Zone API Endpoints documentation

[![](https://mermaid.ink/img/pako:eNqljzFrwzAQhf-KuCkBY2TZUSJtJVkDhW5Bi4jURNTSGVmGusb_vVJp6NSh9E1373087ha4orEgwbtgvB5UICQips3mPPXJPWM_3zBIctbDdqtCiYu-_cdKyMUN5MmYaMfRjj920RExGhd0KsGvBX8p-fcNKkAF3kavncm_LwVWkO7WWwUyj0bHNwUqrJnTU8KXOVxBpjjZCqbB5J6T07eoPchX3Y_ZHXS4IPoHlFeQC7yD5LzuWsYaxjkVDe2aCmaQTNS7HT-0rGuoaNvsrxV8fBXQWlBRUCb2-wOlnK2fHX58CA?type=png)](https://mermaid.live/edit#pako:eNqljzFrwzAQhf-KuCkBY2TZUSJtJVkDhW5Bi4jURNTSGVmGusb_vVJp6NSh9E1373087ha4orEgwbtgvB5UICQips3mPPXJPWM_3zBIctbDdqtCiYu-_cdKyMUN5MmYaMfRjj920RExGhd0KsGvBX8p-fcNKkAF3kavncm_LwVWkO7WWwUyj0bHNwUqrJnTU8KXOVxBpjjZCqbB5J6T07eoPchX3Y_ZHXS4IPoHlFeQC7yD5LzuWsYaxjkVDe2aCmaQTNS7HT-0rGuoaNvsrxV8fBXQWlBRUCb2-wOlnK2fHX58CA)

| Type | Method              | First Header      | Second Header                          |
|------|---------------------|-------------------|----------------------------------------|
| GET  | RetrieveComparison  | Get/zone/...      | Retrieve 2 multi-polygons to compare   |
| GET  | RetrieveCalculation | Get/zone/...      | Retrieve recalculated multi-polygon    |
| POST | SaveZone            | Get/zone/...      | Save the multi-polygon in the database |



## User Model

| Name          | Type                  | Description                            | Notes                 |
|---------------|-----------------------|----------------------------------------|-----------------------|
| **Id**        | Pointer to **int64**  |                                        | [optional] [readonly] |
| **Uuid**      | Pointer to **string** | unique per user                        | [optional] [readonly] |
| **ReturnJwt** | Pointer to **bool**   | Set to true for requesting a jwt token | [optional] [readonly] |

| Name      | Type                  | Description                            | Notes                 |
|-----------|-----------------------|----------------------------------------|-----------------------|
| **token** | Pointer to **string** |                                        | [optional] [readonly] |

## Polygon Model
| Name          | Type                            | Description                            | Notes      |
|---------------|---------------------------------|----------------------------------------|------------|
| **Id**        | Pointer to **int64**            | Unique ID                              | [optional] |
| **ZoneId**    | Pointer to **int64**            | Unique ID to which the polygon belongs | [optional] |
| **Zip_codes** | Pointer to **[ ] of lists int** | Zone Name                              | [optional] |

## Zone Model

| Name               | Type                                                 | Description          | Notes                 |
|--------------------|------------------------------------------------------|----------------------|-----------------------|
| **Id**             | Pointer to **int64**                                 | Unique ID            | [optional]            |
| **Name**           | Pointer to **string**                                | Zone Name            | [optional]            |
| **FilterTagNames** | Pointer to **[]string**                              | Tag names            | [optional]            |
| **FilterTagNames** | Pointer to **[]string**                              | Tag names            | [optional]            |
| **FilterTags**     | Pointer to **map[string]interface{}**                |                      | [optional]            |
| **Links**          | Pointer to [**[]LinkModel**](LinkModel.md)           |                      | [optional]            |
| **MetaData**       | Pointer to [**[]MetaDataModel**](MetaDataModel.md)   |                      | [optional]            |
| **ZoneRanges**     | Pointer to [**[]ZoneRangeModel**](ZoneRangeModel.md) |                      | [optional]            |
| **Brands**         | Pointer to [**[]BrandModel**](BrandModel.md)         |                      | [optional]            |
| **BrandIds**       | Pointer to **[]int32**                               | Brand ID&#39;s       | [optional]            |
| **ZoneCreatedAt**  | Pointer to **time.Time**                             | created_at date time | [optional] [readonly] |
| **ZoneUpdatedAt**  | Pointer to **time.Time**                             | updated_at date time | [optional] [readonly] |
| **ZoneCreatedBy**  | Pointer to **int32**                                 | created_by user id   | [optional] [readonly] |
| **ZoneUpdatedBy**  | Pointer to **int32**                                 | created_by user id   | [optional] [readonly] |


## Zone Configuration Model
| Name             | Type                   | Description | Notes      |
|------------------|------------------------|-------------|------------|
| **Id**           | Pointer to **int64**   | Unique ID   | [optional] |
| **ZoneId**       | Pointer to **int64**   | -           | [optional] |
| **ActivityTime** | Pointer to **int64**   | -           | [optional] |
| **FuelCost**     | Pointer to **float64** | -           | [optional] |
| **DrivingTime**  | Pointer to **int64**   | -           | [optional] |

## Activities Model
| Name               | Type                  | Description                               | Notes      |
|--------------------|-----------------------|-------------------------------------------|------------|
| **Id**             | Pointer to **int64**  | Unique ID                                 | [optional] |
| **ZoneId**         | Pointer to **int64**  | Unique ID to which the Activities belongs | [optional] |
| **NrOfActivities** | Pointer to **int64**  | Number of activities in the zone          | [optional] |
| **ActivityTypes**  | Pointer to **string** | Activity types in the zone                | [optional] |



## Multi-Polygon Configuration Model
| Name             | Type                      | Description                            | Notes      |
|------------------|---------------------------|----------------------------------------|------------|
| **Id**           | Pointer to **int64**      | Unique ID                              | [optional] |
| **Uuid**         | Pointer to **int64**      | Unique user ID                         | [optional] |
| **PolygonsId**   | Pointer to **[] int64**   | Unique ID to which the polygon belongs | [optional] |
| **Timestamp**    | Pointer to **[] int64**   | -                                      | [optional] |
| **ActivityTime** | Pointer to **[] int64**   | -                                      | [optional] |
| **FuelCost**     | Pointer to **[] float64** | -                                      | [optional] |
| **DrivingTime**  | Pointer to **[] int64**   | -                                      | [optional] |


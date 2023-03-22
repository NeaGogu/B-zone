// External dependencies
import { useEffect, useRef } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import randomColor from "randomcolor";

var zipCodes = []

/** 
* Fetches the initial zone configuration a user has from Bumbal, returns promise of the response from Bumbal API.
* @param {string} userToken - The user token retrieved from local storage.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the user's initial zone configuration.
*/
async function getInitialZones() {
    // Gets list of zones.
    // Definition of URl, body values, and userToken.
    const zonesURL = "https://sep202302.bumbal.eu/api/v2/zone"

    const bodyValues = JSON.stringify({
        "options": {
            "include_zone_ranges": true,
            "include_brands": true
        },
        "filters": {}
    })

    const userToken = localStorage.getItem('token')

    let initialZones = []

    initialZones = await fetch(zonesURL, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${userToken}`
        },
        body: bodyValues
    })
        // Testing if response recorded was ok.
        .then((response) => {
            if (!response.ok) {
                console.log("Response was not ok ???")
                alert("Unable to retrieve this zone configuration!")
            }
            return response.json();
        })

        // Dealing with received list of zones.
        .then(async (data) => {
            return data.items
        })

        .catch(error => console.log(error, 'error'))

    return initialZones
}

/** 
* With a list of zones, fetches their zip codes and updates their polygon coordinates to coordinatesFile.
* @param {Array} zoneList - The list of zones whose zip codes are to be fetched.
* @return {Promise<Array>} Returns a promise that resolves to an array of objects, where each object represents a zone's list of zip codes.
*/
async function getZipCodes(zoneList) {
    // When you retrieve a list of zones from Bumbal API from zone with PUT, you retrieve a list of zone configurations, 
    // which itself includes a list of zones in each zone configuration.
    let zipCodes = []

    for (let i = 0; i < zoneList.length; i++) {
        // In order for the structure to be [[zipcode ranges zone1], [zipcode ranges zone2], etc...].
        var zone = []

        for (let j = 0; j < zoneList[i].zone_ranges.length; j++) {
            var curr = {
                zipFrom: zoneList[i].zone_ranges[j].zipcode_from,
                zipTo: zoneList[i].zone_ranges[j].zipcode_to
            }
            zone.push(curr)
        }
        zipCodes.push(zone)
    }
    return zipCodes
}

/** 
* Gets a list of the coordinates given a list of zip codes.
* @param {Array} zoneList - The list of zones whose zip codes are to be fetched.
* @return {Promise<Array>} Returns a promise that resolves to an array of objects, where each object represents a zone's list of coordinates.
*/
async function getCoordinates(zipsList) {
    let coordinatesList = []
    // For each zone area, fetch the coordinates and compile them together.
    for (let i = 0; i < zipsList.length; i++) {
        // Same idea as in getzipcodes().
        var struct = []

        for (let j = 0; j < zipsList[i].length; j++) {
            struct.push(await fetch('http://localhost:4000/test/zip/coordinates?zip_from=' + zipsList[i][j].zipFrom.toString() + '&zip_to=' + zipsList[i][j].zipTo.toString())
                .then((response) => {
                    if (!response.ok) {
                        console.log("Response from our backend is not ok ???")
                    }
                    return response.json()
                })
                .then((data) => {
                    return data
                })

                .catch(error => console.log(error)))
        }
        coordinatesList.push(struct)
    }
    return coordinatesList
}

/** 
* Fetches the calculated zone configuration from the backend, returns promise of the response from the backend.
* @param {string} userToken - The user token retrieved from local storage.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the user's calculated zone configuration.
*/
async function calculateZone(){
    // set token
    const userToken = localStorage.getItem('token')
    
    //request url
    const url = "http://localhost:4000/bumbal/algorithm/kmeans"

    //body
    // FOR NOW HARD CODED
    const bodyValues = JSON.stringify({
        "number_of_clusters":2,
        "number_of_candidate_clusters":1
    })

    // array to hold zone configuration
    let calculatedZones = []

    calculatedZones = await fetch(url, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${userToken}`
        },
        body: bodyValues
    })
        // Testing if response recorded was ok.
        .then((response) => {
            if (!response.ok) {
                console.log("Response was not ok ???")
                alert("Unable to retrieve this zone configuration!")
            }
            return response.json();
        })

        // Dealing with received list of zones.
        .then(async (data) => {
            console.log(data.result)
            return data.result
        })

        .catch(error => console.log(error, 'error'))
    

    // cleaning up array
    console.log(calculatedZones)
    var zoneConfig = []

    // go into each zone
    for (let i =0; i < calculatedZones.length; i++) {
        var zoneRanges = calculatedZones[i].zone_ranges
        var currZoneRange =[]
        // go into each range
        for (let j = 0; j < zoneRanges.length; j++) {
            //convert to format used in other GetCoordinates
            var curr = {
                zipFrom: zoneRanges[j].zipcode_from,
                zipTo: zoneRanges[j].zipcode_to
            }
            currZoneRange.push(curr)
        }
        zoneConfig.push(currZoneRange)
    }    

    return zoneConfig
    


}

// Main function to visualize the polygons on the map.
const PolygonVis = (props) => {
    //selections
    const { zoneId } = props

    //keeps track if this is first rendering
    const renderRef = useRef()

    // Map context.
    const context = useLeafletContext()

    // this use effect is called when instantiating the polygons, making them to be the ones from bumbal's api
    // CHANGE IT TO BE BASED ON LAST VIEWED
    useEffect(() => {
        // maybe not needed
        context.layerContainer.eachLayer(function (layer) {
            context.layerContainer.removeLayer(layer)
        })
        // set it that first render has been done
        renderRef.current = 1;

        // Async function in order to wait for response from API.
        const fetchData = async () => {
            let coordinatesList = []

            // Delete old heat layer if it exists.
            context.layerContainer.eachLayer(function (layer) {
                console.log(layer)
            })

            // Get initial zones from Bumbal.
            let initialZones = await getInitialZones();
            zipCodes = await getZipCodes(initialZones);
            // console.log(zipCodes)
            coordinatesList = await getCoordinates(zipCodes)

            // Iterates through zones.
            for (let i = 0; i < coordinatesList.length; i++) {
                // Tterates through zone ranges inside of zones.
                let color = randomColor({ luminosity: 'dark' });
                for (let j = 0; j < coordinatesList[i].length; j++) {
                    // Iteratres through coordinates in zone ranges.
                    for (let k = 0; k < coordinatesList[i][j].length; k++) {
                        let polygon = L.polygon(coordinatesList[i][j][k].zone_coordinates)
                        polygon.setStyle({ color: color })
                        context.layerContainer.addLayer(polygon)
                    }
                }
            }
        };
        fetchData()
    }, [context.layerContainer])

    // updating the polygon based on selection
    useEffect(() => {
        // prevention from initial render being rendered twice
        if (renderRef.current === 1) {
            // increase the rendering in order for it to render on next update
            renderRef.current += 1;
        } else {
            // remove old polygon layers
            context.layerContainer.eachLayer(function (layer) {
                context.layerContainer.removeLayer(layer)
            })
            
            // Async function in order to wait for response from API.
            const fetchData = async () => {
                var coordinatesList = []
    
                // Delete old heat layer if it exists.
                context.layerContainer.eachLayer(function (layer) {
                    console.log(layer)
                })
                
                // check if bumbal zone is wanted or zone by id?
                // FOR NOW CALCULATE IS ID OF CALCULATED ZONE
                

                if (zoneId === 'calculate') {
                    zipCodes = await calculateZone()
                    coordinatesList = await getCoordinates(zipCodes)
                } 
                else if (zoneId === 'initial') {
                     // Get initial zones from Bumbal.
                    let initialZones = await getInitialZones();
                    zipCodes = await getZipCodes(initialZones);
                    // console.log(zipCodes)
                    coordinatesList = await getCoordinates(zipCodes)

                }
                console.log('inside useeef')
                console.log(zipCodes)
    
                // Iterates through zones.
                for (let i = 0; i < coordinatesList.length; i++) {
                    // Tterates through zone ranges inside of zones.
                    let color = randomColor({ luminosity: 'dark' });
                    for (let j = 0; j < coordinatesList[i].length; j++) {
                        // Iteratres through coordinates in zone ranges.
                        for (let k = 0; k < coordinatesList[i][j].length; k++) {
                            let polygon = L.polygon(coordinatesList[i][j][k].zone_coordinates)
                            polygon.setStyle({ color: color })
                            context.layerContainer.addLayer(polygon)
                        }
                    }
                }
            };
            fetchData()
        }        
    }, [context.layerContainer, zoneId])
    return null
}

export default PolygonVis
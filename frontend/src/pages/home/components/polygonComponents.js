// External dependencies
import { useEffect } from 'react'
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
* Gets a list of the coordinates given a list of zip codes.
* @param {Array} plot - Zone configuration which is to be converted.
* @return {Promise<Array>} Returns a zone configuration in the structed expected by the backend.
*/
function convertToStructure(plot) {
    // variable to hold the plot
    var plt = []

    for (let i = 0; i < plot.length; i++) {
        // variable for to the the zone ranges
        var zone_ranges = []
        for (let j = 0; j < plot[i].length; j++) {
            // variable to hold one zone range
            var curr = {
                zipcode_from: plot[i][j].zipFrom,
                zipcode_to: plot[i][j].zipTo
            }
            // add to zone ranges
            zone_ranges.push(curr)
        }
        // variable to hold the zones ranges with appropriate structure
        var zoneobj = {
            "zone_ranges": zone_ranges
        }
        // add the zone range to the plot
        plt.push(zoneobj)
    }
    // variable to hold the plot in the required structure
    var pltfinal = {
        "plot_zones": plt
    }

    return pltfinal
}

/** 
* Fetches the calculated zone configuration from the backend, returns promise of the response from the backend.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the user's calculated zone configuration.
*/
async function calculateZone(algorithm) {
    // set token
    const userToken = localStorage.getItem('token')

    // array to hold zone configuration
    let calculatedZones = []

    // calculated based on algorithm
    if (algorithm === 1) {
        //request url
        const url = "http://localhost:4000/bumbal/algorithm/kmeans"
        //body
        // FOR NOW HARD CODED
        const bodyValues = JSON.stringify({
            "number_of_clusters": 2,
            "number_of_candidate_clusters": 1
        })


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
                //console.log(data.result)
                return data.result
            })
            .catch(error => console.log(error, 'error'))
    } else {
        //request url
        const url = "http://localhost:4000/bumbal/algorithm/genetic"
        //body
        // for now hard coded
        const bodyValues = JSON.stringify({
            "number_of_zones": 5,
            "number_of_generations": 10000,
            "maximum_runtime" : 10
        })
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
                //console.log(data.result)
                return data.result
            })
            .catch(error => console.log(error, 'error'))

    }


    // cleaning up array
    //console.log(calculatedZones)
    var zoneConfig = []

    // go into each zone
    for (let i = 0; i < calculatedZones.length; i++) {
        var zoneRanges = calculatedZones[i].zone_ranges
        var currZoneRange = []
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

    return [zoneConfig, calculatedZones]



}

/** 
* Fetches saved plot from the backend, returns promise of the response from the backend.
* @param {string} plotID - Id of plot to be looked up.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the zone configuration of the plot.
*/
async function querryDatabase(plotID) {
    const userToken = localStorage.getItem('token')
    var zones = []
    await fetch("http://localhost:4000/plot/" + plotID, {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${userToken}`
        }
    }).then((response) => {
        if (!response.ok) {
            console.log("Response from our backend is not ok ???")
        }
        return response.json()
    }).then((data) => {
        zones = data.plot_zones
    })

    var zoneConfig = []

    // go into each zone
    for (let i = 0; i < zones.length; i++) {
        var zoneRanges = zones[i].zone_ranges
        var currZoneRange = []
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
    const { zoneId, setZipCodes, setComputed, algorithm } = props

    // Map context.
    const context = useLeafletContext()

    // Runs when a polygon is to be generated
    // CHANGE IT TO BE BASED ON LAST VIEWED
    useEffect(() => {
        //console.log('hellobro')
        // maybe not needed
        context.layerContainer.eachLayer(function (layer) {
            context.layerContainer.removeLayer(layer)
        })
        // set it that first render has been done


        // Async function in order to wait for response from API.
        const fetchData = async () => {
            // set the render state to be false
            setComputed(false)
            // variable which holds the coordinates to be displayed
            var coordinatesList = []

            // check if zone is to be calculated
            if (zoneId.startsWith('calculate')) {
                var calculation;
                if (algorithm === 1) {
                    calculation = await calculateZone(algorithm)
                } else {
                    calculation = await calculateZone(algorithm)
                }
                
                convertToStructure(calculation[0])
                setZipCodes(convertToStructure(calculation[0]));
                coordinatesList = await getCoordinates(calculation[0])

            }
            // check if this is the initial run which displays the zone in bumbal
            else if (zoneId === 'initial') {
                let initialZones = await getInitialZones();
                zipCodes = await getZipCodes(initialZones);
                coordinatesList = await getCoordinates(zipCodes)
            }
            // otherwise get zone from our database
            else {
                let querryZone = await querryDatabase(zoneId);
                coordinatesList = await getCoordinates(querryZone)
            }

            // process which draws the zone
            //Iterates through zones.
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
            // set render state to be true
            setComputed(true)
        };
        fetchData()
    }, [context.layerContainer, setZipCodes, zoneId, setComputed])


    return null
}

export default PolygonVis
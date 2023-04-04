// External dependencies
import { useEffect } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import querryDatabase from '../functions/querryDatabase'

import { voronoi, featureCollection, point, intersect } from '@turf/turf'
import nl from '../Data/NL_Boundary.json'
var zipCodes = []
// array of colors to display
const colors = ['#e6194b', '#3cb44b', '#ffe119', '#4363d8', '#f58231', '#911eb4', '#46f0f0', '#f032e6', 
                '#bcf60c', '#fabebe', '#008080', '#e6beff', '#9a6324', '#fffac8', '#800000', '#aaffc3', 
                '#808000', '#ffd8b1', '#000075', '#808080', '#ffffff', '#000000']

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
* Fetches the calculated zone configuration from the backend, returns promise of the response from the backend.
* @param {int} algorithm - The algorithm of choice for the calculation
* @param {int} nrofzones - The number of zones to make
* @return {Promise<Array>} A promise that resolves with an array of objects representing the user's calculated zone configuration.
*/
async function calculateZone(algorithm, nrofzones) {
    // set token
    const userToken = localStorage.getItem('token')

    // array to hold zone configuration
    let calculatedZones = []
    let clusters = []

    // calculated based on algorithm
    if (algorithm === 1) {
        //request url
        const url = "http://localhost:4000/bumbal/algorithm/kmeans"
        //body
        // FOR NOW HARD CODED
        const bodyValues = JSON.stringify({
            "number_of_clusters": parseInt(nrofzones),
            "number_of_candidate_clusters": 1
        })


        let arr = await fetch(url, {
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
                let arr = []
                arr[0] = data.zone_model_result
                arr[1] = data.clusters_result
                //console.log(arr)
                return arr
            })
            .catch(error => console.log(error, 'error'))
        calculatedZones = arr[0]
        clusters = arr[1]
    } else {
        console.log('here')
        //request url
        const url = "http://localhost:4000/bumbal/algorithm/genetic"
        //body
        // for now hard coded
        const bodyValues = JSON.stringify({
            "number_of_zones": parseInt(nrofzones),
            "number_of_generations": 10000,
            "maximum_runtime": 10
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
                console.log(data.zone_model_result)
                return data.zone_model_result
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

    if (algorithm === 1) {
        return [zoneConfig, clusters]
    }
    return [zoneConfig, calculatedZones]



}


/**
* Function to render the non-expanded zones
* @param {Object} plotID - Id of plot to be looked up.
* @param {Arrat} coordinatesList - Coordinates of zones to be rendered
* @return {Void}
*/
function renderZones(context, coordinatesList) {
    //Iterates through zones.
    for (let i = 0; i < coordinatesList.length; i++) {
        for (let j = 0; j < coordinatesList[i].length; j++) {
            // Iteratres through coordinates in zone ranges.
            for (let k = 0; k < coordinatesList[i][j].length; k++) {
                let polygon = L.polygon(coordinatesList[i][j][k].zone_coordinates)
                polygon.setStyle({ color: colors[i] })
                polygon.bindTooltip(`Zone ${i}`)
                context.layerContainer.addLayer(polygon)
            }
        }
    }
}

// Main function to visualize the polygons on the map.
const PolygonVis = (props) => {
    //selections
    const { zoneId, setZipCodes, setComputed, algorithm, nrofzones, setCalculatedZone, voronoib } = props

    // Map context.
    const context = useLeafletContext()

    // Runs when a polygon is to be generated
    useEffect(() => {

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
            var calculation;

            // check if zone is to be calculated
            if (zoneId.startsWith('calculate')) {

                if (algorithm === 1) {

                    calculation = await calculateZone(algorithm, nrofzones)

                    convertToStructure(calculation[0])
                    setZipCodes(convertToStructure(calculation[0]));

                    // if expanded is selected render expanded zones
                    if (voronoib) {
                        // Netherlands bounding box
                        var options = {
                            bbox: [3.0741,50.7368,7.2208,53.749]
                        };
                        // make points in geojson format
                        let points = []
                        for (let i = 0; i < calculation[1].length; i++) {
                            points.push(point([calculation[1][i].Center.Longitude, calculation[1][i].Center.Latitude]))
                        }
                        let featurePoints = featureCollection(points)
                        let voronoiPolygons = voronoi(featurePoints, options);
                        let layers = []
                        // find intersection of polyongs with nl boundary and add them to layer
                        for(let i = 0; i<voronoiPolygons.features.length; i++) {
                            var intersection = intersect(nl.features[0],voronoiPolygons.features[i])
                            layers.push(new L.geoJSON(intersection, {style : {color:colors[i]}}).bindTooltip(`Zone ${i}`))
                        }
                        // add laters to map
                        context.layerContainer.addLayer(L.layerGroup(layers))
                        // set computed to true
                        setComputed(true)
                        return;
                    }

                } else {
                    calculation = await calculateZone(algorithm, nrofzones)
                }

                var zipcodes =  convertToStructure(calculation[0])
                setZipCodes(convertToStructure(calculation[0]));
                setCalculatedZone(zipcodes)
                coordinatesList = await getCoordinates(calculation[0])

            }
            // check if this is the initial run which displays the zone in bumbal
            else if (zoneId === 'initial') {
                let initialZones = await getInitialZones();
                zipCodes = await getZipCodes(initialZones);
                setCalculatedZone(zipCodes)
                coordinatesList = await getCoordinates(zipCodes)
            }
            // otherwise get zone from our database
            else {
                setZipCodes([])
                let querryZone = await querryDatabase(zoneId);
                coordinatesList = await getCoordinates(querryZone)
            }

            renderZones(context,coordinatesList)
            // set render state to be true
            setComputed(true)
        };
        fetchData()

    // warning not usefull
    // eslint-disable-next-line
    }, [context.layerContainer, setZipCodes, zoneId, setComputed])


    return null
}

export default PolygonVis
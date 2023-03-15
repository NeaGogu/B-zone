//imports
import React, { useEffect, useRef, useState } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import {wait} from "@testing-library/user-event/dist/utils";


//create zipFrom and zipTo object
var zips = {
    zipFrom: "",
    zipTo: ""
};
var zipCodes = []
var coordinatesList = []

//fetches the initial zone configuration a user has from Bumbal, returns promise of the response from Bumbal API
async function getInitialZones() {
    //gets list of zones
    //definition of URl, body values, and userToken
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
        //testing if response recorded was ok
        .then((response) => {
            if(!response.ok) {
                console.log("Response was not ok ???")
                alert("Unable to retrieve this zone configuration!")
            }
            return response.json();
        })
        //dealing with received list of zones
        .then(async (data) => {
            console.log("returning data items from bumbal (this is from getInitialZones)")
            return data.items
        })
        .catch(error => console.log(error, 'error'))

    return initialZones
}

//with a list of zones, fetches their zip codes and updates their polygon coordinates to coordinatesFile
async function getZipCodes(zoneList) {
    //when you retrieve a list of zones from Bumbal API from zone with PUT, you retrieve a list of zone configs which itself includes a list of zones in each zone config

    let zipCodes = []
    for (let i = 0; i < zoneList.length; i++) {
        zipCodes[i] = await getAreas(zoneList[i].zone_ranges); //per each zone, there are a list of areas with from and to
    }

    return zipCodes

}

async function getAreas(zoneAreas) {
    //create array of zipcode ranges
    let zoneAreaZips = [zips]
    //loop through range items to find zip code ranges
    for (let j = 0; j < zoneAreas.length; j++) {
        zoneAreaZips[j].zipFrom = zoneAreas[j].zipcode_from;
        zoneAreaZips[j].zipTo = zoneAreas[j].zipcode_to;

    }
    return zoneAreaZips
}

async function getCoordinates(zipsList) {
    let coordinatesList = []
    //for each zone area, fetch the coordinates and compile them together
    for(let j = 0; j < zipsList.length; j++) {
        coordinatesList[j] = await fetch('http://localhost:4000/test/zip/coordinates?zip_from=' + zipsList[j].zipFrom.toString() + '&zip_to=' + zipsList[j].zipTo.toString())
            .then((response) => {
                if(!response.ok) {
                    console.log("Response from our backend is not ok ???")
                }
                return response.json()
            })
            .then((data) => {
                return data
            }).catch(error => console.log(error))
        console.log("verifying output of await for coordinatesList (from getCoordinates)")
        console.log(coordinatesList[j])
    }

    return coordinatesList
}


//main function
const PolygonVis = () => {
    // map context
    const context = useLeafletContext()
    //const zoneList = getInitialZones()

    useEffect(() => {
        // async function in order to wait for response from api
        const fetchData = async () => {
            // delete old heat layer if it exists
            context.layerContainer.eachLayer(function (layer) {
                console.log(layer)
            })

            //get initial zones from Bumbal
            let initialZones = await getInitialZones();
            zipCodes = await getZipCodes(initialZones);
            console.log("Getting coordinates for first zone area... (from fetchData)")
            coordinatesList = await getCoordinates(zipCodes[0]);
            console.log(coordinatesList[0])
            return zipCodes

            // map those points to something interpretable for the heatmap
            // const points = addressPoints
            //     ? addressPoints.map((p) => {
            //         // if activity time is selected
            //         if (value === 1) {
            //             return [p[0], p[1], p[2]];
            //         }
            //         // if location is selected
            //         return [p[0], p[1], intensity];
            //     })
            //     : [];

            const points = [
                [51.515, -0.09],
                [51.52, -0.1],
                [51.52, -0.12],
            ]
            // create new layer and add it to the map context
            context.layerContainer.addLayer(L.polygon(points))

        };

        fetchData()

    })
}

export default PolygonVis
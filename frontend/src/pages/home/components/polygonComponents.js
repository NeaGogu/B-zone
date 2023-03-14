//imports
import React, { useEffect, useRef, useState } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'


//set of zipCodes
var zipCodes
var coordinatesFile

//with a user token, fetches initial zones list
function getInitialZones(userToken) {
    //get list of zones
    //definition of URl, body values, other parameters
    const zonesURL = "https://sep202302.bumbal.eu/api/v2/zone"
    const bodyValues = JSON.stringify({
        "options": {
            "include_zone_ranges": true,
            "include_brands": true
        },
        "filters": {}
    })

    //sending fetch request to receive list of user's zones and updates their postal codes to zipCodes
    fetch(zonesURL, {
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
        .then((data) => {
            //retrieving the zipCodes and coordinates from list of zones
            zipCodes = getZipCodes(data.items)
        })
        .catch(error => console.log(error, 'error'))
}

//with a list of zones, fetches their zip codes and updates their polygon coordinates to coordinatesFile
async function getZipCodes(zoneList) {
    //when you retrieve a list of zones from Bumbal API from zone with PUT, you retrieve a list of zone configs which itself includes a list of zones in each zone config
    let zipCodes = []
    console.log(zoneList[0].zone_ranges.length)
    console.log("Coordinates file is: ")
    console.log(coordinatesFile)

    for (let i = 0; i < zoneList.length; i++) {
        zipCodes[i] = getAreas(zoneList[i].zone_ranges, i); //per each zone, there are a list of areas with from and to
        //for each zone and an area, fetch the coordinates and compile them together
        for(let j = 0; j < zoneList[i].zone_ranges.length; j++) {
            await fetch('http://localhost:4000/test/zip/coordinates?zip_from=' + zipCodes[i][j].zipFrom.toString() + '&zip_to=' + zipCodes[i][j].zipTo.toString())
                .then((response) => {
                    if(!response.ok) {
                        console.log("Reponse from our backend is not ok ???")
                    }
                    return response.json()
                })
                .then((data) => {
                    console.log(data)
                    return data
                }).catch(error => console.log(error))
        }

        console.log("Set the jsonFile var to be the received zipcode data from our backend")

    }

    return zipCodes

}

function getAreas(zoneAreas, index) {
    //create zipFrom and zipTo object
    var zips = {
        zipFrom: "",
        zipTo: ""
    };
    //create array of zipcode ranges
    let zoneAreaZips = [zips]
    //loop through range items to find zip code ranges
    for (let j = 0; j < zoneAreas.length; j++) {
        zoneAreaZips[j].zipFrom = zoneAreas[j].zipcode_from;
        zoneAreaZips[j].zipTo = zoneAreas[j].zipcode_to;


    }
    return zoneAreaZips
}


//main fucntion
const PolygonVis = () => {

}

export default PolygonVis
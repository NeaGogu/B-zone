//imports
import React, { useEffect, useRef, useState } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import randomColor from "randomcolor";
import {wait} from "@testing-library/user-event/dist/utils";


//create zipFrom and zipTo object
var zips = {
    zipFrom: "",
    zipTo: ""
};
var zipCodes = []


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
        // did this in order for the structure to be [[zipcode ranges zone1], [zipcode ranges zone2], etc...]
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

async function getCoordinates(zipsList) {
    // console.log('zipsgetcoords')
    // console.log(zipsList)
    let coordinatesList = []
    //for each zone area, fetch the coordinates and compile them together

    for (let i = 0; i < zipsList.length; i ++){
        // same idea as in getzipcodes()
        var struct = []
        
        for (let j=0; j <zipsList[i].length; j++) {
            struct.push(await fetch('http://localhost:4000/test/zip/coordinates?zip_from=' + zipsList[i][j].zipFrom.toString() + '&zip_to=' + zipsList[i][j].zipTo.toString())
            .then((response) => {
                if(!response.ok) {
                    console.log("Response from our backend is not ok ???")
                }
                return response.json()
            })
            .then((data) => {
                return data
            }).catch(error => console.log(error)))
        }
        coordinatesList.push(struct)
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
        context.layerContainer.eachLayer(function(layer){
            context.layerContainer.removeLayer(layer)
        })
        
        const fetchData = async () => {
            var coordinatesList = []
            
            // delete old heat layer if it exists
            context.layerContainer.eachLayer(function (layer) {
                console.log(layer)
            })

            //get initial zones from Bumbal
            let initialZones = await getInitialZones();
            // console.log('init')
            // console.log(initialZones)
            zipCodes = await getZipCodes(initialZones);
            console.log('zipcodes')
            //console.log(zipCodes)
            //console.log("Getting coordinates for first zone area... (from fetchData)")
            coordinatesList = await getCoordinates(zipCodes)
            //console.log(coordinatesList)
            
            // itterates through zones
            var arr = []
            for (let i = 0; i < coordinatesList.length; i++) {
                //itterates through zone ranges inside of zones
                let color = randomColor();
                for (let j =0; j < coordinatesList[i].length; j++){
                    // itteratres through coordinates in zone ranges
                    for (let k =0; k<coordinatesList[i][j].length; k++){
                        let polygon = L.polygon(coordinatesList[i][j][k].zone_coordinates)
                        polygon.setStyle({color:color})
                        context.layerContainer.addLayer(polygon)
                    }
                }
            }
            // var group = L.layerGroup(arr);
            // context.layerContainer.addLayer(group)
            
            

         
            // create new layer and add it to the map context
            //context.layerContainer.addLayer(L.polygon(points))

        };

        fetchData()

    })
}

export default PolygonVis
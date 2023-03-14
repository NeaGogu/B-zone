import React, { useEffect, useRef, useState } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import 'leaflet.heat'

/**
Updates the token for the user authentication
@param {string} user - The user email
@param {string} passw - The user password
*/
async function updateToken(user, passw) {
    const response = await fetch("https://sep202302.bumbal.eu/api/v2/authenticate/sign-in", {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "email": user,
            "password": passw,
            "return_jwt": true
        })
    })
    const data = await response.json();
    console.log('new token ' + data.token)
    localStorage.setItem('token', data.token)
  }

  /**
Sends a request to the Bumbal API to retrieve a list of activities.
@returns {Promise<Response>} - The response from the API containing a list of activities.
*/

async function getActivities() {
    const token = localStorage.getItem('token')
    console.log('token ' + token)
    const requestOptions = {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`, // add token to Bearer Authorization when sending GET signOut request
        },
        body: JSON.stringify({
            "options": {
              "include_address": true
            },
            "limit": 10,
            "offset": 0,
            "sorting_column": "id",
            "sorting_direction": "asc",
            "as_list": true,
            "count_only": false
          })
    };
    const response = await fetch('https://sep202302.bumbal.eu/api/v2/activity', requestOptions);
    console.log('getActivities response ' + response.status)
    return response;
  }

/**
Finds the latitude and longitude of each activity address and returns the data as an array.
@returns {Promise<Array>} - The array containing latitude, longitude, and intensity for each address.
*/

async function findAddressesPoints() {
    const response = await getActivities();

    if (response.status === 401) {
        // FIXME Delete after token updated correctly
        await updateToken( "sep@bumbal.eu", "cW$#Qbph0524");
        response = await getActivities();
    }

    const data = await response.json();

    console.log(data) 
    let newData = data.items.map((i) => {
        return [i.address.latitude, i.address.longitude, 500]; // lat lng intensity
        })
    console.log(newData) 
    return newData;
}
  
/**
Renders a Leaflet Heatmap based on the data retrieved from the Bumbal API.
@returns {JSX.Element} - The Leaflet Heatmap component.
*/

const Heatmap = () => {
  const context = useLeafletContext() 
  //var intense = props
  

  useEffect(() => {
    const fetchData = async () => {
        console.log('ran')
        let addressPoints = await findAddressesPoints();
        
        const points = addressPoints
        ? addressPoints.map((p) => {
        return [p[0], p[1], 300]; // lat lng intensity
        })
        : [];

        const ellipse = new L.heatLayer(points)
        const container = context.layerContainer || context.map
        // console.log(context.map.eachLayer(function(layer){
        //   console.log("hello")
        //   console.log(layer._heat + 'hello')
        // }))
        
        container.addLayer(ellipse)
    
        return () => {
          container.removeLayer(ellipse)
        }      
      };
  
      fetchData();

      
    })
  
    
}

export default Heatmap
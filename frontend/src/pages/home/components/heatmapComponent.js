import { useEffect, useRef } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import 'leaflet.heat'

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
      'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
    },
    body: JSON.stringify({
      "options": {
        "include_address": true,
        "include_depot_address": true
      }
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

  // If token is invalid, take the user to log in page.
  if (response.status === 401) {
    alert('Expired or Invalid Token')
    localStorage.removeItem('token')
    window.location.reload()
  }

  const data = await response.json();
  console.log(data)
  var data2 = []
  for (let i = 0; i < data.items.length; i++) {
    
    if (data.items[i].depot_address !== null) {
      data2.push(data.items[i])
    }
  }
console.log(data2)
  let newData = data2.map((i) => {
    return [i.address.latitude, i.address.longitude, i.duration]; // Lat Lng intensity.
  })
  console.log(newData)
  return newData;
}

/**
 Renders a Leaflet Heatmap based on the data retrieved from the Bumbal API.
 @returns {JSX.Element} - The Leaflet Heatmap component.
 */
const Heatmap = (props) => {
  const heatRef = useRef()
  const pointsRef = useRef()
  const renderRef = useRef()
  const { value, intensity } = props;

  // Map context.
  const context = useLeafletContext()

  useEffect(() => {
    // Async function in order to wait for response from API.
    renderRef.current = 1;

    const fetchData = async () => {
      // Set address points.
      let addressPoints = await findAddressesPoints();

      // Map those points to something interpretable for the heat map.
      const points = addressPoints
        ? addressPoints.map((p) => {
          // If activity time is selected.
          if (value === 1) {
            return [p[0], p[1], p[2]];
          }
          // If activity location is selected.
          return [p[0], p[1], intensity];
        })
        : [];

      pointsRef.current = points
      heatRef.current = new L.heatLayer(points)

      // Create new layer and add it to the map context.
      context.layerContainer.addLayer(heatRef.current)
    };

    fetchData();

    return () => {
      context.layerContainer.removeLayer(heatRef.current)
    }
  }, [context.layerContainer, intensity, value])

  useEffect(() => {
    if (renderRef.current === 1) {
      console.log('firstime?')
      renderRef.current += renderRef.current;
    } else {
      const fetchData = async () => {
        // Delete old heat layer if it exists.
        context.layerContainer.eachLayer(function (layer) {
          context.layerContainer.removeLayer(layer)
        })

        // Set address points.
        let addressPoints = await findAddressesPoints();

        // Map those points to something interpretable for the heat map.
        const points = addressPoints
          ? addressPoints.map((p) => {
            // If activity time is selected.
            if (value === 1) {
              return [p[0], p[1], p[2]];
            }
            // If activity location is selected.
            return [p[0], p[1], intensity];
          })
          : [];

        pointsRef.current = points
        heatRef.current = new L.heatLayer(points)

        // Create new layer and add it to the map context.
        context.layerContainer.addLayer(heatRef.current)
      };

      fetchData();
    }
  }, [context.layerContainer, value, intensity])

  return null
}

export default Heatmap
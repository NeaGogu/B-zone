import { useEffect } from 'react'
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

  // if token is invalid take the user to log in page
  if (response.status === 401) {
    alert('Expired or Invalid Token')
    localStorage.removeItem('token')
    window.location.reload()
  }

  const data = await response.json();

  console.log(data)
  let newData = data.items.map((i) => {
    return [i.address.latitude, i.address.longitude, i.duration]; // lat lng intensity
  })
  console.log(newData)
  return newData;
}
  
/**
Renders a Leaflet Heatmap based on the data retrieved from the Bumbal API.
@returns {JSX.Element} - The Leaflet Heatmap component.
*/

// Hatmap component
const Heatmap = ({ value, intensity }) => {
  // map context
  const context = useLeafletContext()

  useEffect(() => {
    // async function in order to wait for response from api
    const fetchData = async () => {
      // delete old heat layer if it exists
      context.layerContainer.eachLayer(function (layer) {
        if (layer._heat != null) {
          context.layerContainer.removeLayer(layer)
        }
      })

      // set address points
      let addressPoints = await findAddressesPoints();
      // map those points to something interpretable for the heatmap
      const points = addressPoints
        ? addressPoints.map((p) => {
          // if activity time is selected
          if (value === 1) {
            return [p[0], p[1], p[2]];
          }
          // if location is selected
          return [p[0], p[1], intensity];
        })
        : [];

      // create new layer and add it to the map context
      context.layerContainer.addLayer(L.heatLayer(points))
    };

    fetchData();

  })

}

export default Heatmap
import { useEffect, useRef } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import 'leaflet.heat'
import getActivities from '../functions/getActivities'



/**
 Finds the latitude and longitude of each activity address and returns the data as an array.
 @returns {Promise<Array>} - The array containing latitude, longitude, and intensity for each address.
 */
async function findAddressesPoints() {
  const token = localStorage.getItem('token')

  // first see how many activities there are
  let  activities =  await fetch('https://sep202302.bumbal.eu/api/v2/activity', {
    method: 'PUT',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
    },
    body: JSON.stringify({
      "count_only": true
    })
  }).then((response)=>{return response.json()})

  let datatop = []
  // for loop and get activities based on offset
  for (let i = 0; i < activities.count_filtered;i = i + 100) {
      const response = await getActivities(i);

    // If token is invalid, take the user to log in page.
    if (response.status === 401) {
      alert('Expired or Invalid Token')
      localStorage.removeItem('token')
      window.location.reload()
    }

    const data = await response.json();
    datatop.push(...data.items)
  }
  console.log(activities.count_filtered)
  

  
  var data2 = []
  for (let i = 0; i < datatop.length; i++) {
    
    if (datatop[i].depot_address !== null) {
      console.log(datatop[i])
      data2.push(datatop[i])
    }
  }
  //console.log(data2)
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
  const { value, intensity, setComputed } = props;

  // Map context.
  const context = useLeafletContext()


  useEffect(() => {
      const fetchData = async () => {
        // Delete old heat layer if it exists.
        context.layerContainer.eachLayer(function (layer) {
          context.layerContainer.removeLayer(layer)
        })

        // Set address points.
        setComputed(false)
        let addressPoints = await findAddressesPoints();
        setComputed(false)
        
        // Map those points to something interpretable for the heat map.
        //console.log(addressPoints, 'hello')
        const points = addressPoints
          ? addressPoints.map((p) => {
            setComputed(false)
            // If activity time is selected.
            if (value === 1) {
              return [p[0], p[1], p[2]];
            }
            // If activity location is selected.
            return [p[0], p[1], intensity];
          })
          : [];

        pointsRef.current = points
        setComputed(false)
        heatRef.current = new L.heatLayer(points)
        setComputed(false)

        // Create new layer and add it to the map context.
        context.layerContainer.addLayer(heatRef.current)
        setComputed(true)
      };
      setComputed(false)
      fetchData();
    
  // warning not usefull
  // eslint-disable-next-line
  }, [context.layerContainer, value, intensity])

  return null
}

export default Heatmap
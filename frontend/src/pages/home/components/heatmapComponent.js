import { useEffect, useRef } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import 'leaflet.heat'
import getAllActivities from '../functions/getAllActivities'



/**
 Finds the latitude and longitude of each activity address and returns the data as an array.
 @returns {Promise<Array>} - The array containing latitude, longitude, and intensity for each address.
 */
async function findAddressesPoints() {

  const activities = await getAllActivities()

  var data2 = []
  for (let i = 0; i < activities.length; i++) {

    if (activities[i].depot_address !== null) {
      data2.push(activities[i])
    }
  }

  let newData = data2.map((i) => {
    return [i.address.latitude, i.address.longitude, i.duration]; // Lat Lng intensity.
  })

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
      let addressPoints = await findAddressesPoints();

      // Map those points to something interpretable for the heat map.
      //console.log(addressPoints, 'hello')
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
      setComputed(true)
    };
    setComputed(false)
    fetchData();

  }, [context.layerContainer, value, intensity, setComputed])

  return null
}

export default Heatmap
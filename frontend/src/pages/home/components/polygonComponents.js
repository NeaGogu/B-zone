    // External dependencies
import { useEffect } from 'react'
import L from 'leaflet'
import { useLeafletContext } from '@react-leaflet/core'
import getInitialZones from '../functions/getInitialZones'
import getZipCodes from '../functions/getZipCodes'
import getCoordinates from '../functions/getCoordinates'
import convertToStructure from '../functions/convertToStructure'
import calculateZone from '../functions/calculateZone'
import querryDatabase from '../functions/querryDatabase'

var zipCodes = []
// array of colors to display
const colors = ['#e6194b', '#3cb44b', '#ffe119', '#4363d8', '#f58231', '#911eb4', '#46f0f0', '#f032e6', 
                '#bcf60c', '#fabebe', '#008080', '#e6beff', '#9a6324', '#fffac8', '#800000', '#aaffc3', 
                '#808000', '#ffd8b1', '#000075', '#808080', '#ffffff', '#000000']


// Main function to visualize the polygons on the map.
const PolygonVis = (props) => {
    //selections
    const { zoneId, setZipCodes, setComputed, algorithm, nrofzones } = props

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

            // check if zone is to be calculated
            if (zoneId.startsWith('calculate')) {
                var calculation;
                if (algorithm === 1) {
                    console.log(nrofzones)
                    calculation = await calculateZone(algorithm, nrofzones)
                } else {
                    calculation = await calculateZone(algorithm, nrofzones)
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
            // set render state to be true
            setComputed(true)
        };
        fetchData()
    }, [context.layerContainer, setZipCodes, zoneId, setComputed])


    return null
}

export default PolygonVis
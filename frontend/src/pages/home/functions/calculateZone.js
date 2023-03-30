/** 
* Fetches the calculated zone configuration from the backend, returns promise of the response from the backend.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the user's calculated zone configuration.
*/
async function calculateZone(algorithm, nrofzones) {
    // set token
    const userToken = localStorage.getItem('token')

    // array to hold zone configuration
    let calculatedZones = []

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
                //console.log(data.result)
                return data.result
            })
            .catch(error => console.log(error, 'error'))
    } else {
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
                //console.log(data.result)
                return data.result
            })
            .catch(error => console.log(error, 'error'))

    }
    // cleaning up array
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
    return [zoneConfig, calculatedZones]
}

export default calculateZone

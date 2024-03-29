/** 
* Fetches saved plot from the backend, returns promise of the response from the backend.
* @param {string} plotID - Id of plot to be looked up.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the zone configuration of the plot.
*/
async function querryDatabase(plotID) {
    const userToken = localStorage.getItem('token')
    var zones = []
    await fetch("http://localhost:4000/plot/" + plotID, {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${userToken}`
        }
    }).then((response) => {
        if (!response.ok) {
            if (response.status === 401){
                alert("Unable to retrieve this zone configuration! " + response.status + ' Error')
                localStorage.removeItem('token')
                localStorage.removeItem('email')
                localStorage.removeItem('id')
                window.location.reload()
                return null
            } 
            alert("Unable to retrieve this zone configuration! " + response.status + ' Error')
        }
        return response.json()
    }).then((data) => {
        zones = data.plot_zones
    })

    var zoneConfig = []

    // Go into each zone.
    for (let i = 0; i < zones.length; i++) {
        var zoneRanges = zones[i].zone_ranges
        var currZoneRange = []
        // Go into each range.
        for (let j = 0; j < zoneRanges.length; j++) {
            // Convert to format used in other GetCoordinates.
            var curr = {
                zipFrom: zoneRanges[j].zipcode_from,
                zipTo: zoneRanges[j].zipcode_to
            }
            currZoneRange.push(curr)
        }
        zoneConfig.push(currZoneRange)
    }
    return zoneConfig
}

export default querryDatabase
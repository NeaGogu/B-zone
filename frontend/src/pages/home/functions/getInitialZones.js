/** 
* Fetches the initial zone configuration a user has from Bumbal, returns promise of the response from Bumbal API.
* @param {string} userToken - The user token retrieved from local storage.
* @return {Promise<Array>} A promise that resolves with an array of objects representing the user's initial zone configuration.
*/
async function getInitialZones() {
    // Gets list of zones.
    // Definition of URl, body values, and userToken.
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
            return data.items
        })

        .catch(error => console.log(error, 'error'))

    return initialZones
}
export default getInitialZones
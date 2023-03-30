/** 
* Gets a list of the coordinates given a list of zip codes.
* @param {Array} zoneList - The list of zones whose zip codes are to be fetched.
* @return {Promise<Array>} Returns a promise that resolves to an array of objects, where each object represents a zone's list of coordinates.
*/
async function getCoordinates(zipsList) {
    let coordinatesList = []
    // For each zone area, fetch the coordinates and compile them together.
    for (let i = 0; i < zipsList.length; i++) {
        // Same idea as in getzipcodes().
        var struct = []

        for (let j = 0; j < zipsList[i].length; j++) {
            struct.push(await fetch('http://localhost:4000/test/zip/coordinates?zip_from=' + zipsList[i][j].zipFrom.toString() + '&zip_to=' + zipsList[i][j].zipTo.toString())
                .then((response) => {
                    if (!response.ok) {
                        console.log("Response from our backend is not ok ???")
                    }
                    return response.json()
                })
                .then((data) => {
                    return data
                })

                .catch(error => console.log(error)))
        }
        coordinatesList.push(struct)
    }
    return coordinatesList
}
export default getCoordinates
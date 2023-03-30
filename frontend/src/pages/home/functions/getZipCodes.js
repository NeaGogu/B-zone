/** 
* With a list of zones, fetches their zip codes and updates their polygon coordinates to coordinatesFile.
* @param {Array} zoneList - The list of zones whose zip codes are to be fetched.
* @return {Promise<Array>} Returns a promise that resolves to an array of objects, where each object represents a zone's list of zip codes.
*/
async function getZipCodes(zoneList) {
    // When you retrieve a list of zones from Bumbal API from zone with PUT, you retrieve a list of zone configurations, 
    // which itself includes a list of zones in each zone configuration.
    let zipCodes = []

    for (let i = 0; i < zoneList.length; i++) {
        // In order for the structure to be [[zipcode ranges zone1], [zipcode ranges zone2], etc...].
        var zone = []

        for (let j = 0; j < zoneList[i].zone_ranges.length; j++) {
            var curr = {
                zipFrom: zoneList[i].zone_ranges[j].zipcode_from,
                zipTo: zoneList[i].zone_ranges[j].zipcode_to
            }
            zone.push(curr)
        }
        zipCodes.push(zone)
    }
    return zipCodes
}
export default getZipCodes
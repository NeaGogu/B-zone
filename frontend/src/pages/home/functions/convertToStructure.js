/** 
* Gets a list of the coordinates given a list of zip codes.
* @param {Array} plot - Zone configuration which is to be converted.
* @return {Promise<Array>} Returns a zone configuration in the structed expected by the backend.
*/
function convertToStructure(plot) {
    // variable to hold the plot
    var plt = []

    for (let i = 0; i < plot.length; i++) {
        // variable for to the the zone ranges
        var zone_ranges = []
        for (let j = 0; j < plot[i].length; j++) {
            // variable to hold one zone range
            var curr = {
                zipcode_from: plot[i][j].zipFrom,
                zipcode_to: plot[i][j].zipTo
            }
            // add to zone ranges
            zone_ranges.push(curr)
        }
        // variable to hold the zones ranges with appropriate structure
        var zoneobj = {
            "zone_ranges": zone_ranges
        }
        // add the zone range to the plot
        plt.push(zoneobj)
    }
    // variable to hold the plot in the required structure
    var pltfinal = {
        "plot_zones": plt
    }

    return pltfinal
}
export default convertToStructure
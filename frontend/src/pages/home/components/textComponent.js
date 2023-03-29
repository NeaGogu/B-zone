//import 'leaflet.heat'

import {useEffect, useState} from "react";

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
            console.log("Response from our backend is not ok ???")
        }
        return response.json()
    }).then((data) => {
        zones = data.plot_zones
    })

    var zoneConfig = []

    // go into each zone
    for (let i = 0; i < zones.length; i++) {
        var zoneRanges = zones[i].zone_ranges
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
    return zoneConfig
}

async function getActivities() {
    const token = localStorage.getItem('token')
    console.log('token ' + token)

    const requestOptions = {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
        },
        body: JSON.stringify({
            "options": {
                "include_address": true,
                "include_depot_address": true
            }
        })
    };

    const response = await fetch('https://sep202302.bumbal.eu/api/v2/activity', requestOptions);
    console.log('getActivities response ' + response.status)
    return response;
}

/**
 Finds the latitude and longitude of each activity address and returns the data as an array.
 @returns {int} - The array containing latitude, longitude, and intensity for each address.
 */
async function totalActivityDurations(settime) {
    const response = await getActivities();

    // If token is invalid, take the user to log in page.
    if (response.status === 401) {
        alert('Expired or Invalid Token')
        localStorage.removeItem('token')
        window.location.reload()
    }

    const data = await response.json();
    console.log(data)
    var data2 = []
    var time = 0;
    for (let i = 0; i < data.items.length; i++) {

        if (data.items[i].depot_address !== null) {
            data2.push(data.items[i])
        }
    }

    for(let i = 0; i< data2.length; i++) {
        time += parseInt(data2[i].duration);
    }
    settime(time);
}

async function activityZoneAllocation(plotID) {
    const db = await querryDatabase(plotID)
    return db
}

function TextComponent(props) {
    const { zoneId, zoneName } = props;

    const [time, setTime] = useState(0)


    useEffect(() =>{
        const initial = async () => {
            totalActivityDurations(setTime)
            const plot = await activityZoneAllocation(zoneId)


            let listOfActivities = await getActivities()
            listOfActivities = await listOfActivities.json() //get json data from response

            // get activities and related zipcode.
            let activites2 = listOfActivities.items.map((i) => {
                return [i.address.latitude, i.address.longitude, i.address.zipcode, -1]; // Lat Lng intensity.
            })
            // get zone model.
            //console.log(activites2)
            // go throught each activity and find matching zone
            for (let i = 0; i < activites2.length; i++){
                let toparr = []
                let zipcode = parseInt(activites2[i][2].slice(0,4))

                for (let j =0; j < plot.length; j++) {
                    let arr = []
                    for (let k = 0; k < plot[j].length; k++) {
                        if (zipcode >= plot[j][k].zipFrom && zipcode <= plot[j][k].zipTo) {

                            activites2[i][3] = j;
                        }
                        // else{
                        //     if(zipcode === 5038){
                        //         console.log(plot[j][k].zipFrom, plot[j][k].zipTo, zipcode )
                        //     }
                        //
                        // }
                    }
                }
            }

            let activities2Filtered = []
            for(let i = 0; i < activites2.length; i++) {
                if(activites2[i][3] !== -1) {
                    activities2Filtered.push(activites2[i])
                }
            }

            console.log(activities2Filtered)


        }
        initial()
    },[zoneId])


    //1. get list of activities from bumbal
    //CASE: WE HAVE ID
    //2. get all zones from bzone backend
    //3. for each activity , see which zone it fits in via zipcode?
    //




    return (
        <div style={{  position: "absolute", color: 'white'}}>
            <p>Title: {zoneName}</p>
            {/*fuel cost = fuel input times driving time*/}
            <p>Total cost:  </p>
            {/* done */}
            <p>Total activity time (hrs): {time}</p>
            {/*driving time = find activities per zone. find driving time in order between those activities using OSRM */}
            <p>Total driving time: </p>
        </div>
    );
}
export default TextComponent
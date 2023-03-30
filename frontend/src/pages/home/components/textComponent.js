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
    return response.json();
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

    const data = await response;
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

async function getDrivingTime(drivingData) {
    //input: a fetch response from OSRM with multiple legs per driving route
    //output: the total sum of the duration over all the legs of the route
    let drivingLegs = drivingData.routes[0].legs
    let sum = 0
    console.log(drivingLegs)
    for(let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].duration
    }
    console.log(sum)
    return sum;
}

function TextComponent(props) {
    const { zoneId, zoneName } = props;

    const [time, setTime] = useState(0)
    const [drivingTime, setDrivingTime] = useState(0)


    useEffect(() =>{
        const initial = async () => {
            totalActivityDurations(setTime) //set the total activity duration to be the time spend on activities
            const plot = await activityZoneAllocation(zoneId) //gets all plots which are saved to b-zone's backend


            let listOfActivities = await getActivities() //gets all activity locations from Bumbal
            //listOfActivities = await listOfActivities.json() //get json data from activity response

            // get activities and related zipcode + add a blank zone field with -1 id
            let activites2 = listOfActivities.items.map((i) => {
                return [i.address.latitude, i.address.longitude, i.address.zipcode, -1]; // Lat Lng intensity.
            })
            // get zone model.
            //console.log(activites2)
            // go throught each activity and find matching zone
            for (let i = 0; i < activites2.length; i++){
                //get the numerical part of a zipcode for each activity
                let zipcode = parseInt(activites2[i][2].slice(0,4))
                for (let j =0; j < plot.length; j++) {
                    //for each possible zone, check all zone ranges
                    for (let k = 0; k < plot[j].length; k++) {
                        //for each zone range, check whether the zipcode of the activity fits into that zone range
                        if (zipcode >= plot[j][k].zipFrom && zipcode <= plot[j][k].zipTo) {

                            activites2[i][3] = j;
                        }
                    }
                }
            }
            //filter out all activities which are depot activities, only driving time activities remain
            let activities2Filtered = []
            for(let i = 0; i < activites2.length; i++) {
                if(activites2[i][3] !== -1) {
                    activities2Filtered.push(activites2[i])
                }
            }
            console.log(activities2Filtered)
            //array to hold driving time per zone
            let drivingTimeActivities = []
            //array to hold fetch requests HTML per zone
            let drivingTimeReqs = []
            //body of the fetch request to be sent out to OSRM
            const drivingTimeBody = {
                method: 'GET'
            };
            //go through all zones to see what the driving time of activities in that zone is
            for(let i = 0; i < plot.length; i++) {
                //go through all activities to find which ones belong to zone i, calculate their driving time
                drivingTimeActivities[i] = 0
                drivingTimeReqs[i] = "http://router.project-osrm.org/route/v1/driving/"
                for(let j = 0; j < activities2Filtered.length; j++) {
                    if(activities2Filtered[j][3] === i) {
                        //update the request for zone i to contain the coordinates of all activities
                        drivingTimeReqs[i] = drivingTimeReqs[i] + activities2Filtered[j][1] + ',' + activities2Filtered[j][0] + ';'

                    }
                }
                drivingTimeReqs[i] = drivingTimeReqs[i].slice(0, -1);
                //now that the request string for zone i is built, send out fetch request
                const response = await fetch(drivingTimeReqs[i], drivingTimeBody);
                const drivingData = await response.json();
                console.log(drivingData)
                //using the data from OSRM, compile over all legs of the activities what the total duration is and store it for zone i
                drivingTimeActivities[i] = await getDrivingTime(drivingData)
            }
            console.log(drivingTimeReqs[1])
            console.log(drivingTimeActivities)

            let totalDrivingTime = 0
            for(let i = 0; i < drivingTimeActivities.length; i++) {
                totalDrivingTime = totalDrivingTime + drivingTimeActivities[i]
            }
            setDrivingTime(totalDrivingTime/3600)

        }

        initial()
        console.log('updated')
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
            <p>Total driving time: {drivingTime}</p>
        </div>
    );
}
export default TextComponent
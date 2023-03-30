//import 'leaflet.heat'
import { Card, Spin } from 'antd'
import { useEffect, useState } from "react";
import querryDatabase from "../functions/querryDatabase";
import getAllActivities from "../functions/getAllActivities";

/**
 Finds the latitude and longitude of each activity address and returns the data as an array.
 @returns {int} - The array containing latitude, longitude, and intensity for each address.
 */
async function totalActivityDurations(settime) {
    const activities = await getAllActivities();



    console.log(activities)
    var data2 = []
    var time = 0;
    for (let i = 0; i < activities.length; i++) {

        if (activities[i].depot_address !== null) {
            data2.push(activities[i])
        }
    }

    for (let i = 0; i < data2.length; i++) {
        time += parseInt(data2[i].duration);
    }
    settime(time);
}

async function getDrivingTime(drivingData) {
    //input: a fetch response from OSRM with multiple legs per driving route
    //output: the total sum of the duration over all the legs of the route
    let drivingLegs = drivingData.routes[0].legs
    let sum = 0
    console.log(drivingLegs)
    for (let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].duration
    }
    console.log(sum)
    return sum;
}

async function getDrivingDistance(drivingData) {
    //input: a fetch response from OSRM with multiple legs per driving route
    //output: the total sum of the duration over all the legs of the route
    let drivingLegs = drivingData.routes[0].legs
    let sum = 0
    console.log(drivingLegs)
    for (let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].distance
    }
    console.log(sum)
    return sum;
}

function TextComponent(props) {
    const { zoneId, zoneName } = props;

    const [loaded, setLoaded] = useState(false)
    const [time, setTime] = useState(0)
    const [drivingTime, setDrivingTime] = useState(0)
    const [fuelCost, setFuelCost] = useState(0)


    useEffect(() => {
        const initial = async () => {
            setLoaded(false)
            let averageFuelCost = 1.8
            let averageFuelConsumption = 0.047 //litres of fuel consumption per km
            totalActivityDurations(setTime) //set the total activity duration to be the time spend on activities
            const plot = await querryDatabase(zoneId) //gets all plots which are saved to b-zone's backend


            let listOfActivities = await getAllActivities() //gets all activity locations from Bumbal

            //listOfActivities = await listOfActivities.json() //get json data from activity response

            // get activities and related zipcode + add a blank zone field with -1 id
            let activites2 = listOfActivities.map((i) => {
                return [i.address.latitude, i.address.longitude, i.address.zipcode, -1]; // Lat Lng intensity.
            })
            // get zone model.
            //console.log(activites2)
            // go throught each activity and find matching zone
            for (let i = 0; i < activites2.length; i++) {
                //get the numerical part of a zipcode for each activity
                let zipcode = parseInt(activites2[i][2].slice(0, 4))
                for (let j = 0; j < plot.length; j++) {
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
            for (let i = 0; i < activites2.length; i++) {
                if (activites2[i][3] !== -1) {
                    activities2Filtered.push(activites2[i])
                }
            }
            console.log(activities2Filtered)
            //array to hold driving time per zone
            let drivingTimeActivities = []
            //array to hold driving distances per zone
            let drivingDistanceActivities = []
            //array to hold fetch requests HTML per zone
            let drivingTimeReqs = []
            //body of the fetch request to be sent out to OSRM
            const drivingTimeBody = {
                method: 'GET'
            };
            //go through all zones to see what the driving time of activities in that zone is
            for (let i = 0; i < plot.length; i++) {
                //go through all activities to find which ones belong to zone i, calculate their driving time
                drivingTimeActivities[i] = 0
                drivingDistanceActivities[i] = 0
                drivingTimeReqs[i] = "http://router.project-osrm.org/route/v1/driving/"
                for (let j = 0; j < activities2Filtered.length; j++) {
                    if (activities2Filtered[j][3] === i) {
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
                drivingDistanceActivities[i] = await getDrivingDistance(drivingData)
            }
            console.log(drivingTimeReqs[1])
            console.log(drivingTimeActivities)
            console.log(drivingDistanceActivities)

            let totalDrivingTime = 0
            let totalDrivingDistance = 0
            for (let i = 0; i < drivingTimeActivities.length; i++) {
                totalDrivingTime = totalDrivingTime + drivingTimeActivities[i]
                totalDrivingDistance = totalDrivingDistance + drivingDistanceActivities[i]
            }
            setDrivingTime(totalDrivingTime / 3600)
            //time to find fuel cost: fuel cost = (litres used * fuel cost)
            //litres used = driving distance * fuel efficiency
            setFuelCost(((totalDrivingDistance / 1000) * averageFuelConsumption) * averageFuelCost)
            setLoaded(true)

        }

        initial()
        console.log('updated')
    }, [zoneId])


    //1. get list of activities from bumbal
    //CASE: WE HAVE ID
    //2. get all zones from bzone backend
    //3. for each activity , see which zone it fits in via zipcode?
    //




    return (
        <Card title={zoneName} style={{ width: 'fit-content', marginLeft: 'auto', marginRight: 'auto' }} >
            <Spin spinning={!loaded} delay={200} tip='Calculating...'>
                <div >
                    {/*fuel cost = fuel input times driving time*/}
                    <p>Total cost: {fuelCost}</p>
                    {/* done */}
                    <p>Total activity time (hrs): {time}</p>
                    {/*driving time = find activities per zone. find driving time in order between those activities using OSRM */}
                    <p>Total driving time: {drivingTime}</p>
                </div>
            </Spin>
        </Card>
    );
}
export default TextComponent
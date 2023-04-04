import { Card, Spin } from 'antd'
import { useEffect, useState } from "react";
import querryDatabase from "../functions/querryDatabase";
import getAllActivities from "../functions/getAllActivities";
import { Collapse } from 'antd';

const { Panel } = Collapse;

/**
 Finds the latitude and longitude of each activity address and returns the data as an array.
 @returns {int} - The array containing latitude, longitude, and intensity for each address.
 */
async function totalActivityDurations(settime) {
    const activities = await getAllActivities();

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
    // Input: a fetch response from OSRM with multiple legs per driving route
    // Output: the total sum of the duration over all the legs of the route
    let drivingLegs = drivingData.routes[0].legs
    let sum = 0

    for (let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].duration
    }
    return sum;
}

async function getDrivingDistance(drivingData) {
    // Input: a fetch response from OSRM with multiple legs per driving route
    // Output: the total sum of the distance over all the legs of the route
    let drivingLegs = drivingData.routes[0].legs
    let sum = 0

    for (let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].distance
    }
    return sum;
}

function TextComponent(props) {
    const { zoneId, zoneName, calculatedZone, averageFuelCost, averageFuelUsage } = props;

    const [name, setName] = useState('')
    const [loaded, setLoaded] = useState(false)
    const [time, setTime] = useState(0)
    const [drivingTime, setDrivingTime] = useState(0)
    const [fuelCost, setFuelCost] = useState(0)
    const [drivingTimeActiv, setDrivingTimeActiv] = useState([])
    const [drivingDistanceActiv, setDrivingDistanceActiv] = useState([])
    let averageFuelConsumption = averageFuelUsage

    useEffect(() => {
        const initial = async () => {
            setLoaded(false)
            let averageFuelCost = 1.8 // Cost in euros per litre of fuel
            let averageFuelConsumption = 0.047 // Litres of fuel consumption per km
            totalActivityDurations(setTime) // Set the total activity duration to be the time spend on activities
            let plot = []; // Variable to hold the object of zones/zone configuration upon which to calculate driving time

            if (zoneId.startsWith('calculate')) { // The plot is a freshly calculated optimized plot
                for (let i = 0; i < calculatedZone.plot_zones.length; i++) {
                    plot[i] = calculatedZone.plot_zones[i].zone_ranges
                    for (let j = 0; j < plot[i].length; j++) {
                        plot[i][j] = { zipFrom: plot[i][j].zipcode_from, zipTo: plot[i][j].zipcode_to }
                    }
                }
                setName('Calculated Zone')
            }

            else {
                if (zoneId.startsWith('initial')) { //the plot is the initial plot from Bumbal
                    setName(zoneName)
                    for (let i = 0; i < calculatedZone.length; i++) {
                        for (let j = 0; j < calculatedZone[i].length; j++) {
                            calculatedZone[i][j].zipFrom = parseInt(calculatedZone[i][j].zipFrom);
                            calculatedZone[i][j].zipTo = parseInt(calculatedZone[i][j].zipTo);
                        }
                    }
                    plot = calculatedZone // Retrieve the current plot and work with it
                } else { // The plot is a saved plot from the backend
                    plot = await querryDatabase(zoneId)
                    setName(zoneName)
                }
            }

            console.log("The plot currently working with is: ")
            console.log(plot)

            let listOfActivities = await getAllActivities() // Get all activity locations from Bumbal

            let activites2 = listOfActivities.map((i) => { // Get activities and related zipcode + add a blank zone field with -1 id
                return [i.address.latitude, i.address.longitude, i.address.zipcode, -1, i.duration]; // Lat Lng intensity.
            })

            // Go through each activity and find matching zone
            for (let i = 0; i < activites2.length; i++) {
                // Get the numerical part of a zipcode for each activity
                let zipcode = parseInt(activites2[i][2].slice(0, 4))
                for (let j = 0; j < plot.length; j++) {
                    // For each possible zone, check all zone ranges
                    for (let k = 0; k < plot[j].length; k++) {
                        // For each zone range, check whether the zipcode of the activity fits into that zone range
                        if (zipcode >= plot[j][k].zipFrom && zipcode <= plot[j][k].zipTo) {
                            activites2[i][3] = j;
                        }
                    }
                }
            }
            // Filter out all activities which are depot activities, only driving time activities remain
            let activities2Filtered = []
            for (let i = 0; i < activites2.length; i++) {
                if (activites2[i][3] !== -1) {
                    activities2Filtered.push(activites2[i])
                }
            }

            // Array to hold driving time per zone
            let drivingTimeActivities = new Array(plot.length);
            // Array to hold driving distances per zone
            let drivingDistanceActivities = new Array(plot.length);
            // Array to hold fetch requests HTML per zone
            let drivingTimeReqs = new Array(plot.length);
            //Driving time per zone, index i = zone i
            let activityTimeZone = []
            // Body of the fetch request to be sent out to OSRM
            const drivingTimeBody = {
                method: 'GET'
            };
            // Introduce activity counter to see how many activities there are per zone
            let countActivities = 0
            // Go through all zones to see what the driving time of activities in that zone is
            for (let i = 0; i < plot.length; i++) {
                // Fo through all activities to find which ones belong to zone i, calculate their driving time
                countActivities = 0
                drivingTimeActivities[i] = 0
                drivingDistanceActivities[i] = 0
                activityTimeZone[i] = 0;
                drivingTimeReqs[i] = "http://router.project-osrm.org/route/v1/driving/"
                for (let j = 0; j < activities2Filtered.length; j++) {
                    if (activities2Filtered[j][3] === i) {
                        // Update the request for zone i to contain the coordinates of all activities
                        drivingTimeReqs[i] = drivingTimeReqs[i] + activities2Filtered[j][1] + ',' + activities2Filtered[j][0] + ';'
                        countActivities = countActivities + 1;
                        activityTimeZone[i] = activityTimeZone[i] + parseInt(activities2Filtered[j][4])

                    }
                }
                
                // Remove last ';' from the request string for driving time in zone i
                drivingTimeReqs[i] = drivingTimeReqs[i].slice(0, -1);
                // Check whether there are less than 2 activities in this zone
                if (countActivities < 2) {
                    drivingTimeActivities[i] = 0 // Minimal driving time needed in this zone
                    drivingDistanceActivities[i] = 0 //Minimal distance to drive in this zone
                } else {
                    // Now that the request string for zone i is built, send out fetch request
                    const response = await fetch(drivingTimeReqs[i], drivingTimeBody);
                    const drivingData = await response.json();
                    // Using the data from OSRM, compile over all legs of the activities what the total duration is and store it for zone i
                    drivingTimeActivities[i] = await getDrivingTime(drivingData)
                    drivingDistanceActivities[i] = await getDrivingDistance(drivingData)
                }
            }
            //FOR TANIA
            console.log(activityTimeZone)

            setDrivingTimeActiv(prevDrivingTimeActiv => []);
            setDrivingDistanceActiv(prevDrivingDistanceActiv => []);

            let totalDrivingTime = 0
            let totalDrivingDistance = 0

            for (let i = 0; i < drivingTimeActivities.length; i++) {
                totalDrivingTime = totalDrivingTime + drivingTimeActivities[i]
                setDrivingTimeActiv(prevDrivingTimeActiv => [...prevDrivingTimeActiv, drivingTimeActivities[i]]);
                totalDrivingDistance = totalDrivingDistance + drivingDistanceActivities[i]
                setDrivingDistanceActiv(prevDrivingDistanceActiv => [...prevDrivingDistanceActiv, drivingDistanceActivities[i]]);
            }
            setDrivingTime(totalDrivingTime.toPrecision(3) / 3600)
            // Time to find fuel cost: fuel cost = (litres used * fuel cost)
            // Litres used = driving distance * fuel efficiency
            setFuelCost(((totalDrivingDistance.toPrecision(4) / 1000) * averageFuelConsumption) * averageFuelCost)
            setLoaded(true)
        }

        initial()
        console.log('updated')
    }, [zoneId, zoneName, calculatedZone])

    return (
        <Card title={name} style={{  marginLeft: 'auto', marginRight: 'auto', flex: '1' }} >
            <Spin spinning={!loaded} delay={200} tip='Calculating...'>
                <div >
                    {/*fuel cost = fuel input times driving time*/}
                    {/*<p>Total cost: ${fuelCost}</p>*/}
                    <Collapse >
                        <Panel header={`Total cost: \u20AC${fuelCost.toPrecision(3)}`} key="1">
                            <ul>
                                <li>Distance over the zones: </li>
                                {drivingDistanceActiv.map((drivingDistance, index) => (
                                    <p key={index}>Zone {index}: {drivingDistance.toFixed(2)} meters</p>
                                ))}
                                <li>Total driving distance: </li>
                                <p>
                                    {drivingDistanceActiv.map((drivingDistance, index) => (
                                        <span key={index}>{drivingDistance.toFixed(2)} {index < drivingDistanceActiv.length - 1 && '+'} </span>
                                    ))}
                                    = {drivingDistanceActiv.reduce((acc, time) => acc + time, 0).toFixed(2)} meters
                                </p>
                                <li>Total cost with fuel consumption = {averageFuelConsumption} euros and fuel cost = {averageFuelCost} euros:</li>
                                <p>(({drivingDistanceActiv.reduce((acc, time) => acc + time, 0).toFixed(2)} / 1000) * {averageFuelConsumption}) * {averageFuelCost} = {fuelCost.toFixed(2)} euros</p>
                            </ul>
                        </Panel>
                    </Collapse>
                    <p> </p>
                    <Collapse>
                        <Panel key={3} header={`Total activity time in hours: ${time}`}>
                            <li>This is computed by adding all of the activity times </li>
                        </Panel>
                    </Collapse>
                    <p> </p>

                    {/*driving time = find activities per zone. find driving time in order between those activities using OSRM */}
                    {/*<p>Total driving time: {drivingTime}</p>*/}
                    <Collapse>
                        <Panel key={2} header={`Total driving time: ${Math.floor(drivingTime)} hours and ${Math.round((drivingTime % 1) * 60)} minutes`}>
                            <ul>
                                <li>Driving time over the zones:</li>
                                {drivingTimeActiv.map((drivingTime, index) => {
                                    const hours = Math.floor(drivingTime / 3600);
                                    const minutes = Math.floor((drivingTime % 3600) / 60);
                                    const seconds = Math.floor(drivingTime % 60);
                                    return (
                                        <p key={index}>
                                            Zone {index}: {hours > 0 && `${hours} hour${hours > 1 ? 's' : ''} `}{minutes > 0 && `${minutes} minute${minutes > 1 ? 's' : ''} `}{seconds} second{seconds !== 1 ? 's' : ''}
                                        </p>
                                    );
                                })}

                                <li>Driving time sum:</li>
                                <p>
                                    {drivingTimeActiv.map((drivingTime, index) => {
                                        const hours = Math.floor(drivingTime / 3600);
                                        const minutes = Math.floor((drivingTime % 3600) / 60);
                                        const seconds = Math.floor(drivingTime % 60);
                                        return (
                                            <span key={index}>
                                                {hours > 0 && `${hours} hour${hours > 1 ? 's' : ''} `}
                                                {minutes > 0 && `${minutes} minute${minutes > 1 ? 's' : ''} `}
                                                {seconds} second{seconds !== 1 ? 's' : ''}
                                                {index < drivingTimeActiv.length - 1 && ' + '}
                                            </span>
                                        );
                                    })}
                                </p>

                                <p>Total driving time: {Math.floor(drivingTime)} hours and {Math.round((drivingTime % 1) * 60)} minutes</p>
                            </ul>
                        </Panel>
                    </Collapse>
                </div>
            </Spin>
        </Card>
    );
}

export default TextComponent
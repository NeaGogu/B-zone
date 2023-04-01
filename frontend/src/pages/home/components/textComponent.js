//import 'leaflet.heat'
import {Card, Dropdown, Spin, Tooltip} from 'antd'
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



    //console.log(activities)
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
    //console.log(drivingLegs)
    for (let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].duration
    }
    //console.log(sum)
    return sum;
}

async function getDrivingDistance(drivingData) {
    //input: a fetch response from OSRM with multiple legs per driving route
    //output: the total sum of the duration over all the legs of the route
    let drivingLegs = drivingData.routes[0].legs
    let sum = 0
    //console.log(drivingLegs)
    for (let i = 0; i < drivingLegs.length; i++) {
        sum = sum + drivingLegs[i].distance
    }
    //console.log(sum)
    return sum;
}

function TextComponent(props) {
    const { zoneId, zoneName, calculatedZone, setAverageFuelUsage,
        averageFuelCost,averageFuelUsage, setAverageFuelCost  } = props;

    const [name, setName] = useState('')
    const [loaded, setLoaded] = useState(false)
    const [time, setTime] = useState(0)
    const [drivingTime, setDrivingTime] = useState(0)
    const [fuelCost, setFuelCost] = useState(0)
    const [drivingTimeActiv, setDrivingTimeActiv] = useState([])
    const [drivingDistanceActiv, setDrivingDistanceActiv] = useState([])
    let averageCost = 1.8
    let averageFuelConsumption = averageFuelUsage
    let averageFuel = 0.047 //litres of fuel consumption per km



    useEffect(() => {
        const initial = async () => {
            setLoaded(false)
            let averageFuelCost = 1.8 //cost in euros per litre of fuel
            let averageFuelConsumption = 0.047 //litres of fuel consumption per km
            totalActivityDurations(setTime) //set the total activity duration to be the time spend on activities
            let plot = []; //variable to hold the object of zones/zone configuration upon which to calculate driving time
            if (zoneId.startsWith('calculate')) { //the plot is a freshly calculated optimized plot
                //plot = calculatedZone.plot_zones[1].zone_ranges //change this object model so that it can be used in the calculations below
                 for (let i = 0; i < calculatedZone.plot_zones.length; i++) {
                     plot[i] = calculatedZone.plot_zones[i].zone_ranges
                     for(let j = 0; j < plot[i].length; j ++) {
                         plot[i][j] = {zipFrom: plot[i][j].zipcode_from, zipTo: plot[i][j].zipcode_to}
                     }

                 }
                setName('Calculated Zone')
                //console.log(plot)
            }
            else {
                if(zoneId.startsWith('initial')) { //the plot is the initial plot from Bumbal
                    setName('Initial Zones')
                    plot = await querryDatabase(zoneId) //retrieve the plot object from Bumbal and change it so that it can be used in the calculations below
                } else { //the plot is a saved plot from our backend
                    plot = await querryDatabase(zoneId)
                    setName(zoneName)
                }

            }
            console.log("The plot currently working with is: ")
            console.log(plot)

            let listOfActivities = await getAllActivities() //gets all activity locations from Bumbal

            let activites2 = listOfActivities.map((i) => { // get activities and related zipcode + add a blank zone field with -1 id
                return [i.address.latitude, i.address.longitude, i.address.zipcode, -1]; // Lat Lng intensity.
            })
            // get zone model

            // go through each activity and find matching zone
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
            //console.log(activities2Filtered)
            //array to hold driving time per zone
            let drivingTimeActivities = new Array(plot.length);
            //array to hold driving distances per zone
            let drivingDistanceActivities = new Array(plot.length);
            //array to hold fetch requests HTML per zone
            let drivingTimeReqs = new Array(plot.length);
            //body of the fetch request to be sent out to OSRM
            const drivingTimeBody = {
                method: 'GET'
            };
            //introduce activity counter to see how many activities there are per zone
            let countActivities = 0
            //go through all zones to see what the driving time of activities in that zone is
            for (let i = 0; i < plot.length; i++) {
                //go through all activities to find which ones belong to zone i, calculate their driving time
                countActivities = 0
                drivingTimeActivities[i] = 0
                drivingDistanceActivities[i] = 0
                drivingTimeReqs[i] = "http://router.project-osrm.org/route/v1/driving/"
                for (let j = 0; j < activities2Filtered.length; j++) {
                    if (activities2Filtered[j][3] === i) {
                        //update the request for zone i to contain the coordinates of all activities
                        drivingTimeReqs[i] = drivingTimeReqs[i] + activities2Filtered[j][1] + ',' + activities2Filtered[j][0] + ';'
                        countActivities = countActivities + 1;

                    }
                }
                //remove last ';' from the request string for driving time in zone i
                drivingTimeReqs[i] = drivingTimeReqs[i].slice(0, -1);
                //check whether there are less than 2 activities in this zone
                if (countActivities < 2) {
                    drivingTimeActivities[i] = 0 //minimal driving time needed in this zone
                    drivingDistanceActivities[i] = 0 //minimal distance to drive in this zone
                } else {
                    //now that the request string for zone i is built, send out fetch request
                    const response = await fetch(drivingTimeReqs[i], drivingTimeBody);
                    const drivingData = await response.json();
                    //console.log(drivingData)
                    //using the data from OSRM, compile over all legs of the activities what the total duration is and store it for zone i
                    drivingTimeActivities[i] = await getDrivingTime(drivingData)
                    drivingDistanceActivities[i] = await getDrivingDistance(drivingData)
                }

            }
           // console.log(drivingTimeReqs[1])
            //console.log(drivingTimeActivities)
            //console.log(drivingDistanceActivities)
            setDrivingTimeActiv(prevDrivingTimeActiv => []);
            setDrivingDistanceActiv(prevDrivingDistanceActiv => []);

            //console.log(drivingTimeReqs[1])
            //console.log(drivingTimeActivities)
            //console.log(drivingDistanceActivities)

            let totalDrivingTime = 0
            let totalDrivingDistance = 0
            for (let i = 0; i < drivingTimeActivities.length; i++) {
                totalDrivingTime = totalDrivingTime + drivingTimeActivities[i]
                setDrivingTimeActiv(prevDrivingTimeActiv => [...prevDrivingTimeActiv, drivingTimeActivities[i]]);
                totalDrivingDistance = totalDrivingDistance + drivingDistanceActivities[i]
                setDrivingDistanceActiv(prevDrivingDistanceActiv => [...prevDrivingDistanceActiv, drivingDistanceActivities[i]]);
            }
            setDrivingTime(totalDrivingTime / 3600)
            //time to find fuel cost: fuel cost = (litres used * fuel cost)
            //litres used = driving distance * fuel efficiency
            setFuelCost(((totalDrivingDistance / 1000) * averageFuelConsumption) * averageFuelCost)
            setLoaded(true)

        }

        initial()
        console.log('updated')
    }, [zoneId, calculatedZone])


    return (
        <Card title={name} style={{ width: 'fit-content', marginLeft: '0', marginRight: 'auto' }} >
            <Spin spinning={!loaded} delay={200} tip='Calculating...'>
                <div >
                    {/*fuel cost = fuel input times driving time*/}
                    {/*<p>Total cost: ${fuelCost}</p>*/}
                    <Collapse>
                        <Panel header = {`Total cost: ${fuelCost}`} key="1">
                            <ul>
                                <li>Distance over the zones</li>
                                {drivingDistanceActiv.map((drivingDistance, index) => (
                                    <p key={index}>Zone {index}: {drivingDistance}</p>
                                ))}
                                <li>Total driving distance</li>
                                <p>
                                    {drivingDistanceActiv.map((drivingDistance, index) => (
                                        <span key={index}>{drivingDistance} {index < drivingDistanceActiv.length - 1 && '+'} </span>
                                    ))}
                                    = {drivingDistanceActiv.reduce((acc, time) => acc + time, 0)}
                                </p>
                                <li>Total cost with fuel consumption = {averageFuelConsumption} and fuel cost = {averageFuelCost}</li>
                                <p>Total driving cost =(({drivingDistanceActiv.reduce((acc, time) => acc + time, 0)} / 1000) * {averageFuelConsumption}) * {averageFuelCost}) = {fuelCost}</p>
                            </ul>
                        </Panel>
                    </Collapse>
                    {/* done */}
                    <Tooltip title = "This is computed by adding all of the activity times ">
                        <p>Total activity time (hrs): {time}</p>
                    </Tooltip>

                    {/*driving time = find activities per zone. find driving time in order between those activities using OSRM */}
                    {/*<p>Total driving time: {drivingTime}</p>*/}
                    <Collapse>
                        <Panel key={2} header={`Total driving time: ${drivingTime}`}>
                            <ul>
                                <li>Driving time over the zones</li>
                                {drivingTimeActiv.map((drivingTime, index) => (
                                    <p key={index}>Zone {index}: {drivingTime}</p>
                                ))}
                                <li>Driving time sum</li>
                                <p>Driving time = {drivingTimeActiv.map((drivingTime, index) => (
                                    <span key={index}>{drivingTime} {index < drivingTimeActiv.length - 1 && '+'} </span>
                                ))}
                                    = {drivingTime * 3600} </p>
                                <li>Driving time  in hrs</li>
                                <p>Total driving time ={drivingTime * 3600} / 3600 = {drivingTime}</p>

                            </ul>
                        </Panel>
                    </Collapse>



                </div>
            </Spin>
        </Card>
    );
}
export default TextComponent
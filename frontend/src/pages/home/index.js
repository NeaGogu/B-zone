// External dependencies
import React, { useState, useEffect } from 'react';
import L from 'leaflet';
import { MapContainer, TileLayer, Marker, Popup,  useMapEvents, useMap, LayersControl, LayerGroup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Breadcrumb, Layout, Menu, theme, Form, Input, Button, Dropdown, Space, ConfigProvider, Select } from 'antd';

// Icons
import { LaptopOutlined, NotificationOutlined, UserOutlined, DownOutlined } from '@ant-design/icons';

// Components
// import Heatmap from './components/heatmapComponent';

// CSS
import './index.css';
import dumbzones from './tempData/allcases.json'

// Helper function
function getItem(label, key, icon, children, type) {
    return {
        key,
        icon,
        children,
        label,
        type,
    };
}

// Function to convert a string to an array
function toArr(str) {
    return Array.isArray(str) ? str : [str];
}

// Context
const MyFormItemContext = React.createContext([]);

const MyFormItemGroup = ({ prefix, children }) => {
    const prefixPath = React.useContext(MyFormItemContext);
    const concatPath = React.useMemo(() => [...prefixPath, ...toArr(prefix)], [prefixPath, prefix]);
    return <MyFormItemContext.Provider value={concatPath}>{children}</MyFormItemContext.Provider>;
};

const MyFormItem = ({ name, ...props }) => {
    const prefixPath = React.useContext(MyFormItemContext);
    const concatName = name !== undefined ? [...prefixPath, ...toArr(name)] : undefined;
    return <Form.Item name={concatName} {...props} />;
};

// Components from Ant Design
const { SubMenu } = Menu;
const { darkAlgorithm } = theme;
const { Header, Content, Sider } = Layout;
const item = [
    getItem('Saved zones', 'sub1', null, [
        getItem('Initial Zone', 'g1', null, null, 'group'),
        getItem('Saved Zone 1', 'g2', null, null, 'group'),
    ]),
    getItem('Filter', 'sub2', null, null)
];

// Form handling
const handleChange = (value) => {
    console.log(`selected ${value}`);
};

function setEmail() {
    document.getElementById('email').innerHTML = user_id //gets email for text in item
}

// User details
const user_id = localStorage.getItem('email')

// User menu items
const user_items = [
    {
        key: '1',
        label: (
            <p target="_blank" rel="noopener noreferrer" id="email" style={{ color: '#ffd369' }}>
                <var>{user_id}</var>

            </p>
        )
    },
    {
        key: '2',
        label: (
            <a style={{ color: '#ffd369' }} target="_blank" rel="noopener noreferrer" onClick={() => signOut()}>
                Log Out
            </a>
        )
    },
];

// Menu items
const menu = [
    getItem('Saved zones', 'sub1', null, [
        getItem('Initial Zone', 'g1', null, null, 'group'),
        getItem('Saved Zone 1', 'g2', null, null, 'group'),
    ]),
    getItem('Filter', 'sub2', null, null)
];

//function for map which shows address, zipcode and coordinates when clicking on the map
function LocationMarker() {
    const [position, setPosition] = useState(null);
    const [address, setAddress] = useState(null);
    const [zipcode, setZipcode] = useState(null);

    const map = useMapEvents({
        click(e) {
            const { lat, lng } = e.latlng;
            setPosition(e.latlng);
            const API_KEY = 'pk.eyJ1IjoidGFuaWFnb2lhMTEiLCJhIjoiY2xleTRrYm02MDlmMTN4bzVsZTR4cWp4OCJ9.hmT59q-Q1IcEjC6mdY2R9w';
            const API_URL = `https://api.mapbox.com/geocoding/v5/mapbox.places/${lng},${lat}.json?access_token=${API_KEY}&types=postcode,address`;
            fetch(API_URL)
                .then(response => response.json())
                .then(data => {
                    const features = data.features;
                    if (features.length > 0) {
                        const address = features[0].place_name;
                        const zipcodes = features.filter(feature => feature.place_type[0] === 'postcode');
                        const zipcode = zipcodes.length > 0 ? zipcodes[0].text : null;
                        setAddress(address);
                        setZipcode(zipcode);
                    } else {
                        setAddress(null);
                        setZipcode(null);
                    }
                });
            map.flyTo(e.latlng, map.getZoom());
        },
    });

    return position === null ? null : (
        <Marker position={position}>
            <Popup>
                <div>
                    <div>Latitude: {position.lat.toFixed(4)}</div>
                    <div>Longitude: {position.lng.toFixed(4)}</div>
                    {address && <div>Address: {address}</div>}
                    {zipcode && <div>Zipcode: {zipcode}</div>}
                </div>
            </Popup>
        </Marker>
    );
}
//signOut function
function signOut() {
    const token = localStorage.getItem('token');
    var valid;

    //1. delete the user token and send sign out GET message
    console.log("Fetching sign out API")
    fetch(`https://sep202302.bumbal.eu/api/v2/authenticate/sign-out?token=${encodeURIComponent(token)}`, {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`, // add token to Bearer Authorization when sending GET signOut request
        }
    })
        .then((response) => {
            console.log("This is the response:")
            console.log(response)
            if (response.ok) {
                valid = true;

            }
        }).then((data) => {
            console.log("This is the data:")
            console.log(data)
            if (valid) {
                localStorage.clear();
                alert('You have been logged out!')
                //2. re-route user to home page
                window.location.reload()
            } else {
                alert("You could not be logged out! Please try again later.")
            }
        })

}

//Input field function -> later on add calculations, for now checks if the two fields are filled and if so, then the button is activated
const ZoneSubMenu = ({ onSubmit }) => {
    const [averageFuelCost, setAverageFuelCost] = useState("");
    const [averageFuelUsage, setAverageFuelUsage] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        const isValid = onSubmit(averageFuelCost, averageFuelUsage);
        if (isValid) {
            // add calculations
        }
    };

    return (
        <Form onFinish={handleSubmit}>
            <Form.Item rules={[{ required: true }]}>
                Average fuel cost
                <Input
                    placeholder="input fuel cost"
                    type="number"
                    step="0.01"
                    value={averageFuelCost}
                    onChange={(e) => setAverageFuelCost(e.target.value)}
                />
            </Form.Item>
            <Form.Item>
                Average fuel usage of car
                <Input
                    placeholder="input fuel usage of car"
                    type="number"
                    step="0.01"
                    value={averageFuelUsage}
                    onChange={(e) => setAverageFuelUsage(e.target.value)}
                />
            </Form.Item>
            <Button
                style={{ width: "100%" }}
                type="primary"
                htmlType="submit"
                disabled={!averageFuelCost || !averageFuelUsage}
            >
                Calculate
            </Button>
        </Form>
    );
};

localStorage.getItem('token')
export default function Home() {

    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const onFinish = (value) => {
        console.log(value);
    };

    // for comparison button to split maps
    const [showComparison, setShowComparison] = useState(false);
    const [showMap, setShowMap] = useState(true);
    // for toggling between zones
    const [selection, setSelection] = useState([])
    const [poly, setPoly] = useState([])
    const [togZone, setTogZone] = useState(true)

    //set of zipCodes
    var zipCodes

    // save which item in menu is selected
    const menuClick = (e) => {
        setSelection(e.key);
    }

    function getInitialZones(userToken) {
        //get list of zones, question remains how to get the last used zone?
        //definition of URl, body values, other parameters
        const zonesURL = "https://sep202302.bumbal.eu/api/v2/zone"
        const bodyValues = JSON.stringify({
            "options": {
                "include_zone_ranges": true,
                "include_brands": true
            },
            "filters": {}
        })

        //sending fetch request to receive list of user's zones
        fetch(zonesURL, {
            method: 'PUT',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userToken}`
            },
            body: bodyValues
        })
            //testing if response recorded was ok
            .then((response) => {
                if(!response.ok) {
                    console.log("Response was not ok ???")
                    alert("Unable to retrieve this zone configuration!")
                }
                return response.json();
            })
            //dealing with received list of zones
            .then((data) => {
                localStorage.setItem('zoneID', data.items[0].id.toString())
                zipCodes = getZipCodes(data.items)
                console.log("ID set to local storage")
                console.log(zipCodes[0])
            })
            .catch(error => console.log(error, 'error'))
    }

    async function getZipCodes(zoneList) {
        //when you retrieve a list of zones from Bumbal API from zone with PUT, you retrieve a list of zone configs which itself includes a list of zones in each zone config
        let zipCodes = []

        for (let i = 0; i < zoneList.length; i++) {
            zipCodes[i] = getAreas(zoneList[i].zone_ranges); //per each zone, there are a list of areas with from and to
        }

        //fetch('http://localhost:4000/test/zip/coordinates?zip_from=' + zipCodes[0].zipFrom.toString() + '&zip_to=' + zipCodes[0].zipTo.toString(), {
        let res = await fetch('http://localhost:4000/test/zip/coordinates?zip_from=5611&zip_to=5613')

        if(!res.ok) {
            console.error(`Response to localhost:4000 was not ok ??? ${res.status}`)
            alert("Unable to retrieve zipcodes from backend!")
            return
        }

        console.log(res.json())

        console.log("gucci")
        // .then((response) => {
        //     if(!response.ok) {
        //         console.error(`Response to localhost:4000 was not ok ??? ${response.}`)
        //         alert("Unable to retrieve zipcodes from backend!")
        //         return
        //     }
        //     return response.json();
        // })
        //     .then((data) => {
        //         zipCodes[0] = data
        //     })
        return zipCodes

    }

    function getAreas(zoneAreas) {
        //create zipFrom and zipTo object
        var object = {
            zipFrom: "",
            zipTo: ""
        };
        //create array of zipcode ranges
        let zoneAreaZips = [object]
        //loop through range items to find zip code ranges
        for (let j = 0; j < zoneAreas.length; j++) {
            zoneAreaZips[j].zipFrom = zoneAreas[j].zipcode_from;
            zoneAreaZips[j].zipTo = zoneAreas[j].zipcode_to;
        }
        return zoneAreaZips
    }

    const toggleComparison = () => {
        if (showComparison) {
            return; // If showComparison is already true, do nothing
        }

        setShowComparison(true);
        setShowMap(false); // reset to singular map
    }

    const toggleMap = () => {
        setShowMap(!showMap);
        setShowComparison(false); // reset comparison state when switching singular map
        if (selection === '5' ) {
            if (togZone) {
                getInitialZones(localStorage.getItem('token'))
                setPoly(dumbzones.nl)
                setTogZone(!togZone)
            } else {
                setPoly([])
                setTogZone(!togZone)
            }
        }
        if (selection === '6' ) {
            if (togZone) {
                setPoly([])
                setTogZone(!togZone)
            } else {
                setPoly([])
                setTogZone(!togZone)
            }
        }
        //console.log(selection)
    };

    return (
        <ConfigProvider
            // theme which we should use
            theme={{
                token: {
                    colorPrimary: "#ffd369",
                    colorBgBase: "#222831",
                    colorPrimaryBg: "#393E46",
                    colorTextBase: "#eeeeee"
                },
                // renders antd components with dark mode
                algorithm: darkAlgorithm
            }}
        >
            <Layout style={{ height: '100vh' }}>
                <Header className="header" >
                    <div style={{ display: 'flex', justifyContent: 'flex-end' }} >
                        <div style={{ margin: 'auto', width: '100%' }} >
                            <img src='./b-zone-logo.png' style={{ display: 'block', width: '11%' }} alt='logo' />
                        </div>
                        <div className="headerText">
                            B-ZONE
                        </div>

                        <Dropdown
                            menu={{
                                items: user_items,
                            }}
                        >
                            <a onClick={(e) => e.preventDefault()}>
                                <Space>
                                    <button type="button" variant="contained" style={{ float: 'right' }} >
                                        <UserOutlined />
                                    </button>
                                    <DownOutlined style={{ color: '#ffd369' }} />
                                </Space>
                            </a>
                        </Dropdown>



                    </div>

                    <Menu theme="dark" mode="horizontal" />

                </Header>
                <Layout>
                    <Sider
                        width={200}
                        style={{
                            background: colorBgContainer,
                        }}
                    >

                        <Menu
                            mode="inline"
                            defaultSelectedKeys={['1']}
                            defaultOpenKeys={['sub1']}
                            style={{
                                height: '100%',
                                borderRight: 0,
                            }}
                            onClick={menuClick}
                        >
                            <SubMenu key="sub3" title="Heat map">
                                <Menu.Item key="5">Location based </Menu.Item>
                                <Menu.Item key="6">Time based </Menu.Item>
                            </SubMenu>


                            <SubMenu key="sub4" title="Zones">
                                <ZoneSubMenu />
                            </SubMenu>
                            <SubMenu key="sub2" title="Saved Zones">
                                <Menu.Item key="5" style={{ height: "80px", padding: 0 }}>
                                    <div style={{ textAlign: "center" }}>Initial Zone</div>
                                    <div style={{ display: "flex", justifyContent: "space-between", width: "100%" }}>
                                        <Button style={{ flex: 1, marginRight: "3px" }} onClick={toggleMap}>View</Button>
                                        <Button style={{ flex: 1, marginLeft: "3px" }} onClick={toggleComparison}>Compare</Button>
                                    </div>
                                </Menu.Item>
                                <Menu.Item key="6" style={{ height: "80px", padding: 0 }}>
                                    <div style={{ textAlign: "center" }}>Saved Zone 1</div>
                                    <div style={{ display: "flex", justifyContent: "space-between", width: "100%" }}>
                                        <Button style={{ flex: 1, marginRight: "3px" }} onClick={toggleMap}>View</Button>
                                        <Button style={{ flex: 1, marginLeft: "3px" }} onClick={toggleComparison}>Compare</Button>
                                    </div>
                                </Menu.Item>
                            </SubMenu>

                        </Menu>
                    </Sider>
                    <Layout style={{

                        padding: 30
                    }}
                    >

                        <Content className="map" id="map"
                            style={{
                                minHeight: 500,
                            }}
                        >
                            {showComparison ? (
                                <div style={{ display: "flex", justifyContent: "space-between", padding: "5px" }}>
                                    <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500, flex: "1" }}>
                                        <TileLayer
                                            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                        />
                                        <LocationMarker />
                                    </MapContainer>
                                    <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500, flex: "1", marginLeft: "20px" }}>
                                        <TileLayer
                                            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                        />
                                        <LocationMarker />
                                    </MapContainer>
                                </div>
                            ) : (
                                <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500, flex: "1", marginLeft: showMap ? "20px" : "0" }}>
                                    <TileLayer
                                        attribution='&copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors'
                                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                    />
                                    <LayersControl position='topright'>
                                        <LayersControl.Overlay name='Heatmap'>
                                            <LayerGroup>
                                                {/*<Heatmap />*/}
                                            </LayerGroup>
                                        </LayersControl.Overlay>
                                    </LayersControl>

                                    <LocationMarker />
                                </MapContainer>
                            )}
                        </Content>
                    </Layout>
                </Layout>
            </Layout>

        </ConfigProvider>


    );
}
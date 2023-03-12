import { LaptopOutlined, NotificationOutlined, UserOutlined, DownOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, Form, Input, Button, Dropdown, Space, ConfigProvider } from 'antd';


import DropdownButton from "antd/es/dropdown/dropdown-button";
import React, {useState, useEffect} from 'react';
import L from 'leaflet';
import { MapContainer, TileLayer, Marker, Popup,  useMapEvents, useMap } from 'react-leaflet';

import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import './index.css';
import "leaflet.heat"

function HeatmapFunction({show}){
  const map = useMap()
  

  if ( show ) {
    const fetchData = async () => {
        let addressPoints = await findAddressesPoints();
        //console.log(addressPoints)
        const points = addressPoints
        ? addressPoints.map((p) => {
        return [p[0], p[1], 500]; // lat lng intensity
        })
        : [];
  
        var heat = L.heatLayer(points);
        map.addLayer(heat)
      };
  
      fetchData();
  } else {
    map.eachLayer(function(layer) {
        if (layer._heat != null) {
            map.removeLayer(layer)
        }
    })
  }
    
    

    

}
/*
async componentDidMount() {
    // POST request using fetch with async/await
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: 'React POST Request Example' })
    };
    const response = await fetch('https://reqres.in/api/posts', requestOptions);
    const data = await response.json();
    this.setState({ postId: data.id });
}
*/
async function findAddressesPoints() {
   
    
    const response = await getActivities();
    
    if (response.status === 401) {
        // FIXME Delete after token updated correctly
        // await updateToken( "sep@bumbal.eu", "cW$#Qbph0524");
        // response = await getActivities();
        alert('Token has expired!');
        localStorage.removeItem('token')
        window.location.reload()
    }

    const data = await response.json();

    console.log(data) 
    let newData = data.items.map((i) => {
        return [i.address.latitude, i.address.longitude]; // lat lng intensity
        })
    console.log(newData)  
    return newData;
    
}

async function getActivities() {
    const token = localStorage.getItem('token')
    console.log('token ' + token)
    const requestOptions = {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`, // add token to Bearer Authorization when sending GET signOut request
        },
        body: JSON.stringify({
            "options": {
              "include_address": true
            },
            "limit": 10,
            "offset": 0,
            "sorting_column": "id",
            "sorting_direction": "asc",
            "as_list": true,
            "count_only": false
          })
    };
    const response = await fetch('https://sep202302.bumbal.eu/api/v2/activity', requestOptions);
    console.log('getActivities response ' + response.status)
    return response;
}

function getItem(label, key, icon, children, type) {
    return {
        key,
        icon,
        children,
        label,
        type,
    };
}

function setEmail() {
    document.getElementById('email').innerHTML = id_user //gets email for text in item
}

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

const id_user = localStorage.getItem('email')

const items = [
    {
        key: '1',
        label: (
            <p target="_blank" rel="noopener noreferrer" id="email" style={{ color: '#ffd369' }}>
                <var>{id_user}</var>

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
    var expired;

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
            if (response.status === 401) {
                expired = true;
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
            if ( expired) {
                alert('Token has expired')
                localStorage.removeItem('token')
            } else {
                alert("You could not be logged out! Please try again later.")
            }
        }
    })


}

const MyFormItemContext = React.createContext([]);
function toArr(str) {
    return Array.isArray(str) ? str : [str];
}
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

localStorage.getItem('token')


export default function Home() {
    const [showResults, setShowResults] = React.useState(false)
   
    const onClick = () => setShowResults(!showResults)

    
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const onFinish = (value) => {
        console.log(value);
    };

    // for comparison button to split maps
    const [showComparison, setShowComparison] = useState(false);
    const [showMap, setShowMap] = useState(true);

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
                                items,
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
                        >
                            <div style={{ width: "100%" }}>
                                <Button style={{ width: "50%" }} type="primary" onClick={onClick}>Heat map</Button>
                                <Button style={{ width: "50%" }} type="primary">Zones</Button>
                            </div>

                            <Form name="form_item_path" layout="vertical" onFinish={onFinish}>
                                <MyFormItemGroup>
                                    <MyFormItemGroup>
                                        <MyFormItem name="fuelCost" label="Average fuel cost">
                                            <Input placeholder="1" />
                                        </MyFormItem>
                                        <MyFormItem name="fuelUsage" label="Average fuel usage of car">
                                            <Input placeholder="1" />
                                        </MyFormItem>
                                    </MyFormItemGroup>
                                </MyFormItemGroup>

                                <Button style={{ width: "100%" }} type="primary">Calculate</Button>
                                &nbsp;
                                <Button style={{ width: "100%" }} type="primary">Save</Button>
                            </Form>

                            <SubMenu key="sub2" title="Saved Zones">
                                <Menu.Item key="5">Initial Zone</Menu.Item>
                                <Menu.Item key="6">Saved Zone 1</Menu.Item>
                                <Button style={{ width: "100%" }} type="primary" onClick={toggleComparison}>Compare</Button>
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
                                        <LocationMarker/>
                                    </MapContainer>
                                    <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500, flex: "1", marginLeft: "20px" }}>
                                        <TileLayer
                                            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                        />
                                        <LocationMarker/>
                                    </MapContainer>
                                </div>
                            ) : (
                                <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500, flex: "1", marginLeft: showMap ? "20px" : "0" }}>
                                    <TileLayer
                                        attribution='&copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors'
                                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                    />
                                    <LocationMarker/>
                                    <HeatmapFunction show={showResults}/>
                                </MapContainer>
                            )}
                        </Content>
                    </Layout>
                </Layout>
            </Layout>

        </ConfigProvider>


    );
}
import { LaptopOutlined, NotificationOutlined, UserOutlined, DownOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, Form, Input, Button, Dropdown, Space, ConfigProvider } from 'antd';
import React from 'react';
import DropdownButton from "antd/es/dropdown/dropdown-button";
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import './index.css';

import { useMap, useMapEvents } from 'react-leaflet'
import { useEffect } from 'react'
import L from 'leaflet'
import "leaflet.heat"

function HeatmapFunction(){
  const map = useMap()
  useEffect(() => {
    var addressPoints = [
      [
          52.141,
          4.457
      ],
      [
          51.853,
          4.377
      ],
      [
          51.851,
          4.329
      ],
      [
          52.044,
          4.728
      ],
      [
          52.154,
          4.662
      ],
      [
          52.434,
          4.647
      ],
      [
          52.642,
          5.003
      ],
      [
          52.013,
          4.746
      ],
      [
          52.301,
          4.547
      ],
      [
          52.605,
          4.689
      ],
      [
          51.893,
          4.543
      ],
      [
          52.092,
          4.327
      ],
      [
          52.236,
          4.451
      ],
      [
          52.444,
          4.635
      ],
      [
          52.046,
          4.731
      ],
      [
          51.919,
          4.564
      ],
      [
          51.832,
          4.733
      ],
      [
          52.0,
          4.443
      ],
      [
          52.586,
          4.704
      ],
      [
          52.067,
          4.237
      ],
      [
          52.14,
          4.534
      ],
      [
          52.333,
          4.634
      ],
      [
          51.937,
          4.331
      ],
      [
          52.05,
          4.644
      ],
      [
          52.657,
          5.108
      ],
      [
          52.281,
          4.626
      ],
      [
          51.956,
          4.604
      ],
      [
          52.141,
          4.457
      ],
      [
          51.853,
          4.377
      ],
      [
          51.793,
          4.481
      ],
      [
          51.851,
          4.329
      ],
      [
          51.793,
          4.481
      ],
      [
          52.373,
          4.528
      ],
      [
          52.154,
          4.662
      ],
      [
          52.434,
          4.647
      ],
      [
          52.642,
          5.003
      ],
      [
          52.013,
          4.746
      ],
      [
          52.301,
          4.547
      ],
      [
          52.605,
          4.689
      ],
      [
          51.795,
          4.047
      ],
      [
          52.235,
          4.459
      ],
      [
          52.036,
          4.718
      ],
      [
          51.793,
          4.712
      ],
      [
          52.349,
          4.585
      ],
      [
          52.044,
          4.728
      ],
      [
          52.088,
          4.311
      ],
      [
          51.893,
          4.543
      ],
      [
          51.86,
          4.506
      ],
      [
          51.849,
          4.488
      ],
      [
          52.236,
          4.451
      ],
      [
          51.909,
          4.171
      ],
      [
          52.435,
          4.655
      ],
      [
          52.431,
          4.656
      ],
      [
          52.633,
          4.665
      ],
      [
          52.444,
          4.635
      ],
      [
          52.046,
          4.731
      ],
      [
          51.919,
          4.564
      ],
      [
          52.034,
          4.3
      ],
      [
          51.832,
          4.733
      ],
      [
          52.027,
          4.696
      ],
      [
          52.169,
          4.579
      ],
      [
          52.317,
          4.644
      ],
      [
          52.0,
          4.443
      ],
      [
          52.224,
          4.669
      ],
      [
          52.586,
          4.704
      ],
      [
          52.067,
          4.237
      ],
      [
          52.14,
          4.534
      ],
      [
          52.162,
          4.469
      ],
      [
          51.53,
          4.452
      ],
      [
          52.238,
          4.431
      ],
      [
          51.835,
          4.16
      ],
      [
          51.91,
          4.168
      ],
      [
          51.904,
          4.232
      ],
      [
          51.842,
          4.155
      ],
      [
          51.846,
          4.496
      ],
      [
          52.333,
          4.634
      ],
      [
          52.306,
          4.865
      ],
      [
          51.937,
          4.331
      ],
      [
          51.954,
          4.611
      ],
      [
          51.993,
          4.512
      ],
      [
          52.315,
          4.874
      ],
      [
          52.036,
          4.644
      ],
      [
          52.03,
          4.616
      ],
      [
          52.05,
          4.644
      ],
      [
          52.657,
          5.108
      ],
      [
          52.484,
          4.762
      ],
      [
          51.997,
          4.482
      ],
      [
          52.281,
          4.626
      ],
      [
          51.835,
          4.353
      ],
      [
          52.64,
          4.813
      ],
      [
          51.956,
          4.604
      ],
      [
          52.436,
          4.805
      ],
      [
          51.923,
          4.612
      ]]
    const points = addressPoints
    ? addressPoints.map((p) => {
      return [p[0], p[1], 500]; // lat lng intensity
      })
    : [];

    L.heatLayer(points).addTo(map);
  }, []);
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
            <p target="_blank" rel="noopener noreferrer" id="email" style={{ color: '#ffd369'}}>
                <var>{id_user}</var>

            </p>
        )
    },
    {
        key: '2',
        label: (
            <a style={{ color: '#ffd369'}} target="_blank" rel="noopener noreferrer" onClick={() => signOut()}>
                Log Out
            </a>
        )
    },
];

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
    const onClick = () => setShowResults(true)
    
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const onFinish = (value) => {
        console.log(value);
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
            <Layout  style={{height:'100vh'}}>
                <Header className="header" >
                    <div style={{ display: 'flex', justifyContent: 'flex-end' }} >
                        <div style={{ margin: 'auto', width: '100%' }} >
                            <img src='./b-zone-logo.png' style={{ display: 'block', width: '11%'}} alt='logo'/>
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
                                    <button type="button" variant="contained" style={{ float: 'right'}} >
                                        <UserOutlined />
                                    </button>
                                    <DownOutlined style={{ color: '#ffd369'}} />
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
                                <Button style={{ width: "100%" }} type="primary">Compare</Button>
                            </SubMenu>

                        </Menu>
                    </Sider>
                    <Layout  style={{
                                
                                padding: 30
                            }}
                    >

                        <Content className="map" id="map"
                            style={{
                                minHeight: 500,
                            }}
                        >
                            <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500 }}>
                                <TileLayer
                                    attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                />
                                 { showResults ? <HeatmapFunction/> : null }
                            </MapContainer>
                        </Content>
                    </Layout>
                </Layout>
            </Layout>

        </ConfigProvider>


    );
}
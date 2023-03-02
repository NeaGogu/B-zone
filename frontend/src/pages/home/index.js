import { LaptopOutlined, NotificationOutlined, UserOutlined, DownOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, Form, Input, Button, Dropdown, Space, ConfigProvider } from 'antd';
import React from 'react';
import DropdownButton from "antd/es/dropdown/dropdown-button";
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";

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
            <p target="_blank" rel="noopener noreferrer" id="email">
                <var>{id_user}</var>

            </p>
        )
    },
    {
        key: '2',
        label: (
            <a target="_blank" rel="noopener noreferrer" onClick={() => signOut()}>
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
            <Layout>
                <Header className="header" >
                    <div style={{ display: 'flex', justifyContent: 'flex-end' }} >

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
                                    <DownOutlined />
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
                                <Button style={{ width: "50%" }} type="primary">Heat map</Button>
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
                    <Layout
                        style={{
                            padding: '0 24px 24px',
                        }}
                    >

                        <Content className="map" id="map"
                            style={{
                                margin: '24px 16px',
                                padding: 24,
                                minHeight: 500,
                                background: colorBgContainer,
                            }}
                        >
                            <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500 }}>
                                <TileLayer
                                    attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                />
                            </MapContainer>
                        </Content>
                    </Layout>
                </Layout>
            </Layout>

        </ConfigProvider>


    );
}
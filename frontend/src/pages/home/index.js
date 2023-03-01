import { LaptopOutlined, NotificationOutlined, UserOutlined, DownOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, Button, Dropdown, Space } from 'antd';
import React from 'react';
import DropdownButton from "antd/es/dropdown/dropdown-button";
import { MapContainer, TileLayer,Marker,Popup } from 'react-leaflet';
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
        key: '3',
        label: (
            <a target="_blank" rel="noopener noreferrer" >
                password
            </a>
        )
    },
];
localStorage.getItem('token')
export default function Home() {

    const {
        token: { colorBgContainer },
    } = theme.useToken();


    return (
        <Layout>
            <Header className="header">
                <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                    className="logo"

                    <Dropdown
                        menu={{
                            items,
                        }}
                    >
                        <a onClick={(e) => {e.preventDefault(); }}>
                            <Space>
                                <button type="button"  variant="contained" style={{float: 'right'}} >
                                    <UserOutlined/>
                                </button>
                                <DownOutlined />
                            </Space>
                        </a>
                    </Dropdown>



                </div>

                <Menu theme="dark" mode="horizontal"/>
                
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
                        <SubMenu key="sub1" title="Input field">
                        </SubMenu>
                        <SubMenu key="sub2"  title="Saved Zones">
                            <Menu.Item key="5">Initial Zone</Menu.Item>
                            <Menu.Item key="6">Saved Zone 1</Menu.Item>
                        </SubMenu>
                        <Button type="primary">Compare</Button>


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
    );
}
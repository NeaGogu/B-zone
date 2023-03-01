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
        key: '3',
        label: (
            <a target="_blank" rel="noopener noreferrer" >
                password
            </a>
        )
    },
];

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

                                <Button style={{ width: "100%" }} type="primary">
                                    Calculate
                                </Button>
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
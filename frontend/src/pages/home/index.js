// External dependencies
import React, { useState, useEffect } from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Layout, Menu, theme, Form, Input, Button, Dropdown, Space, ConfigProvider, Radio, InputNumber } from 'antd';

// Icons
import { UserOutlined, SaveOutlined } from '@ant-design/icons';

// Components
import Map from './components/mapComponent';

// CSS
import './index.css';
import SiderComponent from "./components/siderComponent";
//import dumbzones from './tempData/allcases.json'

// Components from Ant Design
const { SubMenu } = Menu;
const { darkAlgorithm } = theme;
const { Header, Content, Sider } = Layout;

// User email.
const user_id = localStorage.getItem('email')

// User credentials menu items: their email and the sign out button.
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
            <button style={{ color: '#ffd369', textDecoration: 'underline', backgroundColor: 'transparent', border: 'none', cursor: 'pointer' }} onClick={() => signOut()}>
                Sign Out
            </button>
        )
    },
];

/** 
* Signs out the user, so deletes the user token and redirects them to the login page.
* @return {void}
*/
function signOut() {
    const token = localStorage.getItem('token');
    var valid;

    //1. Delete the user token and send sign out GET message.
    console.log("Fetching sign out API")
    fetch(`https://sep202302.bumbal.eu/api/v2/authenticate/sign-out?token=${encodeURIComponent(token)}`, {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
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
                //2. Re-route user to home page.
                window.location.reload()
            } else {
                alert("You could not be logged out! Please try again later.")
            }
        })
}



localStorage.getItem('token')

export default function Home() {
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    // For comparison button to split map into two maps.
    const [showComparison, setShowComparison] = useState(false);

    // For view button to bring two maps back to one map.
    const [showMap, setShowMap] = useState(true);


    // For radio.
    const [value, setValue] = useState(1);
    const onChange = (e) => {
        setValue(e.target.value);
    };

    // For number.
    const [intensity, setIntensity] = useState(500)
    const onChangeNumber = (e) => {
        console.log(intensity)
        setIntensity(e)
    }

    // Contains the name of the zone configuration that the user wants to save.
    const [saveName, setSaveName] = useState('');

    /** 
    * When the user clicks the save button, they are asked to give a name to the zone configuration.
    * @param {string} name - The name to give the saved zone configuration.
    * @return {void}
    */
    function handleSaveClick() {
        const name = window.prompt('Enter a name for the save:');
        if (name) {
            addSavedZone(name);
            setSaveName(name);
        }
    }

    // This makes sure the initial zone configuration is always shown.
    const [savedZones, setSavedZones] = useState([
        { key: 'saved-initial', name: 'Initial Zone' },

    ]);
    localStorage.setItem('saved-initial', 'Initial Zone');
    const handleDeleteZone = (key) => {
        localStorage.removeItem(key);
        const newSavedZones = savedZones.filter((item) => item.key !== key);
        setSavedZones(newSavedZones);
    };

    /** 
    * When the user has entered a name for the zone configuration they want to save, it is shown in the saved zones list.
    * @param {string} name - The name given to the saved zone configuration.
    * @return {void}
    */
    function addSavedZone(name) {
        const key = `saved-${Date.now()}-${Math.random()}`;
        const newZone = { key, name };
        setSavedZones([...savedZones, newZone]);
        localStorage.setItem(key, name);
    }

    useEffect(() => {
        const savedZones = Object.keys(localStorage)
            .filter(key => key.startsWith('saved-'))
            .map(key => ({
                key,
                name: localStorage.getItem(key)
            }));
        setSavedZones(savedZones);
    }, []);

    return (
        <ConfigProvider
            // Theme of the web-app
            theme={{
                token: {
                    colorPrimary: "#ffd369",
                    colorBgBase: "#222831",
                    colorPrimaryBg: "#393E46",
                    colorTextBase: "#eeeeee"
                },
                // Renders AntDesign components with dark mode
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

                        <Space >
                            <Button
                                type="primary"
                                style={{ verticalAlign: 'middle', background: 'transparent' }}
                                onClick={handleSaveClick}
                            >
                                <SaveOutlined style={{ color: '#ffd369' }} />
                            </Button>
                            <Dropdown
                                menu={{
                                    items: user_items,
                                }}
                            >
                                <button type="button" onClick={signOut} style={{ color: '#ffd369', backgroundColor: 'transparent', border: 'none', textDecoration: 'underline', cursor: 'pointer' }}>
                                    <Space>
                                        <UserOutlined style={{ verticalAlign: 'middle' }} />
                                    </Space>
                                </button>
                            </Dropdown>
                        </Space>
                    </div>
                    <Menu theme="dark" mode="horizontal" />
                </Header>

                <Layout>
                    <Sider
                        width={"200"}
                        style={{
                            background: colorBgContainer,
                        }}
                    >
                        //heatmap settings as Menu for testing purposes 
                            {/*<Menu key="sub3" title="Heat map" style={{}}>*/}
                            {/*    <Menu.Item key="5" style={{ padding: 0 }}>*/}
                            {/*        <Radio.Group value={value} onChange={onChange} size='small'  >*/}
                            {/*            <Radio.Button value={1}>*/}
                            {/*                Time based*/}
                            {/*            </Radio.Button>*/}
                            {/*            <Radio.Button value={2}>*/}
                            {/*                Location based*/}
                            {/*            </Radio.Button>*/}
                            {/*        </Radio.Group>*/}
                            {/*    </Menu.Item>*/}
                            {/*    <Menu.Item key="6" style={{ height: "80px", padding: 0 }}>*/}
                            {/*        <div style={{ textAlign: "center" }}>Intensity</div>*/}
                            {/*        <div style={{ paddingLeft: 50 }}>*/}
                            {/*            <InputNumber min={1} max={1000} defaultValue={500} onChange={onChangeNumber} disabled={value === 1 ? true : false} />*/}
                            {/*        </div>*/}
                            {/*    </Menu.Item>*/}
                            {/*</Menu>*/}

                            //sider component
                            <SiderComponent
                                savedZones={savedZones}
                                setSavedZones={setSavedZones}
                                onDeleteZone={handleDeleteZone}
                                showMap={showMap}
                                setShowMap={setShowMap}
                                showComparison={showComparison}
                                setShowComparison={setShowComparison}
                                value={value}
                                onChange={onChange}
                                intensity={intensity}
                                onChangeNumber={onChangeNumber}
                            />
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
                                    <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber}/>
                                    <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber}/>
                                </div>
                            ) : (
                                <Map /> 
                            )}
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    );
}
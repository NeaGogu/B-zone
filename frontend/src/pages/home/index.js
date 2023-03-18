// External dependencies
import React, { useState, useEffect } from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Layout, Menu, theme, ConfigProvider } from 'antd';

// Components
import Map from './components/mapComponent';

// CSS
import './index.css';
import SiderComponent from "./components/sliderComponent";
//import dumbzones from './tempData/allcases.json'
import HeaderComponent from "./components/headerComponent";

// Components from Ant Design
const { darkAlgorithm } = theme;
const { Header, Content, Sider } = Layout;

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
        console.log('home')
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

    useEffect(() => {
        console.log('home')
        console.log(value, intensity)
    }, [value, intensity]);

    return (
        <ConfigProvider
            // Theme of the web-app.
            theme={{
                token: {
                    colorPrimary: "#ffd369",
                    colorBgBase: "#222831",
                    colorPrimaryBg: "#393E46",
                    colorTextBase: "#eeeeee"
                },
                // Renders AntDesign components with dark mode.
                algorithm: darkAlgorithm
            }}
        >
            <Layout style={{ height: '100vh' }}>
                <Header className="header" >
                    <HeaderComponent handleSaveClick={handleSaveClick} savedZones={savedZones} saveName={saveName} />
                    <Menu theme="dark" mode="horizontal" />
                </Header>

                <Layout>
                    <Sider
                        width={"200"}
                        style={{
                            background: colorBgContainer,
                        }}
                    >
                        <SiderComponent
                            savedZones={savedZones}
                            addSavedZone={addSavedZone}
                            setSavedZones={setSavedZones}
                            onDeleteZone={handleDeleteZone}
                            showMap={showMap}
                            setShowMap={setShowMap}
                            showComparison={showComparison}
                            setShowComparison={setShowComparison}
                            values={value}
                            setValue={setValue.bind(this)}
                            setIntensity={setIntensity.bind(this)}
                            intensity={intensity}
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
                                    <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} />
                                    <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} />
                                </div>
                            ) : (
                                <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} />
                            )}
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    );
}
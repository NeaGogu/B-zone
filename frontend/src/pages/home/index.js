// External dependencies
import React, { useState, useEffect } from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Layout, Menu, theme, ConfigProvider, Spin } from 'antd';

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

    //for computed zone on map 1
    const [computed, setComputed] = useState(false)
    const [computed2, setComputed2] = useState(false)


    // For radio.
    const [value, setValue] = useState(1);
    const [zipCodes, setZipCodes] = useState([]);
    const onChange = (e) => {
        setValue(e.target.value);
    };

    //track which zone is selected
    const [zoneId, setZoneId] = useState('initial');

    // For number.
    const [intensity, setIntensity] = useState(500)
    const onChangeNumber = (e) => {
        console.log('home')
        console.log(intensity)
        setIntensity(e)
    }

    /** 
    * When the user clicks the save button, they are asked to give a name to the zone configuration if there is a zone to be saved.
    * @return {void}
    */
    async function handleSaveClick() {
        if (zipCodes.length === 0) {
            alert('nothing to save')
            return null;
        }
        const name = window.prompt('Enter a name for the save:');
        if (name) {
            addSavedZone(name);
        }
    }

    // keep track of the saved zones and update when needed
    const [savedZones, setSavedZones] = useState([]);

    //
     // for zone sync
     const [currentView, setCurrentView] = useState('initial')

    // for now not usefull
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
    async function addSavedZone(name) {
        // user token
        const userToken = localStorage.getItem('token')
        // body of http request
        const bodyValues = JSON.stringify({
            "plot_name": name,
            "plot_zones": zipCodes.plot_zones

        })

        //send request to api to add zone and wait until it is completed
        await fetch("http://localhost:4000/plot/save", {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userToken}`
            },
            body: bodyValues

        }).then(async (response) => {
            if (!response.ok) {
                console.log('error in response to set zone')
                alert('errorSaving zones')
                return null;
            } else {
                // if response is ok then update the zones
                const saved = await getSavedZones();
                setSavedZones(saved)
            }
        })
    }

    /** 
    * Function which gets the saved zones from the Mongo database
    * @return {void}
    */
    async function getSavedZones() {
        // array to hold saved zones
        let saved = []

        await fetch("http://localhost:4000/user/plots", {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
        }).then((response) => {
            if (!response.ok) {
                alert('error retrieving zones')
            }
            return response.json();
        }).then((data) => {
            saved = data
        })

        return saved
    }

    // initially runs when rendering home page
    useEffect(() => {
        const fetchData = async () => {
            // variable used for saved plots in database
            let saved = await getSavedZones()
            // set the plots in the sider to be the saved plots
            setSavedZones(saved)
        }

        fetchData()
    }, []);


    // WAS USED FOR CHECKING PROPER UPDATES
    // useEffect(() => {
    //     console.log('home')
    //     console.log(zoneId)
    // }, [zoneId]);
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
                    <HeaderComponent handleSaveClick={handleSaveClick} savedZones={savedZones} />
                    <Menu theme="dark" mode="horizontal" />
                </Header>

                <Layout >
                    <Sider
                        width={"225"}
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
                            setZoneId={setZoneId}
                            setCurrentView={setCurrentView}

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
                            
                                <div style={{ display: "flex", justifyContent: "space-between", padding: "5px" }}>
                                    
                                        <div style={showComparison? { paddingRight: "5px", width: "50%" } : { paddingRight: "5px", width: "100%" } }>
                                            <Spin spinning={!computed} delay={500}>
                                                <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} zoneId={zoneId} setZipCodes={setZipCodes} setComputed={setComputed} />;
                                            </Spin>
                                        </div>
                                    
                                    
                                    <div style={ showComparison? { paddingRight: "5px", width: "50%" } : { paddingRight: "5px", width: "0%" } }>
                                        <Spin spinning={!computed2} delay={500}>
                                            {
                                                showComparison ? <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} zoneId={currentView} setZipCodes={setZipCodes} setComputed={setComputed2} /> : <></>
                                            }
                                        </Spin>
                                    </div>
                                </div>
                         
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    );
}
// External dependencies
import React, { useState, useEffect } from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Layout, Menu, theme, ConfigProvider, Spin } from 'antd';

// Components
import Map from './components/mapComponent';
import TextComponent from './components/textComponent'

// CSS
import './index.css';
import SiderComponent from "./components/sliderComponent";

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

    const [calculatedZone, setCalculatedZone] = useState([])

    // For view button to bring two maps back to one map.
    const [showMap, setShowMap] = useState(true);

    // Empty state
    const [holder, setHolder] = useState([])

    // For holding render state of map 1
    const [computed, setComputed] = useState(false)
    // For holding render state of heat map 1
    const [computedHeat, setComputedHeat] = useState(false)
    // For holding render state of map 2
    const [computed2, setComputed2] = useState(false)
    // For holding render state of heat map 2
    const [computedHeat2, setComputedHeat2] = useState(true)
    // For keeping track of selected algorithm, by default kmeans
    const [algorithm, setAlgorithm] = useState(1);
    // For keeping track of number of zones
    const [nrofzones, setNrofZones] = useState(1);
    // For setting name of zone in comparison (single)
    const [zoneName, setZoneName] = useState("Initial Zone");
    // For setting name of zone in comparison (single)
    const [zoneName2, setZoneName2] = useState("");


    // For radio
    const [value, setValue] = useState(1);
    const [zipCodes, setZipCodes] = useState([]);
    const [averageFuelCost, setAverageFuelCost] = useState("1.8");
    const [averageFuelUsage, setAverageFuelUsage] = useState("0.047");
    const onChange = (e) => {
        setValue(e.target.value);
    };

    // Track which zone is selected
    const [zoneId, setZoneId] = useState('initial');

    // For intensity
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
        console.log(savedZones);
        if (zipCodes.length === 0) {
            alert('nothing to save')
            return null;
        }
        const name = window.prompt('Enter a name for the save:');
        if (name) {
            for (let i = 0; i < savedZones.length; i++) {
                if (name === savedZones[i].user_plot_name) {
                    alert("you have already saved a zone under this name!")
                    return null;
                }
            }
            addSavedZone(name);
            setZipCodes([])
        }
    }

    // Keep track of the saved zones and update when needed
    const [savedZones, setSavedZones] = useState([]);

    // For zone synchronization
    const [currentView, setCurrentView] = useState('initial')

    /** 
    * When the user clicks the delete button, the zone will be deleted.
    * @return {void}
    */
    const handleDeleteZone = (id, name) => {
        const confirm = window.confirm('Confirm deletion of ' + name);
        // Make sure user actually wants to delete plot
        if (confirm) {
            deleteSavedZone(id)
        }
    };

    /** 
    * When the user has entered a name for the zone configuration they want to save, it is shown in the saved zones list.
    * @param {string} name - The name given to the saved zone configuration.
    * @return {void}
    */
    async function addSavedZone(name) {
        // User token
        const userToken = localStorage.getItem('token')
        // Body of http request
        const bodyValues = JSON.stringify({
            "plot_name": name,
            "plot_zones": zipCodes.plot_zones

        })

        // Send request to api to add zone and wait until it is completed
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
                // If response is ok then update the zones
                const saved = await getSavedZones();
                setSavedZones(saved)
            }
        })
    }

    /** 
    * Deletes plot from database.
    * @param {string} id - The id of the zone configuration to delete.
    * @return {void}
    */
    async function deleteSavedZone(id) {
        const userToken = localStorage.getItem('token')
        await fetch("http://localhost:4000/plot/" + id, {
            method: 'DELETE',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userToken}`
            }
        }).then(async (response) => {
            if (response.ok) {
                const saved = await getSavedZones();
                setSavedZones(saved)
                alert('Plot succesfully deleted')
                return null
            }
            alert('Could not delete plot, server error')
            return null
        })
    }

    /** 
    * Function which gets the saved zones from the Mongo database
    * @return {void}
    */
    async function getSavedZones() {
        // Array to hold saved zones
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

    // Initially runs when rendering home page
    useEffect(() => {
        const fetchData = async () => {
            // Variable used for saved plots in database
            let saved = await getSavedZones()
            // Set the plots in the sider to be the saved plots
            setSavedZones(saved)
        }

        fetchData()
    }, []);

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
                    <Sider width={"225"} style={{ background: colorBgContainer }}>
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
                            currentView={currentView}
                            algorithm={algorithm}
                            setAlgorithm={setAlgorithm}
                            setNrofZones={setNrofZones}
                            setZoneName={setZoneName}
                            setZoneName2={setZoneName2}
                            loadedHeat={computedHeat && computedHeat2}
                            averageFuelCost={averageFuelCost}
                            setAverageFuelCost={setAverageFuelCost}
                            averageFuelUsage={averageFuelUsage}
                            setAverageFuelUsage={setAverageFuelUsage}
                        />
                    </Sider>
                    <Layout style={{ padding: 30 }}>
                        <Content className="map" id="map" style={{ minHeight: '60vh', overflow: 'auto' }}>
                            <div style={{ display: "flex", justifyContent: "space-between", padding: "5px" }}>
                                <div style={showComparison ? { paddingRight: "5px", width: "50%" } : { paddingRight: "5px", width: "100%" }}>
                                    <Spin spinning={!computed || !computedHeat} delay={500} tip={
                                        (() => {
                                            if (!computed && !computedHeat) {
                                                return 'Loading Plot and Heatmap'
                                            }
                                            if (!computed) {
                                                return 'Loading Map'
                                            }
                                            if (!computedHeat) {
                                                return 'Loading Heatmap'
                                            }
                                        })()
                                    }>
                                        <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} zoneId={zoneId} setZipCodes={setZipCodes} setComputed={setComputed} algorithm={algorithm} nrofzones={nrofzones} setComputedHeat={setComputedHeat} setCalculatedZone={setCalculatedZone} />;
                                    </Spin>
                                </div>
                                <div style={showComparison ? { paddingRight: "5px", width: "50%" } : { paddingRight: "5px", width: "0%" }}>
                                    {
                                        showComparison ?
                                            <Spin spinning={!computed2 || !computedHeat2} delay={500} tip={
                                                (() => {
                                                    if (!computed2 && !computedHeat2) {
                                                        return 'Loading Plot and Heatmap'
                                                    }
                                                    if (!computed2) {
                                                        return 'Loading Map'
                                                    }
                                                    if (!computedHeat2) {
                                                        return 'Loading Heatmap'
                                                    }
                                                })()
                                            }>
                                                <Map intensity={intensity} value={value} onChange={onChange} onChangeNumber={onChangeNumber} zoneId={currentView} setZipCodes={setHolder} setComputed={setComputed2} setComputedHeat={setComputedHeat2} />
                                            </Spin>
                                            : <></>
                                    }

                                </div>
                            </div>
                            <div style={{ display: "flex", justifyContent: "space-between", padding: "5px" }}>
                                <div style={showComparison ? { paddingRight: "5px", width: "50%" } : { paddingRight: "5px", width: "100%" }}>
                                    <TextComponent zoneId={zoneId}
                                        zoneName={zoneName}
                                        calculatedZone={calculatedZone}
                                        averageFuelCost={averageFuelCost}
                                        setAverageFuelCost={setAverageFuelCost}
                                        averageFuelUsage={averageFuelUsage}
                                        setAverageFuelUsage={setAverageFuelUsage}
                                    />
                                </div>
                                <div style={showComparison ? { paddingRight: "5px", width: "50%" } : { paddingRight: "5px", width: "0%" }}>
                                    {
                                        showComparison ?
                                            <TextComponent zoneId={currentView} zoneName={zoneName2}
                                                averageFuelCost={averageFuelCost}
                                                setAverageFuelCost={setAverageFuelCost}
                                                averageFuelUsage={averageFuelUsage}
                                                setAverageFuelUsage={setAverageFuelUsage} />
                                            : <></>
                                    }
                                </div>
                            </div>
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    );
}
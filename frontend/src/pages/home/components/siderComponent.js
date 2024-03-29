// External dependencies
import React, { useRef, useState } from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Menu, Form, Input, Button, Radio, InputNumber, Tooltip } from 'antd';

// Icons
import { DeleteOutlined, } from '@ant-design/icons';

// CSS
import '../index.css';

const { SubMenu } = Menu;

// Input field function.
const ZoneSubMenu = ({ onSubmit, setZoneId, toggleMap, algorithm, setAlgorithm, setNrofZones, voronoi, setVoronoi, setAverageFuelUsage, averageFuelCost,averageFuelUsage, setAverageFuelCost}) => {

    const calculate = useRef(0)

    const onChangeAlgo = (e) => {
        setAlgorithm(e.target.value);
    };

    const onChangeVoronoi = (e) => {
        setVoronoi(e.target.value)
    };

    /**
     * Function to disable and enable the Voronoi option for algorithm
     * @param {Int} algorithm - What type of algorithm option was selected by user.
     * @return {Boolean}
     */
    function disabledVornoi(algorithm) {
        if (algorithm === 1){
            return false;
        }
        return true
    }

    const handleSubmit = (e) => {
        const isValid = true;
        if (isValid) {
            toggleMap()
            // Toggles the map to be one map.
            calculate.current += 1;
            setZoneId('calculate' + calculate.current.toString() )    
        }
    };

    // Input fields for the zone calculation (average fuel cost, average fuel usage of car, algorithm choices, number of zones).
    return (
        <Form onFinish={handleSubmit}>
            {/*Form item to hold user input value of average fuel cost*/}
            <Form.Item rules={[{ required: true }]}>
                <div style={{ padding: "0 5px" }}>
                    Average fuel cost
                    <Input
                        placeholder="input fuel cost"
                        type="number"
                        step="0.01"
                        min="0"
                        value={averageFuelCost}
                        onChange={(e) => setAverageFuelCost(e.target.value)}
                    />
                </div>
            </Form.Item>
            {/*Form item to hold user input value of average fuel usage of car*/}
            <Form.Item>
                <div style={{ padding: "0 5px" }}>
                    Average fuel usage of car
                    <Input
                        placeholder="input fuel usage of car"
                        type="number"
                        step="0.001"
                        min="0.0"
                        value={averageFuelUsage}
                        onChange={(e) => setAverageFuelUsage(e.target.value)}
                    />
                </div>
            </Form.Item>

            <div style={{ padding: "0 5px"}}>
                Options for Algorithms
            </div>
            <p> </p>

            {/*Radio group to hold user input for the type of algorithm*/}
            <div style={{ width: '100%', textAlign: 'center' }}>
                <Radio.Group onChange={onChangeAlgo} value={algorithm} style={{ paddingBottom: '10px' }}>
                    <Radio value={1}> Standard </Radio>
                    <Tooltip title="May take up to 10 minutes for result.">
                        <Radio value={2}> Advanced </Radio>
                    </Tooltip>  
                </Radio.Group>
            </div>
            <p> </p>

            <div style={{ padding: "0 5px"}}>
                Options for Standard Algorithm
            </div>
            <p> </p>

            {/*Radio group to hold user input for the options of standard algorithm*/}
            <div style={{ width: '100%', textAlign: 'center' }}>
                <Radio.Group onChange={onChangeVoronoi} value={voronoi} style={{ paddingBottom: '10px' }} disabled={disabledVornoi(algorithm)}>
                    <Radio value={false}> Regular </Radio>
                    <Tooltip title="This is only visual, zones will be saved as standard.">
                        <Radio value={true}> Expanded </Radio>
                    </Tooltip>
                </Radio.Group>
            </div>

            {/*Form item to hold user input for the number of desired zones in calculation*/}
            <Form.Item>
                <div style={{ padding: "0 5px" }}>
                    Number of Zones
                    <Input
                        placeholder="input desired number of zones"
                        type="number"
                        step="1"
                        min="1"

                        onChange={(e) => setNrofZones(e.target.value)}
                    />
                </div>
            </Form.Item>

            {/*Calculation button for requesting a new zone claulcation*/}
            <div style={{ textAlign: "center", padding: 5 }}>
                <Button
                    style={{ width: "95%"}}
                    type="primary"
                    htmlType="submit"
                    disabled={!averageFuelCost || !averageFuelUsage}
                >
                    Calculate
                </Button>
            </div>
        </Form>
    );
};

//main SiderComponent function that renders all elements
function SiderComponent(props) {
    const { values, setShowMap, setShowComparison, showMap, showComparison, onDeleteZone,
        setValue, setIntensity, savedZones, setZoneId, setCurrentView, algorithm,
        setAlgorithm, setNrofZones, currentView, setZoneName, setZoneName2, loadedHeat, setAverageFuelUsage,
        averageFuelCost,averageFuelUsage, setAverageFuelCost, voronoi, setVoronoi   } = props;

    //For toggling map configuration.
    const toggleMap = () => {
        if (showComparison && !showMap) {
            setShowComparison(false);
            setShowMap(true);
        }
    };

    //For toggling comparison configuration.
    const toggleComparison = () => {
        if (!showComparison && showMap) {
            setShowComparison(true);
            setShowMap(false);
        }
    };

    // For intensity change.
    const onChangeNumber = (e) => {
        setIntensity(e)
    }

    // For radio change.
    const onChange = (e) => {
        setValue(e.target.value);
    };

    return (
        <Menu
            mode="inline"
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            style={{
                height: '100%',
                borderRight: 0,
                overflowY: 'auto', overflowX: 'hidden' 
            }}
         
            selectable={false}
        >
            {/*Submenu to contain options and parameter input of the heatmap*/}
            <SubMenu key="sub3" data-testid="heatmap-btn" title="Heat map" style={{}}>
                <div style={{ width: '100%', textAlign: 'center' }}>
                    <Menu.Item key="5" style={{ padding: 0 }}>
                        <Radio.Group value={values} onChange={onChange} size='small' disabled={!loadedHeat}>
                            <Radio.Button data-testid="time-based" value={1} >
                                Time based
                            </Radio.Button>
                            <Radio.Button data-testid="location-based" value={2}>
                                Location based
                            </Radio.Button>
                        </Radio.Group>
                    </Menu.Item>
                </div>

                <Menu.Item key="6" style={{ height: "80px", padding: 0 }}>
                    <div style={{ textAlign: "center" }}>Intensity</div>
                    <div style={{ textAlign: "center" }}>
                        <InputNumber data-testid='intensity-input' min={1} max={1000} defaultValue={500} onChange={onChangeNumber} disabled={(()=>{
                            if (!loadedHeat || values ===1){return true}
                            return false
                        })()} />
                    </div>
                </Menu.Item>
            </SubMenu>

            {/*Zone submenu*/}
            <SubMenu key="sub4" title="Zones">
                <ZoneSubMenu
                    setZoneId={setZoneId}
                    toggleMap={toggleMap}
                    algorithm={algorithm}
                    setAlgorithm={setAlgorithm}
                    setNrofZones={setNrofZones}
                    averageFuelCost = {averageFuelCost}
                    setAverageFuelCost = {setAverageFuelCost}
                    averageFuelUsage = {averageFuelUsage}
                    setAverageFuelUsage = {setAverageFuelUsage}
                    voronoi={voronoi} setVoronoi={setVoronoi}
                />
            </SubMenu>

            {/*Saved zones submenu*/}
            <SubMenu key="sub2" title="Saved Zones" >
                <div style={{'max-height': '50vh', 'overflow': 'auto'}}>
                    {savedZones.map((zone) => (
                        <Menu.Item key={zone.user_plot_id} style={{ height: '80px', padding: 0 }}>
                            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                                <span style={{ paddingLeft: '10px' }}>{zone.user_plot_name}</span>
                                {zone.user_plot_name !== '' && (
                                    <Button
                                        style={{ float: 'right' }}
                                        onClick={() => {
                                            onDeleteZone(zone.user_plot_id, zone.user_plot_name);
                                        }}
                                    >
                                        <DeleteOutlined />
                                    </Button>
                                )}
                            </div>
                            <div style={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}>
                                <Button style={{ flex: 1, marginRight: '3px' }} onClick={() => {
                                    toggleMap()
                                    setCurrentView('')
                                    setZoneId(zone.user_plot_id)
                                    setZoneName(zone.user_plot_name)
                                }}>
                                    View
                                </Button>
                                <Button style={{ flex: 1, marginLeft: '3px' }} onClick={() =>{
                                    if (currentView === zone.user_plot_id){
                                        toggleMap()
                                        setCurrentView('')
                    
                                    } else {
                                        setCurrentView(zone.user_plot_id)
                                        toggleComparison()
                                        setZoneName2(zone.user_plot_name)
                                    }
                                }}>
                                    Compare
                                </Button>
                            </div>
                        </Menu.Item>
                    ))}
                </div>
            </SubMenu>
        </Menu>
    );
}

export default SiderComponent
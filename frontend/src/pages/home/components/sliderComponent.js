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

//Input field function -> later on add calculations, for now checks if the two fields are filled and if so, then the button is activated
const ZoneSubMenu = ({ onSubmit, setZoneId, toggleMap, algorithm, setAlgorithm, setNrofZones }) => {
    const [averageFuelCost, setAverageFuelCost] = useState("");
    const [averageFuelUsage, setAverageFuelUsage] = useState("");
  
    const calculate = useRef(0)

    const onChangeAlgo = (e) => {
        setAlgorithm(e.target.value);
    };

    const handleSubmit = (e) => {
        //e.preventDefault();
        //const isValid = onSubmit(averageFuelCost, averageFuelUsage);
        const isValid = true;
        if (isValid) {
            toggleMap()
            // toggles the map to be one map
            calculate.current += 1;
            setZoneId('calculate' + calculate.current.toString() )    
        }
    };

    // Input fields for the zone calculation (average fuel cost, average fuel usage of car).
    return (
        <Form onFinish={handleSubmit}>
            <Form.Item rules={[{ required: true }]}>
                <div style={{ padding: "0 5px" }}>
                    Average fuel cost
                    <Input
                        placeholder="input fuel cost"
                        type="number"
                        step="0.01"
                        value={averageFuelCost}
                        onChange={(e) => setAverageFuelCost(e.target.value)}
                    />
                </div>
            </Form.Item>
            <Form.Item>
                <div style={{ padding: "0 5px" }}>
                    Average fuel usage of car
                    <Input
                        placeholder="input fuel usage of car"
                        type="number"
                        step="0.01"
                        value={averageFuelUsage}
                        onChange={(e) => setAverageFuelUsage(e.target.value)}
                    />
                </div>
            </Form.Item>
            <div style={{ width: '100%', textAlign: 'center' }}>
                <Radio.Group onChange={onChangeAlgo} value={algorithm} style={{ paddingBottom: '10px' }}>
                    <Radio value={1}> KMeans </Radio>
                    <Tooltip title="May take up to 10 minutes for result.">
                        <Radio value={2}> Genetic </Radio>
                    </Tooltip>  
                </Radio.Group>
            </div>
            <Form.Item>
                <div style={{ padding: "0 5px" }}>
                    Nunber of Zones
                    <Input
                        placeholder="input desired number of zones"
                        type="number"
                        step="1"
                        
                        onChange={(e) => setNrofZones(e.target.value)}
                    />
                </div>
            </Form.Item>
                
            
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

function SiderComponent(props) {
    const { values, setShowMap, setShowComparison, showMap, showComparison, onDeleteZone, setValue, setIntensity, savedZones, setZoneId, setCurrentView, algorithm, setAlgorithm, setNrofZones } = props;


    const toggleMap = () => {
        if (showComparison && !showMap) {
            setShowComparison(false);
            setShowMap(true);
        }
    };

    const toggleComparison = () => {
        if (!showComparison && showMap) {
            setShowComparison(true);
            setShowMap(false);
        }
    };

    // for intensity change
    const onChangeNumber = (e) => {
        setIntensity(e)
    }

    // for radio change
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
            }}
        >

            <SubMenu key="sub3" title="Heat map" style={{}}>
                <div style={{ width: '100%', textAlign: 'center' }}>
                    <Menu.Item key="5" style={{ padding: 0 }}>
                        <Radio.Group value={values} onChange={onChange} size='small'  >
                            <Radio.Button value={1}>
                                Time based
                            </Radio.Button>
                            <Radio.Button value={2}>
                                Location based
                            </Radio.Button>
                        </Radio.Group>
                    </Menu.Item>
                </div>

                <Menu.Item key="6" style={{ height: "80px", padding: 0 }}>
                    <div style={{ textAlign: "center" }}>Intensity</div>
                    <div style={{ textAlign: "center" }}>
                        <InputNumber min={1} max={1000} defaultValue={500} onChange={onChangeNumber} disabled={values === 1 ? true : false} />
                    </div>
                </Menu.Item>
            </SubMenu>

            <SubMenu key="sub4" title="Zones">
                <ZoneSubMenu setZoneId={setZoneId} toggleMap={toggleMap} algorithm={algorithm} setAlgorithm={setAlgorithm} setNrofZones={setNrofZones} />
            </SubMenu>

            <SubMenu key="sub2" title="Saved Zones" style={{'max-height': '30vh', 'overflow': 'auto'}}>
                {savedZones.map((zone) => (
                    <Menu.Item key={zone.user_plot_id} style={{ height: '80px', padding: 0 }}>
                        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                            <span style={{ paddingLeft: '10px' }}>{zone.user_plot_name}</span>
                            {zone.user_plot_name !== 'Initial Zone' && (
                                <Button
                                    style={{ float: 'right' }}
                                    onClick={() => {
                                        onDeleteZone(zone.key);
                                    }}
                                >
                                    <DeleteOutlined />
                                </Button>
                            )}
                        </div>
                        <div style={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}>
                            <Button style={{ flex: 1, marginRight: '3px' }} onClick={() => {
                                toggleMap()
                                setZoneId(zone.user_plot_id)
                            }}>
                                View
                            </Button>
                            <Button style={{ flex: 1, marginLeft: '3px' }} onClick={() =>{
                                setCurrentView(zone.user_plot_id)
                                toggleComparison()
                            }}>
                                Compare
                            </Button>
                        </div>
                    </Menu.Item>
                ))}
            </SubMenu>
        </Menu>
    );
}

export default SiderComponent
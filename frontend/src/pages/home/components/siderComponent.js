// External dependencies
import React, { useState, useEffect } from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Layout, Menu, theme, Form, Input, Button, Dropdown, Space, ConfigProvider, Radio, InputNumber } from 'antd';

// Icons
import { UserOutlined, DeleteOutlined, SaveOutlined } from '@ant-design/icons';
// CSS
import '../index.css';
import PolygonVis from "./polygonComponents";
const { SubMenu } = Menu;
//Input field function -> later on add calculations, for now checks if the two fields are filled and if so, then the button is activated
const ZoneSubMenu = ({ onSubmit }) => {
    const [averageFuelCost, setAverageFuelCost] = useState("");
    const [averageFuelUsage, setAverageFuelUsage] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        const isValid = onSubmit(averageFuelCost, averageFuelUsage);
        if (isValid) {
            // Add calculations
        }
    };
    // Input fields for the zone calculation (average fuel cost, average fuel usage of car).
    return (
        <Form onFinish={handleSubmit}>
            <Form.Item rules={[{ required: true }]}>
                Average fuel cost
                <Input
                    placeholder="input fuel cost"
                    type="number"
                    step="0.01"
                    value={averageFuelCost}
                    onChange={(e) => setAverageFuelCost(e.target.value)}
                />
            </Form.Item>
            <Form.Item>
                Average fuel usage of car
                <Input
                    placeholder="input fuel usage of car"
                    type="number"
                    step="0.01"
                    value={averageFuelUsage}
                    onChange={(e) => setAverageFuelUsage(e.target.value)}
                />
            </Form.Item>
            <Button
                style={{ width: "100%" }}
                type="primary"
                htmlType="submit"
                disabled={!averageFuelCost || !averageFuelUsage}
            >
                Calculate
            </Button>
        </Form>
    );
};
function SiderComponent(props) {
    const {value, intensity, onChange, onChangeNumber, setShowMap, setShowComparison, showMap, showComparison, onDeleteZone} = props;

    const toggleMap = () => {
        setShowMap(!showMap);
        setShowComparison(false);
    };

    const toggleComparison = () => {
        setShowComparison(!showComparison);
        setShowMap(false);
    };

    const [savedZones, setSavedZones] = useState([
        { key: 'saved-initial', name: 'Initial Zone' },
    ]);


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
                <Menu.Item key="5" style={{ padding: 0 }}>
                    <Radio.Group value={value} onChange={onChange} size='small'  >
                        <Radio.Button value={1}>
                            Time based
                        </Radio.Button>
                        <Radio.Button value={2}>
                            Location based
                        </Radio.Button>
                    </Radio.Group>
                </Menu.Item>
                <Menu.Item key="6" style={{ height: "80px", padding: 0 }}>
                    <div style={{ textAlign: "center" }}>Intensity</div>
                    <div style={{ paddingLeft: 50 }}>
                        <InputNumber min={1} max={1000} defaultValue={500} onChange={onChangeNumber} disabled={value === 1 ? true : false} />
                    </div>
                </Menu.Item>
            </SubMenu>
            <SubMenu key="sub4" title="Zones">
                <ZoneSubMenu />
            </SubMenu>
            <SubMenu key="sub2" title="Saved Zones">
                {savedZones.map((zone) => (
                    <Menu.Item key={zone.key} style={{ height: '80px', padding: 0 }}>
                        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                            <span style={{ paddingLeft: '10px' }}>{zone.name}</span>
                            {zone.name !== 'Initial Zone' && (
                                <Button
                                    style={{ float: 'right' }}
                                    onClick={() => {
                                        localStorage.removeItem(zone.key);
                                        onDeleteZone(zone.key);
                                    }}
                                >
                                    <DeleteOutlined />
                                </Button>
                            )}
                        </div>
                        <div style={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}>
                            <Button style={{ flex: 1, marginRight: '3px' }} onClick={toggleMap}>
                                View
                            </Button>
                            <Button style={{ flex: 1, marginLeft: '3px' }} onClick={toggleComparison}>
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
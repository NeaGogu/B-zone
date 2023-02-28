import { LaptopOutlined, NotificationOutlined, UserOutlined,DownOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, Button, Dropdown, Space} from 'antd';
import React from 'react';
import DropdownButton from "antd/es/dropdown/dropdown-button";
function getItem(label, key, icon, children, type) {
    return {
        key,
        icon,
        children,
        label,
        type,
    };
}
const {SubMenu} = Menu;
const { Header, Content, Sider } = Layout;
const item = [
    getItem('Saved zones', 'sub1', null , [
        getItem('Initial Zone', 'g1', null,null, 'group'),
        getItem('Saved Zone 1', 'g2', null, null, 'group'),
    ]),
    getItem('Filter', 'sub2' , null, null)
];

const items = [
    {
        key: '1',
        label: (
            <a target="_blank" rel="noopener noreferrer">
                email
            </a>
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





export default function Home(){
    const {
        token: { colorBgContainer },
    } = theme.useToken();
    return(
        <Layout>
            <Header className="header">
                <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                    className="logo"

                    <Dropdown
                        menu={{
                            items,
                        }}
                    >
                        <a onClick={(e) => e.preventDefault()}>
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


                {/*<Dropdown*/}
                {/*    menu={{*/}
                {/*        items,*/}
                {/*    }}*/}
                {/*>*/}
                {/*    /!*<a onClick={(e) => e.preventDefault()}>*!/*/}
                {/*    /!*    <Space>*!/*/}
                {/*    /!*        Hover me*!/*/}
                {/*    /!*    </Space>*!/*/}
                {/*    /!*</a>*!/*/}
                {/*</Dropdown>*/}




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
                        <SubMenu key="sub1"  title="Input field">
                        </SubMenu>
                        <SubMenu key="sub2"  title="Saved Zones" icon={<NotificationOutlined/>}>
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

                    <Content
                        style={{
                            padding: 24,
                            margin: 0,
                            minHeight: 280,
                            background: colorBgContainer,
                        }}
                    >
                        Content
                    </Content>
                </Layout>
            </Layout>
        </Layout>
    );
}
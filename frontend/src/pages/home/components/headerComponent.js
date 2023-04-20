// External dependencies
import React from 'react';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";
import { Button, Dropdown, Space } from 'antd';

// Icons
import { UserOutlined, SaveOutlined } from '@ant-design/icons';

// CSS
import '../index.css';

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
            <button data-testid="btn" style={{ color: '#ffd369', textDecoration: 'underline', backgroundColor: 'transparent', border: 'none', cursor: 'pointer' }} onClick={() => signOut()}>
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
    //get token from local storage, instantiate boolean variable for validity
    const token = localStorage.getItem('token');
    var valid;

    //send sign-out API request to Bumbal API, GET method, user token encoded in both URL and Bearer's Authorization
    fetch(`https://sep202302.bumbal.eu/api/v2/authenticate/sign-out?token=${encodeURIComponent(token)}`, {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
        }
    })
        //if response is ok, set validity of user signing out to true
        .then((response) => {
            if (response.ok) {
                valid = true;

            }
        //if valid response, clear the local storage and reroute user to login page. otherwise, notify user of error
        }).then((data) => {
        if (valid) {
            localStorage.clear();
            alert('You have been logged out!')
            window.location.reload()
        } else {
            alert("You could not be logged out! Please try again later.")
        }
    })
}

// The main header component, shows the logo, the B-Zone title, the save button and the user credentials menu.
// Includes B-Zone logo, B-Zone title, and Dropdown menu for viewing credentials + signing out
const HeaderComponent = ({handleSaveClick}) => {
    return (
        <div style={{ display: 'flex', justifyContent: 'flex-end' }} >
            <div style={{ margin: 'auto', width: '100%' }} >
                <img src='./b-zone-logo.png' style={{ display: 'block', width: '11%' }} alt='logo' />
            </div>
            <div className="headerText">
                B-Zone
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
                        <Space>
                            <UserOutlined style={{ verticalAlign: 'middle', color: '#ffd369' }} />
                            <button data-testid= 'dropdown-btn' type="button" onClick={signOut} style={{ color: '#ffd369', backgroundColor: 'transparent', border: 'none', textDecoration: 'underline', cursor: 'pointer' }}>
                            </button>
                        </Space>
                </Dropdown>
            </Space>

        </div>
    );
};

export default HeaderComponent
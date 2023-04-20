import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Form, Input, Card, ConfigProvider, theme } from 'antd';
import './index.css'

// Used for dark mode
const { darkAlgorithm } = theme;

export default function Login() {
    // A boolean variable that keeps track of whether the user has been verified or not.
    var verified;

    /** 
    * Authenticates the user's credentials against the server API. 
    * @param {string} user The email address of the user to authenticate.
    * @param {string} passw The password of user to authenticate.
    * @return {undefined} 
    */
    async function signIn(user, passw) {
        //given a username and password, use Fetch API to sign the user in through Bumbal
        //signing in through the Bumbal API requires a POST method, with three values in a JSON body: email, password, and return_jwt
        fetch("https://sep202302.bumbal.eu/api/v2/authenticate/sign-in", {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "email": user,
                "password": passw,
                "return_jwt": true
            })
        })
            //verify if the response from Bumbal is ok or not
            .then((response) => {
                if (!response.ok) {
                    verified = false;
                } else {
                    verified = true;
                }
                return response.json();
            })
            //deal with the received data from Bumbal
            .then((data) => {
                //if user is not correctly verified by Bumbal API, alert them
                if (!verified) {
                    alert('Invalid Credentials')
                } else {
                    //otherwise, synchronize the user's plots with B-Zone's backend
                    //syncing plots in B-Zone's backend requires a PUT method, and to include the user's session token as a Bearer's Authentication
                    fetch("http://localhost:4000/plot/sync",{
                        method: 'PUT',
                        headers: {
                            'Accept': 'application/json',
                            'Content-Type': 'application/json',
                            'Authorization': `Bearer ${data.token}`
                        },
                        //await second response from B-Zone backend
                    }).then((response2) =>{
                        // If sync was successful then redirect by reloading window
                        if(response2.ok){
                            //set all token and user credentials to local storage
                            localStorage.setItem('token', data.token)
                            localStorage.setItem('email', user)
                            localStorage.setItem('id', data.user.id)
                            //reloads window from login page to home
                            window.location.reload()
                        //otherwise notify the user to error
                        } else {
                            alert('server error when syncing zones, please try again')
                        }
                    }).catch((error)=>{
                        alert('server error when syncing zones, please try again')

                    })                    
                }
            })
        return undefined;
    }

    /**
    Handles the form submission event and triggers the signIn() function with the entered email and password.
    @param {object} values - An object containing the entered form values.
    @return {undefined}
    */
    const onFinish = (values) => {
        var email = Object.values(values)[0];
        var pass = Object.values(values)[1];

        //passes values to the signIn function
        signIn(email, pass)
    }

    return (
        // Parent component, essentially acts as a background. 
        <div className='page'>
            <ConfigProvider
                // The theme we use.
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
                {/* Form Implementation */}
                <Card className='form-card' >
                    <Form name='login-form' className='login-form' onFinish={onFinish}>
                        {/* Form item which includes the email input */}
                        {/* Email is a required input from the user, and an error message is given if it is blank */}
                        <Form.Item name='Email' rules={[{ required: true, message: 'Email is required!' }]}>
                            <Input data-testid="username-input" placeholder='Email' prefix={<UserOutlined />} />
                        </Form.Item>

                        {/* Form item which includes the password input of the user */}
                        {/* Password is a required input from the user, and an error message is given if it is blank */}
                        <Form.Item name='Password' rules={[{ required: true, message: 'Password is required!' }]}>
                            <Input.Password data-testid="password-input" placeholder='Password' prefix={<LockOutlined />} />
                        </Form.Item>

                        {/* Form item which includes the login button for the user */}
                        <Form.Item >
                            <Button data-testid="button" type='primary' className='login-button' htmlType='submit'>
                                Log in
                            </Button>

                            {/* Form item which includes the register to Bumbal link redirect, since users of B-Zone need to already have a Bumbal account */}
                            <a href='https://support.bumbal.eu/en/register' style={{ marginLeft: 35 }}> or Register!</a>
                        </Form.Item>
                    </Form>
                </Card>
            </ConfigProvider>

            {/* Bee Image (maybe use a background pre-made instead of an svg as component) */}
            <img src='https://upload.wikimedia.org/wikipedia/commons/c/c3/Bee_-_The_Noun_Project.svg' alt='Bee SVG' className='bee' />
        </div>
    );
}
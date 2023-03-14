import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Form, Input, Card, ConfigProvider, theme, Alert } from 'antd';
import { useNavigate } from "react-router-dom";
import './index.css'

// used for dark mdoe
const { darkAlgorithm } = theme;

export default function Login() {
    /**
    * A hook from React Router that allows for navigation within the app.
    @type {Function}
    */
    const navigate = useNavigate();

    /**
    * A boolean variable that keeps track of whether the user has been verified or not
     @type {boolean}
    */
    var verified;

    /** 
    * Authenticates the user's credentials against the server API. 
    * @param {string} user The email address of the user to authenticate.
    * @param {string} passw The password of user to authenticate.
    * @return {undefined} 
    */
    function signIn(user, passw) {
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
            .then((response) => {
                
                if (!response.ok){
                    verified = false;
                }
                else {
                    verified = true;
                }
                return response.json();


                
              })
            .then((data) => {
                if (!verified) {
                    alert('Invalid Credentials')
                }
                else{
                    localStorage.setItem('token', data.token)
                    localStorage.setItem('email', user)
                    localStorage.setItem('id', data.user.id)
                    window.location.reload()
                    navigate('/home')
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
        signIn(email, pass)
    }

    return (
        // parent component, essentially acts as a background 
        <div className='page'>
            <ConfigProvider
                // theme which we should use
                theme={{
                    token: {
                        colorPrimary: "#ffd369",
                        colorBgBase: "#222831",
                        colorPrimaryBg: "#393E46",
                        colorTextBase: "#eeeeee"
                    },
                    // renders antd components with dark mode
                    algorithm: darkAlgorithm
                }}
            >   
                {/* Form Implementation */}
                <Card className='form-card' >
                    <Form name='login-form' className='login-form' onFinish={onFinish}>
                        <Form.Item name='Email' rules={[{ required: true, message: 'Email is required!' }]}>
                            <Input placeholder='Email' prefix={<UserOutlined />} />
                        </Form.Item>

                        <Form.Item name='Password' rules={[{ required: true, message: 'Password is required!' }]}>
                            <Input.Password placeholder='Password' prefix={<LockOutlined />} />
                        </Form.Item>

                        <Form.Item >
                            <Button type='primary' className='login-button' htmlType='submit'>
                                Log in
                            </Button>
                            <a href='https://support.bumbal.eu/en/register' style={{ marginLeft: 35 }}> or Register!</a>
                        </Form.Item>
                    </Form>
                </Card>
            </ConfigProvider>
            

            {/* Bee Image (maybe use a background pre made instead of an svg as component) */}
            <img src='https://upload.wikimedia.org/wikipedia/commons/c/c3/Bee_-_The_Noun_Project.svg' alt='Bee SVG' className='bee' />
        </div>
    );
            

}

/**
Renders a card component containing a login form.
@param {string} className - The class name for the card component.
@param {string} name - The name of the form.
@param {function} onFinish - The function to call when the form is submitted.
@returns {JSX.Element} - The card component with a login form.
*/
// couldnt put this in line 97 above {/* Form Implementation */} gave me error.

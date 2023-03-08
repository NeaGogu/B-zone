import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Form, Input, Card, ConfigProvider, theme, Alert } from 'antd';
import { useNavigate } from "react-router-dom";
import './index.css'

// used for dark mdoe
const { darkAlgorithm } = theme;

export default function Login() {

    const navigate = useNavigate();
    var verified;

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

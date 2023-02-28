import { CopyOutlined } from '@ant-design/icons';
import {
    Button,
    Input,
} from 'antd';

function signIn(user, passw) {
    console.log("Sign in button clicked")
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
        .then((response) => console.log(response.json()))
        .then((data) => console.log(data));
    return undefined;
}

export default function Login(){
    const username = "sep@bumbal.eu"
    const password = "cW$#Qbph0524";
    const token = "";

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
            .then((response) => response.json())
            .then((data) => console.log(data))
        return undefined;
    }

    return (<Input.Group compact>
        <Input
            style={{
                width: 'calc(100% - 200px)',
            }}
            defaultValue="https://ant.design"
        />
        <Button type="primary" onClick={() => signIn(username, password)}>Submit</Button>
    </Input.Group>);

}

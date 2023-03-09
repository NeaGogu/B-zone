import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import Home from './pages/home'
import Login from './pages/login'


async function isAuthenticated() {
  var token = localStorage.getItem("token")
  if (token === null) {
    return false;
  }

  var result = await isTokenValid(token);

  return result;

}

async function isTokenValid(token) {
  const requestOptions = {
      method: 'GET',
      headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`, // add token to Bearer Authorization when sending GET signOut request
      }
     
  };
  const response = await fetch('https://sep202302.bumbal.eu/api/v2/authenticate/check-token', requestOptions);
  
  console.log('isTokenValid response ' + (response.status === 200)) 

  if ((response.status === 200)) {
    return true;
  } else {
    return false;
  }
}

function App() {
  var authenticated
  if ((localStorage.getItem("token") === null)) {
    authenticated = false;
  } else {
    authenticated = true;
  }
  console.log(authenticated)

  return (

      <Router>
        <Routes>
          <Route path='/home' element={authenticated ? <Home /> : <Navigate replace to={'/'}/>}  />
          <Route path='/login' element={authenticated ? <Navigate replace to={'/'}/> : <Login/>}/>
          <Route path='/' element={

            authenticated ? (<Navigate to="/home" replace />) :
                (<Navigate to="/login" replace />)

          } />
        </Routes>
      </Router>
  );
}

export default App;
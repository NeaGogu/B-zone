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

/**
 * The main component of the application.
 * @return {JSX.Element} The JSX element representing the application.
 */

function App() {
  
  var authenticated
  // check whether a token exists in local storage.
  if ((localStorage.getItem("token") === null)) {
    authenticated = false;
  } else {
    authenticated = true;
  }

  return (
    // If user is authenticated navigates to home page otherwise navigates to login page.
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



/**
 * The Navigate component from the React Router library.
 *
 * @typedef {Object} Navigate
 * @property {string} to The URL to navigate to.
 * @property {boolean} replace Whether to replace the current URL in the history stack.
 */

/**
 * The Router component from the React Router library.
 *
 * @typedef {Object} Router
 * @property {function} Router The Router component.
 */

/**
 * The Route component from the React Router library.
 *
 * @typedef {Object} Route
 * @property {string} path The URL path to match.
 * @property {JSX.Element} element The JSX element to render when the path matches.
 */

/**
 * The Routes component from the React Router library.
 *
 * @typedef {Object} Routes
 * @property {JSX.Element} children The child elements to render.
 */

/**
 * The localStorage object from the API.
 *
 * @typedef {Object} localStorage
 * @property {function} getItem The getItem method to retrieve an item(token) from local storage.
 */ 

export default App;
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import Home from './pages/home'
import Login from './pages/login'

/**
 * The main component of the application.
 * @return {JSX.Element} The JSX element representing the application.
 */
function App() {
  
  var authenticated
  // Check whether a token exists in local storage.
  if ((localStorage.getItem("token") === null)) {
    authenticated = false;
  } else {
    authenticated = true;
  }

  return (
    // If user is authenticated, navigates to home page, otherwise navigates to login page.
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
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import Home from './pages/home'
import Login from './pages/login'

/**
 * The main component of the application.
 * @return {JSX.Element} The JSX element representing the application.
 */
function App() {

  //boolean variable to check whether a user token exists or not
  var authenticated

  // Check whether a token exists in local storage
  if ((localStorage.getItem("token") === null)) {
    //if there is no token, the user must not be authenticated or must be logged out
    authenticated = false;
  } else {
    //otherwise, user could possible me authenticated. validity of token checked further in Home
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

export default App;
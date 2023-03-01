import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import Home from './pages/home'
import Login from './pages/login'

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
        <Route path='/home' element={<Home />} />

        <Route path='/login' element={<Login />} />

        <Route path='/' element={

          authenticated ? (<Navigate to="/home" replace />) :
            (<Navigate to="/login" replace />)

        } />
      </Routes>
    </Router>
  );
}

export default App;
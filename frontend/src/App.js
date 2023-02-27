import {BrowserRouter as Router, Routes, Route, Navigate} from 'react-router-dom' 
import Login from './pages/login/index.js'
import Home from './pages/home/index.js'
import './App.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path='/login' element={<Login/>} />
        <Route path='/home' element={<Home/>}/>
      </Routes>
    </Router>
  );
}

export default App;

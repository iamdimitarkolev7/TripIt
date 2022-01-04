import React from "react";
import { Routes, Route } from "react-router-dom";
import NavBar from './components/common/navigation/NavBar';
import isLoggedIn from './utils/isLoggedIn';
import Home from "./components/common/home/Home";
import Register from "./components/user/Register";
import Login from "./components/user/Login";
import Profile from "./components/user/Profile";

import './App.css';

function App() {
  return (
    <div className="App">
        <NavBar isLoggedIn={isLoggedIn} />
        <Routes>
          <Route exact path="/" element={<Home/>} />
          <Route path="/register" element={<Register/>} />
          <Route path="/login" element={<Login/>} />
          <Route path="/profile" element={<Profile/>} />
        </Routes>
    </div>
  );
}

export default App;

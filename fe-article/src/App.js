import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './components/Home';
import Login from './components/Login';
import AdminDashboard from './components/AdminDashboard';
import CreateArticle from './components/CreateArticle';
import EditArticle from './components/EditArticle';
import PrivateRoute from './components/PrivateRoute';
import ArticleDetails from './components/ArticleDetail';
import Register from './components/Register';
import ForgotPassword from './components/ForgotPassword';
import ResetPassword from './components/ResetPassword';
//change something
function App() {
  const [user, setUser] = useState(JSON.parse(localStorage.getItem('user')));

  useEffect(() => {
    const loggedInUser = localStorage.getItem('user');
    if (loggedInUser) {
      setUser(JSON.parse(loggedInUser));
    }
  }, []);

  const handleLogin = async (email, password) => {
    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/user/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email,
          password,
        }),
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }
      const data = await response.json();
      const userData = data.data.accessToken;
      localStorage.setItem('user', JSON.stringify(data.data));
      setUser(userData);
    } catch (error) {
      console.error('Error logging in:', error);
      alert('Login failed: ' + error.message);
    }
  };

  const handleRegister = async ({ name, email, password, role }) => {
    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/user/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name,
          email,
          password,
          role,
        }),
      });
  
      if (!response.ok) {
        throw new Error('Registration failed');
      }
  
      alert('Registration successful! Please log in.');
    } catch (error) {
      console.error('Error registering:', error);
      alert('Registration failed: ' + error.message);
    }
  };
  

  const handleLogout = () => {
    localStorage.removeItem('user');
    setUser(null); 
  };

  return (
    <Router>
      <div className='container-fluid'>
      <Navbar user={user} setUser={handleLogout} />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path='/articles/:id' element={<ArticleDetails />} />
        <Route path="/login" element={<Login onLogin={handleLogin} />} />
        <Route path="/register" element={<Register onRegister={handleRegister} />} />
        <Route path="/forgot-password" element={<ForgotPassword />} />
        <Route path="/reset-password/:token" element={<ResetPassword />} />

        <Route path="/admin" element={<PrivateRoute user={user}><AdminDashboard /></PrivateRoute>} />
        <Route path="/create-article" element={<PrivateRoute user={user}><CreateArticle /></PrivateRoute>} />
        <Route path="/edit-article/:id" element={<PrivateRoute user={user}><EditArticle /></PrivateRoute>} />

        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
      </div>
    </Router>
  );
}

export default App;

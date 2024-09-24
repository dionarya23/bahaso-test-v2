import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';

const Login = ({ onLogin }) => {
  const [credentials, setCredentials] = useState({ email: '', password: '' });
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCredentials(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      if (!credentials.email || !credentials.password) {
        setError('Both email and password are required');
        return;
      }
      await onLogin(credentials.email, credentials.password);
      navigate('/admin');
    } catch (error) {
      console.error('Login error:', error);
      setError('Invalid credentials, please try again.');
    }
  };

  return (
    <div className="container mt-5">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <h2 className="mb-3">Login</h2>
          <form onSubmit={handleSubmit}>
            <div className="mb-3">
              <label htmlFor="email" className="form-label">Email address</label>
              <input
                type="email"
                className="form-control"
                id="email"
                name="email"
                value={credentials.email}
                onChange={handleChange}
                required
              />
            </div>
            <div className="mb-3">
              <label htmlFor="password" className="form-label">Password</label>
              <input
                type="password"
                className="form-control"
                id="password"
                name="password"
                value={credentials.password}
                onChange={handleChange}
                required
              />
            </div>
            {error && <div className="alert alert-danger">{error}</div>}
            <button type="submit" className="btn btn-primary">Login</button>
          </form>
          <div className="mt-3">
            <p>Don't have an account? <Link to="/register">Register here</Link></p>
            <p>Forgot your password? <Link to="/forgot-password">Reset it here</Link></p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;

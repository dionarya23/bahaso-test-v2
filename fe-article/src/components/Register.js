import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Register({ onRegister }) {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [role, setRole] = useState('visitor');
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    if (password !== confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    try {
      await onRegister({ name, email, password, role });
      navigate('/login');
    } catch (err) {
      setError('Registration failed: ' + err.message);
    }
  };

  return (
    <div className="container mt-5">
      <h2>Register</h2>
      {error && <p className="alert alert-danger">{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="name" className="form-label">Full Name</label>
          <input
            type="text"
            className="form-control"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="email" className="form-label">Email address</label>
          <input
            type="email"
            className="form-control"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="password" className="form-label">Password</label>
          <input
            type="password"
            className="form-control"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="confirmPassword" className="form-label">Confirm Password</label>
          <input
            type="password"
            className="form-control"
            id="confirmPassword"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="role" className="form-label">Role</label>
          <select
            className="form-control"
            id="role"
            value={role}
            onChange={(e) => setRole(e.target.value)}
          >
            <option value="author">Author</option>
            <option value="visitor">Visitor</option>
          </select>
        </div>
        <button type="submit" className="btn btn-primary">Register</button>
      </form>
    </div>
  );
}

export default Register;
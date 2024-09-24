import React, { useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

const ResetPassword = () => {
  const [passwords, setPasswords] = useState({ newPassword: '', confirmPassword: '' });
  const [error, setError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  const navigate = useNavigate();
  const { token } = useParams();

  const handleSubmit = async (event) => {
    event.preventDefault();
    
    // Clear previous messages
    setError('');
    setSuccessMessage('');

    if (passwords.newPassword !== passwords.confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/user/change-password`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
          token,
          password: passwords.newPassword,
        }),
      });

      if (response.ok) {
        setSuccessMessage('Password has been reset successfully!');
        setTimeout(() => {
          navigate('/login'); // Redirect to login after success
        }, 2000);
      } else {
        const data = await response.json();
        setError(data.message || 'Failed to reset password');
      }
    } catch (err) {
      setError('An error occurred while resetting the password. Please try again.');
    }
  };

  return (
    <div className="container mt-5">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card">
            <div className="card-body">
              <h2 className="text-center">Reset Password</h2>
              <form onSubmit={handleSubmit}>
                <div className="form-group mb-3">
                  <label htmlFor="newPassword">New Password:</label>
                  <input
                    type="password"
                    id="newPassword"
                    className="form-control"
                    value={passwords.newPassword}
                    onChange={(e) => setPasswords({ ...passwords, newPassword: e.target.value })}
                    required
                  />
                </div>
                <div className="form-group mb-3">
                  <label htmlFor="confirmPassword">Confirm Password:</label>
                  <input
                    type="password"
                    id="confirmPassword"
                    className="form-control"
                    value={passwords.confirmPassword}
                    onChange={(e) => setPasswords({ ...passwords, confirmPassword: e.target.value })}
                    required
                  />
                </div>
                {error && <div className="alert alert-danger">{error}</div>}
                {successMessage && <div className="alert alert-success">{successMessage}</div>}
                <button type="submit" className="btn btn-primary btn-block">Submit</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ResetPassword;

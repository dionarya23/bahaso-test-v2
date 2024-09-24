import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

const NavBar = ({ user, setUser }) => {
  const navigate = useNavigate();

  const handleLogout = (event) => {
    event.preventDefault();
    setUser(null);
    localStorage.removeItem('user'); 
    navigate('/login');
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <Link className="navbar-brand" to="/">Blog</Link>
      <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>
      <div className="collapse navbar-collapse" id="navbarNav">
        <ul className="navbar-nav ml-auto">
          {user ? (
            <>
              <li className="nav-item">
                <Link className="nav-link" to="/admin">Admin</Link>
              </li>
              <li className="nav-item">
                <button className="nav-link btn btn-link" onClick={handleLogout}>Logout</button>  {/* Changed to a button to handle logout */}
              </li>
            </>
          ) : (
            <li className="nav-item">
              <Link className="nav-link" to="/login">Login</Link>
            </li>
          )}
        </ul>
      </div>
    </nav>
  );
};

export default NavBar;

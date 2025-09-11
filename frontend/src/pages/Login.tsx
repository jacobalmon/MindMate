import { useState } from 'react';
import '../styles/Login.css';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleLogin = () => {
    if (!email || !password) {
      setMessage('Please enter email and password');
      return;
    }
    setMessage('Logging in...');
    // TODO: Call your backend API here
  };

  const handleForgotPassword = () => {
    alert('Redirect to password reset page (implement this)');
  };

  const handleSignUp = () => {
    alert('Redirect to sign up page (implement this)');
  }

  return (
    <div className="login-container">
      {/* Top header */}
      <div className="top-card">
        <h2>MindMate</h2>
      </div>

      
      {/* Center Login Card */}
      <div className="login-box">
        <h1>Login</h1>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button onClick={handleLogin}>Login</button>

        {/* Links container */}
        <div className="links-container">
          <p className="forgot" onClick={handleForgotPassword}>
            Forgot password?
          </p>
          <p className="signup" onClick={handleSignUp}>
            Don't have an account? Sign Up
          </p>
        </div>

        {message && <p className="message">{message}</p>}
      </div>

    </div>
  );
}

export default Login;

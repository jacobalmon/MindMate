import { useState } from 'react';
import '../styles/Login.css';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleLogin = async () => {
    if (!email || !password) {
      setMessage('Please enter email and password');
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage('Login Successful');
        localStorage.setItem('token', data.token);
      } else {
        setMessage(data.error || 'Login Failed');
      }
    } catch (error) {
      console.error('Error Logging In', error);
      setMessage('Server Error, Please Try Again Later.');
    }
  };

  const handleForgotPassword = () => {
    navigate('/forgot-password'); 
  };

  const handleSignUp = () => {
    navigate('/signup'); 
  };

  return (
    <div className="login-container">
      <div className="top-card">
        <h2>MindMate</h2>
      </div>

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

        <p className="forgot" onClick={handleForgotPassword}>
          Forgot password?
        </p>
        <p className="signup" onClick={handleSignUp}>
          Don&apos;t have an account? Sign Up
        </p>

        {message && <p className="message">{message}</p>}
      </div>
    </div>
  );
}

export default Login;

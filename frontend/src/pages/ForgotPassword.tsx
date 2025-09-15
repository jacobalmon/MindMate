import { useState } from 'react';
import '../styles/Login.css';

function ForgotPassword() {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');

  const handleReset = async () => {
    if (!email) {
      setMessage('Please enter your email address.');
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/auth/forgot-password', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email }),
      });

      if (response.ok) {
        setMessage('Password reset instructions sent to your email.');
      } else {
        setMessage('Error: Unable to send reset instructions.');
      }
    } catch (err) {
      setMessage('Server error. Please try again later.');
    }
  };

  return (
    <div className="login-container">
      {/* Top header */}
      <div className="top-card">
        <h2>MindMate</h2>
      </div>

      {/* Forgot Password Box */}
      <div className="login-box">
        <h1>Forgot Password</h1>
        <input
          type="email"
          placeholder="Enter your registered email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <button onClick={handleReset}>Send Reset Link</button>
        {message && <p className="message">{message}</p>}
      </div>
    </div>
  );
}

export default ForgotPassword;

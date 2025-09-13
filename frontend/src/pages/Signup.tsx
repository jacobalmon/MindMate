import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/Login.css';

function SignUp() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleSignUp = async () => {
    if (!email || !password) {
      setMessage('Please enter email and password');
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/auth/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage('Signup successful! Redirecting to login...');
        setTimeout(() => navigate('/'), 1500);
      } else {
        setMessage(data.error || 'Signup failed');
      }
    } catch (error) {
      console.error(error);
      setMessage('Server error');
    }
  };

  return (
    <div className="login-container">
      <div className="top-card">
        <h2>MindMate</h2>
      </div>

      <div className="login-box">
        <h1>Sign Up</h1>
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
        <button onClick={handleSignUp}>Create Account</button>

        <p className="signup" onClick={() => navigate('/')}>
          Already have an account? Login
        </p>

        {message && <p className="message">{message}</p>}
      </div>
    </div>
  );
}

export default SignUp;

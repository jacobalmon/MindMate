import { useState } from 'react';
import '../styles/Login.css';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  return (
    <div className="login-container">
      <div className="login-box">
        <h1>Welcome Back</h1>
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
        <button>Login</button>
        <p className="message">{message}</p>
      </div>
    </div>
  );
}

export default Login;

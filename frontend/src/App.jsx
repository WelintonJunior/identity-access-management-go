import { useEffect, useState } from 'react';
import '../styles/App.css';
import ListUsers from '../components/listUsers';

function App() {
  const API_URL = import.meta.env.VITE_API_URL;

  const [loginMessage, setLoginMessage] = useState('');
  const [registerMessage, setRegisterMessage] = useState('');
  const [users, setUsers] = useState('');
  const [token, setToken] = useState(localStorage.getItem('token'));

  useEffect(() => {
    const fetchUsers = async () => {
      if (!token) return;

      try {
        const response = await fetch(`${API_URL}/admin/users`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });

        const result = await response.json();

        if (!response.ok) {
          throw new Error(result.error || 'Failed to fetch users');
        }

        setUsers(result.data);
      } catch (error) {
        setUsers(`No users found: ${error.message}`);
      }
    };

    fetchUsers();
  }, [token]);

  async function handleSubmitLogin(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());

    try {
      const response = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
      });

      const result = await response.json();

      if (!response.ok) {
        throw new Error(result.error || 'Login failed');
      }

      setLoginMessage('Login successful!');
      localStorage.setItem('token', result.access_token);
      setToken(result.access_token);
      e.target.reset();
    } catch (error) {
      setLoginMessage(`Error: ${error.message}`);
    }
  }

  async function handleSubmitRegister(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());

    try {
      const response = await fetch(`${API_URL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
      });

      const result = await response.json();

      if (!response.ok) {
        throw new Error(result.error || 'Registration failed');
      }

      setRegisterMessage('Registration successful!');
      e.target.reset();
    } catch (error) {
      setRegisterMessage(`Error: ${error.message}`);
    }
  }

  return (
    <>
      <div className="container">
        <form onSubmit={handleSubmitLogin}>
          <h2>Login</h2>
          <input type="email" name="email" placeholder="Email" required />
          <input type="password" name="password" placeholder="Password" required />
          <input type="submit" value="Sign in" />
          <div className="message">{loginMessage}</div>
        </form>

        <form onSubmit={handleSubmitRegister}>
          <h2>Register</h2>
          <input type="text" name="full_name" placeholder="Full name" required />
          <input type="email" name="email" placeholder="Email" required />
          <input type="password" name="password" placeholder="Password" required />
          <input type="password" name="repeat_password" placeholder="Repeat password" required />
          <input type="submit" value="Register" />
          <div className="message">{registerMessage}</div>
        </form>
      </div>

      <ListUsers users={users} />
    </>
  );
}

export default App;

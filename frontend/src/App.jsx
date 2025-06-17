import { useState } from 'react'
import './App.css'

function App() {
  const API_URL = import.meta.env.VITE_API_URL;

  const [loginMessage, setLoginMessage] = useState('');
  const [registerMessage, setRegisterMessage] = useState('');

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
        throw new Error(result.error || 'Erro no login');
      }

      setLoginMessage('Login realizado com sucesso!');
      console.log('Login:', result);
      e.target.reset(); 
    } catch (error) {
      setLoginMessage(`Erro: ${error.message}`);
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
        throw new Error(result.error || 'Erro no cadastro');
      }

      setRegisterMessage('Cadastro realizado com sucesso!');
      e.target.reset(); 
      console.log('Register:', result);
    } catch (error) {
      setRegisterMessage(`Erro: ${error.message}`);
    }
  }

  return (
    <div className="container">
      <form onSubmit={handleSubmitLogin}>
        <h2>Login</h2>
        <input type="email" name="email" id="login-email" placeholder="Email" required />
        <input type="password" name="password" id="login-password" placeholder="Senha" required />
        <input type="submit" value="Entrar" />
        <div className="message">{loginMessage}</div>
      </form>

      <form onSubmit={handleSubmitRegister}>
        <h2>Cadastro</h2>
        <input type="full_name" name="full_name" id="full_name" placeholder="full_name" required />
        <input type="email" name="email" id="register-email" placeholder="Email" required />
        <input type="password" name="password" id="register-password" placeholder="Senha" required />
        <input type="password" name="repeat_password" id="register-repeatPassword" placeholder="Repetir senha" required />
        <input type="submit" value="Cadastrar" />
        <div className="message">{registerMessage}</div>
      </form>
    </div>
  );
}

export default App;

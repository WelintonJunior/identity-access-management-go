import '../styles/listUsers.css';

export default function ListUsers({ users }) {
  const isLoggedIn = !!localStorage.getItem('token');

  const handleLogout = () => {
    localStorage.removeItem('token');
    window.location.reload();
  };

  if (!isLoggedIn) return null;

  return (
    <div className="users-container">
      <div className="users-header">
        <h2>User List</h2>
        <button onClick={handleLogout} className="logout-btn">
          Logout
        </button>
      </div>

      {typeof users === 'string' ? (
        <div className="users-message">{users}</div>
      ) : (
        <ul className="users-list">
          {users.map((user) => (
            <li key={user.id} className="user-item">
              <strong>{user.full_name}</strong>
              <br />
              <span>{user.email}</span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

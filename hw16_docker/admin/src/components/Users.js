import React, { useEffect, useState } from 'react';

function Users() {
  const [users, setUsers] = useState([]);
  const [newUser, setNewUser] = useState({ name: '', email: '', password: '' });
  const [refresh, setRefresh] = useState(false);

  useEffect(() => {
    fetch('http://localhost:8080/users')
      .then((res) => res.json())
      .then((data) => setUsers(data))
      .catch((err) => console.error("Error fetching users:", err));
  }, [refresh]);

  const handleChange = (e) => {
    setNewUser({ ...newUser, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    fetch('http://localhost:8080/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newUser)
    })
      .then((res) => res.json())
      .then((data) => {
        console.log("User created with ID:", data.id);
        setNewUser({ name: '', email: '', password: '' });
        setRefresh(!refresh);
      })
      .catch((err) => console.error("Error creating user:", err));
  };

  return (
    <div>
      <h2>Users</h2>
      <ul className="list-group">
        {users.map((u) => (
          <li key={u.id} className="list-group-item">
            {u.id}: {u.name} ({u.email})
          </li>
        ))}
      </ul>

      <h3 className="mt-4">Add New User</h3>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label>Name:</label>
          <input 
            type="text" 
            className="form-control" 
            name="name" 
            value={newUser.name} 
            onChange={handleChange} 
            required />
        </div>
        <div className="mb-3">
          <label>Email:</label>
          <input 
            type="email" 
            className="form-control" 
            name="email" 
            value={newUser.email} 
            onChange={handleChange} 
            required />
        </div>
        <div className="mb-3">
          <label>Password:</label>
          <input 
            type="password" 
            className="form-control" 
            name="password" 
            value={newUser.password} 
            onChange={handleChange} 
            required />
        </div>
        <button type="submit" className="btn btn-primary">Add User</button>
      </form>
    </div>
  );
}

export default Users;

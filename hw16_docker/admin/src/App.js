import React, { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  const [users, setUsers] = useState([]);
  const [newUser, setNewUser] = useState({ name: "", email: "" });

  useEffect(() => {
    fetch("http://localhost:8081/users")
      .then((res) => res.json())
      .then((data) => setUsers(data))
      .catch((err) => console.error("Ошибка загрузки пользователей:", err));
  }, []);

  const handleCreateUser = () => {
    fetch("http://localhost:8081/users/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(newUser),
    })
      .then(() => window.location.reload())
      .catch((err) => console.error("Ошибка при добавлении пользователя:", err));
  };
const updateUser = async (id, newPassword) => {
    await fetch("http://localhost:8081/users/update", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ id, password: newPassword }),
    });
    fetchUsers();
};

const deleteUser = async (id) => {
    await fetch("http://localhost:8081/users/delete", {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ id }),
    });
    fetchUsers();
};
  const handleDeleteUser = (id) => {
    fetch("http://localhost:8081/users/delete", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id }),
    })
      .then(() => window.location.reload())
      .catch((err) => console.error("Ошибка при удалении пользователя:", err));
  };

  return (
    <div className="container mt-5">
      <h1>Панель управления пользователями</h1>
      <div className="mb-3">
        <input
          className="form-control mb-2"
          placeholder="Имя"
          value={newUser.name}
          onChange={(e) => setNewUser({ ...newUser, name: e.target.value })}
        />
        <input
          className="form-control mb-2"
          placeholder="Email"
          value={newUser.email}
          onChange={(e) => setNewUser({ ...newUser, email: e.target.value })}
        />
        <button className="btn btn-primary" onClick={handleCreateUser}>
          Добавить пользователя
        </button>
      </div>
      <ul className="list-group">
        {users.map((user) => (
          <li key={user.id} className="list-group-item d-flex justify-content-between align-items-center">
            {user.name} ({user.email})
            <button className="btn btn-danger" onClick={() => handleDeleteUser(user.id)}>
              Удалить
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;

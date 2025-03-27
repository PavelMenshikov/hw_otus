import React, { useEffect, useState } from 'react';

function Orders() {
  const [orders, setOrders] = useState([]);
  const userId = 1;

  useEffect(() => {
    fetch(`http://localhost:8080/orders?user_id=${userId}`)
      .then((res) => res.json())
      .then((data) => setOrders(data))
      .catch((err) => console.error("Error fetching orders:", err));
  }, [userId]);

  return (
    <div>
      <h2>Orders for User {userId}</h2>
      <ul className="list-group">
        {orders.map((o) => (
          <li key={o.id} className="list-group-item">
            Order {o.id}: Total ${o.totalAmount} on {o.orderDate}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Orders;

import React, { useEffect, useState } from 'react';

function Stats() {
  const [stats, setStats] = useState(null);
  const userId = 1;

  useEffect(() => {
    fetch(`http://localhost:8080/stats?user_id=${userId}`)
      .then((res) => res.json())
      .then((data) => setStats(data))
      .catch((err) => console.error("Error fetching stats:", err));
  }, [userId]);

  return (
    <div>
      <h2>Stats for User {userId}</h2>
      {stats ? (
        <div>
          <p>Total Spent: ${stats.totalSpent}</p>
          <p>Average Product Price: ${stats.avgProductPrice}</p>
        </div>
      ) : (
        <p>Loading stats...</p>
      )}
    </div>
  );
}

export default Stats;

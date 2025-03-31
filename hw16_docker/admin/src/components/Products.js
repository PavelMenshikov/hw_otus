import React, { useEffect, useState } from 'react';

function Products() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/products')
      .then((res) => res.json())
      .then((data) => setProducts(data))
      .catch((err) => console.error("Error fetching products:", err));
  }, []);

  return (
    <div>
      <h2>Products</h2>
      <ul className="list-group">
        {products.map((p) => (
          <li key={p.id} className="list-group-item">
            {p.id}: {p.name} - ${p.price}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Products;

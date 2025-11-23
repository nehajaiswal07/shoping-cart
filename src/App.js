import React, { useState } from "react";

function App() {
  const [cart, setCart] = useState([]);
  const [orders, setOrders] = useState([]);

  const items = [
    { id: 1, name: "Item A", price: 200 },
    { id: 2, name: "Item B", price: 350 },
    { id: 3, name: "Item C", price: 120 },
  ];

  const addToCart = (item) => {
    setCart([...cart, item]);
    alert(`${item.name} added to cart`);
  };

  const checkout = () => {
    setOrders([...orders, { id: Date.now(), cart }]);
    setCart([]);
    alert("Order placed successfully!");
  };

  return (
    <div style={{ padding: 20 }}>
      <h1>Simple Shopping Cart</h1>

      <h2>Items</h2>
      {items.map((item) => (
        <div key={item.id}>
          <strong>{item.name}</strong> - ₹{item.price}
          <button onClick={() => addToCart(item)} style={{ marginLeft: 10 }}>
            Add to Cart
          </button>
        </div>
      ))}

      <h2>Cart</h2>
      {cart.length === 0 ? (
        <p>No items in cart</p>
      ) : (
        <>
          {cart.map((item, index) => (
            <p key={index}>{item.name} - ₹{item.price}</p>
          ))}
          <button onClick={checkout}>Checkout</button>
        </>
      )}

      <h2>Order History</h2>
      {orders.length === 0 ? (
        <p>No orders yet</p>
      ) : (
        orders.map((order) => (
          <div key={order.id} style={{ marginBottom: 10 }}>
            <strong>Order ID: {order.id}</strong>
            {order.cart.map((item, index) => (
              <p key={index}>• {item.name} - ₹{item.price}</p>
            ))}
          </div>
        ))
      )}
    </div>
  );
}

export default App;

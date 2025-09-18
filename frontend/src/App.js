import React, { useEffect, useState } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    // Fetch from Go backend (change URL if using Docker network)
    fetch('http://backend:4000/api/test')
      .then(res => res.json())
      .then(data => setMessage(data.message))
      .catch(err => setMessage('Could not reach backend'));
  }, []);

  return (
    <div style={{textAlign: 'center', marginTop: '50px'}}>
      <h1>React Frontend</h1>
      <p>{message || 'Loading from backend...'}</p>
    </div>
  );
}

export default App;

import React, { useEffect, useState } from "react";
import { fetchTestData } from "./api";

function App() {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetchTestData().then(result => setData(result));
  }, []);

  return (
    <div>
      <h1>React Frontend</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

export default App;

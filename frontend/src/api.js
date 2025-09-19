// frontend/src/api.js

// Base URL for backend API (Kubernetes service)
export const API_URL = "http://backend-service:4000/api";

// Example function to fetch data from backend
export async function fetchTestData() {
  try {
    const response = await fetch(`${API_URL}/test`);
    if (!response.ok) throw new Error('Network response was not ok');
    return await response.json();
  } catch (error) {
    console.error('Error fetching data:', error);
    return null;
  }
}

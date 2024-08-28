import React, { useState, useEffect } from 'react';
import axios from 'axios';

function KidsConsole() {
  const [chores, setChores] = useState([]);

  useEffect(() => {
    fetchChores();
  }, []);

  const fetchChores = async () => {
    try {
      const response = await axios.get('/api/chores', {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      });
      setChores(response.data);
    } catch (error) {
      console.error('Failed to fetch chores:', error);
    }
  };

  const completeChore = async (choreId) => {
    try {
      await axios.post(`/api/chores/${choreId}/complete`, {}, {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      });
      fetchChores(); // Refresh the chore list
    } catch (error) {
      console.error('Failed to complete chore:', error);
    }
  };

  return (
    <div>
      <h1>My Chores</h1>
      <ul>
        {chores.map(chore => (
          <li key={chore.id}>
            {chore.title} - {chore.description}
            {!chore.completed && (
              <button onClick={() => completeChore(chore.id)}>Complete</button>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default KidsConsole;
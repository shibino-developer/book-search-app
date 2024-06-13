import React, { useState } from 'react';
import axios from 'axios';
import SearchBar from './components/SearchBar';
import BookList from './components/BookList';
import './App.css';

const App = () => {
  const [books, setBooks] = useState([]);
  const [error, setError] = useState('');

  const fetchBooks = async (query) => {
    try {
      const response = await axios.get(`http://localhost:8080/books`, {
        params: { q: query }
      });
      setBooks(response.data);
      setError('');
    } catch (error) {
      console.error('Error fetching books:', error);
      setError('Failed to fetch books');
    }
  };

  return (
    <div className="App">
      <h1>Book Search</h1>
      <SearchBar onSearch={fetchBooks} />
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <BookList books={books} />
    </div>
  );
};

export default App;

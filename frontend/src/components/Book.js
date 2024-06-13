// src/components/Book.js

import React from 'react';
import './Book.css'; // Optional: Create a Book.css file for specific styling

const Book = ({ title, author, description }) => (
  <div className="book">
    <h2>{title}</h2>
    <h3>{author}</h3>
    <p>{description}</p>
  </div>
);

export default Book;

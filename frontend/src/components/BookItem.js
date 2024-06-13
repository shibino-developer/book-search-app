import React from 'react';

const BookItem = ({ book }) => {
  return (
    <div>
      <h2>{book.title}</h2>
      <h3>{book.authors.join(', ')}</h3>
      <p>{book.description}</p>
    </div>
  );
};

export default BookItem;

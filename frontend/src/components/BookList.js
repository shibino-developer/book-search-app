import React from 'react';
import BookItem from './BookItem';

const BookList = ({ books }) => {
  return (
    <div>
      {books.map((book, index) => (
        <BookItem key={index} book={book} />
      ))}
    </div>
  );
};

export default BookList;

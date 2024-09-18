import React from 'react';

const CopyrightX: React.FC = () => {
  const currentYear = new Date().getFullYear();
  
  return (
    <footer>
    <p>&copy; {currentYear} fred test。保留所有权利。</p>
    </footer>
  );
};

export default CopyrightX;
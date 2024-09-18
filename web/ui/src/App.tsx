import { Outlet } from 'react-router-dom';
import CopyrightX from './components/Copyright';

export default function App() {
  return (
    <>
      <Outlet />
      <CopyrightX />
    </>
  );
}
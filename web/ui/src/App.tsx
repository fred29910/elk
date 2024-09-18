import { Outlet } from 'react-router-dom';
import CopyrightX from './components/Copyright';
import Box from '@mui/material/Box';

export default function App() {
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <Box sx={{ flex: 1 }}>
        <Outlet />
      </Box>
      <Box component="footer" sx={{ py: 3, px: 2, mt: 'auto', backgroundColor: 'background.paper' }}>
        <CopyrightX />
      </Box>
    </Box>
  );
}
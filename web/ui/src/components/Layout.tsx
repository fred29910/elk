import { Outlet } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { useAuthStore } from '../store/authStore';

export default function Layout() {
  const { isAuthenticated, username, logout } = useAuthStore();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            我的应用
          </Typography>
          <Button color="inherit" component={RouterLink} to="/">首页</Button>
          {isAuthenticated && (
            <Button color="inherit" component={RouterLink} to="/parse">解析</Button>
          )}
          <Button color="inherit" component={RouterLink} to="/about">关于</Button>
          {isAuthenticated ? (
            <>
              <Typography variant="body1" sx={{ mr: 2 }}>
                欢迎, {username}
              </Typography>
              <Button color="inherit" onClick={handleLogout}>登出</Button>
            </>
          ) : (
            <Button color="inherit" component={RouterLink} to="/login">登录</Button>
          )}
        </Toolbar>
      </AppBar>
      <Container maxWidth="lg">
        <Box sx={{ mt: 2, mb: 4 }}>
          <Outlet />
        </Box>
      </Container>
    </>
  );
}
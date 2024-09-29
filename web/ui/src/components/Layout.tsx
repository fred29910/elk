import { useState } from 'react';
import { Outlet } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { useAuthStore } from '../store/authStore';

export default function Layout() {
  const { isAuthenticated, username, logout } = useAuthStore();
  const navigate = useNavigate();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = () => {
    logout();
    handleClose();
    navigate('/');
  };

  return (
    <>
      <AppBar position="static">
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between' }}>
          <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="h6" component="div" sx={{ mr: 2 }}>
              我的应用
            </Typography>
            <Button color="inherit" component={RouterLink} to="/">首页</Button>
            {isAuthenticated && (
              <Button color="inherit" component={RouterLink} to="/parse">解析</Button>
            )}
            <Button color="inherit" component={RouterLink} to="/about">关于</Button>
          </Box>
          <Box>
            {isAuthenticated ? (
              <>
                <Button color="inherit" onClick={handleMenu}>
                  {username}
                </Button>
                <Menu
                  anchorEl={anchorEl}
                  open={Boolean(anchorEl)}
                  onClose={handleClose}
                >
                  <MenuItem component={RouterLink} to="/users" onClick={handleClose}>个人中心</MenuItem>
                  <MenuItem component={RouterLink} to="/systems" onClick={handleClose}>系统设置</MenuItem>
                  <MenuItem onClick={handleLogout}>登出</MenuItem>
                </Menu>
              </>
            ) : (
              <Button color="inherit" component={RouterLink} to="/login">登录</Button>
            )}
          </Box>
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
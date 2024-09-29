import React, { useState } from 'react';
import { Box, TextField, Button, Typography, Link } from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
// import { useAuthStore } from '../store/authStore';

export default function Register() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (password !== confirmPassword) {
      alert('密码不匹配');
      return;
    }
    // 这里添加注册逻辑
    console.log('注册:', { username, password });
    
    // 注册成功后跳转到登录页
    navigate('/login');
  };

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        gap: 2,
        maxWidth: 300,
        margin: 'auto',
        mt: 4,
      }}
    >
      <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
        注册
      </Typography>
      <TextField
        label="用户名"
        variant="outlined"
        fullWidth
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        required
      />
      <TextField
        label="密码"
        type="password"
        variant="outlined"
        fullWidth
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      />
      <TextField
        label="确认密码"
        type="password"
        variant="outlined"
        fullWidth
        value={confirmPassword}
        onChange={(e) => setConfirmPassword(e.target.value)}
        required
      />
      <Button type="submit" variant="contained" fullWidth>
        注册
      </Button>
      <Link component={RouterLink} to="/login" variant="body2">
        已有账号？登录
      </Link>
    </Box>
  );
}
import React, { useState } from 'react';
import { Box, TextField, Button, Typography, Link } from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import { useAuthStore } from '../store/authStore';

export default function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const login = useAuthStore((state) => state.login);
  const navigate = useNavigate();

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // 这里应该有实际的登录逻辑，比如API调用
    // 为了演示，我们直接调用login
    login(username);
    navigate('/');
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
        登录
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
      <Button type="submit" variant="contained" fullWidth>
        登录
      </Button>
      <Link component={RouterLink} to="/" variant="body2">
        返回主页
      </Link>
    </Box>
  );
}
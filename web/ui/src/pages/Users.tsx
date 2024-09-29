import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

export default function Users() {
  return (
    <Box>
      <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
        个人中心
      </Typography>
      <Typography variant="body1">
        这里是用户个人中心页面。您可以在这里查看和编辑您的个人信息。
      </Typography>
    </Box>
  );
}
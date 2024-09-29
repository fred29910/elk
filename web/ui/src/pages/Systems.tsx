import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

export default function Systems() {
  return (
    <Box>
      <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
        系统设置
      </Typography>
      <Typography variant="body1">
        这里是系统设置页面。您可以在这里管理系统相关的配置和选项。
      </Typography>
    </Box>
  );
}
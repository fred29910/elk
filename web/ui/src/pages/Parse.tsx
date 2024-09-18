import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

export default function Parse() {
  return (
    <Box>
      <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
        解析页面
      </Typography>
      <Typography variant="body1">
        这里是解析功能的内容。
      </Typography>
    </Box>
  );
}
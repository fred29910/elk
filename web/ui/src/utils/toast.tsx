import React from 'react';
import { createRoot } from 'react-dom/client';
import { Snackbar, Alert, AlertColor } from '@mui/material';

const ToastComponent: React.FC<ToastProps> = ({ message, type, duration }) => {
    const [open, setOpen] = React.useState(true);

    const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
      if (reason === 'clickaway') {
        console.log(event, reason)
        return;
      }
      setOpen(false);
    };
    return (
      <Snackbar
        open={open}
        autoHideDuration={duration}
        onClose={handleClose}
        anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
      >
        <Alert onClose={handleClose} severity={type} sx={{ width: '100%' }}>
          {message}
        </Alert>
      </Snackbar>
    );
};

let toastRoot: HTMLDivElement | null = null;

export const showToast = (message: string, type: AlertColor = 'info', duration: number = 3000) => {
  if (!toastRoot) {
    toastRoot = document.createElement('div');
    toastRoot.id = 'toast-root';
    document.body.appendChild(toastRoot);
  }

  const root = createRoot(toastRoot);
  root.render(<ToastComponent message={message} type={type} duration={duration} />);

  // 清理函数
  setTimeout(() => {
    root.unmount();
  }, duration + 1000); // 给一些额外时间以确保动画完成
};

interface ToastProps {
    message: string;
    type: AlertColor;
    duration: number;
}



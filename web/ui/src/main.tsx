import { StrictMode, lazy, Suspense } from 'react'
import { createRoot } from 'react-dom/client'
import { ThemeProvider } from '@mui/material/styles';
import { CssBaseline, CircularProgress } from '@mui/material';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import App from './App.tsx'
import Layout from './components/Layout.tsx'
import ProtectedRoute from './components/ProtectedRoute.tsx'
import Home from './pages/Home.tsx'
import Parse from './pages/Parse.tsx'
import About from './pages/About.tsx'
import Login from './pages/Login.tsx'
import theme from './theme';
import './index.css'
import Register from './pages/Register.tsx';

// 懒加载 Users 和 Systems 组件
const Users = lazy(() => import('./pages/Users.tsx'))
const Systems = lazy(() => import('./pages/Systems.tsx'))

// 创建一个加载中的组件
const LoadingComponent = () => (
  <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
    <CircularProgress />
  </div>
)

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline enableColorScheme={true} />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />}>
            <Route path="login" element={<Login />} />
            <Route path="register" element={
              <Suspense fallback={<LoadingComponent />}>
                <Register />
              </Suspense>
            } />
            <Route element={<Layout />}>
              <Route index element={<Home />} />
              <Route element={<ProtectedRoute />}>
                <Route path="parse" element={<Parse />} />
                <Route path="users" element={
                  <Suspense fallback={<LoadingComponent />}>
                    <Users />
                  </Suspense>
                } />
                <Route path="systems" element={
                  <Suspense fallback={<LoadingComponent />}>
                    <Systems />
                  </Suspense>
                } />
              </Route>
              <Route path="about" element={<About />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  </StrictMode>,
)

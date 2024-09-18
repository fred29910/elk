import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { ThemeProvider } from '@mui/material/styles';
import { CssBaseline } from '@mui/material';
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

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline enableColorScheme={true} />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />}>
            <Route path="login" element={<Login />} />
            <Route element={<Layout />}>
              <Route index element={<Home />} />
              <Route element={<ProtectedRoute />}>
                <Route path="parse" element={<Parse />} />
              </Route>
              <Route path="about" element={<About />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  </StrictMode>,
)

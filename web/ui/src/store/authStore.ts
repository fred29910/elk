import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { LoginApi } from '../openapi/apis/LoginApi'
import { Configuration } from '../openapi/runtime'
import { LoginRequestBody } from '../openapi/models/LoginRequestBody'

interface AuthState {
  isAuthenticated: boolean
  username: string | null
  token: string | null
  login: (username: string, password: string) => Promise<void>
  logout: () => Promise<void>
}

// 使用环境变量或默认值
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000'

const createApi = (token: string | null) => new LoginApi(
  new Configuration({ 
    basePath: API_BASE_URL,
    accessToken: token || undefined
  })
)

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      isAuthenticated: false,
      username: null,
      token: null,
      login: async (username: string, password: string) => {
        try {
          const api = createApi(null)
          const loginRequestBody: LoginRequestBody = { username, password }
          const result = await api.loginLogin({ loginRequestBody })
          if (result.token) {
            set({ isAuthenticated: true, username, token: result.token })
          } else {
            throw new Error('登录失败')
          }
        } catch (error) {
          console.error('登录错误:', error)
          throw error
        }
      },
      logout: async () => {
        try {
          const api = createApi(get().token)
          await api.loginLogout()
          set({ isAuthenticated: false, username: null, token: null })
        } catch (error) {
          console.error('登出错误:', error)
          throw error
        }
      },
    }),
    {
      name: 'auth-storage',
    }
  )
)
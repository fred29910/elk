import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { AuthApi } from '../openapi/apis/AuthApi'
import { Configuration } from '../openapi/runtime'
import { OvLoginRequest } from '../openapi/models/OvLoginRequest'

interface AuthState {
  isAuthenticated: boolean
  username: string | null
  token: string | null
  login: (username: string, password: string) => Promise<void>
  logout: () => Promise<void>
}

// 使用环境变量或默认值
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000'

let apiInstance: AuthApi | null = null

const createApi = () => {
  if (!apiInstance) {
    apiInstance = new AuthApi(
      new Configuration({ 
        basePath: API_BASE_URL,
        accessToken: async () => {
          const state = useAuthStore.getState()
          return state.token || ""
        }
      })
    )
  }
  return apiInstance
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      isAuthenticated: false,
      username: null,
      token: null,
      login: async (username: string, password: string) => {
        try {
          const api = createApi()
          const loginRequestBody: OvLoginRequest = { username, password, captcha: '' }
          const result = await api.apiAuthLoginPost({ body: loginRequestBody })
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
          const api = createApi()
          await api.apiAuthLogoutPost()
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
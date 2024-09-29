import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'
import { AuthApi } from '../openapi/apis/AuthApi'
import { Configuration } from '../openapi/runtime'
import { OvLoginRequest } from '../openapi/models/OvLoginRequest'
import { showToast } from '../utils/toast'
import CryptoJS from 'crypto-js'
import { OvRefreshTokenRequest } from '../openapi/models/OvRefreshTokenRequest'

// 加密密钥，应该存储在安全的地方，比如环境变量
const ENCRYPTION_KEY = import.meta.env.VITE_ENCRYPTION_KEY || 'your-secret-key'

interface AuthState {
  isAuthenticated: boolean
  username: string | null
  token: string | null
  login: (username: string, password: string) => Promise<void>
  logout: () => Promise<void>
}

// 使用环境变量或默认值
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8081'

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

// 加密函数
const encrypt = (text: string): string => {
  return CryptoJS.AES.encrypt(text, ENCRYPTION_KEY).toString()
}

// 解密函数
const decrypt = (ciphertext: string): string => {
  const bytes = CryptoJS.AES.decrypt(ciphertext, ENCRYPTION_KEY)
  return bytes.toString(CryptoJS.enc.Utf8)
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
          const loginRequestBody: OvLoginRequest = { username, password, captcha: '123456' }
          const result = await api.apiAuthLoginPost({ body: loginRequestBody })
          if (result.token) {
            set({ 
              isAuthenticated: true, 
              username: encrypt(username), 
              token: encrypt(result.token) 
            })
          } else {
            throw new Error('登录失败')
          }
        } catch (error) {
          console.error('登录错误:', error)
          showToast('登录失败，请检查您的用户名和密码', 'error')
          throw error
        }
      },
      register: async (username: string, password: string) => {
        try {
          const api = createApi()
          const registerRequestBody: OvRefreshTokenRequest = { refreshToken: '123456', token: '123456' }
          const result = await api.apiAuthRefreshTokenPost({ body: registerRequestBody })
          if (result.token) {
            set({ 
              isAuthenticated: true, 
              username: encrypt(username),
              token: encrypt(result.token)
            })
          } else {
            throw new Error('注册失败')
          }
        } catch (error) {
          console.error('注册错误:', error)
          showToast('注册失败，请稍后重试', 'error')
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
          showToast('登出失败，请稍后重试', 'error')
          throw error
        }
      },
    }),
    {
      name: 'auth-storage',
      storage: createJSONStorage(() => localStorage),
      partialize: (state) => ({
        isAuthenticated: state.isAuthenticated,
        username: state.username,
        token: state.token,
      }),
      // 在存储之前加密，在读取之后解密
      serialize: (state: AuthState) => JSON.stringify(state),
      deserialize: (str: string) => {
        const state = JSON.parse(str)
        if (state.username) {
          state.username = decrypt(state.username)
        }
        if (state.token) {
          state.token = decrypt(state.token)
        }
        return state
      },
    }
  )
)
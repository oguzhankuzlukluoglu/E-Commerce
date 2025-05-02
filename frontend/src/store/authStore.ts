import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface User {
  id: string
  email: string
  name: string
  role: 'user' | 'admin'
}

interface AuthState {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  login: (email: string, password: string) => Promise<void>
  register: (name: string, email: string, password: string) => Promise<void>
  logout: () => void
}

export const useAuth = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      token: null,
      isAuthenticated: false,
      login: async (email: string, password: string) => {
        try {
          const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
          })

          if (!response.ok) {
            throw new Error('Giriş başarısız')
          }

          const data = await response.json()
          set({
            user: data.user,
            token: data.token,
            isAuthenticated: true,
          })
        } catch (error) {
          console.error('Login error:', error)
          throw error
        }
      },
      register: async (name: string, email: string, password: string) => {
        try {
          const response = await fetch('/api/auth/register', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, email, password }),
          })

          if (!response.ok) {
            throw new Error('Kayıt başarısız')
          }

          const data = await response.json()
          set({
            user: data.user,
            token: data.token,
            isAuthenticated: true,
          })
        } catch (error) {
          console.error('Register error:', error)
          throw error
        }
      },
      logout: () => {
        set({
          user: null,
          token: null,
          isAuthenticated: false,
        })
      },
    }),
    {
      name: 'auth-storage',
    }
  )
) 
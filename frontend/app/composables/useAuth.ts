import { ref, computed } from 'vue'

export interface User {
  id: string
  name: string
  email: string
  role: 'admin' | 'customer'
  address: string
  phone: string
}

export const useAuth = () => {
  const token = useState<string | null>('auth_token', () => null)
  const user = useState<User | null>('auth_user', () => null)
  const loading = useState<boolean>('auth_loading', () => false)

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 'admin')

  const config = {
    apiUrl: 'http://localhost:5000'
  }

  const initAuth = async () => {
    if (import.meta.client) {
      const storedToken = localStorage.getItem('fourmart_token')
      if (storedToken) {
        token.value = storedToken
        await fetchUser()
      }
    }
  }

  const fetchUser = async () => {
    if (!token.value) return
    loading.value = true
    try {
      const res = await fetch(`${config.apiUrl}/api/auth/me`, {
        headers: {
          'Authorization': `Bearer ${token.value}`
        }
      })
      if (res.ok) {
        user.value = await res.json()
      } else {
        logout()
      }
    } catch (e) {
      console.error('Failed to fetch user profile:', e)
    } finally {
      loading.value = false
    }
  }

  const login = async (email: string, password: string) => {
    loading.value = true
    try {
      const res = await fetch(`${config.apiUrl}/api/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      })

      const data = await res.json()
      if (!res.ok) {
        throw new Error(data.error || 'Login gagal. Coba lagi.')
      }

      token.value = data.token
      user.value = data.user
      
      if (import.meta.client) {
        localStorage.setItem('fourmart_token', data.token)
      }
      return data.user
    } finally {
      loading.value = false
    }
  }

  const register = async (name: string, email: string, password: string, address: string, phone: string) => {
    loading.value = true
    try {
      const res = await fetch(`${config.apiUrl}/api/auth/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, email, password, address, phone })
      })

      const data = await res.json()
      if (!res.ok) {
        throw new Error(data.error || 'Registrasi gagal. Coba lagi.')
      }

      token.value = data.token
      user.value = data.user

      if (import.meta.client) {
        localStorage.setItem('fourmart_token', data.token)
      }
      return data.user
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    if (token.value) {
      try {
        await fetch(`${config.apiUrl}/api/auth/logout`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token.value}`
          }
        })
      } catch (e) {
        console.error('Logout error on server:', e)
      }
    }

    token.value = null
    user.value = null
    if (import.meta.client) {
      localStorage.removeItem('fourmart_token')
    }
  }

  return {
    token,
    user,
    loading,
    isLoggedIn,
    isAdmin,
    initAuth,
    fetchUser,
    login,
    register,
    logout
  }
}

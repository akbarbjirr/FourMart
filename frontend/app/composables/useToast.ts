import { ref } from 'vue'

export interface Toast {
  id: string
  message: string
  type: 'success' | 'error' | 'info'
  duration: number
}

export const useToast = () => {
  const toasts = useState<Toast[]>('global_toasts', () => [])

  const showToast = (message: string, type: 'success' | 'error' | 'info' = 'success', duration: number = 3000) => {
    const id = Math.random().toString(36).substring(2, 9)
    toasts.value.push({ id, message, type, duration })

    setTimeout(() => {
      removeToast(id)
    }, duration)
  }

  const removeToast = (id: string) => {
    toasts.value = toasts.value.filter(t => t.id !== id)
  }

  return {
    toasts,
    showToast,
    removeToast
  }
}

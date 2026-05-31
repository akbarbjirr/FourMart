import { ref, computed } from 'vue'

export interface CartItem {
  product_id: string
  name: string
  price: number
  image: string
  quantity: number
  stock: number
}

export const useCart = () => {
  const items = useState<CartItem[]>('cart_items', () => [])

  const initCart = () => {
    if (import.meta.client) {
      const stored = localStorage.getItem('fourmart_cart')
      if (stored) {
        try {
          items.value = JSON.parse(stored)
        } catch (e) {
          console.error('Failed to parse cart items:', e)
          items.value = []
        }
      }
    }
  }

  const saveCart = () => {
    if (import.meta.client) {
      localStorage.setItem('fourmart_cart', JSON.stringify(items.value))
    }
  }

  const addToCart = (product: any, quantity: number = 1) => {
    const existing = items.value.find(i => i.product_id === product.id)
    
    if (existing) {
      const newQty = existing.quantity + quantity
      if (newQty <= product.stock) {
        existing.quantity = newQty
      } else {
        existing.quantity = product.stock
      }
    } else {
      items.value.push({
        product_id: product.id,
        name: product.name,
        price: product.price,
        image: product.image,
        quantity: Math.min(quantity, product.stock),
        stock: product.stock
      })
    }
    saveCart()
  }

  const removeFromCart = (productId: string) => {
    items.value = items.value.filter(i => i.product_id !== productId)
    saveCart()
  }

  const updateQuantity = (productId: string, quantity: number) => {
    const item = items.value.find(i => i.product_id === productId)
    if (item) {
      if (quantity <= 0) {
        removeFromCart(productId)
      } else if (quantity <= item.stock) {
        item.quantity = quantity
        saveCart()
      }
    }
  }

  const clearCart = () => {
    items.value = []
    saveCart()
  }

  const cartCount = computed(() => {
    return items.value.reduce((total, item) => total + item.quantity, 0)
  })

  const cartTotal = computed(() => {
    return items.value.reduce((total, item) => total + (item.price * item.quantity), 0)
  })

  return {
    items,
    initCart,
    addToCart,
    removeFromCart,
    updateQuantity,
    clearCart,
    cartCount,
    cartTotal
  }
}

<template>
  <div class="app-layout">
    <header class="md-top-bar" :class="{ 'md-top-bar--scrolled': isScrolled }">
      <div class="md-top-bar-container">
        <div class="md-brand" v-show="!isMobileSearchVisible" @click="navigateTo('/')" style="cursor: pointer;">
          <span class="material-symbols-outlined">school</span>
          <span>FourMart</span>
        </div>

        <!-- Desktop Search bar -->
        <div class="header-search no-print desktop-only" v-if="$route.path === '/'">
          <span class="material-symbols-outlined search-icon">search</span>
          <input 
            type="text" 
            placeholder="Cari peralatan sekolah..." 
            v-model="searchQuery"
            @input="onSearch"
            class="search-input"
          />
        </div>

        <!-- Mobile Search Input (Visible when toggled) -->
        <div class="header-search mobile-search-active" v-if="isMobileSearchVisible">
          <button class="md-btn md-btn--icon" @click="isMobileSearchVisible = false">
            <span class="material-symbols-outlined">arrow_back</span>
          </button>
          <input 
            type="text" 
            placeholder="Cari produk..." 
            v-model="searchQuery"
            @input="onSearch"
            class="search-input"
            autoFocus
          />
        </div>

        <div class="md-nav-links no-print" v-show="!isMobileSearchVisible">
          <!-- Desktop Navigation -->
          <div class="desktop-only" style="display: flex; gap: 8px;">
          <button class="md-btn md-btn--text" @click="scrollToAnchor('katalog')">Katalog</button>

          <button 
            v-if="isLoggedIn && isAdmin" 
            class="md-btn md-btn--tonal"
            @click="navigateTo('/admin')"
          >
            <span class="material-symbols-outlined">dashboard</span>
            Admin Panel
          </button>
          </div>

          <!-- Mobile Search Toggle -->
          <button 
            v-if="$route.path === '/'" 
            class="md-btn md-btn--icon mobile-only" 
            @click="isMobileSearchVisible = true"
          >
            <span class="material-symbols-outlined">search</span>
          </button>

          <div class="md-badge-container">
            <button class="md-btn md-btn--icon" @click="isCartOpen = true" aria-label="Buka Keranjang">
              <span class="material-symbols-outlined">shopping_cart</span>
            </button>
            <span v-if="cartCount > 0" class="md-badge">{{ cartCount }}</span>
          </div>

          <div v-if="isLoggedIn" class="user-menu-container">
            <button class="md-btn md-btn--outlined" @click="toggleUserMenu">
              <span class="material-symbols-outlined">account_circle</span>
              {{ user?.name.split(' ')[0] }}
            </button>
            
            <div v-if="isUserMenuOpen" class="user-dropdown md-card md-card--elevated">
              <div class="dropdown-header">
                <strong>{{ user?.name }}</strong>
                <span class="text-muted" style="font-size: 0.75rem;">{{ user?.email }}</span>
              </div>
              <hr class="dropdown-divider" />
              <button class="dropdown-item" @click="goToOrders">
                <span class="material-symbols-outlined">receipt_long</span>
                Pesanan Saya
              </button>
              <button class="dropdown-item text-error" @click="handleLogout">
                <span class="material-symbols-outlined">logout</span>
                Keluar
              </button>
            </div>
          </div>

          <button 
            v-else 
            class="md-btn md-btn--filled" 
            @click="navigateTo('/login')"
          >
            Masuk
          </button>
        </div>
      </div>
    </header>

    <main class="main-content">
      <slot />
    </main>

    <template v-if="route.path === '/'">
      <AppWhyChooseUs />
      <AppTestimonials />
      <AppContactSupport />
    </template>

    <AppFooter class="no-print" />

    <div v-if="isCartOpen" class="md-drawer-overlay no-print" @click.self="isCartOpen = false">
      <div class="md-drawer">
        <div class="md-drawer-header">
          <div style="display: flex; align-items: center; gap: 8px;">
            <span class="material-symbols-outlined text-primary">shopping_cart</span>
            <h2 style="font-size: 1.25rem; font-weight: 500;">Keranjang Belanja</h2>
          </div>
          <button class="md-btn md-btn--icon" @click="isCartOpen = false">
            <span class="material-symbols-outlined">close</span>
          </button>
        </div>

        <div class="md-drawer-content">
          <div v-if="cartItems.length === 0" class="empty-cart-state flex-center" style="flex-direction: column; height: 100%; gap: 16px;">
            <span class="material-symbols-outlined" style="font-size: 64px; color: var(--md-sys-color-outline-variant);">shopping_cart_off</span>
            <p class="text-muted">Keranjang belanja Anda masih kosong</p>
            <button class="md-btn md-btn--filled" @click="isCartOpen = false">Mulai Belanja</button>
          </div>

          <div v-else class="cart-items-list">
            <div v-for="item in cartItems" :key="item.product_id" class="cart-item md-card md-card--outlined">
              <img :src="item.image" :alt="item.name" class="cart-item-img" />
              <div class="cart-item-info">
                <h4 class="cart-item-title">{{ item.name }}</h4>
                <p class="cart-item-price text-primary">Rp{{ item.price.toLocaleString('id-ID') }}</p>
                
                <div class="cart-item-controls">
                  <div class="quantity-selector">
                    <button 
                      class="qty-btn" 
                      @click="updateQuantity(item.product_id, item.quantity - 1)"
                    >-</button>
                    <span class="qty-val">{{ item.quantity }}</span>
                    <button 
                      class="qty-btn" 
                      @click="updateQuantity(item.product_id, item.quantity + 1)"
                      :disabled="item.quantity >= item.stock"
                    >+</button>
                  </div>
                  <button class="remove-btn text-error" @click="removeFromCart(item.product_id)">
                    <span class="material-symbols-outlined" style="font-size: 20px;">delete</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="cartItems.length > 0" class="md-drawer-footer">
          <div class="cart-summary-row">
            <span>Total Pembayaran</span>
            <strong class="text-primary" style="font-size: 1.25rem;">
              Rp{{ cartTotal.toLocaleString('id-ID') }}
            </strong>
          </div>
          <button class="md-btn md-btn--filled" style="width: 100%; height: 48px; margin-top: 16px;" @click="goToCheckout">
            Lanjut ke Checkout
          </button>
        </div>
      </div>
    </div>

    <div class="md-snackbar-container no-print">
      <div 
        v-for="toast in toasts" 
        :key="toast.id" 
        class="md-snackbar"
        :style="toast.type === 'error' ? 'border-left: 4px solid var(--md-sys-color-error)' : 'border-left: 4px solid var(--md-sys-color-primary)'"
      >
        <span class="md-snackbar-text">{{ toast.message }}</span>
        <span class="md-snackbar-action" @click="removeToast(toast.id)">TUTUP</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'
import { useCart } from '~/composables/useCart'
import { useToast } from '~/composables/useToast'
import { useRouter, useRoute } from 'vue-router'

const { isLoggedIn, isAdmin, user, initAuth, logout } = useAuth()
const { items: cartItems, initCart, updateQuantity, removeFromCart, cartCount, cartTotal } = useCart()
const { toasts, removeToast, showToast } = useToast()

const router = useRouter()
const route = useRoute()

const isScrolled = ref(false)
const isCartOpen = ref(false)
const isUserMenuOpen = ref(false)
const searchQuery = ref('')
const isMobileSearchVisible = ref(false)

const navigateTo = (path) => {
  if (route.path === '/' && path === '/') {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
  router.push(path)
}

const scrollToAnchor = (id) => {
  if (route.path === '/') {
    const el = document.getElementById(id)
    if (el) {
      el.scrollIntoView({ behavior: 'smooth' })
    }
  } else {
    router.push({ path: '/', hash: `#${id}` })
  }
}

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const onSearch = () => {
  if (route.path !== '/') {
    router.push({ path: '/', query: { search: searchQuery.value } })
  } else {
    router.replace({ query: { ...route.query, search: searchQuery.value || undefined } })
  }
}

watch(() => route.query.search, (newVal) => {
  searchQuery.value = newVal || ''
})

const toggleUserMenu = () => {
  isUserMenuOpen.value = !isUserMenuOpen.value
}

const closeMenuAndNavigate = (path) => {
  isUserMenuOpen.value = false
  navigateTo(path)
}

const handleLogout = async () => {
  await logout()
  isUserMenuOpen.value = false
  showToast('Anda telah keluar dari akun')
  router.push('/login')
}

const goToOrders = () => {
  isUserMenuOpen.value = false
  router.push('/orders')
}

const goToCheckout = () => {
  isCartOpen.value = false
  router.push('/checkout')
}

const handleScroll = () => {
  isScrolled.value = window.scrollY > 10
}

onMounted(async () => {
  await initAuth()
  initCart()
  searchQuery.value = route.query.search || ''
  window.addEventListener('scroll', handleScroll)
})
</script>

<style scoped>
/* Tambahkan smooth scroll secara global untuk layout ini */
:global(html) {
  scroll-behavior: smooth;
  overflow-x: hidden;
}

.app-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  width: 100%;
}

.md-top-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  width: 100%;
  height: 64px;
  z-index: 110;
  background-color: var(--md-sys-color-surface);
  transition: background-color 0.3s, box-shadow 0.3s, border-color 0.3s;
}

.md-top-bar--scrolled {
  background-color: var(--md-sys-color-surface-container-low);
  border-bottom: 1px solid var(--md-sys-color-outline-variant);
  box-shadow: var(--md-sys-shadow-1);
}

.desktop-only {
  display: flex;
}

.mobile-only {
  display: none;
}

.main-content {
  flex: 1;
  padding-top: 64px;
}

.header-search {
  display: flex;
  align-items: center;
  position: relative;
  max-width: 480px;
  width: 100%;
  margin: 0 20px;
}

.search-icon {
  position: absolute;
  left: 12px;
  color: var(--md-sys-color-on-surface-variant);
  pointer-events: none;
}

.search-input {
  width: 100%;
  height: 40px;
  padding: 0 16px 0 44px;
  background-color: var(--md-sys-color-surface-container-high);
  border: none;
  border-radius: var(--md-shape-corner-full);
  outline: none;
  font-size: 0.95rem;
  color: var(--md-sys-color-on-surface);
  transition: background-color 0.2s, box-shadow 0.2s;
}

.search-input:focus {
  background-color: var(--md-sys-color-surface-container-lowest);
  box-shadow: 0 0 0 2px var(--md-sys-color-primary);
}

.user-menu-container {
  position: relative;
}

.user-dropdown {
  position: absolute;
  right: 0;
  top: 48px;
  width: 220px;
  padding: 8px 0;
  z-index: 200;
  display: flex;
  flex-direction: column;
}

.dropdown-header {
  padding: 8px 16px 12px;
  display: flex;
  flex-direction: column;
}

.dropdown-divider {
  border: none;
  border-top: 1px solid var(--md-sys-color-outline-variant);
  margin: 4px 0;
}

.dropdown-item {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 16px;
  font-size: 0.9rem;
  text-align: left;
  transition: background-color 0.2s;
}

.dropdown-item:hover {
  background-color: rgba(103, 80, 164, 0.08);
}

.dropdown-item span {
  font-size: 20px;
}

/* Cart Drawer Styling helper */
.cart-items-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.cart-item {
  flex-direction: row !important;
  gap: 12px;
  padding: 12px !important;
  align-items: center;
}

.cart-item-img {
  width: 72px;
  height: 72px;
  object-fit: cover;
  border-radius: var(--md-shape-corner-medium);
  background-color: var(--md-sys-color-surface-container);
}

.cart-item-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.cart-item-title {
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface);
  display: -webkit-box;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.cart-item-price {
  font-size: 0.85rem;
  font-weight: 700;
}

.cart-item-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
}

.quantity-selector {
  display: flex;
  align-items: center;
  border: 1px solid var(--md-sys-color-outline-variant);
  border-radius: var(--md-shape-corner-small);
  overflow: hidden;
  height: 28px;
}

.qty-btn {
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--md-sys-color-surface-container-high);
  font-weight: bold;
}
.qty-btn:hover {
  background-color: var(--md-sys-color-outline-variant);
}
.qty-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.qty-val {
  padding: 0 10px;
  font-size: 0.85rem;
  font-weight: 600;
}

.remove-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  transition: background-color 0.2s;
}
.remove-btn:hover {
  background-color: rgba(186, 26, 26, 0.08);
}

.cart-summary-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.95rem;
}

@media (max-width: 768px) {
  .desktop-only {
    display: none !important;
  }
  .mobile-only {
    display: inline-flex !important;
  }
  .header-search.mobile-search-active {
    margin: 0;
    max-width: none;
    gap: 8px;
  }
  .md-badge-container {
    margin-left: 0;
  }
  .user-dropdown {
    width: 200px;
    top: 40px;
  }
  .md-top-bar-container {
    gap: 8px;
  }
}
</style>

<template>
  <div class="admin-layout">
    <aside class="admin-sidebar no-print" :class="{ 'admin-sidebar--open': isSidebarOpen }">
      <div class="sidebar-brand" @click="navigateTo('/')">
        <span class="material-symbols-outlined">school</span>
        <span>FourMart Admin</span>
      </div>

      <nav class="sidebar-nav">
        <button 
          class="nav-link" 
          :class="{ 'nav-link--active': isActiveRoute('/admin') }"
          @click="closeSidebarAndNavigate('/admin')"
        >
          <span class="material-symbols-outlined">dashboard</span>
          <span>Ringkasan</span>
        </button>

        <button 
          class="nav-link" 
          :class="{ 'nav-link--active': isActiveRoute('/admin/products') }"
          @click="closeSidebarAndNavigate('/admin/products')"
        >
          <span class="material-symbols-outlined">inventory_2</span>
          <span>Kelola Produk</span>
        </button>

        <button 
          class="nav-link" 
          :class="{ 'nav-link--active': isActiveRoute('/admin/orders') }"
          @click="closeSidebarAndNavigate('/admin/orders')"
        >
          <span class="material-symbols-outlined">receipt_long</span>
          <span>Kelola Pesanan</span>
        </button>

        <hr class="nav-divider" />

        <button 
          class="nav-link" 
          @click="closeSidebarAndNavigate('/')"
        >
          <span class="material-symbols-outlined">shopping_basket</span>
          <span>Lihat Toko</span>
        </button>
      </nav>

    </aside>

    <div v-if="isSidebarOpen" class="sidebar-overlay no-print" @click="isSidebarOpen = false"></div>

    <div class="admin-main">
      <header class="admin-header no-print">
        <div style="display: flex; align-items: center; gap: 12px;">
          <button class="md-btn md-btn--icon menu-toggle" @click="isSidebarOpen = !isSidebarOpen" aria-label="Menu">
            <span class="material-symbols-outlined">menu</span>
          </button>
          <h1 class="page-title">{{ pageTitle }}</h1>
        </div>
        
        <div style="display: flex; align-items: center; gap: 16px;">
          <span class="text-muted text-email-header">{{ user?.email }}</span>
          <button class="md-btn md-btn--outlined logout-btn-header" @click="handleLogout">
            <span class="material-symbols-outlined" style="font-size: 18px;">logout</span>
            <span>Keluar</span>
          </button>
        </div>
      </header>

      <div class="admin-content">
        <slot />
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
import { ref, computed, onMounted } from 'vue'
import { useAuth } from '~/composables/useAuth'
import { useToast } from '~/composables/useToast'
import { useRouter, useRoute } from 'vue-router'

const { user, initAuth, logout, isAdmin } = useAuth()
const { toasts, removeToast, showToast } = useToast()
const router = useRouter()
const route = useRoute()

const isSidebarOpen = ref(false)

const handleLogout = async () => {
  await logout()
  showToast('Anda telah keluar dari akun')
  router.push('/login')
}

const isActiveRoute = (path) => {
  return route.path === path
}

const closeSidebarAndNavigate = (path) => {
  isSidebarOpen.value = false
  router.push(path)
}

const pageTitle = computed(() => {
  if (route.path === '/admin') return 'Dashboard Ringkasan'
  if (route.path === '/admin/products') return 'Manajemen Produk'
  if (route.path === '/admin/orders') return 'Manajemen Pesanan'
  return 'Dashboard Admin'
})

onMounted(async () => {
  await initAuth()
  
  if (!isAdmin.value) {
    showToast('Akses ditolak. Halaman khusus Administrator!', 'error')
    router.push('/')
  }
})
</script>

<style scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
  background-color: var(--md-sys-color-background);
  width: 100%;
}

.admin-sidebar {
  width: 260px;
  background-color: var(--md-sys-color-surface-container);
  border-right: 1px solid var(--md-sys-color-outline-variant);
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  z-index: 200; /* Sidebar aman di posisi paling depan */
}

.sidebar-brand {
  height: 84px;
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 0 24px;
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--md-sys-color-primary);
  letter-spacing: 0.5px;
  border-bottom: 1px solid var(--md-sys-color-outline-variant);
  cursor: pointer;
}

.sidebar-brand span {
  font-size: 24px;
}

.sidebar-nav {
  padding: 32px 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 12px;
  height: 48px;
  padding: 0 16px;
  border-radius: var(--md-shape-corner-full);
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface-variant);
  text-align: left;
  transition: all 0.2s;
}

.nav-link:hover {
  background-color: rgba(103, 80, 164, 0.08);
  color: var(--md-sys-color-on-surface);
}

.nav-link--active {
  background-color: var(--md-sys-color-primary-container);
  color: var(--md-sys-color-on-primary-container);
}

.nav-divider {
  border: none;
  border-top: 1px solid var(--md-sys-color-outline-variant);
  margin: 16px 0;
}

/* Mengunci konten agar bergeser ke kanan sejauh 260px (Lebar sidebar) */
.admin-main {
  flex: 1;
  margin-left: 260px;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  box-sizing: border-box;
  width: calc(100% - 260px);
}

.admin-header {
  height: 64px;
  position: fixed;
  top: 0;
  right: 0;
  left: 260px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
  background-color: var(--md-sys-color-surface-container-low);
  border-bottom: 1px solid var(--md-sys-color-outline-variant);
  z-index: 90;
  transition: left 0.3s cubic-bezier(0.2, 0, 0, 1);
}

.page-title {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.menu-toggle {
  display: none;
}

/* Jarak vertikal bawaan layout admin */
.admin-content {
  padding: 96px 32px 32px 32px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.sidebar-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.4);
  z-index: 150;
  backdrop-filter: blur(2px);
}

@media (max-width: 768px) {
  .admin-sidebar {
    transform: translateX(-100%);
    transition: transform 0.3s cubic-bezier(0.2, 0, 0, 1);
  }
  
  .admin-sidebar--open {
    transform: translateX(0);
  }
  
  .menu-toggle {
    display: inline-flex;
  }
  
  .admin-header {
    left: 0;
  }
  
  .admin-main {
    margin-left: 0;
    width: 100%;
  }
  
  .admin-header {
    padding: 0 16px;
  }
  
  .text-email-header {
    display: none;
  }
  
  .admin-content {
    padding: 20px 16px;
  }
  
  .logout-btn-header span:not(.material-symbols-outlined) {
    display: none;
  }
  .logout-btn-header {
    width: 40px;
    height: 40px;
    padding: 0;
    border-radius: 50%;
  }
}
</style>
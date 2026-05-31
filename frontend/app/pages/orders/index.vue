<template>
  <NuxtLayout name="default">
    <div class="container orders-container">
      <div class="orders-header">
        <h1 style="font-size: 1.75rem; font-weight: 700;">Riwayat Pesanan Saya</h1>
        <p class="text-muted">Pantau status pengiriman dan riwayat belanja peralatan sekolah Anda.</p>
      </div>

      <div v-if="loading" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="loader"></span>
        <p class="text-muted">Memuat daftar pesanan...</p>
      </div>

      <div v-else-if="orders.length === 0" class="empty-orders flex-center" style="flex-direction: column; padding: 64px 0; gap: 16px;">
        <span class="material-symbols-outlined" style="font-size: 64px; color: var(--md-sys-color-outline-variant);">receipt_long</span>
        <h3>Belum Ada Pesanan</h3>
        <p class="text-muted">Anda belum melakukan pemesanan produk apapun.</p>
        <button class="md-btn md-btn--filled" @click="navigateTo('/')">Mulai Belanja Sekarang</button>
      </div>

      <div v-else class="orders-list">
        <div 
          v-for="order in orders" 
          :key="order.id" 
          class="md-card md-card--outlined order-card"
        >
          <div class="order-card-header">
            <div class="header-left">
              <span class="order-id">ID Pesanan: #{{ order.id.slice(0, 8).toUpperCase() }}</span>
              <span class="order-date">{{ formatDate(order.created_at) }}</span>
            </div>
            
            <span class="status-badge" :class="getStatusClass(order.status)">
              {{ getStatusLabel(order.status) }}
            </span>
          </div>

          <hr class="card-divider" />

          <div class="order-card-body">
            <div class="order-items-preview">
              <span class="item-preview-text">
                {{ order.items[0]?.product_name }} 
                <strong v-if="order.items[0]?.quantity > 1">({{ order.items[0]?.quantity }}x)</strong>
                <span v-if="order.items.length > 1" class="text-muted"> 
                  dan {{ order.items.length - 1 }} produk lainnya...
                </span>
              </span>
            </div>
            <div class="order-total-preview">
              <span class="text-muted" style="font-size: 0.8rem;">Total Belanja</span>
              <strong class="text-primary" style="font-size: 1.1rem;">
                Rp{{ order.total_amount.toLocaleString('id-ID') }}
              </strong>
            </div>
          </div>

          <div class="order-card-actions">
            <button 
              class="md-btn md-btn--outlined" 
              @click="goToDetail(order.id)"
            >
              <span class="material-symbols-outlined">receipt</span>
              Lihat Detail & Invoice
            </button>
          </div>
        </div>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useToast } from '~/composables/useToast'

const router = useRouter()
const { isLoggedIn, token, initAuth } = useAuth()
const { showToast } = useToast()

const orders = ref([])
const loading = ref(false)

const fetchOrders = async () => {
  loading.value = true
  try {
    const res = await fetch('http://localhost:5000/api/orders', {
      headers: {
        'Authorization': `Bearer ${token.value}`
      }
    })
    if (res.ok) {
      orders.value = await res.json()
    } else {
      console.error('Failed to fetch orders')
    }
  } catch (e) {
    console.error('Error fetching orders:', e)
  } finally {
    loading.value = false
  }
}

const goToDetail = (orderId) => {
  router.push(`/orders/${orderId}`)
}

const formatDate = (dateString) => {
  const d = new Date(dateString)
  return d.toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusLabel = (status) => {
  switch (status) {
    case 'Pending': return 'Menunggu Pembayaran'
    case 'Paid': return 'Sudah Dibayar'
    case 'Shipped': return 'Dalam Pengiriman'
    case 'Cancelled': return 'Dibatalkan'
    default: return status
  }
}

const getStatusClass = (status) => {
  switch (status) {
    case 'Pending': return 'badge--pending'
    case 'Paid': return 'badge--paid'
    case 'Shipped': return 'badge--shipped'
    case 'Cancelled': return 'badge--cancelled'
    default: return ''
  }
}

onMounted(async () => {
  await initAuth()
  if (!isLoggedIn.value) {
    showToast('Masuk terlebih dahulu untuk melihat riwayat pesanan Anda', 'error')
    router.push('/login')
    return
  }
  fetchOrders()
})
</script>

<style scoped>
.orders-container {
  padding-top: 32px;
  padding-bottom: 64px;
  max-width: 800px !important;
}

.orders-header {
  margin-bottom: 32px;
}

.orders-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.order-card {
  padding: 20px !important;
}

.order-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
}

.header-left {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.order-id {
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.order-date {
  font-size: 0.8rem;
  color: var(--md-sys-color-on-surface-variant);
}

.status-badge {
  font-size: 0.75rem;
  font-weight: 700;
  padding: 6px 12px;
  border-radius: var(--md-shape-corner-full);
}

.badge--pending {
  background-color: var(--md-sys-color-tertiary-container);
  color: var(--md-sys-color-on-tertiary-container);
}

.badge--paid {
  background-color: #d1e7dd;
  color: #0f5132;
}

.badge--shipped {
  background-color: #cff4fc;
  color: #055160;
}

.badge--cancelled {
  background-color: var(--md-sys-color-error-container);
  color: var(--md-sys-color-on-error-container);
}

.card-divider {
  border: none;
  border-top: 1px solid var(--md-sys-color-outline-variant);
  margin: 16px 0;
}

.order-card-body {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 16px;
}

.item-preview-text {
  font-size: 0.95rem;
  color: var(--md-sys-color-on-surface);
}

.order-total-preview {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

.order-card-actions {
  display: flex;
  justify-content: flex-end;
}

.loader {
  width: 48px;
  height: 48px;
  border: 5px solid var(--md-sys-color-primary-container);
  border-bottom-color: var(--md-sys-color-primary);
  border-radius: 50%;
  display: inline-block;
  box-sizing: border-box;
  animation: rotation 1s linear infinite;
}

@keyframes rotation {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>

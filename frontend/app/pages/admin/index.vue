<template>
  <NuxtLayout name="admin">
    <div class="admin-dashboard-container">
      <div v-if="loading" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="loader"></span>
        <p class="text-muted">Memuat data statistik...</p>
      </div>

      <div v-else>
        <div class="stats-grid">
          <div class="md-card md-card--outlined stats-card">
            <div class="stats-card-header">
              <span class="text-muted">Total Pendapatan (Lunas)</span>
              <span class="material-symbols-outlined icon-box icon-box--primary">payments</span>
            </div>
            <strong class="stats-val">Rp{{ stats.total_revenue?.toLocaleString('id-ID') || 0 }}</strong>
            <span class="stats-subtext">Menunggu: Rp{{ stats.pending_sales?.toLocaleString('id-ID') || 0 }}</span>
          </div>

          <div class="md-card md-card--outlined stats-card">
            <div class="stats-card-header">
              <span class="text-muted">Total Pesanan</span>
              <span class="material-symbols-outlined icon-box icon-box--secondary">receipt_long</span>
            </div>
            <strong class="stats-val">{{ stats.total_orders || 0 }}</strong>
            <span class="stats-subtext">Dari seluruh pelanggan</span>
          </div>

          <div class="md-card md-card--outlined stats-card">
            <div class="stats-card-header">
              <span class="text-muted">Item Produk</span>
              <span class="material-symbols-outlined icon-box icon-box--tertiary">inventory_2</span>
            </div>
            <strong class="stats-val">{{ stats.total_products || 0 }}</strong>
            <span class="stats-subtext">Katalog aktif</span>
          </div>

          <div class="md-card md-card--outlined stats-card">
            <div class="stats-card-header">
              <span class="text-muted">Total Pelanggan</span>
              <span class="material-symbols-outlined icon-box icon-box--info">group</span>
            </div>
            <strong class="stats-val">{{ stats.total_customers || 0 }}</strong>
            <span class="stats-subtext">Terdaftar di sistem</span>
          </div>
        </div>

        <div class="dashboard-main-grid">
          <div class="md-card md-card--outlined dashboard-table-card">
            <div class="card-header">
              <h3 style="font-size: 1.1rem; font-weight: 700;">Pesanan Terbaru</h3>
              <button class="md-btn md-btn--text" @click="navigateTo('/admin/orders')">Lihat Semua</button>
            </div>
            
            <div v-if="recentOrders.length === 0" class="flex-center" style="padding: 32px; flex-direction: column; gap: 8px;">
              <span class="material-symbols-outlined" style="font-size: 40px; color: var(--md-sys-color-outline-variant);">receipt</span>
              <p class="text-muted">Belum ada transaksi pesanan.</p>
            </div>

            <div v-else class="md-table-container" style="border: none; margin-bottom: 0;">
              <table class="md-table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Pelanggan</th>
                    <th>Tanggal</th>
                    <th>Status</th>
                    <th style="text-align: right;">Total</th>
                    <th style="text-align: center;">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="order in recentOrders" :key="order.id">
                    <td style="font-family: monospace; font-weight: bold;">
                      #{{ order.id.slice(0, 6).toUpperCase() }}
                    </td>
                    <td>
                      <div>
                        <strong>{{ order.customer_name }}</strong>
                        <div class="text-muted" style="font-size: 0.75rem;">{{ order.customer_phone }}</div>
                      </div>
                    </td>
                    <td>{{ formatDate(order.created_at) }}</td>
                    <td>
                      <span class="status-badge" :class="getStatusClass(order.status)">
                        {{ getStatusLabel(order.status) }}
                      </span>
                    </td>
                    <td style="text-align: right; font-weight: 600;">
                      Rp{{ order.total_amount.toLocaleString('id-ID') }}
                    </td>
                    <td style="text-align: center;">
                      <button 
                        class="md-btn md-btn--icon" 
                        @click="navigateTo(`/orders/${order.id}`)"
                        title="Detail Invoice"
                      >
                        <span class="material-symbols-outlined">receipt</span>
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <div class="md-card md-card--outlined dashboard-side-card">
            <h3 style="font-size: 1.1rem; font-weight: 700; margin-bottom: 16px;">Distribusi Kategori</h3>
            
            <div v-if="!stats.category_stats || Object.keys(stats.category_stats).length === 0" class="text-muted">
              Belum ada data kategori.
            </div>
            
            <div v-else class="category-bars-list">
              <div v-for="(count, category) in stats.category_stats" :key="category" class="category-bar-row">
                <div class="category-bar-info">
                  <span>{{ category }}</span>
                  <strong>{{ count }} Produk</strong>
                </div>
                <div class="category-bar-bg">
                  <div 
                    class="category-bar-fill" 
                    :style="{ width: `${(count / stats.total_products) * 100}%` }"
                  ></div>
                </div>
              </div>
            </div>
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

const router = useRouter()
const { token, isLoggedIn, isAdmin, initAuth } = useAuth()

const stats = ref({})
const recentOrders = ref([])
const loading = ref(false)

const fetchDashboardData = async () => {
  loading.value = true
  try {
    const statsRes = await fetch('http://localhost:5000/api/dashboard/stats', {
      headers: { 'Authorization': `Bearer ${token.value}` }
    })
    if (statsRes.ok) {
      stats.value = await statsRes.json()
    }


    const ordersRes = await fetch('http://localhost:5000/api/orders', {
      headers: { 'Authorization': `Bearer ${token.value}` }
    })
    if (ordersRes.ok) {
      const allOrders = await ordersRes.json()
      recentOrders.value = allOrders.slice(0, 5)
    }
  } catch (e) {
    console.error('Error fetching dashboard statistics:', e)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateString) => {
  const d = new Date(dateString)
  return d.toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'short',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusLabel = (status) => {
  switch (status) {
    case 'Pending': return 'Pending'
    case 'Paid': return 'Lunas'
    case 'Shipped': return 'Dikirim'
    case 'Cancelled': return 'Batal'
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
    router.push('/login')
    return
  }
  if (!isAdmin.value) {
    router.push('/')
    return
  }
  fetchDashboardData()
})
</script>

<style scoped>
.admin-dashboard-container {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 20px;
}

.stats-card {
  padding: 24px !important;
  display: flex;
  flex-direction: column;
  gap: 8px;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.stats-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.stats-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.icon-box {
  width: 38px;
  height: 38px;
  border-radius: var(--md-shape-corner-medium);
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-box--primary {
  color: var(--md-sys-color-primary);
  background-color: var(--md-sys-color-primary-container);
}

.icon-box--secondary {
  color: var(--md-sys-color-secondary);
  background-color: var(--md-sys-color-secondary-container);
}

.icon-box--tertiary {
  color: var(--md-sys-color-tertiary);
  background-color: var(--md-sys-color-tertiary-container);
}

.icon-box--info {
  color: #0288d1;
  background-color: #e1f5fe;
}

.stats-val {
  font-size: 1.75rem;
  font-weight: 800;
  color: var(--md-sys-color-on-surface);
}

.stats-subtext {
  font-size: 0.75rem;
  color: var(--md-sys-color-outline);
}

.dashboard-main-grid {
  display: grid;
  grid-template-columns: 1.8fr 1fr;
  gap: 24px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--md-sys-color-outline-variant);
  margin-bottom: 16px;
}

.dashboard-table-card {
  padding: 20px !important;
}

.dashboard-side-card {
  padding: 20px !important;
}

.category-bars-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.category-bar-row {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.category-bar-info {
  display: flex;
  justify-content: space-between;
  font-size: 0.85rem;
}

.category-bar-bg {
  width: 100%;
  height: 8px;
  border-radius: var(--md-shape-corner-full);
  background-color: var(--md-sys-color-surface-container-highest);
  overflow: hidden;
}

.category-bar-fill {
  height: 100%;
  border-radius: var(--md-shape-corner-full);
  background-color: var(--md-sys-color-primary);
}

.status-badge {
  font-size: 0.75rem;
  font-weight: 700;
  padding: 4px 8px;
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

@media (max-width: 992px) {
  .dashboard-main-grid {
    grid-template-columns: 1fr;
  }
}
</style>

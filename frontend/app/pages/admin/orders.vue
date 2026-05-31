<template>
  <NuxtLayout name="admin">
    <div class="orders-admin-container">
      <div class="md-chip-group">
        <button 
          v-for="status in statusTabs" 
          :key="status"
          class="md-chip"
          :class="{ 'md-chip--selected': selectedStatus === status }"
          @click="selectStatusFilter(status)"
        >
          {{ getStatusLabel(status) }}
        </button>
      </div>

      <div v-if="loading" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="loader"></span>
        <p class="text-muted">Memuat daftar pesanan...</p>
      </div>

      <div v-else-if="filteredOrders.length === 0" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="material-symbols-outlined" style="font-size: 64px; color: var(--md-sys-color-outline-variant);">receipt</span>
        <h3>Tidak Ada Pesanan</h3>
        <p class="text-muted">Tidak ada pesanan dengan status "{{ getStatusLabel(selectedStatus) }}".</p>
      </div>

      <div v-else class="md-table-container">
        <table class="md-table">
          <thead>
            <tr>
              <th style="width: 100px;">ID Pesanan</th>
              <th>Penerima & Kontak</th>
              <th>Tanggal</th>
              <th>Status</th>
              <th style="text-align: right;">Total</th>
              <th style="text-align: center; width: 220px;">Aksi & Status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="order in filteredOrders" :key="order.id">
              <td style="font-family: monospace; font-weight: bold;">
                #{{ order.id.slice(0, 8).toUpperCase() }}
              </td>
              <td>
                <div class="customer-info-cell">
                  <strong>{{ order.customer_name }}</strong>
                  <span class="text-muted" style="font-size: 0.8rem;">Telp: {{ order.customer_phone }}</span>
                  <span class="text-muted text-truncate-custom" :title="order.customer_address">
                    Alamat: {{ order.customer_address }}
                  </span>
                </div>
              </td>
              <td>{{ formatDate(order.created_at) }}</td>
              <td>
                <span class="status-badge" :class="getStatusClass(order.status)">
                  {{ getStatusLabel(order.status) }}
                </span>
              </td>
              <td style="text-align: right; font-weight: 600; color: var(--md-sys-color-primary);">
                Rp{{ order.total_amount.toLocaleString('id-ID') }}
              </td>
              <td style="text-align: center;">
                <div class="order-action-cell">
                  <button 
                    class="md-btn md-btn--icon" 
                    @click="navigateTo(`/orders/${order.id}`)"
                    title="Cetak Invoice"
                  >
                    <span class="material-symbols-outlined">print</span>
                  </button>

                  <div class="status-transitions">
                    <button 
                      v-if="order.status === 'Pending'"
                      class="md-btn md-btn--filled status-trans-btn"
                      style="background-color: #2e7d32;"
                      @click="updateStatus(order.id, 'Paid')"
                      :disabled="updatingId === order.id"
                    >
                      Konfirmasi Bayar
                    </button>

                    <button 
                      v-if="order.status === 'Paid'"
                      class="md-btn md-btn--filled status-trans-btn"
                      style="background-color: #0288d1;"
                      @click="updateStatus(order.id, 'Shipped')"
                      :disabled="updatingId === order.id"
                    >
                      Kirim Barang
                    </button>

                    <button 
                      v-if="order.status === 'Pending' || order.status === 'Paid'"
                      class="md-btn md-btn--outlined status-trans-btn text-error"
                      style="height: 32px; padding: 0 10px;"
                      @click="updateStatus(order.id, 'Cancelled')"
                      :disabled="updatingId === order.id"
                    >
                      Batalkan
                    </button>
                    
                    <span v-if="order.status === 'Shipped'" class="completed-text">
                      <span class="material-symbols-outlined" style="font-size: 16px;">done_all</span>
                      Selesai
                    </span>
                    
                    <span v-if="order.status === 'Cancelled'" class="cancelled-text">
                      Batal
                    </span>
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useToast } from '~/composables/useToast'

const router = useRouter()
const { token, isLoggedIn, isAdmin, initAuth } = useAuth()
const { showToast } = useToast()

const orders = ref([])
const loading = ref(false)
const updatingId = ref(null)

const selectedStatus = ref('Semua')
const statusTabs = ['Semua', 'Pending', 'Paid', 'Shipped', 'Cancelled']

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
      console.error('Failed to load orders list')
    }
  } catch (e) {
    console.error('Error fetching orders:', e)
  } finally {
    loading.value = false
  }
}

const selectStatusFilter = (status) => {
  selectedStatus.value = status
}

const filteredOrders = computed(() => {
  if (selectedStatus.value === 'Semua') return orders.value
  return orders.value.filter(o => o.status === selectedStatus.value)
})

const updateStatus = async (orderId, newStatus) => {
  updatingId.value = orderId
  try {
    const res = await fetch(`http://localhost:5000/api/orders/${orderId}/status`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token.value}`
      },
      body: JSON.stringify({ status: newStatus })
    })

    const data = await res.json()

    if (res.ok) {
      showToast(`Status pesanan berhasil diubah menjadi: ${getStatusLabel(newStatus)}`)
      fetchOrders()
    } else {
      showToast(data.error || 'Gagal merubah status pesanan', 'error')
    }
  } catch (e) {
    showToast('Terjadi kesalahan jaringan', 'error')
  } finally {
    updatingId.value = null
  }
}

const formatDate = (dateString) => {
  const d = new Date(dateString)
  return d.toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusLabel = (status) => {
  switch (status) {
    case 'Semua': return 'Semua'
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
    router.push('/login')
    return
  }
  if (!isAdmin.value) {
    router.push('/')
    return
  }
  fetchOrders()
})
</script>

<style scoped>
.orders-admin-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.customer-info-cell {
  display: flex;
  flex-direction: column;
  gap: 2px;
  max-width: 250px;
}

.text-truncate-custom {
  font-size: 0.75rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.order-action-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
}

.status-transitions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-trans-btn {
  height: 32px;
  padding: 0 12px;
  font-size: 0.8rem;
  border-radius: var(--md-shape-corner-small);
}

.completed-text {
  font-size: 0.8rem;
  color: #2e7d32;
  font-weight: 700;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}

.cancelled-text {
  font-size: 0.8rem;
  color: var(--md-sys-color-error);
  font-weight: 700;
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
</style>

<template>
  <NuxtLayout name="default">
    <div class="container invoice-wrapper">
      <div class="invoice-actions no-print">
        <button class="md-btn md-btn--text" @click="goBack">
          <span class="material-symbols-outlined">arrow_back</span>
          Kembali ke Pesanan Saya
        </button>
        <button class="md-btn md-btn--filled" @click="printInvoice">
          <span class="material-symbols-outlined">print</span>
          Cetak Invoice
        </button>
      </div>

      <div v-if="loading" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="loader"></span>
        <p class="text-muted">Memuat data invoice...</p>
      </div>

      <div v-else-if="error" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="material-symbols-outlined text-error" style="font-size: 64px;">error</span>
        <h3>Gagal Memuat Invoice</h3>
        <p class="text-muted">{{ error }}</p>
        <button class="md-btn md-btn--filled" @click="navigateTo('/orders')">Kembali</button>
      </div>

      <div v-else-if="order" class="md-card md-card--outlined invoice-card">
        <div class="invoice-header-grid">
          <div>
            <div class="invoice-brand">
              <span class="material-symbols-outlined">school</span>
              <span>FourMart</span>
            </div>
            <p class="text-muted" style="font-size: 0.8rem; margin-top: 4px;">Penyedia Peralatan Sekolah Berkualitas</p>
          </div>
          <div class="invoice-title-block">
            <h1 class="invoice-title">INVOICE</h1>
            <span class="invoice-id">#{{ order.id.toUpperCase() }}</span>
          </div>
        </div>

        <hr class="invoice-divider" />

        <div class="invoice-meta-grid">
          <div>
            <h4 class="meta-section-title">Penerbit:</h4>
            <strong>FourMart School Supplies</strong>
            <p class="meta-text">Gedung Rektorat Lt. 2, Jakarta</p>
            <p class="meta-text">Telp: 0812-3456-7890</p>
            <p class="meta-text">Email: support@fourmart.com</p>
          </div>

          <div>
            <h4 class="meta-section-title">Tujuan Pengiriman:</h4>
            <strong>{{ order.customer_name }}</strong>
            <p class="meta-text">Telp: {{ order.customer_phone }}</p>
            <p class="meta-text">Email: {{ order.customer_email }}</p>
            <p class="meta-text address-text">Alamat: {{ order.customer_address }}</p>
          </div>

          <div class="order-meta-info">
            <h4 class="meta-section-title">Informasi Pesanan:</h4>
            <div class="meta-row">
              <span>Tanggal:</span>
              <strong>{{ formatDate(order.created_at) }}</strong>
            </div>
            <div class="meta-row" style="margin-top: 4px;">
              <span>Status:</span>
              <span class="status-badge" :class="getStatusClass(order.status)">
                {{ getStatusLabel(order.status) }}
              </span>
            </div>
            <div class="meta-row" style="margin-top: 8px;">
              <span>Metode:</span>
              <strong>Transfer Bank</strong>
            </div>
          </div>
        </div>

        <div class="md-table-container invoice-table-container">
          <table class="md-table">
            <thead>
              <tr>
                <th style="width: 60px; text-align: center;">No.</th>
                <th>Nama Produk</th>
                <th style="width: 140px; text-align: right;">Harga Satuan</th>
                <th style="width: 100px; text-align: center;">Jumlah</th>
                <th style="width: 160px; text-align: right;">Subtotal</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in order.items" :key="item.product_id">
                <td style="text-align: center;">{{ index + 1 }}</td>
                <td>
                  <strong>{{ item.product_name }}</strong>
                </td>
                <td style="text-align: right;">Rp{{ item.price.toLocaleString('id-ID') }}</td>
                <td style="text-align: center;">{{ item.quantity }}</td>
                <td style="text-align: right; font-weight: 600;">
                  Rp{{ (item.price * item.quantity).toLocaleString('id-ID') }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="invoice-bottom-grid">
          <div class="invoice-instructions">
            <h4 style="font-size: 0.9rem; font-weight: 700; margin-bottom: 6px;">Catatan Pembayaran:</h4>
            <p style="font-size: 0.8rem; line-height: 1.4; color: var(--md-sys-color-on-surface-variant);">
              1. Transfer total pembayaran ke rekening BCA 8830192831 a/n FourMart.<br/>
              2. Simpan struk transfer Anda dan hubungi kami jika status pesanan tidak berubah dalam 1x24 jam.<br/>
              3. Barang yang sudah dibeli tidak dapat ditukar kecuali ada video unboxing saat paket diterima.
            </p>
          </div>

          <div class="invoice-totals">
            <div class="total-calc-row">
              <span>Subtotal Produk:</span>
              <span>Rp{{ order.total_amount.toLocaleString('id-ID') }}</span>
            </div>
            <div class="total-calc-row">
              <span>Ongkos Kirim:</span>
              <span>Rp0</span>
            </div>
            <hr style="border: none; border-top: 1px dotted var(--md-sys-color-outline-variant); margin: 8px 0;" />
            <div class="total-calc-row grand-total-row">
              <span>Total Tagihan:</span>
              <span class="text-primary">Rp{{ order.total_amount.toLocaleString('id-ID') }}</span>
            </div>
          </div>
        </div>

        <div class="invoice-footer-message">
          <p>Terima kasih telah mempercayakan kebutuhan sekolah Anda di FourMart!</p>
        </div>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useToast } from '~/composables/useToast'

const route = useRoute()
const router = useRouter()
const { token, isLoggedIn, user, initAuth } = useAuth()
const { showToast } = useToast()

const order = ref(null)
const loading = ref(false)
const error = ref('')

const fetchOrderDetails = async () => {
  loading.value = true
  error.value = ''
  try {
    const orderId = route.params.id
    const res = await fetch(`http://localhost:5000/api/orders/${orderId}`, {
      headers: {
        'Authorization': `Bearer ${token.value}`
      }
    })

    const data = await res.json()
    if (res.ok) {
      order.value = data
    } else {
      error.value = data.error || 'Gagal memuat detail pesanan'
    }
  } catch (e) {
    error.value = 'Terjadi kesalahan jaringan saat memuat invoice.'
    console.error(e)
  } finally {
    loading.value = false
  }
}

const goBack = () => {
  if (user.value?.role === 'admin') {
    router.push('/admin/orders')
  } else {
    router.push('/orders')
  }
}

const printInvoice = () => {
  window.print()
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
    showToast('Masuk terlebih dahulu untuk melihat invoice pesanan', 'error')
    router.push('/login')
    return
  }
  fetchOrderDetails()
})
</script>

<style scoped>
.invoice-wrapper {
  padding-top: 32px;
  padding-bottom: 64px;
  max-width: 900px !important;
}

.invoice-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.invoice-card {
  background-color: var(--md-sys-color-surface-container-lowest) !important;
  padding: 40px !important;
  box-shadow: 0 4px 16px var(--md-sys-color-shadow);
}

.invoice-header-grid {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  flex-wrap: wrap;
  gap: 16px;
}

.invoice-brand {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--md-sys-color-primary);
  display: flex;
  align-items: center;
  gap: 8px;
}

.invoice-brand span {
  font-size: 32px;
}

.invoice-title-block {
  text-align: right;
}

.invoice-title {
  font-size: 2rem;
  font-weight: 800;
  letter-spacing: 1px;
  color: var(--md-sys-color-on-surface);
}

.invoice-id {
  font-size: 0.95rem;
  font-family: monospace;
  color: var(--md-sys-color-on-surface-variant);
  font-weight: 600;
}

.invoice-divider {
  border: none;
  border-top: 2px solid var(--md-sys-color-outline-variant);
  margin: 28px 0;
}

.invoice-meta-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 24px;
  margin-bottom: 32px;
}

.meta-section-title {
  font-size: 0.85rem;
  color: var(--md-sys-color-outline);
  text-transform: uppercase;
  margin-bottom: 8px;
  font-weight: 700;
}

.meta-text {
  font-size: 0.85rem;
  color: var(--md-sys-color-on-surface-variant);
  line-height: 1.4;
}

.address-text {
  max-width: 250px;
  display: block;
}

.order-meta-info {
  display: flex;
  flex-direction: column;
}

.meta-row {
  display: flex;
  justify-content: space-between;
  font-size: 0.85rem;
}

.status-badge {
  font-size: 0.75rem;
  font-weight: 700;
  padding: 2px 8px;
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

.invoice-table-container {
  margin-bottom: 32px;
  border-radius: var(--md-shape-corner-small);
}

.invoice-bottom-grid {
  display: grid;
  grid-template-columns: 1.4fr 1fr;
  gap: 32px;
  align-items: start;
}

.invoice-totals {
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-size: 0.9rem;
}

.total-calc-row {
  display: flex;
  justify-content: space-between;
}

.grand-total-row {
  font-size: 1.25rem;
  font-weight: 800;
}

.invoice-footer-message {
  text-align: center;
  margin-top: 48px;
  font-size: 0.85rem;
  color: var(--md-sys-color-outline);
  border-top: 1px dashed var(--md-sys-color-outline-variant);
  padding-top: 24px;
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

/* Printing styles adjustments */
@media print {
  .invoice-card {
    border: none !important;
    padding: 0 !important;
    box-shadow: none !important;
    background-color: transparent !important;
  }
  .invoice-meta-grid {
    grid-template-columns: 1fr 1fr 1fr !important;
    gap: 12px !important;
  }
  .invoice-divider {
    margin: 16px 0 !important;
  }
}

@media (max-width: 768px) {
  .invoice-meta-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }
  .invoice-bottom-grid {
    grid-template-columns: 1fr;
  }
  .invoice-card {
    padding: 20px !important;
  }
}
</style>

<template>
  <NuxtLayout name="default">
    <div class="container checkout-container">
      <div class="checkout-header">
        <button class="md-btn md-btn--text" @click="router.back()">
          <span class="material-symbols-outlined">arrow_back</span>
          Kembali
        </button>
        <h1 style="font-size: 1.75rem; font-weight: 700; margin-top: 12px;">Checkout Pesanan</h1>
      </div>

      <div v-if="cartItems.length === 0" class="empty-checkout flex-center" style="flex-direction: column; padding: 64px 0; gap: 16px;">
        <span class="material-symbols-outlined" style="font-size: 64px; color: var(--md-sys-color-outline-variant);">shopping_cart_off</span>
        <p class="text-muted">Keranjang belanja Anda kosong, silakan pilih produk terlebih dahulu.</p>
        <button class="md-btn md-btn--filled" @click="navigateTo('/')">Kembali ke Katalog</button>
      </div>

      <div v-else class="checkout-layout">
        <div class="checkout-form-section md-card md-card--outlined">
          <h2 class="section-subtitle">Alamat Pengiriman & Kontak</h2>
          <p class="text-muted" style="margin-bottom: 24px; font-size: 0.85rem;">Harap periksa kembali detail kontak dan alamat pengiriman Anda.</p>

          <form @submit.prevent="handlePlaceOrder" class="checkout-form">
            <div class="md-field-group">
              <input 
                type="text" 
                id="custName" 
                placeholder=" " 
                v-model="custName" 
                required 
                class="md-field"
              />
              <label for="custName" class="md-field-label">Nama Penerima</label>
            </div>

            <div class="md-field-group">
              <input 
                type="tel" 
                id="custPhone" 
                placeholder=" " 
                v-model="custPhone" 
                required 
                class="md-field"
              />
              <label for="custPhone" class="md-field-label">Nomor Telepon Penerima</label>
            </div>

            <div class="md-field-group" style="margin-bottom: 24px;">
              <textarea 
                id="custAddress" 
                placeholder=" " 
                v-model="custAddress" 
                required 
                class="md-field"
                style="height: 100px; padding-top: 16px; resize: none;"
              ></textarea>
              <label for="custAddress" class="md-field-label" style="top: 24px;">Alamat Pengiriman Lengkap</label>
            </div>

            <button 
              type="submit" 
              class="md-btn md-btn--filled submit-order-btn" 
              :disabled="submitting"
            >
              <span v-if="submitting" class="btn-spinner"></span>
              <span v-else>Bayar & Konfirmasi Pesanan</span>
            </button>
          </form>
        </div>

        <div class="checkout-summary-section">
          <div class="md-card md-card--elevated summary-card">
            <h2 class="section-subtitle">Ringkasan Pesanan</h2>
            <hr style="border: none; border-top: 1px solid var(--md-sys-color-outline-variant); margin: 16px 0;" />

            <div class="summary-items">
              <div v-for="item in cartItems" :key="item.product_id" class="summary-item">
                <div class="summary-item-left">
                  <span class="item-qty">{{ item.quantity }}x</span>
                  <span class="item-name">{{ item.name }}</span>
                </div>
                <span class="item-subtotal">Rp{{ (item.price * item.quantity).toLocaleString('id-ID') }}</span>
              </div>
            </div>

            <hr style="border: none; border-top: 1px solid var(--md-sys-color-outline-variant); margin: 16px 0;" />

            <div class="summary-totals">
              <div class="total-row">
                <span>Subtotal Produk</span>
                <span>Rp{{ cartTotal.toLocaleString('id-ID') }}</span>
              </div>
              <div class="total-row">
                <span>Ongkos Kirim</span>
                <span class="text-primary" style="font-weight: 600;">GRATIS</span>
              </div>
              <hr style="border: none; border-top: 1px dotted var(--md-sys-color-outline-variant); margin: 12px 0;" />
              <div class="total-row final-total-row">
                <span>Total Pembayaran</span>
                <span class="text-primary">Rp{{ cartTotal.toLocaleString('id-ID') }}</span>
              </div>
            </div>
            
            <div class="payment-instructions md-card md-card--filled" style="margin-top: 24px; padding: 12px !important; font-size: 0.8rem; line-height: 1.4;">
              <strong>Metode Pembayaran: Transfer Bank (Manual)</strong>
              <p class="text-muted" style="margin-top: 4px;">Transfer ke Rekening BCA 8830192831 a/n FourMart. Unggah bukti pembayaran saat dihubungi admin via WA.</p>
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
import { useCart } from '~/composables/useCart'
import { useToast } from '~/composables/useToast'

const router = useRouter()
const { isLoggedIn, user, token, initAuth } = useAuth()
const { items: cartItems, cartTotal, clearCart, initCart } = useCart()
const { showToast } = useToast()

const custName = ref('')
const custPhone = ref('')
const custAddress = ref('')
const submitting = ref(false)

const handlePlaceOrder = async () => {
  submitting.value = true
  try {
    const formattedItems = cartItems.value.map(item => ({
      product_id: item.product_id,
      quantity: item.quantity
    }))

    const res = await fetch('http://localhost:5000/api/orders', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token.value}`
      },
      body: JSON.stringify({
        customer_name: custName.value,
        customer_phone: custPhone.value,
        customer_address: custAddress.value,
        items: formattedItems
      })
    })

    const data = await res.json()

    if (!res.ok) {
      throw new Error(data.error || 'Gagal membuat pesanan')
    }

    showToast('Pesanan Anda berhasil dibuat!')
    clearCart()
    
    router.push(`/orders/${data.id}`)
  } catch (err) {
    showToast(err.message || 'Terjadi kesalahan saat memproses pesanan', 'error')
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  await initAuth()
  initCart()
  
  if (!isLoggedIn.value) {
    showToast('Silakan masuk terlebih dahulu untuk melakukan checkout', 'error')
    router.push('/login')
    return
  }

  if (user.value) {
    custName.value = user.value.name || ''
    custPhone.value = user.value.phone || ''
    custAddress.value = user.value.address || ''
  }
})
</script>

<style scoped>
.checkout-container {
  padding-top: 32px;
  padding-bottom: 64px;
}

.checkout-header {
  margin-bottom: 24px;
}

.checkout-layout {
  display: grid;
  grid-template-columns: 1.4fr 1fr;
  gap: 32px;
  align-items: start;
}

.section-subtitle {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.checkout-form {
  display: flex;
  flex-direction: column;
}

.submit-order-btn {
  height: 48px;
  margin-top: 8px;
}

.summary-card {
  background-color: var(--md-sys-color-surface-container-low) !important;
}

.summary-items {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.9rem;
}

.summary-item-left {
  display: flex;
  gap: 8px;
  max-width: 75%;
}

.item-qty {
  font-weight: 700;
  color: var(--md-sys-color-primary);
}

.item-name {
  color: var(--md-sys-color-on-surface);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.item-subtotal {
  font-weight: 600;
}

.summary-totals {
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-size: 0.9rem;
}

.total-row {
  display: flex;
  justify-content: space-between;
}

.final-total-row {
  font-size: 1.15rem;
  font-weight: 800;
}

.btn-spinner {
  width: 20px;
  height: 20px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: #fff;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

@media (max-width: 768px) {
  .checkout-layout {
    grid-template-columns: 1fr;
  }
}
</style>

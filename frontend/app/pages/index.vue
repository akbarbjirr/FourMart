<template>
  <NuxtLayout name="default">
    <section class="store-hero">
      <div class="container hero-container">
        <div class="hero-text-content">
          <span class="hero-badge">Tahun Ajaran Baru 2026</span>
          <h1 class="hero-title">Persiapkan Belajar Terbaikmu Hari Ini</h1>
          <p class="hero-subtitle">FourMart menyediakan peralatan sekolah, tulis-menulis, tas premium, hingga seragam sekolah dengan kualitas terbaik dan harga bersahabat.</p>
          <div class="hero-actions">
            <a href="#katalog" class="md-btn md-btn--filled">
              Jelajahi Produk
              <span class="material-symbols-outlined">arrow_downward</span>
            </a>
          </div>
        </div>
        <div class="hero-visual">
          <div class="visual-card-wrapper">
            <img src="https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=600&auto=format&fit=crop&q=80" alt="School Supplies" class="hero-img" />
            <div class="floating-promo md-card md-card--elevated">
              <span class="material-symbols-outlined text-primary" style="font-size: 32px;">verified</span>
              <div>
                <strong>Kualitas Premium</strong>
                <p class="text-muted" style="font-size: 0.75rem;">100% Produk Original</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section id="katalog" class="catalog-section container">
      <div class="catalog-header">
        <h2 class="section-title">Katalog Peralatan Sekolah</h2>
        <p class="text-muted">Temukan semua kebutuhan sekolah dalam satu tempat.</p>
      </div>
      <div class="md-chip-group">
        <button 
          v-for="cat in categories" 
          :key="cat"
          class="md-chip"
          :class="{ 'md-chip--selected': selectedCategory === cat }"
          @click="selectCategory(cat)"
        >
          {{ cat }}
        </button>
      </div>
      <div v-if="loading" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="loader"></span>
        <p class="text-muted">Memuat katalog produk...</p>
      </div>
      <div v-else-if="products.length === 0" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="material-symbols-outlined" style="font-size: 64px; color: var(--md-sys-color-outline-variant);">search_off</span>
        <h3>Produk Tidak Ditemukan</h3>
        <p class="text-muted">Coba gunakan kata kunci pencarian atau kategori lain.</p>
        <button class="md-btn md-btn--outlined" @click="resetFilters">Reset Filter</button>
      </div>
      <div v-else class="grid-products">
        <div 
          v-for="product in products" 
          :key="product.id" 
          class="md-card md-card--outlined product-card"
        >
          <div class="product-img-wrapper" @click="openDetail(product)">
            <img :src="product.image" :alt="product.name" class="product-img" />
            <span class="product-category-tag">{{ product.category }}</span>
          </div>

          <div class="product-info" @click="openDetail(product)">
            <h3 class="product-title" :title="product.name">{{ product.name }}</h3>
            <p class="product-desc">{{ product.description }}</p>
            <div class="product-stock-status">
              <span 
                class="stock-dot" 
                :class="product.stock > 10 ? 'stock-dot--in' : product.stock > 0 ? 'stock-dot--low' : 'stock-dot--out'"
              ></span>
              <span class="text-muted" style="font-size: 0.75rem;">
                {{ product.stock > 10 ? 'Stok Tersedia' : product.stock > 0 ? `Sisa ${product.stock} Pcs` : 'Habis' }}
              </span>
            </div>
          </div>

          <div class="product-footer">
            <span class="product-price">Rp{{ product.price.toLocaleString('id-ID') }}</span>
            <button 
              class="md-btn md-btn--filled md-btn-add" 
              :disabled="product.stock <= 0"
              @click="handleAddToCart(product)"
            >
              <span class="material-symbols-outlined">add_shopping_cart</span>
              {{ product.stock > 0 ? 'Beli' : 'Habis' }}
            </button>
          </div>
        </div>
      </div>
    </section>
    <div v-if="selectedProduct" class="md-dialog-overlay" @click.self="selectedProduct = null">
      <div class="md-dialog product-detail-dialog">
        <div class="md-dialog-header-custom">
          <h3>Detail Produk</h3>
          <button class="md-btn md-btn--icon" @click="selectedProduct = null">
            <span class="material-symbols-outlined">close</span>
          </button>
        </div>

        <div class="md-dialog-content product-detail-layout">
          <img :src="selectedProduct.image" :alt="selectedProduct.name" class="detail-dialog-img" />
          <div class="detail-dialog-info">
            <span class="md-chip md-chip--selected" style="align-self: flex-start;">{{ selectedProduct.category }}</span>
            <h2 class="detail-title">{{ selectedProduct.name }}</h2>
            <p class="detail-price text-primary">Rp{{ selectedProduct.price.toLocaleString('id-ID') }}</p>
            
            <hr style="border: none; border-top: 1px solid var(--md-sys-color-outline-variant); margin: 12px 0;" />
            
            <p class="detail-desc-label">Deskripsi Produk:</p>
            <p class="detail-desc">{{ selectedProduct.description }}</p>
            
            <div class="detail-stock-section">
              <span class="text-muted">Stok: <strong>{{ selectedProduct.stock }} pcs</strong></span>
            </div>

            <div v-if="selectedProduct.stock > 0" class="detail-action-bar">
              <div class="quantity-selector">
                <button class="qty-btn" @click="detailQty = Math.max(1, detailQty - 1)">-</button>
                <span class="qty-val">{{ detailQty }}</span>
                <button class="qty-btn" @click="detailQty = Math.min(selectedProduct.stock, detailQty + 1)">+</button>
              </div>
              <button class="md-btn md-btn--filled" @click="handleAddToCart(selectedProduct, detailQty)">
                <span class="material-symbols-outlined">add_shopping_cart</span>
                Tambah ke Keranjang
              </button>
            </div>
            <div v-else class="text-error" style="font-weight: bold; margin-top: 16px;">
              Maaf, produk ini sedang habis stok.
            </div>
          </div>
        </div>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCart } from '~/composables/useCart'
import { useToast } from '~/composables/useToast'

const route = useRoute()
const router = useRouter()
const { addToCart } = useCart()
const { showToast } = useToast()

const products = ref([])
const loading = ref(false)
const selectedCategory = ref('Semua')
const categories = ['Semua', 'Tulis-Menulis', 'Tas & Kotak Pensil', 'Seragam & Aksesoris', 'Elektronik & Belajar']

const selectedProduct = ref(null)
const detailQty = ref(1)

const fetchProducts = async () => {
  loading.value = true
  try {
    let url = 'http://localhost:5000/api/products'
    const params = []
    
    if (selectedCategory.value !== 'Semua') {
      params.push(`category=${encodeURIComponent(selectedCategory.value)}`)
    }
    
    if (route.query.search) {
      params.push(`search=${encodeURIComponent(route.query.search)}`)
    }

    if (params.length > 0) {
      url += '?' + params.join('&')
    }

    const res = await fetch(url)
    if (res.ok) {
      products.value = await res.json()
    } else {
      console.error('Failed to fetch products')
    }
  } catch (e) {
    console.error('Error fetching products:', e)
  } finally {
    loading.value = false
  }
}

const selectCategory = (cat) => {
  selectedCategory.value = cat
  router.push({
    query: {
      ...route.query,
      category: cat === 'Semua' ? undefined : cat
    }
  })
}

const resetFilters = () => {
  selectedCategory.value = 'Semua'
  router.push({ query: {} })
}

const handleAddToCart = (product, qty = 1) => {
  addToCart(product, qty)
  showToast(`Berhasil menambahkan ${qty}x ${product.name} ke keranjang!`)
  if (selectedProduct.value) {
    selectedProduct.value = null
  }
}

const openDetail = (product) => {
  selectedProduct.value = product
  detailQty.value = 1
}

watch(() => [route.query.search, route.query.category], () => {
  selectedCategory.value = route.query.category || 'Semua'
  fetchProducts()
})

onMounted(() => {
  selectedCategory.value = route.query.category || 'Semua'
  fetchProducts()
})
</script>

<style scoped>
.store-hero {
  background: linear-gradient(135deg, var(--md-sys-color-primary-container) 0%, var(--md-sys-color-surface-container) 100%);
  padding: 60px 0;
  border-bottom: 1px solid var(--md-sys-color-outline-variant);
}

.hero-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 40px;
}

.hero-text-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 16px;
  max-width: 580px;
}

.hero-badge {
  background-color: var(--md-sys-color-tertiary-container);
  color: var(--md-sys-color-on-tertiary-container);
  padding: 6px 16px;
  border-radius: var(--md-shape-corner-full);
  font-size: 0.85rem;
  font-weight: 700;
  align-self: flex-start;
}

.hero-title {
  font-size: 3rem;
  line-height: 1.15;
  font-weight: 800;
  color: var(--md-sys-color-on-primary-container);
  letter-spacing: -0.5px;
}

.hero-subtitle {
  font-size: 1.1rem;
  color: var(--md-sys-color-on-surface-variant);
  line-height: 1.6;
}

.hero-visual {
  flex: 1;
  display: flex;
  justify-content: center;
}

.visual-card-wrapper {
  position: relative;
  max-width: 420px;
  width: 100%;
}

.hero-img {
  width: 100%;
  height: 300px;
  object-fit: cover;
  border-radius: var(--md-shape-corner-extra-large);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
}

.floating-promo {
  position: absolute;
  bottom: -20px;
  left: -20px;
  flex-direction: row !important;
  align-items: center;
  gap: 12px;
  padding: 12px 20px !important;
  background-color: var(--md-sys-color-surface-container-lowest) !important;
  border-radius: var(--md-shape-corner-large);
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.12) !important;
}
.catalog-section {
  padding-top: 48px;
  padding-bottom: 64px;
}

.catalog-header {
  margin-bottom: 24px;
}

.section-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}
.product-card {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 0 !important;
  cursor: pointer;
}

.product-img-wrapper {
  position: relative;
  width: 100%;
  height: 180px;
  overflow: hidden;
  background-color: var(--md-sys-color-surface-container);
}

.product-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s cubic-bezier(0.2, 0, 0, 1);
}

.product-card:hover .product-img {
  transform: scale(1.08);
}

.product-category-tag {
  position: absolute;
  top: 12px;
  left: 12px;
  background-color: rgba(255, 255, 255, 0.9);
  color: var(--md-sys-color-on-surface);
  font-size: 0.75rem;
  font-weight: 700;
  padding: 4px 8px;
  border-radius: var(--md-shape-corner-small);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.product-info {
  padding: 16px 16px 8px 16px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.product-title {
  font-size: 1rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface);
  line-height: 1.35;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.product-desc {
  font-size: 0.85rem;
  color: var(--md-sys-color-on-surface-variant);
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  margin-bottom: auto;
}

.product-stock-status {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-top: 4px;
}

.stock-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}
.stock-dot--in { background-color: #2e7d32; }
.stock-dot--low { background-color: #f57c00; }
.stock-dot--out { background-color: #d32f2f; }

.product-footer {
  padding: 4px 16px 16px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}

.product-price {
  font-size: 1.15rem;
  font-weight: 700;
  color: var(--md-sys-color-primary);
}

.md-btn-add {
  height: 36px;
  padding: 0 14px;
}

.md-dialog-header-custom {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.md-dialog-header-custom h3 {
  font-size: 1.25rem;
  font-weight: 600;
}

.product-detail-dialog {
  max-width: 780px !important;
}

.product-detail-layout {
  display: grid;
  grid-template-columns: 1fr 1.2fr;
  gap: 24px;
  margin-bottom: 0 !important;
  max-height: 70vh;
}

.detail-dialog-img {
  width: 100%;
  height: 320px;
  object-fit: cover;
  border-radius: var(--md-shape-corner-large);
  background-color: var(--md-sys-color-surface-container);
}

.detail-dialog-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
  overflow-y: auto;
  max-height: 320px;
  padding-right: 8px;
}

.detail-title {
  font-size: 1.5rem;
  font-weight: 700;
  line-height: 1.2;
}

.detail-price {
  font-size: 1.5rem;
  font-weight: 800;
}

.detail-desc-label {
  font-size: 0.9rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.detail-desc {
  font-size: 0.9rem;
  line-height: 1.6;
  color: var(--md-sys-color-on-surface-variant);
}

.detail-stock-section {
  font-size: 0.9rem;
}

.detail-action-bar {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-top: 12px;
}

.quantity-selector {
  display: flex;
  align-items: center;
  border: 1px solid var(--md-sys-color-outline);
  border-radius: var(--md-shape-corner-medium);
  overflow: hidden;
  height: 40px;
}

.qty-btn {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--md-sys-color-surface-container);
  font-weight: bold;
}
.qty-btn:hover {
  background-color: var(--md-sys-color-outline-variant);
}

.qty-val {
  padding: 0 16px;
  font-weight: 600;
  min-width: 40px;
  text-align: center;
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

@media (max-width: 768px) {
  .hero-container {
    flex-direction: column-reverse;
    text-align: center;
    padding-top: 24px;
    padding-bottom: 24px;
  }
  .hero-title {
    font-size: 2rem;
  }
  .hero-actions {
    align-self: center;
  }
  .product-detail-layout {
    grid-template-columns: 1fr;
  }
  .detail-dialog-img {
    height: 180px;
  }
}
</style>

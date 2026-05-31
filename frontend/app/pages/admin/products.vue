<template>
  <NuxtLayout name="admin">
    <div class="products-crud-container">
      <div class="crud-header">
        <div class="header-search">
          <span class="material-symbols-outlined search-icon">search</span>
          <input 
            type="text" 
            placeholder="Cari produk..." 
            v-model="searchQuery"
            @input="onSearch"
            class="search-input"
          />
        </div>
        
        <button class="md-fab" @click="openAddModal">
          <span class="material-symbols-outlined">add</span>
          Tambah Produk
        </button>
      </div>

      <div v-if="loading" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="loader"></span>
        <p class="text-muted">Memuat data produk...</p>
      </div>

      <div v-else-if="products.length === 0" class="flex-center" style="padding: 64px 0; flex-direction: column; gap: 16px;">
        <span class="material-symbols-outlined" style="font-size: 64px; color: var(--md-sys-color-outline-variant);">inventory</span>
        <h3>Produk Tidak Ditemukan</h3>
        <p class="text-muted">Tambahkan produk baru ke dalam katalog FourMart.</p>
        <button class="md-btn md-btn--filled" @click="openAddModal">Tambah Produk Pertama</button>
      </div>

      <div v-else class="md-table-container">
        <table class="md-table">
          <thead>
            <tr>
              <th style="width: 80px; text-align: center;">Gambar</th>
              <th>Nama Produk</th>
              <th>Kategori</th>
              <th style="text-align: right;">Harga</th>
              <th style="text-align: center; width: 100px;">Stok</th>
              <th style="text-align: center; width: 120px;">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in products" :key="p.id">
              <td style="text-align: center;">
                <img :src="p.image" :alt="p.name" class="table-product-img" />
              </td>
              <td>
                <div class="table-product-info">
                  <strong>{{ p.name }}</strong>
                  <p class="text-muted text-truncate-custom" :title="p.description">{{ p.description }}</p>
                </div>
              </td>
              <td>
                <span class="category-chip">{{ p.category }}</span>
              </td>
              <td style="text-align: right; font-weight: 600; color: var(--md-sys-color-primary);">
                Rp{{ p.price.toLocaleString('id-ID') }}
              </td>
              <td style="text-align: center;">
                <span :class="p.stock <= 5 ? 'stock-warning' : ''" style="font-weight: 600;">
                  {{ p.stock }}
                </span>
              </td>
              <td style="text-align: center;">
                <div class="action-buttons-cell">
                  <button class="md-btn md-btn--icon edit-btn" @click="openEditModal(p)" title="Edit Produk">
                    <span class="material-symbols-outlined">edit</span>
                  </button>
                  <button class="md-btn md-btn--icon delete-btn text-error" @click="confirmDelete(p)" title="Hapus Produk">
                    <span class="material-symbols-outlined">delete</span>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="isFormModalOpen" class="md-dialog-overlay" @click.self="closeFormModal">
        <div class="md-dialog form-dialog">
          <h2 class="md-dialog-title">
            <span class="material-symbols-outlined text-primary">
              {{ isEditing ? 'edit_note' : 'add_box' }}
            </span>
            {{ isEditing ? 'Edit Informasi Produk' : 'Tambah Produk Baru' }}
          </h2>

          <form @submit.prevent="saveProduct" class="dialog-form-layout">
            <div class="md-dialog-content">
              <div class="md-field-group">
                <input 
                  type="text" 
                  id="prodName" 
                  placeholder=" " 
                  v-model="form.name" 
                  required 
                  class="md-field"
                />
                <label for="prodName" class="md-field-label">Nama Produk</label>
              </div>

              <div class="md-field-group">
                <select 
                  id="prodCategory" 
                  v-model="form.category" 
                  required 
                  class="md-field md-select"
                >
                  <option value="" disabled selected>Pilih Kategori</option>
                  <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
                </select>
                <label for="prodCategory" class="md-field-label">Kategori</label>
              </div>

              <div class="form-row-2">
                <div class="md-field-group">
                  <input 
                    type="number" 
                    id="prodPrice" 
                    placeholder=" " 
                    v-model.number="form.price" 
                    required 
                    min="1"
                    class="md-field"
                  />
                  <label for="prodPrice" class="md-field-label">Harga (Rupiah)</label>
                </div>

                <div class="md-field-group">
                  <input 
                    type="number" 
                    id="prodStock" 
                    placeholder=" " 
                    v-model.number="form.stock" 
                    required 
                    min="0"
                    class="md-field"
                  />
                  <label for="prodStock" class="md-field-label">Jumlah Stok</label>
                </div>
              </div>

              <div class="md-field-group">
                <input 
                  type="url" 
                  id="prodImage" 
                  placeholder=" " 
                  v-model="form.image"
                  class="md-field"
                />
                <label for="prodImage" class="md-field-label">URL Gambar Produk (Opsional)</label>
                <span class="md-field-helper">Kosongkan untuk menggunakan gambar default.</span>
              </div>

              <div class="md-field-group" style="margin-bottom: 0;">
                <textarea 
                  id="prodDesc" 
                  placeholder=" " 
                  v-model="form.description" 
                  required 
                  class="md-field"
                  style="resize: none;"
                ></textarea>
                <label for="prodDesc" class="md-field-label">Deskripsi Lengkap Produk</label>
              </div>
            </div>

            <div class="md-dialog-actions">
              <button type="button" class="md-btn md-btn--text" @click="closeFormModal">Batal</button>
              <button type="submit" class="md-btn md-btn--filled" :disabled="formSubmitting">
                <span v-if="formSubmitting" class="btn-spinner"></span>
                <span v-else>{{ isEditing ? 'Simpan Perubahan' : 'Tambah Produk' }}</span>
              </button>
            </div>
          </form>
        </div>
      </div>

      <div v-if="isDeleteModalOpen" class="md-dialog-overlay" @click.self="isDeleteModalOpen = false">
        <div class="md-dialog">
          <h2 class="md-dialog-title text-error">
            <span class="material-symbols-outlined">warning</span>
            Hapus Produk?
          </h2>
          <div class="md-dialog-content">
            Apakah Anda yakin ingin menghapus produk <strong>"{{ selectedProduct?.name }}"</strong> dari katalog? Tindakan ini tidak dapat dibatalkan.
          </div>
          <div class="md-dialog-actions">
            <button class="md-btn md-btn--text" @click="isDeleteModalOpen = false">Batal</button>
            <button class="md-btn md-btn--filled text-error" style="background-color: var(--md-sys-color-error); color: white;" @click="deleteProduct" :disabled="formSubmitting">
              <span v-if="formSubmitting" class="btn-spinner"></span>
              <span v-else>Hapus</span>
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
const { token, isLoggedIn, isAdmin, initAuth } = useAuth()
const { showToast } = useToast()

const products = ref([])
const loading = ref(false)
const searchQuery = ref('')

const categories = ['Tulis-Menulis', 'Tas & Kotak Pensil', 'Seragam & Aksesoris', 'Elektronik & Belajar']

const isFormModalOpen = ref(false)
const isEditing = ref(false)
const formSubmitting = ref(false)
const isDeleteModalOpen = ref(false)
const selectedProduct = ref(null)

const form = ref({
  id: '',
  name: '',
  description: '',
  price: 0,
  stock: 0,
  image: '',
  category: ''
})

const fetchProducts = async () => {
  loading.value = true
  try {
    let url = 'http://localhost:5000/api/products'
    if (searchQuery.value) {
      url += `?search=${encodeURIComponent(searchQuery.value)}`
    }
    const res = await fetch(url)
    if (res.ok) {
      products.value = await res.json()
    } else {
      console.error('Failed to load products list')
    }
  } catch (e) {
    console.error('Error fetching products:', e)
  } finally {
    loading.value = false
  }
}

const onSearch = () => {
  fetchProducts()
}

const openAddModal = () => {
  isEditing.value = false
  form.value = {
    id: '',
    name: '',
    description: '',
    price: '',
    stock: '',
    image: '',
    category: ''
  }
  isFormModalOpen.value = true
}

const openEditModal = (product) => {
  isEditing.value = true
  form.value = { ...product }
  isFormModalOpen.value = true
}

const closeFormModal = () => {
  isFormModalOpen.value = false
}

const saveProduct = async () => {
  formSubmitting.value = true
  try {
    const url = isEditing.value 
      ? `http://localhost:5000/api/products/${form.value.id}` 
      : 'http://localhost:5000/api/products'
    
    const method = isEditing.value ? 'PUT' : 'POST'

    const res = await fetch(url, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token.value}`
      },
      body: JSON.stringify(form.value)
    })

    const data = await res.json()

    if (res.ok) {
      showToast(isEditing.value ? 'Produk berhasil diubah' : 'Produk berhasil ditambahkan!')
      closeFormModal()
      fetchProducts()
    } else {
      showToast(data.error || 'Gagal menyimpan produk', 'error')
    }
  } catch (e) {
    showToast('Terjadi kesalahan jaringan', 'error')
    console.error(e)
  } finally {
    formSubmitting.value = false
  }
}

const confirmDelete = (product) => {
  selectedProduct.value = product
  isDeleteModalOpen.value = true
}

const deleteProduct = async () => {
  formSubmitting.value = true
  try {
    const res = await fetch(`http://localhost:5000/api/products/${selectedProduct.value.id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token.value}`
      }
    })
    
    const data = await res.json()

    if (res.ok) {
      showToast('Produk berhasil dihapus')
      isDeleteModalOpen.value = false
      fetchProducts()
    } else {
      showToast(data.error || 'Gagal menghapus produk', 'error')
    }
  } catch (e) {
    showToast('Terjadi kesalahan jaringan', 'error')
  } finally {
    formSubmitting.value = false
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
  fetchProducts()
})
</script>

<style scoped>
.products-crud-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.crud-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 16px;
}

.header-search {
  position: relative;
  max-width: 320px;
  width: 100%;
}

.search-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--md-sys-color-on-surface-variant);
}

.search-input {
  width: 100%;
  height: 40px;
  padding: 0 16px 0 44px;
  background-color: var(--md-sys-color-surface-container-high);
  border: none;
  border-radius: var(--md-shape-corner-full);
  outline: none;
  font-size: 0.9rem;
}

.search-input:focus {
  background-color: var(--md-sys-color-surface-container-lowest);
  box-shadow: 0 0 0 2px var(--md-sys-color-primary);
}

.table-product-img {
  width: 48px;
  height: 48px;
  object-fit: cover;
  border-radius: var(--md-shape-corner-small);
  background-color: var(--md-sys-color-surface-container);
}

.table-product-info {
  display: flex;
  flex-direction: column;
  max-width: 320px;
}

.text-truncate-custom {
  font-size: 0.75rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.category-chip {
  font-size: 0.75rem;
  background-color: var(--md-sys-color-surface-container-high);
  color: var(--md-sys-color-on-surface-variant);
  padding: 4px 8px;
  border-radius: var(--md-shape-corner-small);
  font-weight: 500;
}

.stock-warning {
  color: var(--md-sys-color-error);
}

.action-buttons-cell {
  display: flex;
  justify-content: center;
  gap: 4px;
}

/* Modal Form Styles */
.form-dialog {
  max-width: 500px !important;
}

.form-row-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
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

@media (max-width: 576px) {
  .crud-header {
    flex-direction: column-reverse;
    align-items: stretch;
  }
  .header-search {
    max-width: none;
  }
  .md-fab {
    width: 100%;
    justify-content: center;
  }
  .form-row-2 {
    grid-template-columns: 1fr;
    gap: 0;
  }
}
</style>

<template>
  <NuxtLayout name="admin">
    <div class="dashboard-container">

      <div class="dashboard-grid">
        <div class="md-card summary-card">
          <div class="card-info">
            <span class="card-label">Total Pendapatan (Lunas)</span>
            <h3 class="card-value">Rp0</h3>
            <span class="card-sub text-muted">Menunggu: Rp0</span>
          </div>
          <div class="card-icon icon-purple">
            <span class="material-symbols-outlined">payments</span>
          </div>
        </div>

        <div class="md-card summary-card">
          <div class="card-info">
            <span class="card-label">Total Pesanan</span>
            <h3 class="card-value">0</h3>
            <span class="card-sub text-muted">Dari seluruh pelanggan</span>
          </div>
          <div class="card-icon icon-indigo">
            <span class="material-symbols-outlined">receipt_long</span>
          </div>
        </div>

        <div class="md-card summary-card">
          <div class="card-info">
            <span class="card-label">Item Produk</span>
            <h3 class="card-value">8</h3>
            <span class="card-sub text-muted">Katalog aktif</span>
          </div>
          <div class="card-icon icon-pink">
            <span class="material-symbols-outlined">inventory_2</span>
          </div>
        </div>

        <div class="md-card summary-card">
          <div class="card-info">
            <span class="card-label">Total Pelanggan</span>
            <h3 class="card-value">1</h3>
            <span class="card-sub text-muted">Terdaftar di sistem</span>
          </div>
          <div class="card-icon icon-blue">
            <span class="material-symbols-outlined">group</span>
          </div>
        </div>
      </div>

      <div class="dashboard-row">
        <div class="md-card table-card">
          <div class="card-header-inline">
            <h3 class="section-title">Pesanan Terbaru</h3>
            <button class="text-link" @click="navigateTo('/admin/orders')">Lihat Semua</button>
          </div>
          
          <hr class="card-divider" />
          
          <div class="empty-state">
            <span class="material-symbols-outlined empty-icon">receipt</span>
            <p>Belum ada transaksi pesanan.</p>
          </div>
        </div>

        <div class="md-card distribution-card">
          <h3 class="section-title" style="margin-bottom: 20px;">Distribusi Kategori</h3>
          <div class="category-list">
            <div class="category-item">
              <div class="category-info">
                <span>Elektronik & Belajar</span>
                <span>1 Produk</span>
              </div>
              <div class="progress-bar"><div class="progress" style="width: 12.5%;"></div></div>
            </div>
            <div class="category-item">
              <div class="category-info">
                <span>Seragam & Aksesoris</span>
                <span>1 Produk</span>
              </div>
              <div class="progress-bar"><div class="progress" style="width: 12.5%;"></div></div>
            </div>
            <div class="category-item">
              <div class="category-info">
                <span>Tas & Kotak Pensil</span>
                <span>2 Produk</span>
              </div>
              <div class="progress-bar"><div class="progress" style="width: 25%;"></div></div>
            </div>
            <div class="category-item">
              <div class="category-info">
                <span>Tulis-Menulis</span>
                <span>4 Produk</span>
              </div>
              <div class="progress-bar"><div class="progress" style="width: 50%;"></div></div>
            </div>
          </div>
        </div>
      </div>

    </div>
  </NuxtLayout>
</template>

<script setup>
definePageMeta({
  layout: false
})
</script>

<style scoped>
.dashboard-container {
  display: flex;
  flex-direction: column;
  gap: 48px !important;
  width: 100%;
}

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 24px;
  width: 100%;
}

.dashboard-row {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 24px;
  width: 100%;
  align-items: start;
}

.summary-card, .table-card, .distribution-card {
  background-color: var(--md-sys-color-surface);
  border: 1px solid var(--md-sys-color-outline-variant);
  border-radius: 16px;
  padding: 24px;
}

.summary-card {
  display: flex;
  justify-content: space-between;
  align-items: start;
}

.card-info {
  display: flex;
  flex-direction: column;
}

.card-label {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--md-sys-color-on-surface-variant);
  margin-bottom: 8px;
}

.card-value {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.card-sub {
  font-size: 0.75rem;
  margin-top: 4px;
}

.card-icon {
  padding: 12px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-purple { background-color: rgba(103, 80, 164, 0.08); color: #6750a4; }
.icon-indigo { background-color: rgba(63, 81, 181, 0.08); color: #3f51b5; }
.icon-pink { background-color: rgba(233, 30, 99, 0.08); color: #e91e63; }
.icon-blue { background-color: rgba(33, 150, 243, 0.08); color: #2196f3; }

.card-header-inline {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

/* STYLING GARIS ABU-ABU UNTUK KOTAK PESANAN */
.card-divider {
  border: none;
  border-top: 1px solid var(--md-sys-color-outline-variant);
  margin: 0 0 24px 0;
}

.section-title {
  font-size: 1rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.text-link {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--md-sys-color-primary);
  background: none;
  border: none;
  cursor: pointer;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: var(--md-sys-color-on-surface-variant);
  padding: 48px 0;
  gap: 12px;
}

.empty-icon {
  font-size: 36px;
}

.category-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.category-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.category-info {
  display: flex;
  justify-content: space-between;
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface);
}

.progress-bar {
  width: 100%;
  background-color: var(--md-sys-color-surface-variant);
  height: 8px;
  border-radius: 4px;
  overflow: hidden;
}

.progress {
  background-color: var(--md-sys-color-primary);
  height: 100%;
  border-radius: 4px;
}

@media (max-width: 1024px) {
  .dashboard-row {
    grid-template-columns: 1fr;
  }
}
</style>
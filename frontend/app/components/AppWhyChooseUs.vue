<script setup>
// Komponen untuk section "Mengapa Memilih Kami?"
</script>

<template>
  <section id="mengapa" class="why-choose-us-section">
    <div class="container">
      <h2 class="section-title">Mengapa Memilih FourMart?</h2>
      <p class="section-subtitle">Kami berkomitmen untuk menyediakan pengalaman belanja terbaik demi mendukung pendidikan Anda.</p>

      <div class="features-grid">
        <div class="feature-item md-card md-card--elevated">
          <span class="material-symbols-outlined feature-icon">school</span>
          <h3 class="feature-title">Kualitas Premium Terjamin</h3>
          <p class="feature-description">Setiap produk dipilih dengan cermat untuk memastikan kualitas terbaik, mendukung daya tahan dan efektivitas belajar.</p>
        </div>

        <div class="feature-item md-card md-card--elevated">
          <span class="material-symbols-outlined feature-icon">inventory_2</span>
          <h3 class="feature-title">Pilihan Produk Terlengkap</h3>
          <p class="feature-description">Dari alat tulis hingga seragam, temukan semua kebutuhan sekolah Anda di satu tempat, menghemat waktu dan tenaga.</p>
        </div>

        <div class="feature-item md-card md-card--elevated">
          <span class="material-symbols-outlined feature-icon">support_agent</span>
          <h3 class="feature-title">Layanan Pelanggan Responsif</h3>
          <p class="feature-description">Tim kami siap membantu Anda dengan ramah dan cepat, memastikan setiap pertanyaan dan masalah terselesaikan dengan baik.</p>
        </div>

        <div class="feature-item md-card md-card--elevated">
          <span class="material-symbols-outlined feature-icon">local_shipping</span>
          <h3 class="feature-title">Pengiriman Cepat & Aman</h3>
          <p class="feature-description">Nikmati pengiriman yang efisien dan aman langsung ke pintu Anda, sehingga Anda bisa fokus pada hal yang terpenting: belajar.</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.why-choose-us-section {
  padding: 80px 0;
  background-color: var(--md-sys-color-surface-container-low);
  text-align: center;
  margin-top: 48px; /* Memberi jarak dari section sebelumnya */
}

.section-title {
  font-size: 2.2rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
  margin-bottom: 16px;
}

.section-subtitle {
  font-size: 1.1rem;
  color: var(--md-sys-color-on-surface-variant);
  margin-bottom: 48px;
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.feature-item {
  padding: 24px 20px;
  border-radius: var(--md-shape-corner-large);
  background-color: var(--md-sys-color-surface-container);
  text-align: center;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.feature-item:hover {
  transform: translateY(-5px);
  box-shadow: var(--md-sys-shadow-2); /* Menggunakan level shadow Material 3 */
}

.feature-icon {
  font-size: 40px;
  color: var(--md-sys-color-primary);
  margin-bottom: 16px;
}

.feature-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface);
  margin-bottom: 12px;
}

.feature-description {
  font-size: 0.95rem;
  color: var(--md-sys-color-on-surface-variant);
  line-height: 1.6;
}

@media (max-width: 1024px) {
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 24px;
  }
}

@media (max-width: 768px) {
  .section-title { font-size: 2rem; }
  .section-subtitle { font-size: 1rem; margin-bottom: 32px; }
}

@media (max-width: 480px) {
  .features-grid {
    grid-template-columns: 1fr;
  }
  .feature-item {
    padding: 24px;
  }
}

</style>
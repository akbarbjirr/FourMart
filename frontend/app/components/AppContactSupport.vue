<script setup>
// Komponen Support & Contact untuk Tim Pengembang Proyek
</script>

<template>
  <section id="kontak" class="contact-support-section">
    <div class="container">
      <div class="contact-card-wrapper md-card md-card--outlined">
        <div class="contact-grid">
          <!-- Kolom Visual/Gambar -->
          <div class="contact-visual">
            <img 
              src="https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=800&auto=format&fit=crop&q=60" 
              alt="Tim Support FourMart" 
              class="visual-image"
            />
            <div class="visual-overlay">
              <div class="overlay-badge">
                <span class="material-symbols-outlined">group</span>
                <span>Tim Proyek Kelompok</span>
              </div>
            </div>
          </div>

          <!-- Kolom Konten & Kontak -->
          <div class="contact-content">
            <h2 class="contact-title">Butuh Bantuan atau Ingin Support Kami?</h2>
            <p class="contact-description">
              Kami adalah tim siswa yang berdedikasi untuk mempermudah akses peralatan sekolah. Jika Anda memiliki kendala atau ingin memberikan apresiasi pada proyek kami, tim kami siap merespons dengan senang hati.
            </p>

            <div class="contact-methods">
              <a href="https://wa.me/628991043020" target="_blank" class="method-item">
                <div class="method-icon-box primary">
                  <span class="material-symbols-outlined">chat</span>
                </div>
                <div class="method-text">
                  <span class="method-label">WhatsApp Support</span>
                  <span class="method-value">0899-1043-020</span>
                </div>
              </a>

              <a href="mailto:support@fourmart.com" class="method-item">
                <div class="method-icon-box secondary">
                  <span class="material-symbols-outlined">mail</span>
                </div>
                <div class="method-text">
                  <span class="method-label">Email Resmi</span>
                  <span class="method-value">support@fourmart.com</span>
                </div>
              </a>
            </div>

            <div class="support-footer">
              <span class="material-symbols-outlined icon-small">info</span>
              <p>Aktif Senin - Jumat, pukul 08.00 - 17.00 WIB</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.contact-support-section {
  padding: 64px 0;
  background-color: var(--md-sys-color-surface);
}

.contact-card-wrapper {
  overflow: hidden;
  border-radius: var(--md-shape-corner-extra-large);
  background-color: var(--md-sys-color-surface-container-low) !important;
  border: 1px solid var(--md-sys-color-outline-variant) !important;
}

.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1.2fr;
}

.contact-visual {
  position: relative;
  min-height: 350px;
  overflow: hidden;
  border-radius: var(--md-shape-corner-extra-large) 0 0 var(--md-shape-corner-extra-large);
}

.visual-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.visual-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to top, rgba(0,0,0,0.6), transparent);
  display: flex;
  align-items: flex-end;
  padding: 24px;
}

.overlay-badge {
  background-color: var(--md-sys-color-primary-container);
  color: var(--md-sys-color-on-primary-container);
  padding: 8px 16px;
  border-radius: var(--md-shape-corner-full);
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  font-size: 0.85rem;
}

.contact-content {
  padding: 48px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.contact-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
  margin-bottom: 16px;
  line-height: 1.3;
}

.contact-description {
  color: var(--md-sys-color-on-surface-variant);
  line-height: 1.6;
  margin-bottom: 32px;
}

.contact-methods {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-bottom: 32px;
}

.method-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px;
  text-decoration: none;
  border-radius: var(--md-shape-corner-large);
  transition: background-color 0.2s;
}

.method-item:hover {
  background-color: var(--md-sys-color-surface-container-high);
}

.method-icon-box {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.method-icon-box.primary {
  background-color: var(--md-sys-color-primary-container);
  color: var(--md-sys-color-on-primary-container);
}

.method-icon-box.secondary {
  background-color: var(--md-sys-color-secondary-container);
  color: var(--md-sys-color-on-secondary-container);
}

.method-text {
  display: flex;
  flex-direction: column;
}

.method-label {
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface-variant);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.method-value {
  font-size: 1rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
}

.support-footer {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--md-sys-color-outline);
  font-size: 0.8rem;
}

.icon-small {
  font-size: 16px;
}

@media (max-width: 992px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
  
  .contact-visual {
    min-height: 250px;
    border-radius: var(--md-shape-corner-extra-large) var(--md-shape-corner-extra-large) 0 0;
  }

  .contact-content {
    padding: 32px;
  }
}

@media (max-width: 480px) {
  .contact-title {
    font-size: 1.5rem;
  }
  
  .method-item {
    gap: 12px;
  }
}
</style>
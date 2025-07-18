<script setup lang="ts">
  import ItemsList from './components/ItemsList.vue'
  import PageHeader from './components/PageHeader.vue'
  import ItemsRecommendation from './components/ItemsRecommendation.vue'
  import { useItemStore } from './store/itemStore'

  import { onMounted } from 'vue'; // Importa ref para crear una referencia reactiva

  const store = useItemStore()

  onMounted(() => { // Hook de ciclo de vida que se ejecuta cuando el componente se monta
    store.fetchItems(); // Llama a la función para obtener los items al montar el componente
  });
</script>

<template>
  <PageHeader />
  <div class="mb-8 flex justify-between space-x-6 mt-4">
    <button class="border border-purple-800 rounded px-4 py-1.5 bg-purple-700 hover:bg-cyan-800"
      @click="store.changePage(store.page - 1)"
      :disabled="store.page <= 1"
      :class="{ 'opacity-50 cursor-not-allowed': store.page <= 1 }"
    >Previous</button>
    <button class="border border-purple-800 rounded px-4 py-1.5 bg-purple-700 hover:bg-cyan-800"
      @click="store.changePage(store.page + 1)"
      :disabled="store.page >= store.totalItems"
      :class="{ 'opacity-50 cursor-not-allowed': store.page >= store.totalItems }"
    >Next</button>
  </div>
  <ItemsList :items="store.items" />
  <ItemsRecommendation :items="store.itemsRecommendation" />
</template>


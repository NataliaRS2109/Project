<script setup lang="ts">
  import ItemsList from './components/ItemsList.vue'
  import PageHeader from './components/PageHeader.vue'
  import ItemsRecommendation from './components/ItemsRecommendation.vue'
  import { useItemStore } from './store/itemStore'

  import { onMounted, watch, ref } from 'vue'; // Importa ref para crear una referencia reactiva

  const store = useItemStore()
  const search = ref(""); // Declaración de una referencia de estado reactiva para almacenar el término de búsqueda

  onMounted(() => { // Hook de ciclo de vida que se ejecuta cuando el componente se monta
    store.fetchItems(); // Llama a la función para obtener los items al montar el componente
  });

  watch(search, (newText) => {
    store.filterItems(newText)
  })
</script>

<template>
  <div class="m-auto p-8 text-center">
    <PageHeader />
    <div class="my-8">
      <input type="text" class="border border-gray-300 rounded w-full p-2 bg-white" placeholder="Type to search..." 
        v-model="search"
      />
    </div>
    <div class="mb-8 flex justify-between space-x-6 mt-4 text-white">
      <button class="border border-purple-800 rounded px-4 py-1.5 bg-purple-700 hover:bg-purple-900"
        @click="store.changePage(store.page - 1)"
        :disabled="store.page <= 1"
        :class="{ 'opacity-50 cursor-not-allowed': store.page <= 1 }"
      >Previous</button>
      <button class="border border-purple-800 rounded px-4 py-1.5 bg-purple-700 hover:bg-purple-900"
        @click="store.changePage(store.page + 1)"
        :disabled="store.page >= store.totalItems"
        :class="{ 'opacity-50 cursor-not-allowed': store.page >= store.totalItems }"
      >Next</button>
    </div>

    <ItemsList :item="store.items" />
    <ItemsRecommendation :items="store.itemsRecommendation" />
  </div>

</template>


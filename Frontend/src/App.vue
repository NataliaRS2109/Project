<script setup lang="ts">
  import ItemsList from './components/ItemsList.vue'
  import PageHeader from './components/PageHeader.vue'
  import ItemsRecommendation from './components/ItemsRecommendation.vue'
  //import axiosClient from './utils/axios'; // Importación de una instancia de axios para realizar peticiones HTTP
  import type { Item } from './models/item.model'; // Importa el tipo Item desde el modelo
  import { api } from './utils/axios.ts'; // Importa la instancia de axios para realizar peticiones HTTP
  import { ref, onMounted, watch } from 'vue'; // Importa ref para crear una referencia reactiva
  const page = ref(1); // Declaración de una referencia para manejar la paginación (aunque no se usa en este código)
  const totalItems = ref(0); // Declaración de una referencia para manejar el total de elementos

  const items = ref<Item[]>([]); // Declaración de una referencia reactiva para almacenar los items obtenidos de la API
  const itemsRecommendation = ref<Item[]>([]); // Declaración de referencias reactivas para almacenar los items y las recomendaciones

  const fetchItems = async() => {
    try {
      const {data} = await api.get("/items?page=" + page.value); // Realiza una petición GET a la API
      const {data: recommendedData} = await api.get("/items/recommended"); // Realiza una petición GET a la API
      items.value = data; // Asigna los datos obtenidos a la referencia de items
      itemsRecommendation.value = recommendedData; // Asigna los datos recomendados a la referencia de itemsRecommendation
      totalItems.value = data.length; // Actualiza el total de elementos con la longitud del array de items
    } catch (error) {
      console.error("Error fetching items:", error); // Manejo de errores en caso de que la petición falle
    }
  };

  const changePage = (newPage: number) => {
    page.value = newPage; // Cambia la página actual
  }

  onMounted(() => { // Hook de ciclo de vida que se ejecuta cuando el componente se monta
    fetchItems(); // Llama a la función para obtener los items al montar el componente
  });

  watch(page, () => { // Observa los cambios en la página y vuelve a obtener los items
    fetchItems();
  });
</script>

<template>
  <PageHeader />
  <div class="mb-8 flex justify-between space-x-6 mt-4">
    <button class="border border-purple-800 rounded px-4 py-1.5 bg-purple-700 hover:bg-cyan-800"
      @click="changePage(page - 1)"
      :disabled="page <= 1"
      :class="{ 'opacity-50 cursor-not-allowed': page <= 1 }"
    >Previous</button>
    <button class="border border-purple-800 rounded px-4 py-1.5 bg-purple-700 hover:bg-cyan-800"
      @click="changePage(page + 1)"
      :disabled="page >= totalItems"
      :class="{ 'opacity-50 cursor-not-allowed': page >= totalItems }"
    >Next</button>
  </div>
  <ItemsList :items="items" />
  <ItemsRecommendation :items="itemsRecommendation" />
</template>


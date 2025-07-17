<script setup lang="ts">
import ItemsList from './components/ItemsList.vue'
//import axiosClient from './utils/axios'; // Importación de una instancia de axios para realizar peticiones HTTP
import type { Item } from './models/item.model'; // Importa el tipo Item desde el modelo
import { api } from './utils/axios.ts'; // Importa la instancia de axios para realizar peticiones HTTP
import { ref, onMounted } from 'vue'; // Importa ref para crear una referencia reactiva

const items = ref<Item[]>([]);

const fetchItems = async() => {
  try {
    const {data} = await api.get("/items"); // Realiza una petición GET a la API
    items.value = data; // Asigna los datos obtenidos a la referencia de items
  } catch (error) {
    console.error("Error fetching items:", error); // Manejo de errores en caso de que la petición falle
  }
};

onMounted(() => { // Hook de ciclo de vida que se ejecuta cuando el componente se monta
  fetchItems(); // Llama a la función para obtener los items al montar el componente
});
</script>

<template>
  <h1>List Example</h1>
  <ItemsList :items="items" />
</template>


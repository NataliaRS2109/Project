<script setup lang="ts">
    import type { Item } from '../models/item.model'; // Importa el tipo Item desde el modelo
    // Define las propiedades del componente usando defineProps
    import ItemCard from './ItemCard.vue';

    // Define una interfaz Props para tipar las propiedades del componente.
    interface Props { // Define una interfaz Props para tipar las propiedades del componente.
        items: Item[];
    }

    defineProps<Props>(); // Define las propiedades del componente usando la interfaz Props. Esto permite que el componente reciba un array de objetos Country como propiedad.
</script>

<template>
    <div class="m-15">
        <h2 class="text-4xl font-semibold text-white">Stock Recommendation</h2>
        <p class="text-gray-400 m-2 text-xl">
            Here are the best stocks for you:    
        </p>        
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-10">
        <div v-for="(item, index) in items" class="bg-white border rounded-3xl p-3 text-black relative">
            <div class="absolute -top-8 -left-4 bg-purple-700 rounded-full w-14 h-14 flex items-center justify-center text-white text-xl font-bold shadow-md">
                <p class="text-xl">{{ index + 1 }}</p> <!-- El index es el enumerador de los items recomendados -->
            </div>
            <ItemCard :item="item" />
        </div>
    </div>
</template>



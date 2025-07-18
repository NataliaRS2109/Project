<script setup lang="ts">
    import type { Item } from '../models/item.model'; // Importa el tipo Item desde el modelo
    import ItemDetail from './ItemDetail.vue'
    import { ref } from 'vue'; // Importa ref para crear una referencia reactiva

    // Define una interfaz Props para tipar las propiedades del componente.
    interface Props { // Define una interfaz Props para tipar las propiedades del componente.
        item: Item;
    }

    const showDetails = ref(0);

    const changeShowDetails = (show: number) => {
        showDetails.value = show
    }

    defineProps<Props>(); // Define las propiedades del componente usando la interfaz Props. Esto permite que el componente reciba un array de objetos Country como propiedad.
</script>


<template>
    <div class="h-fit text-left flex flex-col justify-around p-4">
        <h4 class="text-2xl font-bold">{{ item.ticker }} - {{ item.company }}</h4>
        <p><span class="font-bold">Brokerage:</span> {{ item.brokerage }}</p>
        <button 
        class="border border-purple-800 rounded px-2 py-1 mt-4 w-20 text-white bg-purple-900 hover:bg-purple-700"
        @click="changeShowDetails(showDetails===1 ? 0 : 1)">
            Details
        </button>
        <div :class="{ 'hidden': showDetails === 0 }" class="mt-4">
            <ItemDetail :item="item"/>
        </div>
    </div>
</template>
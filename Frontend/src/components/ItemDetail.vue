<script setup lang="ts">
    import type { Item } from '../models/item.model'; // Importa el tipo Item desde el modelo

    // Define una interfaz Props para tipar las propiedades del componente.
    interface Props { // Define una interfaz Props para tipar las propiedades del componente.
        item: Item;
    }

    defineProps<Props>(); // Define las propiedades del componente usando la interfaz Props. Esto permite que el componente reciba un array de objetos Country como propiedad.
</script>

<template>
    <p><span class="font-bold">Action:</span> {{ item.action }}</p>
    <p><span class="font-bold">Rating From:</span> {{ item.rating_from }}</p>
    <p><span class="font-bold">Rating To:</span> {{ item.rating_to }}</p>
    <p><span class="font-bold">Target From:</span> {{ item.target_from }}</p>
    <p><span class="font-bold">Target To:</span> {{ item.target_to }}</p>
</template>